/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-06-02 11:51:27
 * @LastEditors: Symphony zhangleping@cezhiqiu.com
 * @LastEditTime: 2025-03-07 14:27:26
 * @FilePath     : /v2/go-common-v2-dh-http/dhHttp_test.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package dhHttp

import (
	"io/ioutil"
	"log"
	"testing"

	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
)

func TestPostJSON(t *testing.T) {
	// 定义要POST的URL
	url := "http://192.168.31.11:19528/oauth2/code2_token"

	// 定义要发送的数据
	data := map[string]interface{}{
		"code": "VnY2PRCTf41-afiFzBbIOg==",
	}

	// 发送POST请求
	resp, err := PostJSON(url, data)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// 打印响应状态码
	log.Println("Response Status:", resp.Status)

	// 读取响应体
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// 打印响应体内容
	log.Println("Response Body:", string(respBody))
}

func TestPostJSON2Map(t *testing.T) {
	// 定义要POST的URL
	url := "http://192.168.31.11:19528/oauth2/code2_token"

	// 定义要发送的数据
	data := map[string]interface{}{
		"code": "WbORs5YJp5EiSbwVpwuSTQ==",
	}

	// 发送POST请求
	resp, err := PostJSON2Map(url, data)
	if err != nil {
		log.Fatal(err)
	}

	dhlog.DebugAny(resp)
	dhlog.DebugAny(resp["status"])
	d := resp["data"].(map[string]any)
	dhlog.DebugAny(d["access_token"].(string))
}

func TestGetJSON2Map(t *testing.T) {
	// // 创建一个模拟的HTTP服务器
	// mockResponse := `{"message": "success", "data": {"id": 123, "name": "test"}}`
	// server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	// 检查是否是GET请求
	// 	if r.Method != http.MethodGet {
	// 		t.Errorf("Expected GET request, got %s", r.Method)
	// 	}
	// 	// 设置响应头
	// 	w.Header().Set("Content-Type", "application/json")
	// 	// 写入模拟的JSON响应
	// 	w.Write([]byte(mockResponse))
	// }))
	// defer server.Close()

	// 调用GetJSON函数
	url := `http://192.168.31.11:32880/ApiDefinition/testGetJson?apiId=yyyzz`
	params := map[string]interface{}{
		"apiId": "yyyzz11",
	}
	result, err := GetJSON2Map(url, params)
	dhlog.DebugAny(result)
	dhlog.DebugAny(err)
	// if err != nil {
	// 	t.Fatalf("GetJSON failed: %v", err)
	// }

	// // 检查响应中的数据
	// if result["message"] != "success" {
	// 	t.Errorf("Expected message to be 'success', got '%v'", result["message"])
	// }

	// data, ok := result["data"].(map[string]interface{})
	// if !ok {
	// 	t.Fatalf("Expected data to be a map, got %T", result["data"])
	// }

	// if data["id"] != float64(123) { // JSON unmarshals numbers as float64
	// 	t.Errorf("Expected id to be 123, got %v", data["id"])
	// }

	// if data["name"] != "test" {
	// 	t.Errorf("Expected name to be 'test', got '%v'", data["name"])
	// }
}

func TestPostJSON2MapWithBearer(t *testing.T) {
	// 定义要POST的URL
	url := "https://ark.cn-beijing.volces.com/api/v3/chat/completions"

	// 定义要发送的数据
	data := map[string]interface{}{
		"model":       "ep-20250228090136-kzqjj",
		"messages":    []interface{}{map[string]string{"role": "user", "content": "你好"}},
		"bearerToken": "c1ae5365-a0c4-4333-aa92-b3b2f7b23f7e",
	}

	// 发送POST请求
	resp, err := PostJSON2Map(url, data)
	if err != nil {
		log.Fatal(err)
	}

	dhlog.DebugAny(resp)
}
