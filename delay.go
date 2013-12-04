package chanutil

import (
	"sync"
	"time"
)

// Never create this type yourself.
type Delay struct {
	sender   chan interface{}
	receiver chan interface{}
	delay    time.Duration
	c        sync.WaitGroup
}

// Use this function to create a delay channel.
// A delay channel will make sure that all value are sent before closing will happen.
// This means that Receiver channal will not close after the Sender is closed and will wait till all values are sent.
func NewDelay(delay time.Duration, buffer_len int) *Delay {
	d := &Delay{}
	d.delay = delay

	d.sender = make(chan interface{}, buffer_len)
	d.receiver = make(chan interface{}, buffer_len)

	go func() {
		for s := range d.sender {
			d.c.Add(1)
			time.AfterFunc(delay, func() {
				d.receiver <- s
				d.c.Done()
			})
		}

		d.c.Wait()
		close(d.receiver)
	}()

	return d
}

// Get the delay sender
func (d *Delay) Sender() chan<- interface{} {
	return d.sender
}

// Get the delay receiver
func (d *Delay) Receiver() <-chan interface{} {
	return d.receiver
}
