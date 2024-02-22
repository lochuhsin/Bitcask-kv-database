package server

import (
	"encoding/json"
	"fmt"
	"net"
	"rebitcask/discovery/setting"

	"github.com/sirupsen/logrus"
)

type ConnectionTuple struct {
	Ip   string `json:"Ip"`
	Port string `json:"Port`
}

type PeerList struct {
	Peers []ConnectionTuple `json:"Peers"`
}

var RemoteAddrHolder = make(map[string]ConnectionTuple)

func broadcastConnections(conn *net.UDPConn) {
	for addrI := range RemoteAddrHolder {

		connList := make([]ConnectionTuple, len(RemoteAddrHolder)-1)
		j := 0
		for addrJ, connTuple := range RemoteAddrHolder {
			// Avoid sending self addresses
			if addrI == addrJ {
				continue
			}
			connList[j] = connTuple
			j++
		}
		peers := PeerList{Peers: connList}
		packet, err := json.Marshal(peers)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		targetAddr, err := net.ResolveUDPAddr("udp", addrI)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		n, err := conn.WriteToUDP(packet, targetAddr)
		if n != len(packet) || err != nil {
			fmt.Println(err.Error(), "packet size mismatch")
			continue
		}
	}
}

func RunUdp() error {
	logrus.Info("resolving local address and port")
	laddr, err := net.ResolveUDPAddr(
		"udp", setting.Config.UDP_SERVER_PORT,
	)
	if err != nil {
		return UdpError{
			field: "Error while resolving udp address",
			msg:   err.Error(),
		}
	}
	logrus.Info("Starting up UDP server...")
	conn, err := net.ListenUDP("udp", laddr)
	if err != nil {
		return UdpError{
			field: "Error while listening on udp address",
			msg:   err.Error(),
		}
	}
	for {
		buf := make([]byte, 65535)
		_, remoteAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		raddr := remoteAddr.String()
		host, port, err := net.SplitHostPort(raddr)
		if err != nil {
			logrus.Error(PacketDataError{
				field: "failed to split host ip",
				msg:   err.Error(),
			}.Error())
			continue
		}
		logrus.Info(fmt.Sprintf("Recieved host: %v, port: %v", host, port))
		RemoteAddrHolder[raddr] = ConnectionTuple{
			Ip:   host,
			Port: port,
		}
		if len(RemoteAddrHolder) >= setting.Config.CLUSTER_MEMBER_COUNT {
			logrus.Info("Cluster member count reached, broadcasting peers")
			go broadcastConnections(conn)

			// break the for loop to avoid sending information
			// handle this properly
		}
	}
}
