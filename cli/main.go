package main

import (
	"fmt"
	"net"
	"os"
	"rebitcask/internal"
	"strings"
	"time"
)

// TODO: Convert using envFile
const (
	ServerHost = "127.0.0.1"
	ServerPort = 6666
	ServerType = "tcp"
)

var ConnQueue ConnectionQueue

func init() {

	// create connection queue as buffer
	ConnQueue = CreateQueue()
	// start handler
	go connectionHandler()
}

/*
this server uses tcp/ip connection
format:
Get:<key>
Set:<key>,<val>
Delete:<key>
*/

func main() {
	fmt.Println("Server starts to run ...")

	server, err := net.Listen(ServerType, fmt.Sprintf("%v:%v", ServerHost, ServerPort))
	if err != nil {
		fmt.Println("Starting server error, shutting down", err.Error())
		os.Exit(1)
	}

	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			continue
		}
		// TODO: convert to goroutine
		go insertQueueHandler(conn)
	}
}

func insertQueueHandler(conn net.Conn) {
	ConnQueue.Enqueue(conn)
}

func connectionHandler() {
	for {
		if qSize := ConnQueue.GetSize(); qSize > 0 {
			conn, status := ConnQueue.Dequeue()
			if status != true {
				// figure out handling this
			}

			buffer := make([]byte, 1024)
			_, _ = conn.Read(buffer)

			operations := strings.Split(string(buffer), ":")
			if len(operations) != 2 {
				conn.Write([]byte("400"))
			} else {
				method, values := operations[0], operations[1]
				fmt.Println(method, values)
				if method == "GET" {
					outputs, status := internal.Get(values)
					fmt.Println("Get status: ", status, method, values)
					if !status {
						conn.Write([]byte("400"))
					} else {
						conn.Write([]byte(outputs))
					}
				}
				if method == "DELETE" {
					err := internal.Delete(values)
					if err != nil {
						conn.Write([]byte("400"))
					} else {
						conn.Write([]byte("200"))
					}
				}
				if method == "SET" {
					val := strings.Split(values, ",")
					if len(val) != 2 {
						conn.Write([]byte("400"))
						continue
					}
					err := internal.Set(val[0], val[1])
					if err != nil {
						conn.Write([]byte("400"))
					} else {
						conn.Write([]byte("200"))
					}
				}
			}
			conn.Close()
		}
		time.Sleep(1 * time.Second)
	}
}

// this is none stopping method of keep reading file
//scanner := bufio.NewScanner(conn)
//for scanner.Scan() {
//	ln := scanner.Text()
//	fmt.Println(ln)
//}
//conn.Close()
