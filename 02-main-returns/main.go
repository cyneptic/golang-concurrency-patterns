package main

import (
	"fmt"
	"time"
)

func main() {
	go boring("boring!") // shouldn't run because main returns
}

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Second)
	}
}
