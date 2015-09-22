package go4redis

import (
	"bufio"
	"container/list"
	"errors"
	"net"
	"strconv"
	"strings"
	"fmt"
)

type Response struct {
	val interface{}
	err error
}

type Client struct {
	conn            net.Conn
	reader          *bufio.Reader

	subCount        int //Active subsciptions. Needed to determine number of response in UnSubscribe etc
	subActive       bool // Is Subscription Active ?
	readChannel     chan *Response // Channel for communication between recvr and sender
	chanMap map[string] chan string // Map of redis channel vs go channel to send back published messages
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
	readChannel := make (chan *Response)

	c := &Client{
				conn: conn,
				reader: reader,
				readChannel: readChannel,
				chanMap : make(map[string] chan string),
			}
	go readConnection(c)
	return c, nil
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

func closeConnection(c *Client) {
	// TODO: Any clean up that needs to be done when
	// reader stops reading
}

func sendSubMessage(c *Client, channel string, msg string) {
	c.chanMap[channel] <- msg
}

func readConnection(c *Client) {
	defer closeConnection(c)

	// TODO: infinite loop till
	// reader encounters error
	//for readErr == nil {
	// while reader is valid...
	for {
		val, readErr := readType(c.reader)
		if isSubMessage(val, c) {
			// This is pub sub response...
			// Assume no error....
			// Dont know what to do with the message error..
			_, channel, _, msg, _ := parsePubSubResp(val)
			go sendSubMessage(c, channel, msg)
		} else {
			c.readChannel <- &Response{val: val, err: readErr}
		}
	}
}

func isSubMessage(resp interface{}, c *Client) bool {
	if c.subActive == false {
		return false
	}

	l, ok := resp.(*list.List)
	if ok == false {
		return false
	}

	if l.Len() != 3 {
		return false
	}

	command, ok := l.Front().Value.(string)
	if ok == false {
		return false
	}

	if strings.ToUpper(command) == "MESSAGE" {
		return true
	}
	return false
}

func (c *Client) sendRequest(cmd string, args ...interface{}) (interface{}, error) {
	request, err := createRequest(cmd, args...)
	if err != nil {
		return nil, err
	}
	fmt.Fprintf(c.conn, request)
	s := <-c.readChannel
	return s.val, s.err
}

func (c *Client) sendRequestN(consolidatedRequest string, n int) ([]interface{}, error) {
	fmt.Fprintf(c.conn, consolidatedRequest)
	resp := []interface{}{}
	for i := 0; i < n; i++ {
		s := <-c.readChannel
		if s.err != nil {
			return resp, s.err
		} else {
			resp = append(resp, s.val)
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
