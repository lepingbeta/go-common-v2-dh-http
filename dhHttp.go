package dhHttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	"go.mongodb.org/mongo-driver/bson"
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
	var bearerToken string // 用于临时存储提取的Token

	// 提取Bearer Token的逻辑
	extractBearerToken := func() {
		// 尝试从map或bson.M类型中提取bearerToken
		if params, ok := data.(map[string]interface{}); ok {
			if tokenVal, exists := params["bearerToken"]; exists {
				if tokenStr, ok := tokenVal.(string); ok {
					bearerToken = tokenStr
					delete(params, "bearerToken") // 删除键避免出现在请求中
				}
			}
		} else if params, ok := data.(bson.M); ok {
			if tokenVal, exists := params["bearerToken"]; exists {
				if tokenStr, ok := tokenVal.(string); ok {
					bearerToken = tokenStr
					delete(params, "bearerToken") // 删除键避免出现在请求中
				}
			}
		}
	}

	// GET请求处理
	if reqType == "GET" {
		// 先提取可能的Token
		extractBearerToken()

		// 处理请求参数
		params, ok1 := data.(map[string]interface{})
		params2, ok2 := data.(bson.M)
		var queryParams url.Values

		if ok1 || ok2 {
			queryParams = url.Values{}
			var source map[string]interface{}
			if ok1 {
				source = params
			} else {
				source = params2
			}
			for key, value := range source {
				if strValue, ok := value.(string); ok {
					queryParams.Add(key, strValue)
				} else {
					dhlog.Error("Skipping non-string value for key: %s", key)
				}
			}
			// 构造带参数的URL
			if strings.Contains(urlStr, "?") {
				urlStr += "&" + queryParams.Encode()
			} else {
				urlStr += "?" + queryParams.Encode()
			}
		}

		// 创建请求
		req, err = http.NewRequest("GET", urlStr, nil)

		// 非GET请求处理
	} else {
		// 先提取可能的Token
		extractBearerToken()

		// 序列化JSON（已删除bearerToken的data）
		jsonData, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}

		// 创建请求
		jsonBuffer := bytes.NewBuffer(jsonData)
		req, err = http.NewRequest(reqType, urlStr, jsonBuffer)
	}

	if err != nil {
		return nil, err
	}

	// 设置Authorization头（如果存在Token）
	if bearerToken != "" {
		req.Header.Set("Authorization", "Bearer "+bearerToken)
	}

	// 设置Content-Type
	if reqType != "GET" {
		req.Header.Set("Content-Type", "application/json")
	}

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 检查状态码
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("请求失败，状态码: %d", resp.StatusCode)
	}

	// 解析响应
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
func GetJSON2Map(urlStr string, params interface{}) (map[string]interface{}, error) {
	return ReqJSON2Map("GET", urlStr, params)
}
