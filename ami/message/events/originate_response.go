package events

import (
	"errors"
	"fmt"
)

type OriginateResponseEvent struct {
	ActionID     string
	Response     string
	Channel      string
	Context      string
	Exten        string
	Application  string
	Data         string
	Reason       string
	Uniqueid     string
	CallerIDNum  string
	CallerIDName string
}

func (OriginateResponseEvent) EventTypeName() string {
	return "OriginateResponse"
}

func (evt OriginateResponseEvent) Err() error {
	if evt.Response == "Failure" {
		return errors.New(fmt.Sprintf("Originate failed: exten %v reason %v", evt.Exten, evt.Reason))
	}
	return nil
}
