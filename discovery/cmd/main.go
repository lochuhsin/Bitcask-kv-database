package main

import "rebitcask/discovery/server"

func main() {
	/**
	 * Setting up server as
	 * http://localhost:8765/_rebitcask/........
	 */
	configSetup()
	server.RunUdp()
}
