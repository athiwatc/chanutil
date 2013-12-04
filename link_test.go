package chanutil

import (
	"testing"
)

func TestLinker(t *testing.T) {
	c1 := make(chan interface{}, 10)
	c2 := make(chan interface{}, 10)

	l := NewLink(c1, c2)

	c2 <- 20
	if <-c1 != 20 {
		t.Fail()
	}

	l.Unlink()
}

func TestLinkerUnlink(t *testing.T) {
	c1 := make(chan interface{}, 10)
	c2 := make(chan interface{}, 10)

	l := NewLink(c1, c2)
	l.Unlink()

	c2 <- 20
	if len(c1) != 0 {
		t.Fail()
	}
}

func TestLinkerLink(t *testing.T) {
	c1 := make(chan interface{}, 10)
	c2 := make(chan interface{}, 10)

	l := NewLink(c1, c2)

	c2 <- 20
	if len(c1) != 0 {
		t.Fail()
	}

	l.Unlink()
}

func TestLinkerUnbuffed(t *testing.T) {
	c1 := make(chan interface{})
	c2 := make(chan interface{})

	l := NewLink(c1, c2)

	// Should not block as the value will get transfered to c1
	c2 <- 20

	if <-c1 != 20 {
		t.Fail()
	}

	l.Unlink()
}

func TestLinkerClose(t *testing.T) {
	c1 := make(chan interface{})
	c2 := make(chan interface{})

	NewLink(c1, c2)

	close(c2)

	if _, ok := <-c1; ok == true {
		t.Fail()
	}
}
