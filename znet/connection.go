package znet

import (
	"context"
	"fmt"
	"net"

	"github.com/Modestying/zinx/ziface"
)

/**
* connection is the bridge between callerA and CallerB
* it exists both server and client
* it contains: net.conn,local/remote addr
* it do: sendData
 */
type Connection struct {
	conn       net.Conn
	connID     uint64
	alive      bool
	remoteAddr net.Addr
	localAddr  net.Addr
	ctx        context.Context
	cancel     context.CancelFunc
}

var _ ziface.IConnection = (*Connection)(nil)

func NewServerConnection(conn net.Conn, server ziface.IServer, connID uint64) ziface.IConnection {
	serverConn := &Connection{
		conn:       conn,
		connID:     connID,
		alive:      true,
		remoteAddr: conn.RemoteAddr(),
		localAddr:  conn.LocalAddr(),
	}
	fmt.Printf("A new Client Connected!\n%s", serverConn)
	return serverConn
}

func NewClientConnection(conn net.Conn, client ziface.IClient) ziface.IConnection {
	clientConn := &Connection{
		conn:       conn,
		alive:      true,
		remoteAddr: conn.RemoteAddr(),
		localAddr:  conn.LocalAddr(),
	}
	return clientConn
}
func (c *Connection) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *Connection) RemoteAddrString() string {
	return c.conn.RemoteAddr().String()
}

func (c *Connection) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

func (c *Connection) LocalAddrString() string {
	return c.conn.LocalAddr().String()
}

func (c *Connection) GetConnection() net.Conn {
	return c.conn
}

func (c *Connection) GetTcpConnection() net.Conn {
	return c.conn
}

func (c *Connection) Send(data []byte) error {
	_, err := c.conn.Write(data)
	return err
}

func (c *Connection) IsAlive() bool {
	return c.alive
}

func (c *Connection) Start() {
	c.ctx, c.cancel = context.WithCancel(context.Background())
	go c.StartRead()

	select {
	case <-c.ctx.Done():
		// 接收到结束信号，清理资源
		c.Clean()
		return
	}
}

func (c *Connection) Stop() {
	c.cancel()
}

func (c *Connection) Clean() {

}

func (c *Connection) StartRead() {
	defer c.Stop()
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("StartReadErr:%s\n", err)
		}
	}()
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			msgData := make([]byte, 512)
			cnt, err := c.conn.Read(msgData)
			if err != nil {
				fmt.Printf("Read data from server failed:%s\n", err.Error())
				panic(err)
			} else {
				fmt.Printf("Read data from server:%s\n", msgData[:cnt])
				c.Send(msgData[:cnt])
			}
		}
	}
}

func (c *Connection) GetConnectionID() uint64 {
	return c.connID
}

func (c *Connection) String() string {
	return fmt.Sprintf("ConnectionID:%d\nRemoteAddr:%s\n", c.connID, c.RemoteAddrString())
}
