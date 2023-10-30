package znet

import (
	"encoding/hex"
	"fmt"
	"net"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	addr := &net.TCPAddr{
		IP:   net.ParseIP("0.0.0.0"),
		Port: 8089,
	}
	conn, err := net.DialTCP("tcp4", nil, addr)
	if err != nil {
		panic(err)
	}
	data := "I LOVE YOU"
	//49204c4f564520594f55
	hexString := hex.EncodeToString([]byte(data))
	//[73 32 76 79 86 69 32 89 79 85]
	
	hexBytes, _ := hex.DecodeString(hexString)
	if err != nil {
		panic(err)
	}
	for {
		conn.Write(hexBytes)
		time.Sleep(time.Second * 1)
		data := make([]byte, 512)
		cnt, err := conn.Read(data)
		if err != nil {
			fmt.Printf("Read data from server failed:%s\n", err.Error())
			break
		} else {
			fmt.Printf("Read data from server:%s\n", data[:cnt])
		}
	}

}
