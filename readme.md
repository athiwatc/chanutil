# chanutil
[![Build Status](https://travis-ci.org/athiwatc/chanutil.png?branch=master)](https://travis-ci.org/athiwatc/chanutil)
Everything that I use in my projects that has pattern that deals with channels.
All of this is the pattern I needed in each of my projects which might not be ideal for general uses.
All pattern are self generating, this means that it's not flexible but allows a constant way of using the library through out.

## How to use

### Broadcast
This broadcast all value sent to `Sender()` to all Receivers that has been created by `CreateReceiver()`.

Create a broadcaster using `b := NewBroadcaster(buffer_len int)`

To get the sender you use `s := b.Sender()`. This gives you a send only channel.

To create a new reciever you call `r := b.CreateReceiver()`.

You can also close the sender channel `close(s)`. It will also close the reciver channel for you. You should no longer attempt to use this after using s.

### Filter
This filter all value sent to `Sender()` and output it to `Receiver()` if the method `Want` return true.

To create a filter new `f := NewFilter(fil Filter, buffer_len int)`.

To get the sender use `s := f.Sender()`.

To get the receiver use `r := f.Receiver()`.

To close the receiver, close the sender like `close(s)`.