package go4redis


const (
	READY_TO_START = 0 // Sub handler notifes that it is ready to start
	START          = 1 // Ask the sub handler to start processing
	REQUEST_ACCESS = 2 // Req sub handler to reqlinquish reading from connection
	QUIT           = 3 // Ask the sub handler to quit (unsubscribe)
	END            = 8 // Sub handler quits and sends notification...
)

func cleanUpSubscribe(c *Client) {
	c.subActive = false
	c.subCount = 0
}


func (c *Client) Subscribe(channels ...string) (int, error, chan string) {
	n := len(channels)
	consolidatedRequest, err := createRequest("SUBSCRIBE", stringsToIfaces(channels)...)
	if err != nil {
		return -1, err, nil
	}
	resp, err := c.sendRequestN(consolidatedRequest, n)
	if err != nil {
		return -1, err, nil
	}
	pubsubResp := resp[len(resp)-1]
	_, _, i, _, err := parsePubSubResp(pubsubResp)
	callbackChannel := make (chan string)

	for _, channel := range channels {
		c.chanMap[channel] = callbackChannel
	}
	c.subCount = 0
	if i > 0 {
		c.subActive = true
		c.subCount = i
	}
	return i, err, callbackChannel
}

func (c *Client) UnSubscribe(channels ...string) (int, error) {
	n := len(channels)
	if n == 0 {
		n = c.subCount
	}
	consolidatedRequest, err := createRequest("UNSUBSCRIBE", stringsToIfaces(channels)...)
	if err != nil {
		return 0, err
	}
	resp, err := c.sendRequestN(consolidatedRequest, n)
	if err != nil {
		return -1, err
	}
	pubsubResp := resp[len(resp)-1]
	_, _, i, _, err := parsePubSubResp(pubsubResp)
	if err != nil {
		return 0, err
	}

	c.subCount = i
	if i <= 0 {
		c.subActive = false
		c.subCount = 0
	}
	return i, err
}

func (c *Client) Publish(channel string, message string) (int, error) {
	val, err := c.sendRequest("PUBLISH", channel, stringify(message))
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}

func (c *Client) Channels(pattern string) ([]string, error) {
	val, err := c.sendRequest("PUBSUB CHANNELS", pattern)
	if err != nil {
		return nil, err
	}
	channels, err := ifaceToStrings(val)
	return channels, err
}
