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

func (c *Client) Get(key string)(string, error) {
  val, err := c.sendRequest("GET", key)
	if err != nil {
		return EMPTY_STRING, err
	}
	i, err := ifaceToString(val)
	return i, err
}


func (c *Client) BitOp(operation string, destkey string, keys []string)(int, error) {
  args := append([]string{destkey}, keys...)
  val, err := c.sendRequest("BITOP", stringsToIfaces(args))
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}

func (c *Client) BitPos(key string, bit uint8)(int, error) {
  val, err := c.sendRequest("BITPOS", key, bit)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}

func (c *Client) BitPosWithStartRange(key string, bit uint8, start int)(int, error) {
  val, err := c.sendRequest("BITPOS", bit, start)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}
func (c *Client) BitPosWithRange(key string, bit uint8, start int, end int)(int, error) {
  val, err := c.sendRequest("BITPOS", bit, start, end)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}

func (c *Client) GetBit(key string, offset int)(int, error) {
  val, err := c.sendRequest("GETBIT", key, offset)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}
func (c *Client) GetRange(key string, start int, end int)(int, error) {
  val, err := c.sendRequest("GETRANGE", key, start, end)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}

func (c *Client) GetSet(key string, value string)(int, error) {
  val, err := c.sendRequest("GETSET", key, value)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}

func (c *Client) Incr(key string)(int, error) {
  val, err := c.sendRequest("INCR", key)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}


func (c *Client) IncrBy(key string, increment int)(int, error) {
  val, err := c.sendRequest("INCRBY", key, increment)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}

func (c *Client) IncrByFloat(key string, increment float64)(int, error) {
  val, err := c.sendRequest("INCRBYFLOAT", key, increment)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}

func (c *Client) Mget(keys []string)([]string, error) {

  val, err := c.sendRequest("MGET", stringsToIfaces(keys))
	if err != nil {
		return nil, err
	}
	arr, err := ifaceToStrings(val)
	return arr, err
}

func (c *Client) Mset(key_values map[string] string)(bool, error) {
  args := mapToIfaces(key_values)
  _, err := c.sendRequest("MSET", args)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *Client) MsetNX(key_values map[string] string)(int, error) {
  args := mapToIfaces(key_values)
  _, err := c.sendRequest("MSETNX", args)
	if err != nil {
		return 0, err
	}
  i, err := ifaceToInteger(args)
	return i, err
}

func (c *Client) Psetx(key string, milliseconds uint, value string)(string, error) {
  val, err := c.sendRequest("PSETEX", key, milliseconds, value)
	if err != nil {
		return EMPTY_STRING, err
	}
	str, err := ifaceToString(val)
	return str, err
}

func (c *Client) Set(key string, value string)(string, error) {
  val, err := c.sendRequest("SET", key, value)
	if err != nil {
		return EMPTY_STRING, err
	}
	str, err := ifaceToString(val)
	return str, err
}

func (c *Client) Setbit(key string, offset int, value int)(int, error) {
  val, err := c.sendRequest("SETBIT", key, offset, value)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}

func (c *Client) Strlen(key string)(int, error) {
  val, err := c.sendRequest("STRLEN", key)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}

func (c *Client) SetRange(key string, offset int, value string)(int, error) {
  val, err := c.sendRequest("SETRANGE", key, offset, value)
	if err != nil {
		return -1, err
	}
	i, err := ifaceToInteger(val)
	return i, err
}


/*

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



*/
