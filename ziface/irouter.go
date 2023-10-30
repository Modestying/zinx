package ziface

import "net"

type IRoute func(conn *net.Conn, data []byte)
