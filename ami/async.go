package ami

import (
	"context"

	"github.com/parsidev/asterisk/ami/message"
)

type asyncMsg struct {
	id     string
	msg    *message.Message
	resp   *message.Message
	err    error
	result chan *asyncMsg
	cb     func(v *asyncMsg)
	ctx    context.Context
	done   bool
}

func (async *asyncMsg) complete() {
	if async.done {
		return
	}
	if async.result != nil {
		async.result <- async
		close(async.result)
	}
	if async.cb != nil {
		go func() {
			async.cb(async)
		}()
	}
	async.done = true
}
