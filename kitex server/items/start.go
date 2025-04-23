package items

import (
	"github.com/cloudwego/kitex/server"
	"kitex-server/inits"
	items "kitex-server/items/kitex_gen/items/itemservice"
	"log"
	"net"
)

func Start() {
	//1.连接数据库
	err := inits.ConnectDB()
	if err != nil {
		log.Fatalf("init.ConnectDB error: %v", err)
	}
	//2.连接redis
	err = inits.InitRedis()
	if err != nil {
		log.Fatalf("init.InitRedis error: %v", err)
	}
	//3.启动服务
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8890")
	if err != nil {
		log.Fatalf("net.ResolveTCPAddr error: %v", err)
	}
	svr := items.NewServer(new(ItemServiceImpl), server.WithServiceAddr(addr))
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
