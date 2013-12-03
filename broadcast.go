package chanutil

import (
	"sync"
)

type broadcast struct {
	sender    chan interface{}
	recievers []chan interface{}
	buff      int
	mu        sync.Mutex
}

func NewBroadcaster(buffer_len int) *broadcast {
	//fmt.Println("test")
	b := &broadcast{}

	b.buff = buffer_len

	b.sender = make(chan interface{}, b.buff)

	go func() {
		for s := range b.sender {
			for _, r := range b.recievers {
				r <- s
			}
		}

		// The sender channel is closed
		for _, r := range b.recievers {
			close(r)
		}
	}()
	return b
}

func (b *broadcast) Sender() chan<- interface{} {
	return b.sender
}

func (b *broadcast) CreateReciever() <-chan interface{} {
	new_reciever := make(chan interface{}, b.buff)

	b.mu.Lock()
	b.recievers = append(b.recievers, new_reciever)
	b.mu.Unlock()

	return new_reciever
}
