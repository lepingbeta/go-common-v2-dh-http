/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-06-02 11:51:27
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-06-02 12:25:24
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
