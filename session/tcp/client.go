package tcp

import (
	"log"
	"net"
	"time"
)

// Client tcp client
type Client struct {
	session   *Session
	autoRetry bool
}

// NewClient first new Client object
func NewClient(h IHandler, autoRetry bool) *Client {
	return &Client{
		session: &Session{
			id:      0,
			handler: h,
			packLen: 0,
		},
		autoRetry: autoRetry,
	}
}

//Start second start client object to connection server
func (c *Client) Start(addr string) {
	// try connetion
	for {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			log.Printf("client connection err %v", err.Error())
			if c.autoRetry {
				<-time.After(time.Second * 3)
				continue
			} else {
				return
			}
		}
		c.session.conn = conn
		break
	}

	c.session.handler.OnOpen(c.session)

	go c.session.scan()

	log.Printf("connect %s established. ", addr)
}

//Stop session stop
func (c *Client) Stop() {
	c.session.conn.Close()
}
