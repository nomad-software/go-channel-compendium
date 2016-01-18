/*
In this example, an array of web URL's are iterated upon and passed
individually to separate goroutines. Each goroutine executes asyncronously and
queries the passed URL. Each query response is passed into the first channel,
which (of course) ensures the first query to respond is the first passed into
the channel. We can then read this response from the channel and act
accordingly.
*/
package main

import "net/http"

type response struct {
	resp *http.Response
	url  string
}

func get(url string, r chan response) {
	if resp, err := http.Get(url); err == nil {
		r <- response{resp, url}
	}
}

func main() {
	first := make(chan response)

	for _, url := range []string{"http://code.jquery.com/jquery-1.9.1.min.js",
		"http://cdnjs.cloudflare.com/ajax/libs/jquery/1.9.1/jquery.min.js",
		"http://ajax.googleapis.com/ajax/libs/jquery/1.9.1/jquery.min.js",
		"http://ajax.aspnetcdn.com/ajax/jQuery/jquery-1.9.1.min.js"} {
		go get(url, first)
	}

	r := <-first
	// ... do something
}
