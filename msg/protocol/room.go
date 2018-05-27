package protocol

type C_Create struct {
	RoomId,ServerId string
}
type C_Enter struct {
	RoomId,ServerId string
}
type C_Quit struct {
	RoomId,ServerId string
}
type C_QuickEnter struct {

}