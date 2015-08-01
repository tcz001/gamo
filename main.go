package main

import (
	"fmt"

	"github.com/tcz001/gamo/knife"
)

func main() {
	fmt.Println("Welcome to gamo kitchen")

	fmt.Println("Server starting...")
	s := knife.Server{}
	s.Init()
	fmt.Println("Server started...")
}
