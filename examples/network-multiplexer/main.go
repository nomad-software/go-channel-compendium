/*
This example demonstrates a simple network multiplexer. Within the main
goroutine, a channel is created to handle transfering messages and a network
connection is established. A hundred goroutines are then spawned to generate
strings (to act as our messages) and sent along this channel. Each message is
read from the channel inside a continuous loop and sent to the network
connection.

This example doesn't run (because we are trying to connect to an example
domain) but it expresses how easy it is to have many asyncronous processes
sending messages to a single network connection.
*/
package main

import "net"

func worker(messages chan string) {
	for {
		var msg string // ... generate a message
		messages <- msg
	}
}

func main() {

	messages := make(chan string)
	conn, _ := net.Dial("tcp", "example.com")

	for i := 0; i < 100; i++ {
		go worker(messages)
	}

	for {
		msg := <-messages
		conn.Write([]byte(msg))
	}
}
