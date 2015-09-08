package go4redis

import (
  "errors"
)

func toInteger (intf interface{}) (int, error) {
  length, ok := intf.(int)
	if (ok == false) {
		return 0, errors.New("Cannot convert response to interger")
	} else {
    return length, nil
  }
}
