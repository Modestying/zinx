package ziface

type IClient interface {
	SendMessage(data []byte)
	Connect() bool
	//Start()
}
