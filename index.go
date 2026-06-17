package main

import (
	"fmt"
	"log"
	"net/http"

	"go-demo/handlers"
	"go-demo/utils"
)

func main() {
	utils.ParseUrlParams("/users/:id/list", "/users/2/list?pageNo=1&pageSize=2&timestamp")

	http.HandleFunc("/users", handlers.Users)

	addr := ":8080"
	fmt.Printf("服务器启动，监听地址：http://localhost%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
