/*
In this example, the w channel is created to transfer a unit of work to a
goroutine. This unit of work is received and a GET request is made to the
contained URL. As part of this work, a response channel is also passed. Once
the GET request is performed, the query response is sent back along the
response channel. This allows this goroutine to process work and send a
response back on different channels configured for each unit of work.
*/
package main

import "net/http"

type work struct {
	url  string
	resp chan *http.Response
}

func getter(w chan work) {
	for {
		do := <-w
		resp, _ := http.Get(do.url)
		do.resp <- resp
	}
}

func main() {

	w := make(chan work)

	go getter(w)

	resp := make(chan *http.Response)
	w <- work{"http://cdnjs.cloudflare.com/jquery/1.9.1/jquery.min.js", resp}
	r := <-resp
	// ... do something
}
