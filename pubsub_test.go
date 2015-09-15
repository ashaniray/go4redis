package go4redis

import (
	"testing"
  "fmt"
)

func callBack(msg string, err error) {
    fmt.Println(msg)
    fmt.Println(err)
}

func TestSubscribe(t *testing.T) {
	c, err := DialAndFlush()

	if err != nil {
		t.Errorf("Expected no error while dialing and flushing but got %s", err)
	}

	l, err := c.Subscribe(callBack, "chan1", "chan2")

	if err != nil {
		t.Errorf("expected no error while subscribe command, but got %s", err)
	}

  if l != 2 {
    t.Errorf("expect 2 but got %d", l)
  }
  
}
