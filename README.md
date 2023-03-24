[doc-img]: http://img.shields.io/badge/GoDoc-Reference-blue.svg

[doc]: https://pkg.go.dev/github.com/parsidev/asterisk

[report-card-img]: https://goreportcard.com/badge/github.com/parsidev/asterisk

[report-card]: https://goreportcard.com/report/github.com/parsidev/asterisk

[license-img]: https://img.shields.io/github/license/parsidev/asterisk.svg

[license]: https://github.com/parsidev/asterisk/blob/main/LICENSE

[last-commit]: https://img.shields.io/github/last-commit/parsidev/asterisk

[last-release]: https://img.shields.io/github/release-date/parsidev/asterisk

[downloads]: https://img.shields.io/github/downloads/parsidev/asterisk/total
[go-version]: https://img.shields.io/github/go-mod/go-version/parsidev/asterisk

# Asterisk library for GO

![downloads]
![last-release]
![last-commit]
![go-version]
[![License][license-img]][license]

Libs for Golang to work with Asterisk

* AMI
* AGI (coming soon)


## AMI

* Async Action
* Sync Action
* Event Subscribe
* Auto Reconnect

```go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/parsidev/asterisk/ami"
	"github.com/parsidev/asterisk/ami/actions"
)

var msg chan *ami.Message

func main() {
	msg = make(chan *ami.Message)

	go func() {
		for {
			select {
			case d := <-msg:
				fmt.Println(d)
			}
		}
	}()

	conn, err := ami.Connect("172.16.202.129:5038",
		ami.WithAuth("parsa", "password"),
		ami.WithOnConnected(onConnected), ami.WithOnConnectError(onConnectError),
		ami.WithSubscribe(ami.SubscribeChan(msg, "FullyBooted", "StatusComplete", "SuccessfulAuth", "SessionTimeout")))

	if err != nil {
		panic(err)
	}

	response, err := conn.Request(actions.StatusAction{}, ami.RequestTimeout(2*time.Second))
	if err != nil {
		panic(err)
	}

	fmt.Println(response.Format())

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	<-exit
}

func onConnected(c *ami.Conn) {

}

func onConnectError(c *ami.Conn, err error) {
	fmt.Println(err)
}

```


