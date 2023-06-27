package main

import (
	"fmt"

	"github.com/degreane/ezekod.com/server"
)

func init() {
	fmt.Println("Init from main")
}

func main() {
	server.Start()
	fmt.Print("Main Func called")
	// log.Fatal()
}
