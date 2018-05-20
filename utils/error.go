package utils




//err struct
type errcode struct {
	Error    string  `json:"Error"`
	Result   string  `json:"Result"`
	Code     int     `json:"Code"`   //code
	/*
		200 success
		-1 一般错误
		-2 致命错误，在房间中,强制用户退出房间
		*/
}



func ErrGame(err ,result string ,code int) *errcode {
	e := &errcode{
		Error:err,
		Result:result,
		Code:code,
	}

	return 	e
}
