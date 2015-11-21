package go4redis

func (c *Client) Del(keys ...string) (int, error) {

	args := stringsToIfaces(keys)

	val, err := c.sendRequest("DEL", args...)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}

// Returns if key exists
// The number of keys existing among the ones specified as arguments.
// Keys mentioned multiple times and existing are counted multiple times.
func (c *Client) Exists(keys ...string) (int, error) {
	args := stringsToIfaces(keys)
	val, err := c.sendRequest("EXISTS", args...)

	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}

func (c *Client) Keys(pattern string) ([]string, error) {
	val, err := c.sendRequest("KEYS", pattern)

	if err != nil {
		return nil, err
	}

	arr, err := ifaceToStrings(val)

	return arr, err
}

/*
func (c *Client) Dump(key string) (string, error) { }

func (c *Client) Expire( key seconds) (int, error) { }
func (c *Client) Expireat( key timestamp) (int, error) { }

func (c *Client) Migrate( host port key destination-db timeout [COPY] [REPLACE]) (int, error) { }
func (c *Client) Move( key db) (int, error) { }
func (c *Client) Object( subcommand [arguments [arguments ...]]) (int, error) { }
func (c *Client) Persist( key) (int, error) { }
func (c *Client) Pexpire( key milliseconds) (int, error) { }
func (c *Client) Pexpireat( key milliseconds-timestamp) (int, error) { }
func (c *Client) Pttl( key) (int, error) { }
func (c *Client) Randomkey( ) (int, error) { }
func (c *Client) Rename( key newkey) (int, error) { }
func (c *Client) Renamenx( key newkey) (int, error) { }
func (c *Client) Restore( key ttl serialized-value [REPLACE]) (int, error) { }
func (c *Client) Sort( key [BY pattern] [LIMIT offset count] [GET pattern [GET pattern ...]] [ASC|DESC] [ALPHA] [STORE destination]) (int, error) { }
func (c *Client) Ttl( key) (int, error) { }
func (c *Client) Type( key) (int, error) { }
func (c *Client) Wait( numslaves timeout) (int, error) { }
func (c *Client) Scan( cursor [MATCH pattern] [COUNT count]) (int, error) { }
*/
