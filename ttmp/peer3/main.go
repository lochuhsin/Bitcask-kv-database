package main

import (
	"fmt"
	"net"
)

func main() {
	// tcpListener, _ := net.Listen("tcp", ":8080")
	// mux := cmux.New(tcpListener)
	// // grpcL := mux.Match(cmux.HTTP1HeaderField("content-type", "application/grpc"))
	// httpL := mux.Match(cmux.HTTP1Fast())
	// httpS := gin.Default()
	// httpS.RunListener(httpL)

	// file, _ := os.Open("main.go")
	// fd := file.Fd()
	// os.Stderr.Read()
	// fmt.Println(fd)

	// _, err := http.Get("http://hello-world:8000")
	_, err := net.ResolveUDPAddr("udp", ":80")
	fmt.Println(err)
}
