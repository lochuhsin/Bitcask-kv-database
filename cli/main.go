package main

import (
	"fmt"
	"log"
	"net/http"
	"rebitcask/cli/Handler"
)

func helloServer(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	fmt.Println(r.URL)
	fmt.Println(r.Method)
	fmt.Println(r.URL.Scheme)
	fmt.Println(r.URL.RawQuery)
	fmt.Println("form", r.Form)
	fmt.Fprintf(w, "Welcome to rebitcask") // send message back
}

func main() {
	fmt.Println("Start server")
	http.HandleFunc("/", helloServer)
	http.Handle("/query", new(handler.QueryHandler))
	err := http.ListenAndServe(":6666", nil) //設定監聽的埠
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
