package main

import (
	"kitex-server/items"
	"kitex-server/users"
)

func main() {
	go users.Start()
	go items.Start()
	select {}
}
