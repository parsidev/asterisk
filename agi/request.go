package agi

import (
	"bufio"
	"fmt"
)

type Request struct {
	srv           *Server
	writer        *bufio.Writer
	reader        *bufio.Reader
	args          map[string]string
	network       string
	networkString string
	request       string
	channel       string
	language      string
	aType         string
	uniqueID      string
	version       string
	callerID      string
	callerIdName  string
	calleringPres string
	calleringAni2 string
	calleringTon  string
	calleringTns  string
	dnID          string
	rdnis         string
	context       string
	extension     string
	priority      string
	enhanced      string
	accountCode   string
	threadID      string
}

func (r *Request) SendCommand(command string) {
	_, err := r.writer.WriteString(fmt.Sprintf("%s\n", command))
	if err != nil {
		r.srv.logger.Error(err.Error())
	}
}

func (r *Request) Busy() {
	r.SendCommand("EXEC Hangup 17")
}
