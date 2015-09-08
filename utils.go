package go4redis

import (
  "errors"
  "strconv"
)

func intfToStringFmt (anything interface{}) (string, error) {
  switch anything.(type) {
      case string:
          return intfToString(anything)
      default:
        val, err := intfToInteger(anything)
        if err != nil {
          return "", err
        }
        str := strconv.Itoa(val)
        return str, nil
		}
}

func intfToInteger (intf interface{}) (int, error) {
  val, ok := intf.(int)
	if (ok == false) {
		return 0, errors.New("Cannot convert response to interger")
	} else {
    return val, nil
  }
}

func intfToString (intf interface{}) (string, error) {
  val, ok := intf.(string)
	if (ok == false) {
		return EMPTY_STRING, errors.New("Cannot convert response to string")
	} else {
    return val, nil
  }
}
