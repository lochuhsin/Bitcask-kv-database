package main

import "rebitcask/discovery/server"

func main() {
	configSetup()
	server.RunUdp()
}
