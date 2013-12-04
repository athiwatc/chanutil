package chanutil

// Never create this type yourself.
type Link struct {
	done chan struct{}
}

// Create a new linker.
func NewLink(dst, src chan interface{}) *Link {
	l := &Link{}
	l.done = make(chan struct{})

	go func() {
	outter:
		for {
			select {
			case <-l.done:
				break outter
			case v, ok := <-src:
				if ok {
					dst <- v
				} else {
					close(dst)
					close(l.done)
					break outter
				}
			}
		}
	}()

	return l
}

// Closing of b will automatically Unlink for you.
// Never call Unlink if you intend to close b or call Unlink() twice.
func (l *Link) Unlink() {
	close(l.done)
}
