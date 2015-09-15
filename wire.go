package go4redis

import (
	"bufio"
	"container/list"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

type Client struct {
	conn net.Conn
	reader *bufio.Reader
	chnl chan int
	subActive bool
	reqSuspendToSub bool
	subCount int
	reqQuitToSub bool
}

type SimpleString struct {
	value   string
	success bool
}

const (
	NEWLINE      = "\r\n"
	EMPTY_STRING = ""
)

func Dial(addr string) (*Client, error) {

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(conn)
	return &Client{conn: conn, reader: reader}, nil
}


func ReadLine(r *bufio.Reader) (string, error) {
	line, err := r.ReadString('\n')
	if err != nil {
		return EMPTY_STRING, err
	}

	return strings.Trim(line, NEWLINE), nil
}

func readArray(r *bufio.Reader) (interface{}, error) {
	_, err := r.ReadByte()
	if err != nil {
		return EMPTY_STRING, err
	}
	countAsStr, err := ReadLine(r)
	if err != nil {
		return EMPTY_STRING, err
	}

	arrLen, err := strconv.Atoi(countAsStr)
	if err != nil {
		return EMPTY_STRING, err
	}

	l := list.New()
	for i := 0; i < arrLen; i++ {
		val, err := readType(r)
		if err != nil {
			return EMPTY_STRING, err
		}
		l.PushBack(val)
	}
	return l, nil
}

func readNumber(r *bufio.Reader) (interface{}, error) {
	_, err := r.ReadByte()
	if err != nil {
		return EMPTY_STRING, err
	}
	value, err := ReadLine(r)
	if err != nil {
		return EMPTY_STRING, err
	}
	return strconv.Atoi(value)
}

func readSimpleString(r *bufio.Reader) (interface{}, error) {
	c, err := r.ReadByte()
	if err != nil {
		return EMPTY_STRING, err
	}
	line, err := ReadLine(r)
	if err != nil {
		return EMPTY_STRING, err
	}
	success := true
	if c == '-' {
		success = false
	}
	return SimpleString{
		value:   (strings.Trim(line, NEWLINE)),
		success: success,
	}, nil
}

func readBulkString(r *bufio.Reader) (interface{}, error) {
	_, err := r.ReadByte()
	if err != nil {
		return EMPTY_STRING, err
	}
	countAsStr, err := ReadLine(r)
	if err != nil {
		return EMPTY_STRING, err
	}
	count, err := strconv.Atoi(countAsStr)
	if err != nil {
		return EMPTY_STRING, nil
	}
	line, err := ReadLine(r)
	if err != nil {
		return EMPTY_STRING, err
	}

	if len(line) != count {
		return EMPTY_STRING, errors.New("Expected " + countAsStr + " characters in string and got " + line)
	}
	return line, nil
}

func readType(r *bufio.Reader) (interface{}, error) {
	c, err := r.Peek(1)
	if err != nil {
		return EMPTY_STRING, err
	}
	switch c[0] {
	case '+', '-':
		return readSimpleString(r)
	case ':':
		return readNumber(r)
	case '$':
		return readBulkString(r)
	case '*':
		return readArray(r)
	default:
		return EMPTY_STRING, errors.New("Invalid first token in response")
	}
}

func parseResp(r *bufio.Reader) (interface{}, error) {
	return readType(r)
}

func (c *Client) readResp2() (interface{}, error) {
	return parseResp(c.reader)
}

func (c *Client) readResp() (string, error) {
	r := c.reader
	respType, _ := r.ReadByte()
 	switch string(respType) {
	case "+":
		return ReadLine(r)
	case "-":
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

func sendRequestDone(c *Client) {
	c.chnl <- START
}

func (c *Client) sendRequest(cmd string, args ...interface{}) (interface{}, error) {
	c.prepareRequest()
	request, err := createRequest(cmd, args...)
	if err != nil {
		return nil, err
	}
	fmt.Fprintf(c.conn, request)
	return c.readResp2()
}

func (c *Client) prepareRequest() () {
	if c.subActive {
		c.chnl<-REQUEST_ACCESS
		<-c.chnl
		defer sendRequestDone(c)
	}
}

func (c *Client) sendRequestN(consolidatedRequest string, n int) ([]interface{}, error) {
	c.prepareRequest()
	fmt.Fprintf(c.conn, consolidatedRequest)
	resp := []interface{}{}
	for i := 0; i < n; i++ {
		val, err := c.readResp2()
		if err != nil {
			return resp, err
		} else {
			resp = append(resp, val)
		}
	}

	return resp, nil
}


func createRequest(cmd string, args ...interface{}) (string, error) {
	request := cmd
	for _, arg := range args {
		val, err := ifaceToStringFmt(arg)
		if err != nil {
			return EMPTY_STRING, err
		}
		request = request + " " + val
	}
	request = request + NEWLINE
	return request, nil
}

func BulkString(args ...string) string {
	result := fmt.Sprintf("*%d\r\n", len(args))
	for _, arg := range args {
		result += fmt.Sprintf("$%d\r\n%s\r\n", len(arg), arg)
	}

	return result
}
