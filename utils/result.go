package utils



type Result struct {
	Code int
	Error string
	Result interface{}

}


func Resp_result(code int,inter interface{},err string) *Result {
	r := &Result{
		Code:code,
		Error:err,
		Result:inter,
	}
	return r
}
