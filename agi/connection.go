package agi

import (
	"bufio"
	"net"
	"time"
)

type conn struct {
	net.Conn

	IdleTimeout   time.Duration
	MaxReadBuffer int64

	reader *bufio.Reader
	writer *bufio.Writer
}

func (c *conn) Write(p []byte) (n int, err error) {
	c.updateDeadline()
	n, err = c.writer.Write(p)
	if err != nil {
		return 0, nil
	}
	err = c.writer.Flush()
	return
}

func (c *conn) Close() (err error) {
	err = c.Conn.Close()
	return
}

func (c *conn) updateDeadline() {
	idleDeadline := time.Now().Add(c.IdleTimeout)
	_ = c.Conn.SetDeadline(idleDeadline)
}
