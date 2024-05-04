//go:build amd64 && serenity

package syscall

// Kernel/API/POSIX/time.h
type Timespec struct {
	Sec  int64
	Nsec int64 // TODO: long in C - is that int64 in Go?
}

// copied from ztypes_linux_amd64.go
type Timeval struct {
	Sec  int64
	Usec int64
}
