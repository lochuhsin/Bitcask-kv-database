package main

import (
	"encoding/json"
	"fmt"
	"net"
	"rebitcask/api/chore"
	"rebitcask/api/core"
	_ "rebitcask/docs"
	"rebitcask/internal/setting"
	"rebitcask/server/rebitcaskpb"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
)

func clusterSetup() {
	msgCh := make(chan []byte, 1)
	go udpServer(msgCh)

	buff := <-msgCh
	logrus.Info(string(buff))
	peerList := setting.PeerList{}
	err := json.Unmarshal(buff, &peerList)
	if err != nil {
		logrus.Error("Unable to unmarshal data")
		panic(err)
	}
	option := setting.SetPeerList(peerList)
	option(&setting.Config)
}

func httpServerSetup(port string) {
	r := gin.Default()
	core.Routes(r)
	chore.Routes(r)
	// starts swagger at localhost:port/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(port) // listen and serve on 0.0.0.0:8080
	/**
	 * Change this to run listner to reuse port on both grpc and http
	 */
	// r.RunListener()
}

func grpcServerSetup(port string) {
	rbServer := grpcServer{}
	lst, err := net.Listen("tcp", port)
	if err != nil {
		panic("unable to listen port " + port)
	}
	s := grpc.NewServer()
	rebitcaskpb.RegisterRebitcaskServiceServer(s, &rbServer)
	fmt.Println("Starting grpc server on " + port)
	s.Serve(lst)
}
