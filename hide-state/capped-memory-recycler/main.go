/*
In this example, a single buffered channel is used as a store for our memory
buffers. The channel is set to buffer five entries at any given time. This
means the channel will not halt the current goroutine if there is capacity in
the channel to accept another entry.

The select statements provide non-blocking access to this channel in the case
that it would be full. The first select creates a new buffer if it's unable to
get a buffer from the store. The second select defaults to nothing if it's
unable to place one on the store, which invokes the garbage collector to
deallocate the given buffer.
*/
package main

func get(store chan []byte) []byte {
	select {
	case b := <-store:
		return b
	default:
		return make([]byte, 100)
	}
}

func give(store chan []byte, b []byte) {
	select {
	case store <- b:
	default:
		return
	}
}

func main() {

	// Create a store for the buffers.
	store := make(chan []byte, 5)

	// Gets a new buffer from the store.
	buffer := get(store)

	// Give it back to the store.
	give(store, buffer)

	// Get the recycled buffer again from the store.
	buffer = get(store)
}
