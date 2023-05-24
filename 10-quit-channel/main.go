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
	quit := make(chan bool)
	c := boring("Joe", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		msg := <-c
		fmt.Println(msg.str)
	}
	quit <- true
}

func boring(msg string, quit <-chan bool) <-chan Message { // Returns receive-only channel of strings.
	c := make(chan Message)

	waitForIt := make(chan bool) // Shared between all messages.

	go func() { // We launch the goroutine from inside the function.
		for i := 0; ; i++ {

			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			select {
			case c <- Message{fmt.Sprintf("%s: %d", msg, i), waitForIt}:
			case <-quit:
				return
			}
		}
	}()
	return c // Return the channel to the caller.
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
