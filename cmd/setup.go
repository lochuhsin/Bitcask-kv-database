package main

import (
	"encoding/json"
	"net"
	"rebitcask/api/chore"
	"rebitcask/api/core"
	_ "rebitcask/docs"
	"rebitcask/internal/setting"
	"rebitcask/server/rebitcaskpb"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/soheilhy/cmux"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
)

func serverSetup() {
	if setting.Config.MODE == setting.CLUSTER {
		clusterSetup()
		logrus.Info("Cluster setup complete")
	}

	tcpListener, _ := net.Listen("tcp", setting.Config.HTTP_PORT)
	mux := cmux.New(tcpListener)
	grpcL := mux.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	httpL := mux.Match(cmux.HTTP1Fast())
	tcpL := mux.Match(cmux.Any())

	go grpcServerSetup(grpcL)
	go httpServerSetup(httpL)
	AnonymousTCPSetup(tcpL)
}

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

func httpServerSetup(l net.Listener) {
	r := gin.Default()
	core.Routes(r)
	chore.Routes(r)
	// starts swagger at localhost:port/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.RunListener(l) // listen and serve on 0.0.0.0:8080
}

func grpcServerSetup(l net.Listener) {
	rbServer := grpcServer{}
	s := grpc.NewServer()
	rebitcaskpb.RegisterRebitcaskServiceServer(s, &rbServer)
	s.Serve(l)
}

func AnonymousTCPSetup(l net.Listener) {
	for {
		conn, err := l.Accept()
		if err != nil {
			logrus.Error(err)
		}
		logrus.Info(conn.RemoteAddr().String())
	}
}
