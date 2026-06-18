package utils

import (
	"fmt"
	"net/url"
	"strings"
)

// 传入/users/:id/list,  /users/2/list?pageNo=1&pageSize=2
// 解析出{ id: 2, pageNo: 1, pageSize: 2}
func ParseUrlParams(pattern string, urlStr string) (map[string]string, bool) {
	urlParts := strings.Split(urlStr, "?")
	pathStr := urlParts[0]
	var queryStr string
	if len(urlParts) > 1 {
		queryStr = urlParts[1]
	}
	patternParts := strings.Split(strings.Trim(pattern, "/"), "/")
	pathParts := strings.Split(strings.Trim(pathStr, "/"), "/")
	params := map[string]string{}

	// 解析/users/:id/list里面的参数
	if len(pathParts) == len(patternParts) {
		for idx, part := range pathParts {
			patternPart := patternParts[idx]
			if part != patternPart {
				paramKey := strings.Trim(patternPart, ":")
				params[paramKey] = part
			}
		}
	}
	// 解析pageNo=1&pageSize=2&timestamp的参数
	queryParts := strings.Split(queryStr, "&")
	for _, part := range queryParts {
		p := strings.Split(part, "=")
		key := p[0]
		value := key
		if len(p) > 1 {
			value, _ = url.QueryUnescape(p[1])
		}
		params[key] = value
	}
	fmt.Println("params: ", params)
	return params, true
}
