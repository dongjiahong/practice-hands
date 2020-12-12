package main

import (
	"fmt"
	"time"

	log "github.com/alecthomas/log4go"

	"log4/utils"
)

func main() {
	log.LoadConfiguration("./log.xml")
	fmt.Println("==>  i'am fmt println")
	log.Error("=+> i'm log4go println")
	utils.Hello()
	time.Sleep(time.Second * 2)
}
