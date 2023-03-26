package agi

import (
	"net"
)

type conn struct {
	net.Conn

	MaxReadBuffer int64
}
