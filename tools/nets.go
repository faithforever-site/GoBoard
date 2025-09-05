package tools

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// GetAllParams 获取所有请求参数，包括 GET, POST 表单, JSON, XML
func GetAllParams(r *http.Request) map[string][]string {
	params := make(map[string][]string)

	// 1. GET 参数
	for key, vals := range r.URL.Query() {
		params[key] = append(params[key], vals...)
	}

	// 2. POST 表单参数
	if strings.HasPrefix(r.Header.Get("Content-Type"), "application/x-www-form-urlencoded") ||
		strings.HasPrefix(r.Header.Get("Content-Type"), "multipart/form-data") {
		_ = r.ParseForm()
		for key, vals := range r.PostForm {
			params[key] = append(params[key], vals...)
		}
	}

	// 3. 备份 Body，JSON 和 XML 可能都需要读取
	bodyBytes, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // 恢复 Body

	contentType := r.Header.Get("Content-Type")

	// 4. JSON 参数
	if strings.HasPrefix(contentType, "application/json") && len(bodyBytes) > 0 {
		var data map[string]any
		if err := json.Unmarshal(bodyBytes, &data); err == nil {
			for key, val := range data {
				params[key] = append(params[key], fmt.Sprintf("%v", val))
			}
		}
	}

	// 5. XML 参数
	if (strings.HasPrefix(contentType, "application/xml") || strings.HasPrefix(contentType, "text/xml")) && len(bodyBytes) > 0 {
		var data map[string]string
		if err := xml.Unmarshal(bodyBytes, &data); err == nil {
			for key, val := range data {
				params[key] = append(params[key], val)
			}
		}
	}

	return params
}
