package go4redis


// LLEN key
// Get the length of a list
func (c *Client) llen(key string) (int, error) {
	val, err := c.sendRequest("LLEN", key)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}

// LPUSH key value [value ...]
// Prepend one or multiple values to a list
func (c *Client) lpush(key string, values ...interface{}) (int, error) {
	args := append([]interface{}{}, key)
	args = append(args, values...)

	val, err := c.sendRequest("LPUSH", args...)
	if err != nil {
		return -1, err
	}

	i, err := ifaceToInteger(val)
	return i, err
}

// BLPOP key [key ...] timeout
// Remove and get the first element in a list, or block until one is available

// BRPOP key [key ...] timeout
// Remove and get the last element in a list, or block until one is available

// BRPOPLPUSH source destination timeout
// Pop a value from a list, push it to another list and return it; or block until one is available

// LINDEX key index
// Get an element from a list by its index

func (c *Client) lindex(key string, index int) (string, error) {
	val, err := c.sendRequest("LINDEX", key, index)
	if err != nil {
		return EMPTY_STRING, err
	}
	i, err := ifaceToString(val)
	return i, err
}

// LINSERT key BEFORE|AFTER pivot value
// Insert an element before or after another element in a list

// LPOP key
// Remove and get the first element in a list

func (c *Client) lpop(key string) (string, error) {
	val, err := c.sendRequest("LPOP", key)
	if err != nil {
		return EMPTY_STRING, err
	}
	i, err := ifaceToString(val)
	return i, err
}

// LPUSHX key value
// Prepend a value to a list, only if the list exists

// LRANGE key start stop
// Get a range of elements from a list

// LREM key count value
// Remove elements from a list

// LSET key index value
// Set the value of an element in a list by its index

func (c *Client) lset(key string, idx int, value string) error {
	val, err := c.sendRequest("LSET", key, idx, value)
	if err != nil {
		return err
	}
	if err == nil {
		err = getErrorFromResp(val)
	}
	return err
}

// LTRIM key start stop
// Trim a list to the specified range

// RPOP key
// Remove and get the last element in a list
func (c *Client) rpop(key string) (string, error) {
	val, err := c.sendRequest("RPOP", key)
	if err != nil {
		return EMPTY_STRING, err
	}
	i, err := ifaceToString(val)
	return i, err
}

// RPOPLPUSH source destination
// Remove the last element in a list, prepend it to another list and return it

// RPUSH key value [value ...]
// Append one or multiple values to a list
func (c *Client) rpush(key string, values ...string) (int, error) {

	args := append([]interface{}{}, key)
	args = append(args, stringsToIfaces(values)...)

	val, err := c.sendRequest("RPUSH", args...)
	if err != nil {
		return -1, err
	}

	i, err := ifaceToInteger(val)
	return i, err

}

// RPUSHX key value
// Append a value to a list, only if the list exists
