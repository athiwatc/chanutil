package chanutil

// Implement this for use with the NewFilter function.
type FilterInterface interface {
	// Return true if you want the value to get sent to the output channel
	Want(interface{}) bool
}

// Never create this type yourself
type Filter struct {
	sender   chan interface{}
	receiver chan interface{}
}

// Use this function to setup a new Filter
func NewFilter(filterFunc FilterInterface, buffer_len int) *Filter {
	f := &Filter{}
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

// Get the Sender
func (f *Filter) Sender() chan<- interface{} {
	return f.sender
}

// Get the Receiver
func (f *Filter) Receiver() <-chan interface{} {
	return f.receiver
}
