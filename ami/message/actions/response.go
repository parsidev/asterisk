package actions

import "errors"

type Response struct {
	Response  string
	ActionID  string
	Message   string
	Timestamp string
	Headers   map[string]string
	Error     error
}

func (res *Response) Err() error {
	if res.Error != nil {
		return res.Error
	}
	if res.Response == "Error" {
		return errors.New(res.Message)
	}
	return nil
}
