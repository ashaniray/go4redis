package go4redis



func (c *Client) FlushDB() error {
	val, err := c.sendRequest("FLUSHDB")
	if err != nil {
		return err
	}
	_, err = ifaceToString(val)
	return err
}
