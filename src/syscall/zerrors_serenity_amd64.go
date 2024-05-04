//go:build amd64 && serenity

package syscall

// Kernel/API/POSIX/errno.h

const (
	ESUCCESS = Errno(iota)
	EPERM
	ENOENT
	ESRCH
	EINTR
	EIO
	ENXIO
	E2BIG
	ENOEXEC
	EBADF
	ECHILD
	EAGAIN
	ENOMEM
	EACCES
	EFAULT
	ENOTBLK
	EBUSY
	EEXIST
	EXDEV
	ENODEV
	ENOTDIR
	EISDIR
	EINVAL
	ENFILE
	EMFILE
	ENOTTY
	ETXTBSY
	EFBIG
	ENOSPC
	ESPIPE
	EROFS
	EMLINK
	EPIPE
	ERANGE
	ENAMETOOLONG
	ELOOP
	EOVERFLOW
	EOPNOTSUPP
	ENOSYS
	ENOTIMPL
	EAFNOSUPPORT
	ENOTSOCK
	EADDRINUSE
	ENOTEMPTY
	EDOM
	ECONNREFUSED
	EHOSTDOWN
	EADDRNOTAVAIL
	EISCONN
	ECONNABORTED
	EALREADY
	ECONNRESET
	EDESTADDRREQ
	EHOSTUNREACH
	EILSEQ
	EMSGSIZE
	ENETDOWN
	ENETUNREACH
	ENETRESET
	ENOBUFS
	ENOLCK
	ENOMSG
	ENOPROTOOPT
	ENOTCONN
	ESHUTDOWN
	ETOOMANYREFS
	ESOCKTNOSUPPORT
	EPROTONOSUPPORT
	EDEADLK
	ETIMEDOUT
	EPROTOTYPE
	EINPROGRESS
	ENOTHREAD
	EPROTO
	ENOTSUP
	EPFNOSUPPORT
	EDIRINTOSELF
	EDQUOT
	ENOTRECOVERABLE
	ECANCELED
	EPROMISEVIOLATION
	ESTALE
	EMAXERRNO
)

var errors = [...]string{
	"Success (not an error)",
	"Operation not permitted",
	"No such file or directory",
	"No such process",
	"Interrupted syscall",
	"I/O error",
	"No such device or address",
	"Argument list too long",
	"Exec format error",
	"Bad fd number",
	"No child processes",
	"Try again",
	"Out of memory",
	"Permission denied",
	"Bad address",
	"Block device required",
	"Device or resource busy",
	"File already exists",
	"Cross-device link",
	"No such device",
	"Not a directory",
	"Is a directory",
	"Invalid argument",
	"File table overflow",
	"Too many open files",
	"Not a TTY",
	"Text file busy",
	"File too large",
	"No space left on device",
	"Illegal seek",
	"Read-only filesystem",
	"Too many links",
	"Broken pipe",
	"Range error",
	"Name too long",
	"Too many symlinks",
	"Overflow",
	"Operation not supported",
	"No such syscall",
	"Not implemented",
	"Address family not supported",
	"Not a socket",
	"Address in use",
	"Directory not empty",
	"Math argument out of domain",
	"Connection refused",
	"Host is down",
	"Address not available",
	"Already connected",
	"Connection aborted",
	"Connection already in progress",
	"Connection reset",
	"Destination address required",
	"Host unreachable",
	"Illegal byte sequence",
	"Message size",
	"Network down",
	"Network unreachable",
	"Network reset",
	"No buffer space",
	"No lock available",
	"No message",
	"No protocol option",
	"Not connected",
	"Transport endpoint has shutdown",
	"Too many references",
	"Socket type not supported",
	"Protocol not supported",
	"Resource deadlock would occur",
	"Timed out",
	"Wrong protocol type",
	"Operation in progress",
	"No such thread",
	"Protocol error",
	"Not supported",
	"Protocol family not supported",
	"Cannot make directory a subdirectory of itself",
	"Quota exceeded",
	"State not recoverable",
	"Operation cancelled",
	"The process has a promise violation",
	"Stale network file handle",
	"The highest errno +1 :^)",
}
