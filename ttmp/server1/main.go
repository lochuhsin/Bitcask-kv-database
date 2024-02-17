package main

import (
	"net"
)

func main() {
	lport := ":8888"
	laddr, _ := net.ResolveTCPAddr("tcp", lport)
	listener, _ := net.ListenTCP("tcp", laddr)

	// var remoteAddr string
	// var remoteNetwork string

	for {
		connFd, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		bufferSize := 1024
		buffer := make([]byte, bufferSize)
		outputs := make([]byte, bufferSize)

		for {
			n, err := connFd.Read(buffer)
			if err != nil {
				panic(err)
			}

			outputs = append(outputs, buffer...)
			if n != bufferSize {
				break
			}
		}
		remoteAddr = connFd.RemoteAddr().String()
		remoteNetwork = connFd.RemoteAddr().Network()
		connFd.Write(outputs)
		connFd.Close()
		break
	}

	// saddr, _ := net.ResolveTCPAddr(remoteNetwork, remoteAddr)
	// conn, _ = net.DialTCP("tcp", laddr, saddr)
	// conn.
}
