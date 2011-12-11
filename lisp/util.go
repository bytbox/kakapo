package lisp

import (
	"errors"
	"io"
	"unicode/utf8"
)

func GetRuneScanner(r io.Reader) io.RuneScanner {
	sr := &scanRune{reader: r}
	return sr
}

// readRune is a structure to enable reading UTF-8 encoded code points
// from an io.Reader.  It is used if the Reader given to the scanner does
// not already implement io.RuneReader.
type scanRune struct {
	reader  io.Reader
	buf     [utf8.UTFMax]byte // used only inside ReadRune
	pending int               // number of bytes in pendBuf; only >0 for bad UTF-8
	pendBuf [utf8.UTFMax]byte // bytes left over
	last    rune
	lastSz  int
	hasLast bool
	hasUnr  bool
}

// readByte returns the next byte from the input, which may be
// left over from a previous read if the UTF-8 was ill-formed.
func (r *scanRune) readByte() (b byte, err error) {
	if r.pending > 0 {
		b = r.pendBuf[0]
		copy(r.pendBuf[0:], r.pendBuf[1:])
		r.pending--
		return
	}
	_, err = r.reader.Read(r.pendBuf[0:1])
	return r.pendBuf[0], err
}

// unread saves the bytes for the next read.
func (r *scanRune) unread(buf []byte) {
	copy(r.pendBuf[r.pending:], buf)
	r.pending += len(buf)
}

// ReadRune returns the next UTF-8 encoded code point from the
// io.Reader inside r.
func (r *scanRune) ReadRune() (rr rune, size int, err error) {
	if r.hasUnr {
		r.hasUnr = false
		return r.last, r.lastSz, nil
	}
	rr, size, err = r.readRune()
	if err != nil {
		return
	}
	r.last = rr
	r.lastSz = size
	r.hasLast = true
	r.hasUnr = false
	return
}

func (r *scanRune) readRune() (rr rune, size int, err error) {
	r.buf[0], err = r.readByte()
	if err != nil {
		return 0, 0, err
	}
	if r.buf[0] < utf8.RuneSelf { // fast check for common ASCII case
		rr = rune(r.buf[0])
		return
	}
	var n int
	for n = 1; !utf8.FullRune(r.buf[0:n]); n++ {
		r.buf[n], err = r.readByte()
		if err != nil {
			if err == io.EOF {
				err = nil
				break
			}
			return
		}
	}
	rr, size = utf8.DecodeRune(r.buf[0:n])
	if size < n { // an error
		r.unread(r.buf[size:n])
	}
	return
}

func (r *scanRune) UnreadRune() error {
	if !r.hasLast {
		return errors.New("Bad UnreadRune")
	}
	r.hasLast = false
	r.hasUnr = true
	return nil
}
