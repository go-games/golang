package protocol

type Resp struct{
    Error string
    Code int
}
type LoginSuccess struct{
	Error string
	Code int
    Result struct{
        UserId,UserName string
        IsRobt,IsGuest bool
        State int
    }

}
type RegisterSuccess struct {
    Error string
    Code int
    Result struct{
        UserId,UserName string
        IsRobt,IsGuest bool
        State int
    }
}
type Failed struct {
    Error string
    Code int
    Result string
}
type RoomCreateSuccess struct {
    Error string
    Code int
    Result struct{
        RoomMasterId,RoomId string
        State,Limit,ServerId int
    }
}
type RoomEnterSuccess struct {
    Error string
    Code int
    Result struct{
        UserId,RoomId string
        ServerId,Sit int

    }
}
type RoomQuitSuccess struct {
    Error string
    Code int
    Result struct{
        UserId,RoomId string
        ServerId int
    }
}
type RoomQuickEnterSuccess struct {
    RoomEnterSuccess
}