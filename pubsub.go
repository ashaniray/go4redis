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
  c.subCount = 0
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

func (c *Client) Subscribe(f func (message string, err error)(), channels ...string) (int, error){
  n := len(channels)
  consolidatedRequest, err := createRequest("SUBSCRIBE", stringsToIfaces(channels))
  if err != nil {
    return 0, err
  }
  resp, err := c.sendRequestN(consolidatedRequest, n)
	if err != nil {
		return -1, err
	}
	if (c.subActive == false) {
    c.chnl = make(chan int)
    go handleSubscribe(c, f);
    <-c.chnl
    c.chnl <- START
  }
  pubsubResp := resp[len(resp) - 1]
  _, _, i, _, err := parsePubSubResp(pubsubResp)
  c.subCount = i
	return i, err
}


func (c *Client) UnSubscribe(channels ...string)(int, error) {
  n := len(channels)
  if n == 0 {
    n = c.subCount
  }
  consolidatedRequest, err := createRequest("UNSUBSCRIBE", stringsToIfaces(channels))
  if err != nil {
    return 0, err
  }
  resp, err := c.sendRequestN(consolidatedRequest, n)
	if err != nil {
		return -1, err
	}
  pubsubResp := resp[len(resp) - 1]
  _, _, i, _, err := parsePubSubResp(pubsubResp)
  c.chnl <- START
  if err != nil {
    return 0, err
  }
  if i == 0 {
    c.reqQuitToSub = true
    // Wait for subsribe go routine to send END
    <-c.chnl
  }
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
