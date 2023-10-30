package znet

import (
	"fmt"
	"net"
	"sync/atomic"

	"github.com/Modestying/zinx/ziface"
)

const (
	NetProto      string = "tcp4"
	ServerVersion string = "demo"
	MaxConnection uint64 = 4
)

// IServer 实现
type Server struct {
	name            string
	addr            string
	version         string
	IPVersion       string
	Port            int
	stopChannel     chan bool
	maxConnection   uint64
	countConnection uint64 // 连接数量
}

var _ ziface.IServer = (*Server)(nil)

func NewServer(name string, port int) ziface.IServer {
	s := &Server{
		addr:          "0.0.0.0",
		version:       ServerVersion,
		IPVersion:     NetProto,
		name:          name,
		Port:          port,
		maxConnection: MaxConnection,
		stopChannel:   make(chan bool, 1),
	}
	return s
}

func (s *Server) Start() {
	fmt.Printf("ZinxServer Start:\n IP:%s,端口号:%d\n", s.addr, s.Port)
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.addr, s.Port))
	if err != nil {
		fmt.Println("Resolve tcp addr failed:", err.Error())
		panic(err)
	}
	listener, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		fmt.Println("Statr Listen Tcp failed:", err.Error())
		panic(err)
	}

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("TcpListener accept error....")
				continue
			}
			if s.countConnection >= MaxConnection {
				fmt.Printf("Connection is limit....")
				continue
			}
			newConnID := atomic.AddUint64(&s.countConnection, 1)
			serverConn := NewServerConnection(conn, s, newConnID)
			go s.StartConn(serverConn)
		}
	}()
}

func (s *Server) Stop() {
	s.stopChannel <- true
}

func (s *Server) Serve() {
	s.Start()
	<-s.stopChannel
}

func (s *Server) RegisterRoutes() {

}

func (s *Server) StartConn(conn ziface.IConnection) {
	conn.Start()
}
