package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"go-demo/models"
	"go-demo/utils"
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
	fmt.Println("user handler request method:", r.Method, r.URL.String(), r.URL.Path)
	params, _ := utils.ParseUrlParams("/users/:id", r.URL.String())
	var id int
	fmt.Println("id....", id)
	idStr := params["id"] // 因为 map 取不存在的 key 时也会返回空字符串
	if idStr != "" {
		id, _ = strconv.Atoi(idStr)
	}
	switch r.Method {
	case http.MethodGet:
		if id != 0 {
			fmt.Println("根据id查询用户信息", id)
		} else {
			listUsers(w)
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
