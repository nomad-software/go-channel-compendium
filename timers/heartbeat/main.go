/*
In this example, a goroutine is started to do some work. A heartbeat channel is
created to make sure a case statement is executed at regular intervals. The
heartbeat channel is not reset on each iteration to make sure it always
executes on time.
*/
package main

import "time"

func worker() {
	heartbeat := time.Tick(30 * time.Second)
	for {

		select {
		// ... do some stuff

		case <-heartbeat:
			// ... do heartbeat stuff
		}
	}
}

func main() {
	go worker()
}
