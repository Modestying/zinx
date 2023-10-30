package ziface

// 服务接口
type IServer interface {
	// 服务启动
	Start()
	// 开启业务
	Serve()
	// 服务停止
	Stop()
	// 注册路由
	RegisterRoutes()
}
