package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

const (
	telnetNetwork = "tcp"
	telnetPort    = 23
	maxDataBytes  = 135
)

// Client is telnet client for AVR.
type Client struct {
	conn net.Conn
	r    *bufio.Reader
}

// Dial dials up the specified host and return client.
func Dial(host string, timeout time.Duration) (*Client, error) {
	addr := fmt.Sprintf("%v:%v", host, telnetPort)
	c, err := net.DialTimeout(telnetNetwork, addr, timeout)
	if err != nil {
		return nil, err
	}
	return &Client{
		conn: c,
		r:    bufio.NewReaderSize(c, maxDataBytes),
	}, nil
}

// Write writes command and parameter.
func (c *Client) Write(cmd, param string) {
	data := fmt.Sprintf("%v%v\r", cmd, param)
	c.conn.Write([]byte(data))
}

// Read reads data from read buffer.
func (c *Client) Read() ([]byte, error) {
	buf := make([]byte, maxDataBytes)
	for n := 0; n < len(buf); n++ {
		if n > 0 && c.r.Buffered() == 0 {
			return buf[:n], nil
		}
		b, err := c.r.ReadByte()
		if err != nil {
			return nil, err
		}
		buf[n] = b
	}
	return nil, fmt.Errorf("response data exceeded limit of the data length")
}

// Close closes connection.
func (c *Client) Close() error {
	return c.conn.Close()
}
