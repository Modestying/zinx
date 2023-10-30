package znet

import "testing"

func TestNewServer(t *testing.T) {
	server := NewServer("zinx", 8089)
	server.Serve()
}
