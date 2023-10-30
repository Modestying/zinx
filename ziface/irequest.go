package ziface

type IRequest interface {
	GetConnection() IConnection
	GetMessage() []byte
}
