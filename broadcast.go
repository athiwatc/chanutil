package chanutil

import (
	"sync"
)

type broadcast struct {
	sender   chan interface{}
	receiver []chan interface{}
	buff     int
	mu       sync.Mutex
}

func NewBroadcaster(buffer_len int) *broadcast {
	//fmt.Println("test")
	b := &broadcast{}

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

func (b *broadcast) Sender() chan<- interface{} {
	return b.sender
}

func (b *broadcast) CreateReceiver() <-chan interface{} {
	new_receiver := make(chan interface{}, b.buff)

	b.mu.Lock()
	b.receiver = append(b.receiver, new_receiver)
	b.mu.Unlock()

	return new_receiver
}
