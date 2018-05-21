package protocol

type C_Create struct {
	RoomId,ServerId string
}
type C_Enter struct {
     C_Create
}
type C_Quit struct {
	C_Create
}
type C_QuickEnter struct {

}