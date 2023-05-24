package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	str  string
	wait chan bool
}

func main() {
	quit := make(chan string)
	c := boring("Joe", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		msg := <-c
		fmt.Println(msg.str)
	}
	quit <- "Bye!"
	fmt.Println("I said: 'Bye!'")
	fmt.Printf("Joe says: %q\n", <-quit)
}

func boring(msg string, quit chan string) <-chan Message {
	c := make(chan Message)

	waitForIt := make(chan bool)

	go func() {
		for i := 0; ; i++ {

			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			select {
			case c <- Message{fmt.Sprintf("%s: %d", msg, i), waitForIt}:
			case <-quit:
				//cleanup()
				quit <- "See You!"
				return
			}
		}
	}()
	return c
}

func fanIn(inputs ...<-chan Message) <-chan Message {
	c := make(chan Message)
	for i := range inputs {
		input := inputs[i]
		go func() {
			for {
				c <- <-input
			}
		}()
	}
	return c
}
