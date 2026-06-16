package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go-demo/models"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func listUsers() {
	return {
		message: "",
		data: models.Users
	}
}

func Users(w http.ResponseWriter, r *http.Request) {
	fmt.Println("user handler request method:", r.Method)
	var resp Response = {
		 	Code:    200,
		// Message: fmt.Sprintf("数据获取成功，总共有%d个用户", len(users)),
		// Data:    users,
	}
	if r.Method != http.MethodGet {
		http.Error(w, "仅支持Get请求", http.StatusMethodNotAllowed)
		return
	}
	users := models.Users
	// resp := Response{
	// 	Code:    200,
	// 	Message: fmt.Sprintf("数据获取成功，总共有%d个用户", len(users)),
	// 	Data:    users,
	// }
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(resp)
}
