package main

import (
	"fmt"

	"github.com/avinashmk/goTicketSystem/control"
)

func main() {
	fmt.Println("Hello World")
	defer control.Stop()
	control.Init()
	control.Start()
}
