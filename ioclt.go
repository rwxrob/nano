// +build !aix
// +build !js
// +build !nacl
// +build !plan9
// +build !windows
// +build !android
// +build !solaris

package main

import (
	"syscall"
	"unsafe"
)

func getwinsize() winsize {
	ws := winsize{}
	syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(0), uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&ws)))
	return ws
}
