package main

import (
	"fmt"
	"log"
	"net/http"

	"go-demo/handlers"
)

func main() {
	http.HandleFunc("/users", handlers.Users)
	http.HandleFunc("/users/", handlers.Users)

	addr := ":8080"
	fmt.Printf("服务器启动，监听地址：http://localhost%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
