package command

import (
	"fmt"
	"strconv"
	"strings"
)

type Command interface {
	CommandString() string
}

func joinCommand(s []interface{}) string {
	sb := strings.Builder{}
	for _, param := range s {
		switch v := param.(type) {
		case string:
			sb.WriteString(v)
		case int:
			sb.WriteString(strconv.Itoa(v))
		case float64:
			sb.WriteString(fmt.Sprint(v))
		case *string:
			if v == nil {
				goto DONE
			}
			sb.WriteString(*v)
		case *int:
			if v == nil {
				goto DONE
			}
			sb.WriteString(strconv.Itoa(*v))
		case *float64:
			if v == nil {
				goto DONE
			}
			sb.WriteString(fmt.Sprint(*v))
		}
		sb.WriteRune(' ')
	}
DONE:
	return strings.TrimSpace(sb.String())
}
