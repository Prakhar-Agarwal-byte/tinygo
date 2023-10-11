//go:build wasi || wasip1

package net

// LookupHost looks up the given host using the local resolver.
// It returns a slice of that host's addresses.
//
// LookupHost uses context.Background internally; to specify the context, use
// Resolver.LookupHost.
func LookupHost(host string) (addrs []string, err error) {
	panic("unimplemented: net.LookupHost")
}