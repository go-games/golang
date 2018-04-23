package msg

import (
	"github.com/name5566/leaf/log"

	"github.com/name5566/leaf/network/json"
)

var Processor = json.NewProcessor()

var HelloWorld string

type Hello struct {
	Name string
}
type UserLogin struct {
	LoginName string
	LoginPW   string
}

func init() {
	log.Release("Leaf re lasestarting up")

	Processor.Register(&Hello{})
	Processor.Register(&UserLogin{})
}
