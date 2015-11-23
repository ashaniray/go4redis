package go4redis

func (c *Client) HDel(key string, fields ...string) (int, error) {
	args := append([]string{key}, fields...)
	val, err := c.sendRequest("HDEL", stringsToIfaces(args)...)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}

func (c *Client) HExists(key string, field string ) (int, error) {
	val, err := c.sendRequest("HEXISTS", key, field)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}

func (c *Client) HGet(key string, field string ) (string, error) {
	val, err := c.sendRequest("HGET", key, field)
  if err != nil {
		return EMPTY_STRING, err
	}
	str, err := ifaceToString(val)
	return str, err
}

func (c *Client) HGetAll(key string )  ([]string, error) {
	val, err := c.sendRequest("HGETALL", key)
	var args []string
  if err != nil {
		return args, err
	}
	args, err = ifaceToStrings(val)
	return args, err
}

func (c *Client) HIncrBy(key string, field string ,value int) (int, error) {
	val, err := c.sendRequest("HINCRBY", key,field, value)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}

func (c *Client) HIncrByFloat(key string, field string ,value string) (string, error) {
	val, err := c.sendRequest("HINCRBYFLOAT", key,field, value)
	if err != nil {
		return EMPTY_STRING, err
	}
	str, err := ifaceToStringFmt(val)
	return str, err
}

func (c *Client) HKeys(key string )  ([]string, error) {
	val, err := c.sendRequest("HKEYS", key)
	var args []string
  if err != nil {
		return args, err
	}
	args, err = ifaceToStrings(val)
	return args, err
}

func (c *Client) HLen(key string )  (int, error) {
	val, err := c.sendRequest("HLEN", key)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}

func (c *Client) HMGet(key string ,fields ...string)  ([]string, error) {
	args := append([]string{key}, fields...)
	val, err := c.sendRequest("HMGET", stringsToIfaces(args)...)
  if err != nil {
		return args, err
	}
	args, err = ifaceToStrings(val)
	return args, err
}

func (c *Client) HMSet(key string , fieldvalue ...string)  (string, error) {
	args := append([]string{key}, fieldvalue...)
	val, err := c.sendRequest("HMSET", stringsToIfaces(args)...)
	if err != nil {
		return EMPTY_STRING, err
	}
	str, err := ifaceToString(val)
	return str, err
}

func (c *Client) HSet(key string, field string ,value string) (int, error) {
	val, err := c.sendRequest("HSET", key,field, value)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}

func (c *Client) HSetNx(key string, field string ,value string) (int, error) {
	val, err := c.sendRequest("HSETNX", key,field, value)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}

func (c *Client) HVals(key string )  ([]string, error) {
	val, err := c.sendRequest("HVALS", key)
	var args []string
  if err != nil {
		return args, err
	}
	args, err = ifaceToStrings(val)
	return args, err
}

// ERR giving unknown command 'HSTRLEN'
// HSCAN Implementation left
