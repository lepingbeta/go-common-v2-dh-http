/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-08 19:45:35
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-06-10 19:47:56
 * @FilePath     : /hecos-v2-api/data/mycode/dahe/go-common/v2/go-common-v2-dh-http/types/response_data.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package types

import "go.mongodb.org/mongo-driver/bson"

type ResponseData struct {
	Status string      `json:"status"`
	MsgKey string      `json:"msg_key"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

type DataList struct {
	Page  int64    `json:"page"`
	Total int64    `json:"total"`
	List  []bson.M `json:"list"`
}
