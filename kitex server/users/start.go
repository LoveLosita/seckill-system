package users

import (
	"github.com/cloudwego/kitex/server"
	"kitex-server/init"
	user "kitex-server/users/kitex_gen/user/userservice"
	"log"
	"net"
)

func Start() {
	//1.连接数据库
	err := init.ConnectDB()
	if err != nil {
		log.Fatalf("init.ConnectDB error: %v", err)
	}
	//2.连接redis
	err = init.InitRedis()
	if err != nil {
		log.Fatalf("init.InitRedis error: %v", err)
	}
	//3.启动服务
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8889")
	if err != nil {
		log.Fatalf("net.ResolveTCPAddr error: %v", err)
	}
	svr := user.NewServer(new(UserServiceImpl), server.WithServiceAddr(addr))
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
