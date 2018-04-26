package main

import (
	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
	"golang/conf"
	"server/game"
	"server/gate"
	"server/login"
	//"fmt"
	//"time"
)

func main() {
	config := conf.GetInstance()
	config.Load()
	lconf.LogLevel = config.Server.LogLevel
	lconf.LogPath = config.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = config.Server.ConsolePort
	lconf.ProfilePath = config.Server.ProfilePath

	//go func() {
	//	for {
	//	time.Sleep(time.Second*1)
	//	fmt.Println("aaaaaaaaaaaaaaaaa")
	//	fmt.Println(config.Server.LogLevel)
	//	}
	//
	//
	//}()

	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)
}
