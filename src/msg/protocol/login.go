package protocol

type C_Register struct {
	UserName,Pwd,MacId string
}
type C_LoginByPwd struct {
	UserName,Pwd string
}
type C_LoginByGuest struct {

}