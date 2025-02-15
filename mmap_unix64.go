//go:build unix && (amd64 || arm64)
// +build unix
// +build amd64 arm64

package gommap

import "syscall"

func mmap_syscall(addr, length, prot, flags, fd uintptr, offset int64) (uintptr, error) {
	addr, _, err := syscall.Syscall6(syscall.SYS_MMAP, addr, length, prot, flags, fd, uintptr(offset))
	return addr, err
}
