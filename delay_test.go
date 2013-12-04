package chanutil

import (
	"testing"
	"time"
)

func TestDelayClose(t *testing.T) {
	t.Parallel()
	d := NewDelay(50*time.Millisecond, 10)
	d.Sender() <- "Hello"
	close(d.Sender())
	if <-d.Receiver() != "Hello" {
		t.Fail()
	}
}

func TestDelaySend(t *testing.T) {
	t.Parallel()
	d := NewDelay(50*time.Millisecond, 10)
	d.Sender() <- "Hello"
	if <-d.Receiver() != "Hello" {
		t.Fail()
	}
}
