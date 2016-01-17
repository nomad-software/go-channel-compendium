/*
In this example, a goroutine is started, does some work (in this case waiting
for five seconds) then closes the channel. Unbuffered channels always halt the
current goroutine until communication can take place. Closing the channel
signals to the goroutine that it can continue because there is no more data to
be received. Closed channels never halt execution of the goroutine.
*/
package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan bool)

	go func() {
		// ... do some stuff
		time.Sleep(time.Second * 5)
		close(c)
	}()

	// Halt for communication of data via the channel or for it to be closed.
	<-c

	fmt.Println("Done")
}
