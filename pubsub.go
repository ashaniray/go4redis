package go4redis

import (
  "time"
  "bufio"
  )

const (
  READY_TO_START = 0 // Sub handler notifes that it is ready to start
  START = 1 // Ask the sub handler to start processing
  REQUEST_ACCESS = 2 // Req sub handler to reqlinquish reading from connection
  QUIT = 3 // Ask the sub handler to quit (unsubscribe)
  END = 8 // Sub handler quits and sends notification...
)

func cleanUpSubscribe(c *Client) {
  //c.conn.SetReadDeadline(0)
  c.reqQuitToSub = false
  c.reqSuspendToSub = false
  c.subActive = false

  c.chnl <- END
  close(c.chnl)
}

func handleSubscribe(c *Client, f func (message string, err error)()) {
  defer close(c.chnl)
  c.subActive = true
  c.chnl <- READY_TO_START
  <-c.chnl
  r := bufio.NewReader(c.conn)
  for ;c.reqQuitToSub != true; {
    c.conn.SetReadDeadline(time.Now().Add(time.Second))
    //Read till timeout
    val, err := readType(r)
    if err != nil {
      _, _, _, msg, err := parsePubSubResp(val)
      go f(msg, err)
    }
    if (c.reqSuspendToSub == true) {
      c.reqSuspendToSub = false
      c.chnl <- READY_TO_START
      <-c.chnl
    }
  }
}


func (c *Client) subscribeSingleChannel(f func (message string, err error)(), channel string)(int, error) {
  val, err := c.sendRequest("SUBSCRIBE", channel)
	if err != nil {
		return -1, err
	}
	if (c.subActive == false) {
    c.chnl = make(chan int)
    go handleSubscribe(c, f);
    <-c.chnl
    c.chnl <- START
  }
  _, _, i, _, err := parsePubSubResp(val)
	return i, err
}

func (c *Client) Subscribe(f func (message string, err error)(), channels ...string) (int, error){
  var val int;
  var err error;
  for _, channel := range channels {
    val, err = c.subscribeSingleChannel(f, channel)
  }
  return val, err
}


func (c *Client) UnSubscribe(channels ...string)(int, error) {
  val, err := c.sendRequest("UNSUBSCRIBE", stringsToIfaces(channels))
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)

  //if (err != 0 && i )

  // Wait for subsribe go routine to shutdown
  <-c.chnl

	return i, err
}


func (c *Client) Publish(channel string, message string)(int, error) {
  val, err := c.sendRequest("PUBLISH", channel, stringify(message))
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}
