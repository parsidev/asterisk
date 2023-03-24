package message

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/parsidev/asterisk/ami/message/actions"
)

type MessageType string

const (
	MessageTypeAction   MessageType = "Action"
	MessageTypeEvent                = "Event"
	MessageTypeResponse             = "Response"
)

const (
	EOL = "\r\n"
	EOM = "\r\n\r\n"
	EOB = '\n'
	SEP = ":"
	SPC = " "
	EMP = ""
)

type Message struct {
	Type       MessageType
	Name       string
	Attributes map[string]interface{}
}

func (m *Message) Read(r *bufio.Reader) (err error) {
	return ReadMessage(m, r)
}

func (m *Message) Write(w io.Writer) (err error) {
	return WriteMessage(m, w)
}

func (m *Message) Format() string {
	b := bytes.Buffer{}
	_ = m.Write(&b)
	return b.String()
}

func (m *Message) AttrString(name string) string {
	if m.Attributes == nil {
		return ""
	}
	msg := m.Attributes[name]
	if msg == nil {
		return ""
	}
	return fmt.Sprint(msg)
}

func (m *Message) Message() string {
	return m.AttrString("Message")
}

func (m *Message) Success() bool {
	return m.Type == MessageTypeResponse && m.Name == "Success"
}

func (m *Message) Error() error {
	if m.Type == MessageTypeResponse && m.Name == "Error" {
		msg := m.Message()
		if msg == "" {
			msg = "error response"
		}
		return errors.New(msg)
	}
	return nil
}

func (m *Message) SetAttr(name string, val interface{}) {
	if m.Attributes == nil {
		m.Attributes = make(map[string]interface{})
	}
	m.Attributes[name] = val
}

func ConvertToMessage(in interface{}) (msg *Message, err error) {
	msg = &Message{}
	err = nil

	switch a := in.(type) {
	case Message:
		return &a, nil
	case *Message:
		return a, nil
	case actions.Action:
		msg.Type = MessageTypeAction
		msg.Name = a.ActionTypeName()
	default:
		return nil, errors.New(fmt.Sprintf("invalid type: %T", in))
	}

	m := make(map[string]interface{})
	err = mapstructure.Decode(in, &m)

	for k, v := range m {
		rm := v == nil

		switch v := v.(type) {
		case string:
			rm = v == ""
		case int:
			rm = v == 0
		}
		if rm {
			delete(m, k)
		}
	}

	msg.Attributes = m
	return
}

func MustConvertToMessage(in interface{}) (msg *Message, err error) {
	return ConvertToMessage(in)
}

func ReadMessage(m *Message, rd io.Reader) (err error) {
	r := bufio.NewReader(rd)

	m.Attributes = map[string]interface{}{}
	line := EMP
	line, err = r.ReadString(EOB)
	if err != nil {
		return err
	}

AGAIN:
	line = strings.TrimSuffix(line, EOL)
	if line == EMP {
		goto AGAIN
	}

	sp := strings.SplitN(line, SEP, 2)
	if len(sp) != 2 {
		return errors.New(fmt.Sprintf("invalid type line read: %q", line))
	}

	m.Type = MessageType(sp[0])
	m.Name = strings.TrimSpace(sp[1])

	switch m.Type {
	case MessageTypeAction:
	case MessageTypeEvent:
	case MessageTypeResponse:
	default:
		return errors.New(fmt.Sprintf("invalid message type: %q", sp[0]))
	}

	var stack [][]string
	for {
		line, err = r.ReadString(EOB)
		if err != nil {
			return err
		}
		if line == EOL {
			break
		}

		sp = strings.SplitN(line, SEP, 2)
		switch {
		case (len(sp) == 2 && strings.HasSuffix(line, EOL)) || (len(stack) == 0 && len(sp) == 2):
			stack = append(stack, sp)
		case len(stack) != 0:
			stack = append(stack, []string{EMP, line})
		}
	}

	var k, v string
	for _, pair := range stack {
		switch {
		case pair[0] != EMP && k != EMP:
			m.Attributes[k] = strings.TrimSuffix(v, EOL)
			k = EMP
			v = EMP
			fallthrough
		case pair[0] != EMP:
			k = pair[0]
			v = strings.TrimLeft(pair[1], SPC)
		case pair[0] == EMP:
			v += pair[1]
		}
	}

	if k != EMP {
		m.Attributes[k] = strings.TrimSuffix(v, EOL)
	}
	return
}

func WriteMessage(m *Message, w io.Writer) (err error) {
	wr := bufio.NewWriter(w)
	_, _ = wr.WriteString(fmt.Sprintf("%v: %v\r\n", m.Type, m.Name))

	for k, v := range m.Attributes {
		_, _ = wr.WriteString(fmt.Sprintf("%v: %v\r\n", k, v))
	}
	_, _ = wr.WriteString(EOL)
	err = wr.Flush()
	return
}
