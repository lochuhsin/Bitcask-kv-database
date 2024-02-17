package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type ConnectionTuple struct {
	Ip   string `json:"Ip"`
	Port string `json:"Port"`
}

type ConnectionTupleResponse struct {
	Addrs []ConnectionTuple `json:"Addrs"`
}

var ConnectionHolder = make(map[string]ConnectionTuple)

func broadCastAddrLists(conn *net.UDPConn) {
	fmt.Println("Broadcastinnnng")
	for hostPortI := range ConnectionHolder {
		tupleArr := make([]ConnectionTuple, len(ConnectionHolder)-1)
		i := 0
		for hostPortJ, connTupleJ := range ConnectionHolder {
			// Skip sending address to the same client itself
			if hostPortI == hostPortJ {
				continue
			}
			tupleArr[i] = connTupleJ
		}

		fmt.Println("arr list finished")
		response := ConnectionTupleResponse{tupleArr}
		bytes, err := json.Marshal(response)
		if err != nil {
			panic("Invalid connection tuple struct")
		}

		udpAddr, _ := net.ResolveUDPAddr("udp", hostPortI)
		conn.WriteToUDP(bytes, udpAddr)
	}
}

func main() {
	serverPort := ":80"
	addr, err := net.ResolveUDPAddr("udp", serverPort)
	if err != nil {
		panic(err)
	}

	expectClusterMemberCount := 2
	broadCastChanges := true
	conn, err := net.ListenUDP("udp", addr)
	for {
		if err != nil {
			panic(err)
		}
		if conn == nil {
			continue
		}
		buff := make([]byte, 4096) // Reasonable size that is not going to be chopped off
		// n, whoAddr, err := conn.ReadFrom(buff)
		n, whoAddr, err := conn.ReadFromUDP(buff)
		if err != nil {
			panic(err)
		}
		addrString := whoAddr.String()
		fmt.Printf("receive packets from : %v\n", addrString)
		fmt.Printf("receive packets size : %v\n", n)
		Ip, port, _ := net.SplitHostPort(addrString)

		/**
		 * Implement server name parser
		 */
		ConnectionHolder[addrString] = ConnectionTuple{
			Ip:   Ip,
			Port: port,
		}
		fmt.Println(len(ConnectionHolder))
		if len(ConnectionHolder) >= expectClusterMemberCount {
			if broadCastChanges {
				fmt.Println("broadcasting")
				go broadCastAddrLists(conn)
				// broadCastChanges = false
			}
		}
	}
}
