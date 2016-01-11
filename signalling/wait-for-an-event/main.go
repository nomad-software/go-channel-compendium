/*
In this example, a goroutine is started, does some work (in this case waiting
for five seconds) then closes the channel. The closing of the channel signals
the main goroutine to continue. There is no communication of data here, only
syncronisation.
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
