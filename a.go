package main

import (
	"fmt"
	"github.com/satori/go.uuid"
)

func main() {
	// Creating UUID Version 4
	// panic on error
	u1 := uuid.Must(uuid.NewV4())
	fmt.Printf("UUIDv4: %s\n", u1)

	// or error handling
	u2, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}
	fmt.Printf("UUIDv4: %s\n", u2)

	// Parsing UUID from string input
	u2, err = uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
	}
	fmt.Println("Successfully parsed: %s", u2)


	fmt.Println(uuid.FromBytes([]byte("db1golang.comcom")))
	//alist := []string{"db1","db2","db3","db4","db5"}
	fmt.Println(1001%5)
	fmt.Println(1002%5)
	fmt.Println(1003%5)
	fmt.Println(1004%5)
	fmt.Println(1005%5)
}



type aaa struct {
	v := atomic.LoadInt32(&value)
}