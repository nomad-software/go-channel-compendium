/*
In this example, a single buffered channel is used as a store for our memory
buffers. The channel is set to buffer five entries at any given time. The
select statements provide non-blocking access to this channel to receive and
store memory buffers. The first select creates a new buffer if it's unable to
get a buffer from the store. The second select defaults to nothing if it's
unable to place one on the store which invokes the garbage collector to
deallocate the buffer. The rest of the code is sugar to provide a user friendly
API.
*/
package main

type Recycler struct {
	store chan []byte
}

func NewRecycler(size int) *Recycler {
	return &Recycler{make(chan []byte, size)}
}

func (r *Recycler) Get() []byte {
	select {
	case b := <-r.store:
		return b
	default:
		return make([]byte, 100)
	}
}

func (r *Recycler) Give(b []byte) {
	select {
	case r.store <- b:
	default:
		return
	}
}

func main() {

	// Create a new recycler.
	recycler := NewRecycler(5)

	// Gets a new buffer from the recycler.
	buffer := recycler.Get()

	// Give it back to the recycler.
	recycler.Give(buffer)

	// Get the recycled buffer again.
	buffer = recycler.Get()
}
