package utils



type result struct {
	Code int
	Error string
	Result interface{}

}


func Resp_result(code int,inter interface{},err string) *result {
	r := &result{
		Code:code,
		Error:err,
		Result:inter,
	}
	return r
}
