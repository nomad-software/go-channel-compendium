/*
In this example, a goroutine is started, waiting for communication of data on
the die channel (or for it to be closed). In this case, once closed, the
goroutine performs termination tasks, then signals to the main goroutine (via
the same die channel) that it's finished.
*/
package main

func worker(die chan bool) {
	for {
		select {
		// ... do stuff cases
		case <-die:
			// ... do termination tasks
			die <- true
			return
		}
	}
}

func main() {
	die := make(chan bool)
	go worker(die)
	die <- true

	// Wait until the goroutine has terminated.
	<-die
}
