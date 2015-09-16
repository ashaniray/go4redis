package go4redis

import (
	"errors"
	"strconv"
  "container/list"
	"strings"
)


func ifaceToStringFmt(anything interface{}) (string, error) {
	switch anything.(type) {
	case string:
		return ifaceToString(anything)
	default:
		val, err := ifaceToInteger(anything)
		if err != nil {
			return "", err
		}
		str := strconv.Itoa(val)
		return str, nil
	}
}

func ifaceToInteger(iface interface{}) (int, error) {
	val, ok := iface.(int)
	if ok == false {
		valString, okString := iface.(SimpleString)
		if okString == false {
			return 0, errors.New("Cannot convert to integer")
		} else {
			return 0, errors.New(valString.value)
		}
	} else {
		return val, nil
	}
}

func ifaceToString(iface interface{}) (string, error) {
	switch iface.(type) {
	case SimpleString:
		val, ok := iface.(SimpleString)

		if ok == false {
			return EMPTY_STRING, errors.New("Cannot convert from type simple string")
		} else {
			if (val.success == false) {
				return EMPTY_STRING, errors.New(val.value)
			} else {
				return val.value, nil
			}
		}
	case string:
		val, ok := iface.(string)
		if ok == false {
			return EMPTY_STRING, errors.New("Cannot convert type to string")
		}
		return val, nil
	default:
		return EMPTY_STRING, errors.New("Attempt to convert unknown type to string")
	}


}

func ifaceToStrings(iface interface{}) ([]string, error) {
  l, ok := iface.(*list.List)
	if ok == false {
		return []string{}, errors.New("Cannot convert response to array of string")
	}
  var args []string
  for e := l.Front(); e != nil; e = e.Next() {
    str, err := ifaceToString(e.Value)
    if err != nil {
      return nil, err
    }
    args = append(args, str)
  }
  return args, nil
}

func mapToIfaces(key_values map[string] string) []interface{} {
  args := []interface{}{}
  for key := range key_values {
    args = append(args, key, key_values[key])
  }
  return args
}

func stringify(str string) string {
	if str[0] == '"' {
		return str
	} else {
		return "\"" + str + "\""
	}
}

func parsePubSubResp(resp interface{}) (string, string, int, string, error) {
	l, ok := resp.(*list.List)
	if ok == false {
		return EMPTY_STRING, EMPTY_STRING, 0, EMPTY_STRING, errors.New("Cannot convert SUBSCRIBE response to array")
	}

  first := l.Front()
	second := first.Next()
	third := second.Next()

	command, ok := first.Value.(string)
	if ok == false {
		return EMPTY_STRING, EMPTY_STRING, 0, EMPTY_STRING, errors.New("Cannot convert response of PUB/SUB to command string")
	}

	channel, ok := second.Value.(string)
	if ok == false {
		return EMPTY_STRING, EMPTY_STRING, 0, EMPTY_STRING, errors.New("Cannot convert response of PUB/SUB to channel name")
	}

	switch strings.ToUpper (command) {
	case "SUBSCRIBE", "UNSUBSCRIBE":
		channelCount, ok := third.Value.(int)
		if ok == false {
			return EMPTY_STRING, EMPTY_STRING, 0, EMPTY_STRING, errors.New("Cannot convert response of PUB/SUB to channel count")
		}
		return command, channel, channelCount, EMPTY_STRING, nil
	case "MESSAGE":
		message, ok := third.Value.(string)
		if ok == false {
			return EMPTY_STRING, EMPTY_STRING, 0, EMPTY_STRING, errors.New("Cannot convert response of PUB/SUB to channel count")
		}
		return command, channel, 0, message, nil
	default:
			return EMPTY_STRING, EMPTY_STRING, 0, EMPTY_STRING, errors.New("Unknown Response Type")
	}
}

func getErrorFromResp(resp interface{}) (error) {
	val, ok := resp.(SimpleString)
	if ok == false {
		return errors.New("Error converting response to string")
	} else {
		if (val.success == false) {
			return errors.New(val.value)
		} else {
			return nil
		}
	}
}

func stringsToIfaces(xs []string) []interface{} {
	var args []interface{}
	for _, v := range xs {
		args = append(args, v)
	}
	return args
}
