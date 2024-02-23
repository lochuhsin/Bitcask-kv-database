package main

import (
	"rebitcask"
)

func main() {
	flags := ParseFlags()
	rebitcask.Setup(flags.envPaths...)
	serverSetup()
}
