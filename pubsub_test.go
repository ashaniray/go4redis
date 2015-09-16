package go4redis

import (
	"testing"
  "fmt"
)

func callBack(msg string, channel string, err error) {
	fmt.Println("Callback is Called!!!!")
	if (err == nil) {
		fmt.Println("Message for channel", channel, "received:", msg)
	} else {
		fmt.Println("ERROR!!!: ", err)
	}

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
    t.Errorf("In call to Subscribe expected 2 but got %d", l)
  }

	l, err = c.lpush("qqqq", 0)

	const (
		ERR_NOT_ALLOWED = "ERR only (P)SUBSCRIBE / (P)UNSUBSCRIBE / QUIT allowed in this context"
	)
	if err.Error() != ERR_NOT_ALLOWED  {
		t.Errorf("Expected error " + ERR_NOT_ALLOWED + " but got " + err.Error())
	}

  l, err = c.UnSubscribe()

  if err != nil {
    t.Errorf("expected no error while subscribe command, but got %s", err)
  }

  if l != 0 {
    t.Errorf("In call to Unsubscribe expected 0 but got %d", l)
  }


}
