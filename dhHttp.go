package dhHttp

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
)

// PostJSON 发送一个包含JSON数据的POST请求
func PostJSON(url string, data interface{}) (*http.Response, error) {
	// 将数据编码为JSON格式
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// 创建一个包含JSON数据的缓冲区
	jsonBuffer := bytes.NewBuffer(jsonData)

	// 创建一个请求
	req, err := http.NewRequest("POST", url, jsonBuffer)
	if err != nil {
		return nil, err
	}

	// 设置请求头信息
	req.Header.Set("Content-Type", "application/json")

	// 发送请求并获取响应
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// 检查响应状态码
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		resp.Body.Close()
		return resp, http.ErrBodyNotAllowed
	}

	return resp, nil
}

// ResponseToMap 尝试将HTTP响应体解码为map[string]interface{}。
func ResponseToMap(resp *http.Response) (map[string]interface{}, error) {
	// 读取响应体
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 创建一个map来存储解码后的数据
	var m map[string]interface{}

	// 解码响应体到map中
	if err := json.Unmarshal(respBody, &m); err != nil {
		return nil, err
	}

	return m, nil
}

// ReqJSON2Map 发送一个包含JSON数据的HTTP请求并将响应体解码为map[string]interface{}。
func ReqJSON2Map(reqType string, urlStr string, data interface{}) (map[string]interface{}, error) {
	var req *http.Request
	var err error

	// 根据请求类型决定是否需要将data编码为JSON格式
	if reqType == "GET" {
		// 将GET请求参数编码为查询字符串并添加到URL
		if params, ok := data.(map[string]interface{}); ok {
			queryParams := url.Values{}
			for key, value := range params {
				queryParams.Add(key, value.(string))
			}
			if strings.Contains(urlStr, "?") {
				urlStr += "&" + queryParams.Encode()
			} else {
				urlStr += "?" + queryParams.Encode()
			}
		}
		req, err = http.NewRequest("GET", urlStr, nil)
	} else {
		// 对于非GET请求，将data编码为JSON格式
		jsonData, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}

		// 创建一个包含JSON数据的缓冲区
		jsonBuffer := bytes.NewBuffer(jsonData)
		req, err = http.NewRequest(reqType, urlStr, jsonBuffer)
	}

	if err != nil {
		return nil, err
	}

	// 设置请求头信息
	if reqType != "GET" {
		req.Header.Set("Content-Type", "application/json")
	}

	// 发送请求并获取响应
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	dhlog.Info("resp.StatusCode: %v", resp.StatusCode)
	// 检查响应状态码
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		resp.Body.Close()
		return nil, http.ErrBodyNotAllowed
	}

	// 将响应体解码为map[string]interface{}
	return ResponseToMap(resp)
}

// PostJSON2Map 发送一个POST请求并将响应体解码为map[string]interface{}。
func PostJSON2Map(url string, data interface{}) (map[string]interface{}, error) {
	return ReqJSON2Map("POST", url, data)
}

// PutJSON2Map 发送一个PUT请求并将响应体解码为map[string]interface{}。
func PutJSON2Map(url string, data interface{}) (map[string]interface{}, error) {
	return ReqJSON2Map("PUT", url, data)
}

// GetJSON 发送一个GET请求并将响应体解码为map[string]interface{}。
func GetJSON2Map(urlStr string, params map[string]interface{}) (map[string]interface{}, error) {
	return ReqJSON2Map("GET", urlStr, params)
}
