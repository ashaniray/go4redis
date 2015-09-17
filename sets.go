package go4redis


func (c *Client) SAdd(key string, member ...string) (int, error) {
  args := append([]string{key}, member...)
  val, err := c.sendRequest("SADD", stringsToIfaces(args)...);
   if err != nil {
    return -1, err
  }
  i, err := ifaceToInteger(val)
  return i, err
}

func (c *Client) SCard(key string) (int, error) {
  val, err := c.sendRequest("SCARD", key);
   if err != nil {
    return -1, err
  }
  i, err := ifaceToInteger(val)
  return i, err
}

func (c *Client) SDiff(key ...string) (int, error) {
  val, err := c.sendRequest("SDIFF", stringsToIfaces(key)...);
   if err != nil {
    return -1, err
  }
  i, err := ifaceToInteger(val)
  return i, err
}

func (c *Client) SDIFFSTORE(destination string, key ...string) (int, error) {
  args := append([]string{destination}, key...)
  val, err := c.sendRequest("SDIFFSTORE", stringsToIfaces(args)...);
   if err != nil {
    return -1, err
  }
  i, err := ifaceToInteger(val)
  return i, err
}

func (c *Client) SINTER(key ...string) (int, error) {
  val, err := c.sendRequest("SINTER", stringsToIfaces(key)...);
   if err != nil {
    return -1, err
  }
  i, err := ifaceToInteger(val)
  return i, err
}

func (c *Client) SINTERSTORE(destination string, key ...string) (int, error) {
  args := append([]string{destination}, key...)
  val, err := c.sendRequest("SINTERSTORE", stringsToIfaces(args)...);
   if err != nil {
    return -1, err
  }
  i, err := ifaceToInteger(val)
  return i, err
}

func (c *Client) SISMEMBER(key string, member string) (int, error) {
  val, err := c.sendRequest("SISMEMBER", key, member);
   if err != nil {
    return -1, err
  }
  i, err := ifaceToInteger(val)
  return i, err
}

func (c *Client) SMEMBERS(key string) (int, error) {
  val, err := c.sendRequest("SMEMBERS", key);
   if err != nil {
    return -1, err
  }
  i, err := ifaceToInteger(val)
  return i, err
}

func (c *Client) SMOVE(source string, destination string, member string) (int, error) {
  val, err := c.sendRequest("SMOVE", source, destination, member);
   if err != nil {
    return -1, err
  }
  i, err := ifaceToInteger(val)
  return i, err
}

func (c *Client) SPOP(key string, count string) (int, error) {
  val, err := c.sendRequest("SPOP", key, count);
   if err != nil {
    return -1, err
  }
  i, err := ifaceToInteger(val)
  return i, err
}

func (c *Client) SRANDMEMBER(key string, count string) (int, error) {
  val, err := c.sendRequest("SRANDMEMBER", key, count);
   if err != nil {
    return -1, err
  }
  i, err := ifaceToInteger(val)
  return i, err
}

func (c *Client) SREM(key string, member ...string) (int, error) {
  args := append([]string{key}, member...)
  val, err := c.sendRequest("SREM", stringsToIfaces(args)...);
   if err != nil {
    return -1, err
  }
  i, err := ifaceToInteger(val)
  return i, err
}

func (c *Client) SUNION(key ...string) (int, error) {
  val, err := c.sendRequest("SUNION", stringsToIfaces(key)...);
   if err != nil {
    return -1, err
  }
  i, err := ifaceToInteger(val)
  return i, err
}

func (c *Client) SUNIONSTORE(destination string, key ...string) (int, error) {
  args := append([]string{destination}, key...)
  val, err := c.sendRequest("SUNIONSTORE", stringsToIfaces(args)...);
   if err != nil {
    return -1, err
  }
  i, err := ifaceToInteger(val)
  return i, err
}
/*
func (c *Client) SSCAN(key string, cursor string, MATCH string, pattern string, COUNT string, count string) {
  val, err := c.sendRequest("SSCAN", key, cursor, [MATCH, pattern], [COUNT, count]);
   if err != nil {
    return -1, err
  }
  i, err := ifaceToInteger(val)
  return i, err
}
*/
