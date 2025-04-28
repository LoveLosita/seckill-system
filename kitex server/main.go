package main

import (
	"kitex-server/items"
	"kitex-server/seckill"
	"kitex-server/users"
)

func main() {
	go users.Start()
	go items.Start()
	go seckill.Start()
	select {}
}
