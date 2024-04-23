package main

import (
	"fmt"
	"time"
)

func Print(s string) {
	fmt.Println(s)
}

func main() {
	go Print("hello world")

	time.Sleep(5 * time.Second)

	Print("second")
}
