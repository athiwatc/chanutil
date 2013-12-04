package chanutil

// Never create this type yourself.
// Link or a Linker is what links multiple type of utility together. All utility in this pacakge are generators, this means that you can't provide your own channels and this is the only exception type.
// You can use this to filter + broadcast a value for example.
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
// Closing a closed channel is an error.
func (l *Link) Unlink() {
	close(l.done)
}
