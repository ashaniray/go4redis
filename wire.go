package go4redis

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strings"
)

type Client struct {
	conn net.Conn
}

func Dial(addr string) (*Client, error) {

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &Client{conn: conn}, nil
}

func ReadLine(r *bufio.Reader) (string, error) {
	line, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.Trim(line, "\r\n"), nil
}

func (c *Client) readResp() (string, error) {
	r := bufio.NewReader(c.conn)
	respType, _ := r.ReadByte()

	switch string(respType) {
	case "+":
		return ReadLine(r)
	case ":":
		return ReadLine(r)
	case "$":
		_, err := ReadLine(r)
		if err != nil {
			return "", err
		}
		return ReadLine(r)
	default:
		return "", errors.New("Protocol error")

	}

}

func BulkString(args ...string) string {
	result := fmt.Sprintf("*%d\r\n", len(args))
	for _, arg := range args {
		result += fmt.Sprintf("$%d\r\n%s\r\n", len(arg), arg)
	}

	return result
}
