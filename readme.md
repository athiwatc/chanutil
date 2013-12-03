# chanutil
Everything that I use in my projects that has pattern that deals with channels.
All of this is the pattern I needed in each of my projects which might not be ideal for general uses.

## How to use

### Broadcast
Create a broadcaster using `b := NewBroadcaster(buffer_len int)`

To get the sender you use `s := b.Sender()`. This gives you a send only channel.

To create a new reciever you call `r := b.CreateReciever()`.

You can also close the sender channel `close(s)`. It will also close the reciver channel for you. You should no longer attempt to use this after using s.