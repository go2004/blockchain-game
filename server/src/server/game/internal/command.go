package internal

import (
	"fmt"
	"github.com/name5566/leaf/log"
)

func init() {
	skeleton.RegisterCommand("echo", "echo user inputs", commandEcho)
}

func commandEcho(args []interface{}) interface{} {
	log.Release("commandEcho：%v", args)
	return fmt.Sprintf("commandEcho：%v", args[0])
}
