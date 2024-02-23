package main

import (
	"net"
	"rebitcask/internal/setting"

	"github.com/sirupsen/logrus"
)

func udpServer(ch chan<- []byte) {
	logrus.Info("Start setting UDP connection")
	clusterSetupHost := setting.Config.CLUSTER_SETUP_HOST
	localPort := setting.Config.PORT // This is called spoofing

	raddr, err := net.ResolveUDPAddr("udp", clusterSetupHost)
	if err != nil {
		logrus.Error(err)
	}
	laddr, err := net.ResolveUDPAddr("udp", localPort)
	if err != nil {
		logrus.Error(err)
	}

	/**
	 * Add retry mechanism, make sure the cluster runs
	 */

	logrus.Info("Start established UDP connection")
	conn, err := net.DialUDP("udp", laddr, raddr)
	if err != nil {
		logrus.Error(err)
	}
	/**
	 * Handle this goroutine properly
	 */
	go func() {
		for {
			buff := make([]byte, 65532)
			n, err := conn.Read(buff)
			if err != nil {
				logrus.Error(err)
			}
			ch <- buff[:n]
		}
	}()

	sid := []byte(setting.Config.SERVER_ID)
	n, err := conn.Write(sid)
	if n != len(sid) || err != nil {
		panic("Unexpected write error to udp packet")
	}
}
