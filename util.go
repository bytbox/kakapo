package main

import (
	"bufio"
	"fmt"
	"io"
)

// An io.Reader that prompts to os.Stdout at appropriate times.
type PromptingReader struct {
	src   *bufio.Reader
	ready <-chan interface{}
}

func NewPromptingReader(r io.Reader) (PromptingReader, chan<- interface{}) {
	src := bufio.NewReader(r)
	c := make(chan interface{})
	return PromptingReader{src, c}, c
}

func (r PromptingReader) Read(p []byte) (int, error) {
	fmt.Printf("> ")
	tmp := make([]byte, len(p))
	n, err := r.src.Read(tmp)
	if err != nil {
		return n, err
	}
	copy(p, tmp)
	return n, err
}
