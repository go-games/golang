package main

import (
	"fmt"
	"time"
)

func main() {


	var a string
	a = "bbbb"


	go func() {
		a= "cccc"
	}()

	time.Sleep(1* time.Second)

	fmt.Println(a)
	//输出结果 aaaa
}
