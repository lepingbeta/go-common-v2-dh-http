package types

type TResponseStatus struct {
	Success string
	Error   string
}

// 初始化全局变量 ResponseStatus
var ResponseStatus TResponseStatus

func init() {
	ResponseStatus = TResponseStatus{
		Success: "success",
		Error:   "error",
	}
}
