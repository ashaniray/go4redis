package go4redis

import (
	"fmt"
)

func (c *Client) FlushDB() error {
	sc := BulkString("FLUSHDB")
	fmt.Fprintf(c.conn, sc)

	_, err := c.readResp()
	return err
}
