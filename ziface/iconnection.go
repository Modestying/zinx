package ziface

import "net"

type IConnection interface {
	// 连接信息
	RemoteAddr() net.Addr
	RemoteAddrString() string
	LocalAddr() net.Addr
	LocalAddrString() string

	GetConnectionID() uint64
	GetConnection() net.Conn
	GetTcpConnection() net.Conn

	Send(data []byte) error

	// 连接是否用
	IsAlive() bool
	Start()
	Stop()
}
