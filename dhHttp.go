/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-06-02 11:51:27
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-06-02 12:18:57
 * @FilePath     : /v2/go-common-v2-dh-http/dhHttp.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package dhHttp

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
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
	respBody, err := ioutil.ReadAll(resp.Body)
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

func PostJSON2Map(url string, data interface{}) (map[string]interface{}, error) {
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
		return nil, http.ErrBodyNotAllowed
	}

	return ResponseToMap(resp)
}
