package go4redis


func (c *Client) Append(key string, value string)(int, error) {
  val, err := c.sendRequest("APPEND", key, value)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}

func (c *Client) BitCount(key string)(int, error) {
  val, err := c.sendRequest("BITCOUNT", key)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}

func (c *Client) BitCountWithIndex(key string, start int, end int)(int, error) {
  val, err := c.sendRequest("BITCOUNT", key, start, end)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}


func (c *Client) BitOp(operation string, destkey string, key string)(int, error) {
  val, err := c.sendRequest("BITOP", operation, destkey, key)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}

func (c *Client) Decr(key string)(int, error) {
  val, err := c.sendRequest("DECR", key)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}

func (c *Client) DecrBy(key string, decrement int)(int, error) {
  val, err := c.sendRequest("DECRBY", key, decrement)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}

func (c *Client) GET(key string)(string, error) {
  val, err := c.sendRequest("GET", key)
	if err != nil {
		return EMPTY_STRING, err
	}
	i, err := ifaceToString(val)
	return i, err
}

/*
func (c *Client) BITOP(operation, destkey, key key ...)(int, error) {
  val, err := c.sendRequest("BITOP", destkey, key key ...)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}

func (c *Client) BITPOS(key, bit)(int, error) {
  val, err := c.sendRequest("BITPOS", key, bit)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}

func (c *Client) BITPOS(key, bit, start)(int, error) {
  val, err := c.sendRequest("BITPOS", bit, start)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}
func (c *Client) BITPOS(key, bit, start, end)(int, error) {
  val, err := c.sendRequest("BITPOS", bit, start, end)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}


func (c *Client) GETBIT(key, offset)(int, error) {
  val, err := c.sendRequest("GETBIT", key, offset)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}
func (c *Client) GETRANGE(key, start, end)(int, error) {
  val, err := c.sendRequest("GETRANGE", key, start, end)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}
func (c *Client) GETSET(key, value)(int, error) {
  val, err := c.sendRequest("GETSET", key, value)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}
func (c *Client) INCR(key)(int, error) {
  val, err := c.sendRequest("INCR", key)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}
func (c *Client) INCRBY(key, increment)(int, error) {
  val, err := c.sendRequest("INCRBY", key, increment)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}
func (c *Client) INCRBYFLOAT(key, increment)(int, error) {
  val, err := c.sendRequest("INCRBYFLOAT", key, increment)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}
func (c *Client) MGET(key)(int, error) {
  val, err := c.sendRequest("MGET", key)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}
func (c *Client) MGET(key, key ...)(int, error) {
  val, err := c.sendRequest("MGET", key, key ...)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}
func (c *Client) MSET(key, value)(int, error) {
  val, err := c.sendRequest("MSET", key, value)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}
func (c *Client) MSET(key, value key value ...)(int, error) {
  val, err := c.sendRequest("MSET", key, value key value ...)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}
func (c *Client) MSETNX(key, value)(int, error) {
  val, err := c.sendRequest("MSETNX", key, value)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}
func (c *Client) MSETNX(key, value key value ...)(int, error) {
  val, err := c.sendRequest("MSETNX", key, value key value ...)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}
func (c *Client) PSETEX(key, milliseconds, value)(int, error) {
  val, err := c.sendRequest("PSETEX", key, milliseconds, value)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}
func (c *Client) SET(key value)(int, error) {
  val, err := c.sendRequest("SET", key value)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}
func (c *Client) SETBIT(key, offset, value)(int, error) {
  val, err := c.sendRequest("SETBIT", key, offset, value)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}
func (c *Client) SETEX(key, seconds, value)(int, error) {
  val, err := c.sendRequest("SETEX", key, seconds, value)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}
func (c *Client) SETNX(key, value)(int, error) {
  val, err := c.sendRequest("SETNX", key, value)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}
func (c *Client) SETRANGE(key, offset, value)(int, error) {
  val, err := c.sendRequest("SETRANGE", key, offset, value)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}
func (c *Client) STRLEN(key)(int, error) {
  val, err := c.sendRequest("STRLEN", key)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}

*/
