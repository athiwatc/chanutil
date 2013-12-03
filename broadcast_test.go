package chanutil

import (
	"testing"
)

func TestBroadcastCreation(t *testing.T) {
	NewBroadcaster(10)
}

func TestBroadcasting(t *testing.T) {
	b := NewBroadcaster(10)

	r1 := b.CreateReciever()
	r2 := b.CreateReciever()

	b.Sender() <- 100

	if <-r1 != 100 {
		t.Fail()
	}

	if <-r2 != 100 {
		t.Fail()
	}
}

func TestBroadcastClose(t *testing.T) {
	b := NewBroadcaster(10)

	r1 := b.CreateReciever()
	r2 := b.CreateReciever()

	close(b.Sender())

	// None of this should be a deadlock/block
	<-r1
	<-r2
}

func TestBroadcastRace(t *testing.T) {
	b := NewBroadcaster(1000)

	go func() {
		b.CreateReciever()
	}()
	go func() {
		b.CreateReciever()
	}()
	go func() {
		b.CreateReciever()
	}()

	go func() {
		b.Sender()
	}()
	go func() {
		b.Sender()
	}()
	go func() {
		b.Sender()
	}()

}
