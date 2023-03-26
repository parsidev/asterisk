package agi

import (
	"bufio"
	"net"
)

type conn struct {
	net.Conn

	MaxReadBuffer int64
	reader        *bufio.Reader
	writer        *bufio.Writer
}

func (c *conn) Write(p []byte) (n int, err error) {
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
