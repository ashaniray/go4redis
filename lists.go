package go4redis

import (
	"fmt"
	"strconv"
)

// LLEN key
// Get the length of a list
func (c *Client) llen(key string) (int, error) {
	sc := BulkString("LLEN", key)
	fmt.Fprintf(c.conn, sc)

	r, err := c.readResp()

	if err != nil {
		return 0, err
	}

	return strconv.Atoi(r)
}

// LPUSH key value [value ...]
// Prepend one or multiple values to a list
func (c *Client) lpush(key string, values ...string) (int, error) {

	args := []string{"LPUSH", key}

	for _, value := range values {
		args = append(args, value)
	}

	sc := BulkString(args...)
	fmt.Fprintf(c.conn, sc)

	r, err := c.readResp()

	if err != nil {
		return 0, err
	}

	return strconv.Atoi(r)
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
	sc := BulkString("LINDEX", key, strconv.Itoa(index))
	fmt.Fprintf(c.conn, sc)

	r, err := c.readResp()

	if err != nil {
		return "", err
	}

	return r, nil
}

// LINSERT key BEFORE|AFTER pivot value
// Insert an element before or after another element in a list

// LPOP key
// Remove and get the first element in a list

func (c *Client) lpop(key string) (string, error) {
	sc := BulkString("LPOP", key)
	fmt.Fprintf(c.conn, sc)

	r, err := c.readResp()

	if err != nil {
		return "", err
	}

	return r, nil
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
	sc := BulkString("LSET", key, strconv.Itoa(idx), value)
	fmt.Fprintf(c.conn, sc)

	_, err := c.readResp()

	return err
}

// LTRIM key start stop
// Trim a list to the specified range

// RPOP key
// Remove and get the last element in a list
func (c *Client) rpop(key string) (string, error) {
	sc := BulkString("RPOP", key)
	fmt.Fprintf(c.conn, sc)

	r, err := c.readResp()

	if err != nil {
		return "", err
	}

	return r, nil
}

// RPOPLPUSH source destination
// Remove the last element in a list, prepend it to another list and return it

// RPUSH key value [value ...]
// Append one or multiple values to a list
func (c *Client) rpush(key string, values ...string) (int, error) {

	args := []string{"RPUSH", key}

	for _, value := range values {
		args = append(args, value)
	}

	sc := BulkString(args...)
	fmt.Fprintf(c.conn, sc)

	r, err := c.readResp()

	if err != nil {
		return 0, err
	}

	return strconv.Atoi(r)
}

// RPUSHX key value
// Append a value to a list, only if the list exists
