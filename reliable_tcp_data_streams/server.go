package ch03

import (
	"net"
	"testing"
)

func TestListener(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = listener.Close() }()

	t.Logf("bound to %q", listener.Addr())

	/*
		 * Example for how you would then handle an incoming connection:
			for {
				conn, err := listener.Accept()
				if err != nil {
					return err
				}

				go func(c net.Conn) {
					defer c.Close()

					// Code goes here to handle connection
				}(conn)
			}
	*/
}
