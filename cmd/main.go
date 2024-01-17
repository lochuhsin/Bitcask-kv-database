package main

import (
	"fmt"
	"net"
	"rebitcask"
	"rebitcask/api/chore"
	"rebitcask/api/core"
	"rebitcask/server/rebitcaskpb"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"google.golang.org/grpc"
)

func runGRPC() {
	port := ":9090"
	rbServer := Server{}
	lst, err := net.Listen("tcp", port)
	if err != nil {
		panic("unable to listen port " + port)
	}
	s := grpc.NewServer()
	rebitcaskpb.RegisterRebitcaskServiceServer(s, &rbServer)
	fmt.Println("Starting grpc server on " + port)
	s.Serve(lst)
}

func main() {
	rebitcask.Init()
	r := gin.Default()
	core.Routes(r)
	chore.Routes(r)
	port := ":8000"
	ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.URL("http://localhost:8000/swagger/doc.json"))

	// use channel to handle goroutine shut down
	go runGRPC()

	r.Run(port) // listen and serve on 0.0.0.0:8080

}
