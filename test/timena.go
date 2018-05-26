package main

import (
	"fmt"
	"time"
	"strconv"
)

func main() {
	fmt.Println(time.Now().UnixNano())
	//fmt.Println("server-"+strconv.Itoa(int(time.Now().Nanosecond())))
	a:=strconv.FormatInt(time.Now().UnixNano(),10)
	fmt.Println(a)
}
