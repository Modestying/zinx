package znet

import (
	"fmt"
	"net"

	"github.com/Modestying/zinx/ziface"
)

type ZClient struct {
	remoteServerIP   string
	remoteServerPort int
	conn             *net.TCPConn
	closeConnect     chan bool
}

var _ ziface.IClient = (*ZClient)(nil)

func NewZClient(ip string, port int) ziface.IClient {
	c := &ZClient{
		remoteServerIP:   ip,
		remoteServerPort: port,
		closeConnect:     make(chan bool, 1),
	}
	return c
}

func (c *ZClient) Connect() bool {
	addr := &net.TCPAddr{
		IP:   net.ParseIP(c.remoteServerIP),
		Port: c.remoteServerPort,
	}
	var err error
	c.conn, err = net.DialTCP("tcp4", nil, addr)
	if err != nil {
		fmt.Printf("Dial server failed:%s", err.Error())
		return false
	}
	return true
}

func (c *ZClient) SendMessage(data []byte) {
	if _, err := c.conn.Write(data); err != nil {
		fmt.Printf("Send message to Sever %s failed:%s", c.conn.RemoteAddr(), err.Error())
	}
}
