//go:build wasi || wasip1

package syscall

func Setrlimit(resource int, rlim *Rlimit) error {
	panic("unimplemented: syscall.Setrlimit")
}