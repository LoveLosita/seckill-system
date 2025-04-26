package seckill

import (
	"kitex-server/inits"
	"kitex-server/seckill/kafka"
	seckill "kitex-server/seckill/kitex_gen/seckill/seckillservice"
	"log"
)

func main() {
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
	//3.启动kafka
	go kafka.StartKafkaConsumer()
	//4.启动服务
	svr := seckill.NewServer(new(SecKillServiceImpl))
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
