/*
In this example, a goroutine is started to do some work. A timeout channel is
created to make sure a case is executed if the select is halting for too long.
In this case the goroutine is terminated after thirty seconds of being idle.
The timeout is reset on every iteration of the select to make sure that if work
is done, the timeout is reset.
*/
package main

import "time"

func worker() {
	for {
		timeout := time.After(5 * time.Second)

		select {
		// ... do some stuff

		case <-timeout:
			// Close this go routine after the specified timeout.
			return
		}
	}
}

func main() {
	go worker()
}
