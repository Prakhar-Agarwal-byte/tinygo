package net

import (
	"context"
	"time"
)

type Dialer struct {
	Timeout   time.Duration
	Deadline  time.Time
	DualStack bool
	KeepAlive time.Duration
}

func Dial(network, address string) (Conn, error) {
	return nil, ErrNotImplemented
}

func Listen(network, address string) (Listener, error) {
	return nil, ErrNotImplemented
}

func (d *Dialer) DialContext(ctx context.Context, network, address string) (Conn, error) {
	return nil, ErrNotImplemented
}

// DialTimeout acts like Dial but takes a timeout.
//
// The timeout includes name resolution, if required.
// When using TCP, and the host in the address parameter resolves to
// multiple IP addresses, the timeout is spread over each consecutive
// dial, such that each is given an appropriate fraction of the time
// to connect.
//
// See func Dial for a description of the network and address
// parameters.
func DialTimeout(network, address string, timeout time.Duration) (Conn, error) {
	panic("unimplemented: net.DialTimeout")
}

