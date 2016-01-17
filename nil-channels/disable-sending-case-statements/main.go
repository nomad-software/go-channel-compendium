/*
In this example, a goroutine is started to generate random numbers and send
them on the c channel. If a message is sent on the d channel, the c channel is
set to nil, disbling the associated case statement. Once disabled, the
goroutine can no longer send random numbers.
*/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	c := make(chan int)
	d := make(chan bool)

	go func(src chan int) {
		for {
			select {
			case src <- rand.Intn(100):

			case <-d:
				src = nil
			}
		}
	}(c)

	// Print some random numbers.
	fmt.Printf("%d\n", <-c)
	fmt.Printf("%d\n", <-c)

	// Disable random number generation.
	d <- true

	// Halts because c is now nil.
	fmt.Printf("%d\n", <-c)
}
