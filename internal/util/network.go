package util

import (
	"log"
	"net"
)

// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	// https://stackoverflow.com/questions/23558425/how-do-i-get-the-local-ip-address-in-go
	conn, err := net.Dial("udp", "8.8.8.8:80") // Discovery ip address
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr) // on discovery server

	return localAddr.IP
}

/*
curl google:80
google.com => IPGoogle (DNS resolution)
tcp conn IPGoogle:80
locally: opens a socketA: IPLocal+RandomPort1 <-> IPRouter+RandomOrFixedPort2
routerA: open a socketB: IPRouter+RandomOrFixedPort2 <-> IPLocal+RandomPort1
Google: open a socketC: IPGoogle+80 <-> IPRouter+RandomPort2

IPv6: every computer has a unique IPv6 address

IPv4: subnetworks
Local (private IP) <-> Router (external public IP) <-> (external public IP)Router google <-> Google
Local (external public IP) <-> Router Google

Local1 -> RouterA -> Discovery (external IP)
Local2 -> RouterB -> Discovery (RouterA+Port RouterB+Port)
Local1 -> RouterB
Local2 -> RouterA
*/
