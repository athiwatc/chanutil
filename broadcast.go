// Collection of methods to deal with channels
package chanutil

import (
	"sync"
)

// Broadcast allow a value send to the sender channel to be multiplex to all receiver created with CreateReceiver() method.
// If you close the sender, all of the receiver will also be closed.
type Broadcast struct {
	sender   chan interface{}
	receiver []chan interface{}
	buff     int
	mu       sync.Mutex
}

// Use this function to setup a new Broadcaster
func NewBroadcaster(buffer_len int) *Broadcast {
	//fmt.Println("test")
	b := &Broadcast{}

	b.buff = buffer_len

	b.sender = make(chan interface{}, b.buff)

	go func() {
		for s := range b.sender {
			for _, r := range b.receiver {
				r <- s
			}
		}

		// The sender channel is closed
		for _, r := range b.receiver {
			close(r)
		}
	}()
	return b
}

// Get the Sender
func (b *Broadcast) Sender() chan<- interface{} {
	return b.sender
}

// Create and return a new Receiver
func (b *Broadcast) CreateReceiver() <-chan interface{} {
	new_receiver := make(chan interface{}, b.buff)

	b.mu.Lock()
	b.receiver = append(b.receiver, new_receiver)
	b.mu.Unlock()

	return new_receiver
}
