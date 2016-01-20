/*
In this example, a load balancer has been created building upon previous
examples. It handles reading URL's from stdin and starting goroutines to
perform a request for each. Each request is passed through a load balancer to
filter these jobs into a finite number of workers. These workers process the
requests and ultimately return a response to a single channel. The responses
are then printed.

Using a load balancer such as this can take a huge number of requests, balance
them across resources available and process them in an orderly manner.
*/
package main

import (
	"fmt"
	"net/http"
)

type job struct {
	url  string
	resp chan *http.Response
}

type worker struct {
	jobs  chan *job
	count int
}

func (w *worker) getter(done chan *worker) {
	for {
		j := <-w.jobs
		resp, _ := http.Get(j.url)
		j.resp <- resp
		done <- w
	}
}

func get(jobs chan *job, url string, answer chan string) {
	resp := make(chan *http.Response)
	jobs <- &job{url, resp}
	r := <-resp
	answer <- r.Request.URL.String()
}

func balancer(count int, depth int) chan *job {
	jobs := make(chan *job)
	done := make(chan *worker)
	workers := make([]*worker, count)

	for i := 0; i < count; i++ {
		workers[i] = &worker{make(chan *job, depth), 0}
		go workers[i].getter(done)
	}

	go func() {
		for {
			var free *worker
			min := depth
			for _, w := range workers {
				if w.count < min {
					free = w
					min = w.count
				}
			}
			var jobsource chan *job
			if free != nil {
				jobsource = jobs
			}
			select {
			case j := <-jobsource:
				free.jobs <- j
				free.count++

			case w := <-done:
				w.count--
			}
		}

	}()

	return jobs
}

func main() {
	jobs := balancer(10, 10)
	answer := make(chan string)
	for {
		var url string
		if _, err := fmt.Scanln(&url); err != nil {
			break
		}
		go get(jobs, url, answer)
	}
	for u := range answer {
		fmt.Printf("%s\n", u)
	}
}
