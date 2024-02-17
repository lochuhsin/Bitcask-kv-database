package main

import (
	"net"
)

func main() {
	lport := ":9999"
	laddr, _ := net.ResolveTCPAddr("tcp", lport)
	// saddr, _ := net.ResolveTCPAddr("tcp", "192.128.1.1:8080")
	// conn, err := net.DialTCP("tcp", laddr, saddr)
	// conn.Write("abcde")
	// conn.Close()
	sport := ":8888"
	saddr, _ := net.ResolveTCPAddr("tcp", sport)

	conn, _ := net.DialTCP("tcp", laddr, saddr)
	conn.Write([]byte("I'm here"))
	conn.Close()

	listener, _ := net.ListenTCP("tcp", laddr) // fixed

	for {
		connFD, _ := listener.Accept()

	}

	// data + client(ip:randport) + host(ip:port) -> server
	// client issues a request to the server
	/**
	 * client(ip:clientPort) + host(ip:port) -> router
	 * client(ip:clientPort) : router(ip:routerPort) //convert tran
	 * router(ip:routerPort) + host(ip:port) -> server
	 */

	/**
	 * clientPort is randomly assigned by the Operating System
	 * the client port should be fixed
	 */

	// system -> fixed port (8000) socket

	// conn := net.Dial()
	// conn := net.Listen(":8000")

	// 1.
	// 2.
	conn, _ := net.DialTcp()

	/**
	 * WebRTC, Ip Spoofing,
	 */
	- Client 1:
	conn, _ := net.DialUdp("DISCOVERY:DPORT")
	conn => socketA, localIP (234.23.42.12) + randomPort (356777) -> DISCOVERY:DPORT
    router 1 --> ... -> router n -> Discovery
	conn.Read()
	conn.Write()

	- Discovery:

	go func (newConns <-chan net.Conn) {
		for conn := range newConns {
			for al := range addrList {
				conn.Write(conn.RemoteAddr())
			}
		}
	}
	// Alternative: HTTP endpoint,
		for al := range addrList {
			conn.Write(conn.RemoteAddr())
		}

	list, _ := net.UDPListen(":DPORT")
	addrList := []net.Conn
	for {
		conn, _ := list.Accept() // multiple client
		conn.LocalAddr()
		conn.RemoteAddr()
		addrList = append(addrList, conn)

	}

	// local network:
	// 123.123.123.123

	Accept =>
	conn => socketB, routerNI (1.2.3.4) + routerNPort (86868) -> DISCOVERY:DPORT
	conn.Read()
	conn.Write()

	- Client 2:
 
	type AddresssTuple struct {
		ip string,
		port inr,
	}
    peers := map[string] AddressTuple{}
	conn, _ := net.DialUdp("DISCOVERY:PORT")
	
	// web framework
	go net.httpServer(":80") // re implement http server manage connection

	go func () {
		for {
			addrList := conn.Read()  // id, address
			peers[id] = address
		}
	}()

	func WriteToEveryone() {
		for peer := range peers {
			conn := net.DialUDP(address)
		}
	}

	/**
	 * Start normal server -> Client1, Client2 
	 */


    conn => socketC, routerNI (1.2.3.4) + routerNPort (86868) -> Client2IP:Random2port
	conn.Read()
	conn.Write()

	/**
	 * Discovery 
	 */

}

/**
 * -> chan  -> [1, 2, 3, 4, 5, 6] ->  // should be ordered
 * -> workpool listens to this chan 
 * 
 * -> 3 workers [1, 2, 3]   // worker running on 4, 5, 6
 * -> [1, 2], worker running 4, 3, 6 , task 5 is stored (finished)
 * 
 * memory storage [1, 2, 3] -> worker pool -> segment
 * 
 * user request a read operation -> 1 -> 2 -> 3 -> 4 -> 6 (memory), 5 (segment)
 * 
 * [1, 2, 3, 4, 5, 6] -> 3 worker 
 * 3 worker [4, 5, 6] wait these 3 workers were finished
 * 
 * then remove 4, 5, 6 simultaneously (under same lock) from the memory pool
 * 
 * [1, 2, 3] (memory), [4, 5, 6] segment
 * 
 * flaw -> high cause dead lock 
 * 1. chan is blocked
 * def setEntry(entry){
 *     mu.Lock()
 *     defer mu.Unlock()
 *     ....
 *     if memoryCount >= threshold {
 *         chan <- memoryBlockId  // blocked
 *     }
 * }
 *  
 * 2. Workpool (scheduler)
 * 	deleteing -> memory block [4, 5, 6]
 * 	def bulkDeleteMemoryBlock(ids []int) {
 * 		mu.Lock()
 *      defer mu.Unlock()
 *      [4, 5, 6]
 * }
 * 
 */
