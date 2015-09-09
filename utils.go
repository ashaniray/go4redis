package go4redis

import (
	"errors"
	"strconv"
  "container/list"
)

// func stringArrayToInterfaceArray(args []string) ([]interface{}) {
//   ifaceArray := make([]interface{}, len(args))
//   for i, v := range args {
//       ifaceArray[i] = v
//   }
//   return ifaceArray
// }


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
		return 0, errors.New("Cannot convert response to interger")
	} else {
		return val, nil
	}
}

func ifaceToString(iface interface{}) (string, error) {
	val, ok := iface.(string)
	if ok == false {
		return EMPTY_STRING, errors.New("Cannot convert response to string")
	} else {
		return val, nil
	}
}

func ifaceToStrings(iface interface{}) ([]string, error) {
  l, ok := iface.(list.List)
	if ok == false {
		return []string{}, errors.New("Cannot convert response to array of string")
	}
  var args []string
  for e := l.Front(); e != nil; e = e.Next() {
		iface = e.Value
    str, err := ifaceToString(iface)
    if err != nil {
      return nil, err
    }
    args = append(args, str)
  }
  return args, nil
}


func stringsToIfaces(xs []string) []interface{} {
	var args []interface{}

	for _, v := range xs {
		args = append(args, v)
	}
	return args
}
