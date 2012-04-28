package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"syscall"
	"unsafe"

	. "kakapo/lisp"
)

const VERSION = `0.4`

var (
	version = flag.Bool("V", false, "Display version information and exit")
)

func main() {
	flag.Parse()

	if *version {
		fmt.Printf("Kakapo %s\n", VERSION)
		return
	}

	// Expose imports
	for name, pkg := range _go_imports {
		ExposeImport(name, pkg)
	}

	// Expose globals
	ExposeGlobal("-interpreter", "Kakapo")
	ExposeGlobal("-interpreter-version", VERSION)

	if len(flag.Args()) > 0 {
		for _, fname := range flag.Args() {
			f, err := os.Open(fname)
			if err != nil {
				panic(err)
			}
			EvalFrom(f)
		}
		return
	}

	if IsTerminal(int(os.Stdin.Fd())) {
		// Start the read-eval-print loop
		EvalFrom(strings.NewReader(repl))
	} else {
		EvalFrom(os.Stdin)
	}
}

// IsTerminal returns true if the given file descriptor is a terminal.
// http://goo.gl/PbmRK
func IsTerminal(fd int) bool {
        var termios syscall.Termios
        _, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(fd), uintptr(syscall.TCGETS), uintptr(unsafe.Pointer(&termios)), 0, 0, 0)
        return err == 0
}

