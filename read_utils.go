package go4redis

import (
  "bufio"
  "container/list"
	"errors"
	"strconv"
	"strings"
)

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
		return EMPTY_STRING, err
	}
  if count == -1 {
    return countAsStr, nil
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
