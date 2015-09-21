package go4redis

import (
	"fmt"
	"testing"
)

const (
	CHANNEL         = "chan1"
	MESSAGE         = "Message"
	ERR_NOT_ALLOWED = "ERR only (P)SUBSCRIBE / (P)UNSUBSCRIBE / QUIT allowed in this context"
)

var callbackCalled bool = false

func callBack(msg string, channel string, err error) {
	if err == nil {
		if msg == MESSAGE && channel == CHANNEL {
			callbackCalled = true
		}
	} else {
		fmt.Println("ERROR!!!: ", err)
	}
}

func publish(t *testing.T, channel chan int) {
	c, err := DialAndFlush()

	if err != nil {
		t.Errorf("Expected no error while dialing and flushing but got %s", err)
	}

	channels, err := c.Channels("")
	if len(channels) < 2 {
		t.Errorf("Expected at least 2 channels in call to Channels but got %d", len(channels))
	}

	val, err := c.Publish(CHANNEL, MESSAGE)
	if val != 1 {
		t.Errorf("Expected 1 recevier for Publish but got %d", val)
	}

	channel <- 0
}

func TestSubscribe(t *testing.T) {
	c, err := DialAndFlush()

	if err != nil {
		t.Errorf("Expected no error while dialing and flushing but got %s", err)
	}

	l, err := c.Subscribe(callBack, CHANNEL, "chan2")

	if err != nil {
		t.Errorf("expected no error while subscribe command, but got %s", err)
	}

	if l != 2 {
		t.Errorf("In call to Subscribe expected 2 but got %d", l)
	}

	l, err = c.lpush("qqqq", 0)

	if err.Error() != ERR_NOT_ALLOWED {
		t.Errorf("Expected error " + ERR_NOT_ALLOWED + " but got " + err.Error())
	}
	channel := make(chan int)
	go publish(t, channel)
	<-channel

	l, err = c.UnSubscribe()

	if err != nil {
		t.Errorf("expected no error while subscribe command, but got %s", err)
	}

	if l != 0 {
		t.Errorf("In call to Unsubscribe expected 0 but got %d", l)
	}

	if callbackCalled == false {
		t.Errorf("Callback not called for publish")
	}
}
