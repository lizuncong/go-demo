package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"go-demo/models"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func listUsers(w http.ResponseWriter) {
	fmt.Println("获取所有用户列表：", len(models.Users))
	resp := Response{
		Code:    200,
		Message: fmt.Sprintf("成功，总共有%d个用户", len(models.Users)),
		Data:    models.Users,
	}
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	json.NewEncoder(w).Encode(resp)
}

func getUserById(w http.ResponseWriter, r *http.Request) {
}

func Users(w http.ResponseWriter, r *http.Request) {
	fmt.Println("user handler request method:", r.Method, r.URL.Path)
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	idStr = strings.Trim(idStr, "/")
	id, err := strconv.Atoi(idStr)
	switch r.Method {
	case http.MethodGet:
		if err == nil {
			fmt.Println("根据id查询用户信息", id)
		} else {
			listUsers(w, r)
		}
	case http.MethodPost:
		fmt.Println("POST /users，新增用户")
	case http.MethodPut:
		fmt.Println("PUT /users/1，整体更新 id=1 的用户")
	case http.MethodDelete:
		fmt.Println("Delete /users/1，删除 id=1 的用户")
	default:
		http.Error(w, "不支持该请求方法", http.StatusMethodNotAllowed)
	}
}
