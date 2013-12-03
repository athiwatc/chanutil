package chanutil

import (
	"testing"
)

type AFilterFunc struct{}

func (a AFilterFunc) Want(i interface{}) bool {
	return i.(int)%2 == 0
}

func TestFilter(t *testing.T) {
	var a FilterInterface = AFilterFunc{}
	f := NewFilter(a, 20)

	f.Sender() <- 1
	f.Sender() <- 2
	f.Sender() <- 3
	f.Sender() <- 4
	close(f.Sender())

	if <-f.Receiver() != 2 && <-f.Receiver() != 4 {
		t.Fail()
	}

}
