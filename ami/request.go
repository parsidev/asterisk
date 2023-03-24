package ami

import (
	"context"
	"time"

	"github.com/parsidev/asterisk/ami/message"
	"github.com/pkg/errors"
)

type requestOptions struct {
	Timeout    time.Duration
	OnComplete func(ctx context.Context, msg *message.Message, err error)
}
type RequestOption func(o *requestOptions) error

func RequestTimeout(d time.Duration) RequestOption {
	return func(o *requestOptions) error {
		o.Timeout = d
		return nil
	}
}

func RequestResponseCallback(cb func(ctx context.Context, msg *message.Message, err error)) RequestOption {
	return func(o *requestOptions) error {
		o.OnComplete = cb
		return nil
	}
}

func (c *Conn) Request(r interface{}, opts ...RequestOption) (resp *message.Message, err error) {
	var msg *message.Message
	msg, err = message.ConvertToMessage(r)
	if err != nil {
		return nil, err
	}
	if msg.Type != message.MessageTypeAction {
		return nil, errors.Errorf("can only request action: %v", msg.Type)
	}

	async := &asyncMsg{
		id:     c.nextID(),
		msg:    msg,
		result: make(chan *asyncMsg, 1),
		ctx:    context.Background(),
	}
	o := requestOptions{
		Timeout: time.Second * 30,
	}
	for _, opt := range opts {
		if err = opt(&o); err != nil {
			return
		}
	}

	onComplete := o.OnComplete
	if onComplete != nil {
		async.cb = func(v *asyncMsg) {
			onComplete(v.ctx, v.resp, v.err)
		}
	}

	msg.SetAttr(attrActionID, async.id)

	if async.cb == nil {
		var cancel context.CancelFunc
		async.ctx, cancel = context.WithTimeout(async.ctx, o.Timeout)
		defer cancel()
	}

	c.pending <- async

	if async.cb != nil {
		return nil, errors.New("No response yet")
	}

	select {
	case <-async.ctx.Done():
		return nil, async.ctx.Err()
	case <-async.result:
		err = async.err
		if err == nil && async.resp != nil {
			err = async.resp.Error()
		}
		return async.resp, err
	}
}