package chanutil

type Filter interface {
	Want(interface{}) bool
}

type filter struct {
	sender   chan interface{}
	receiver chan interface{}
}

func NewFilter(filterFunc Filter, buffer_len int) *filter {
	f := &filter{}
	f.sender = make(chan interface{}, buffer_len)
	f.receiver = make(chan interface{}, buffer_len)

	go func() {
		for s := range f.sender {
			if filterFunc.Want(s) {
				f.receiver <- s
			}
		}

		// Close the receiver when the sender is closed
		close(f.receiver)
	}()

	return f
}

func (f *filter) Sender() chan<- interface{} {
	return f.sender
}

func (f *filter) Receiver() <-chan interface{} {
	return f.receiver
}
