package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go-demo/handlers"
)

func main() {
	http.HandleFunc("/users", handlers.Users)
	http.HandleFunc("/users/", handlers.Users)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "hello go",
		})
	})

	addr := ":8080"
	fmt.Printf("服务器启动，监听地址：http://localhost%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
