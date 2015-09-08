package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

func abort(format string, a ...interface{}) {
	panic(fmt.Sprintf(format, a))
}

var (
	_kernel32, _             = syscall.LoadLibrary("kernel32.dll")
	_outputdebugstringhdl, _ = syscall.GetProcAddress(_kernel32, "OutputDebugStringW")
)

func OutputDebug(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	s += "\n"

	syscall.Syscall9(uintptr(_outputdebugstringhdl), 1,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s))),
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0)
	return
}

func main() {
	for _, a := range os.Args[1:] {
		OutputDebug(a)
	}
}
