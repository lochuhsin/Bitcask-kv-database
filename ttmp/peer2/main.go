package main

import (
	"fmt"
	"net"

	"github.com/gin-gonic/gin"
)

func main() {
	peerPort := ":2222"
	serverPort := ":80"
	laddr, _ := net.ResolveUDPAddr("udp", peerPort)
	saddr, _ := net.ResolveUDPAddr("udp", serverPort)
	conn, err := net.DialUDP("udp", laddr, saddr)
	if err != nil {
		panic("unable to dial to discovery server")
	}

	conn.Write([]byte("22222222222"))
	buff := make([]byte, 65535) // Reasonable large size for not being chopped of
	_, _, err = conn.ReadFromUDP(buff)
	if err != nil {
		panic(err)
	}
	fmt.Println("Received ip lists from discovery server")
	fmt.Println(string(buff))
	// starts a server on same port
	// After receiving ip lists,
	r := gin.Default()
	r.Run(peerPort)
}
