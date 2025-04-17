package main

import "kitex-server/users"

func main() {
	go users.Start()

	select {}
}
