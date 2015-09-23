package go4redis

import (
	"container/list"
	"errors"
	"net"
	"strings"
	"fmt"
	"bytes"
	"bufio"
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

	pipelineMode	bool
	pipelineChan  chan *Response // Used in pipeline mode
	pipelineBuffer bytes.Buffer
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
			if c.pipelineMode == false {
					c.readChannel <- &Response{val: val, err: readErr}
			} else {
				c.pipelineChan <- &Response{val: val, err: readErr}
			}
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
	if c.pipelineMode == true {
		return nil, errors.New("Cannot execute command in pipeline mode. Use AddToPipeline(string) method instead")
	}

	request, err := createRequest(cmd, args...)
	if err != nil {
		return nil, err
	}
	fmt.Fprintf(c.conn, request)
	s := <-c.readChannel
	return s.val, s.err
}

func (c *Client) sendRequestN(consolidatedRequest string, n int) ([]interface{}, error) {
	if c.pipelineMode == true {
		return nil, errors.New("Cannot execute command in pipeline mode. Use AddToPipeline(string) method instead")
	}
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
