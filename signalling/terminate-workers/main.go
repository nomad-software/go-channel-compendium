/*
In this example, a hundred goroutines are started, waiting for communication of
data on the die channel (or for it to be closed). In this case, once closed,
all goroutines end.
*/
package main

func worker(die chan bool) {
	for {
		select {
		// ... do stuff cases
		case <-die:
			return
		}
	}
}

func main() {
	die := make(chan bool)

	for i := 0; i < 100; i++ {
		go worker(die)
	}

	// Close all workers.
	close(die)
}
