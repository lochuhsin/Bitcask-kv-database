package main

import (
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
	"sync"
	"time"
)

//func sayhelloName(w http.ResponseWriter, r *http.Request) {
//	r.ParseForm()
//	fmt.Println("form", r.Form)
//	fmt.Println("method", r.Method)
//	fmt.Println("path", r.URL.Path)
//	fmt.Println("scheme", r.URL.Scheme)
//	for k, v := range r.Form {
//		fmt.Println("key:", k)
//		fmt.Println("val:", strings.Join(v, ""))
//	}
//	fmt.Fprintf(w, "Hello astaxie!") //這個寫入到 w 的是輸出到客戶端的
//}

type conn struct {
	resp http.ResponseWriter
	req  *http.Request
	id   string
}

type connQueue struct {
	mu         sync.Mutex
	queryQueue []conn
}

type queryHandler struct{}

var Queue = new(connQueue)

func helloServer(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	fmt.Println(r.URL)
	fmt.Println(r.Method)
	fmt.Println(r.URL.Scheme)
	fmt.Println(r.URL.RawQuery)
	fmt.Println("form", r.Form)
	fmt.Fprintf(w, "Welcome to rebitcask") // send message back
}

// / direction of queue   left is output, right is input (opposite from python)
func EventLoop() {
	for {
		if len(Queue.queryQueue) > 0 {
			//_ = Queue.queryQueue[0]
			//Queue.queryQueue = Queue.queryQueue[1:]
			// TODO: Write GET, POST, PATCH, DELETE ..... etc over here
		}

		time.Sleep(3 * time.Second)
		fmt.Println("I'm still running, the length of Queue", len(Queue.queryQueue))
	}
}

func (query *queryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	connection := conn{
		resp: w,
		req:  r,
		id:   uuid.New().String(),
	}
	Queue.mu.Lock()
	Queue.queryQueue = append(Queue.queryQueue, connection)
	Queue.mu.Unlock()

	fmt.Println("Currently have ", len(Queue.queryQueue), "connection waiting")
}

func main() {
	fmt.Println("Start server")
	go EventLoop()

	http.HandleFunc("/", helloServer)
	http.Handle("/query", new(queryHandler))
	err := http.ListenAndServe(":6666", nil) //設定監聽的埠
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
