package types

type ResponseData struct {
	Status string      `json:"status"`
	MsgKey string      `json:"msg_key"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}
