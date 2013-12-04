# chanutil
[![Build Status](https://drone.io/github.com/athiwatc/chanutil/status.png)](https://drone.io/github.com/athiwatc/chanutil/latest)
[![Coverage Status](https://coveralls.io/repos/athiwatc/chanutil/badge.png?branch=master)](https://coveralls.io/r/athiwatc/chanutil?branch=master)

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

### Delay
This will delay the output from `Sender()` to the `Receiver()` with the amount of time given.

To create a channel delay use `d := NewDelay(t time.Duration, buffer_len int)`.

Then send to it `d.Sender() <- "My Value"`.

After a given amount of time `<-d.Receiver()` the value can be retrive.

### Link
As the design of these functions are self contain and you can't use your own channels, this means that you need a way to pull from a channel and put it in another. It works like a pipe.
It will close the destination channel also if you closed the source channel. There's no way to make sure to know when will unlink happen.

Use it by `l := NewLink(des, src chan interface{})`

And unlink it by `l.Unlink()`

Note that if you intend to close the source then you can't call `Unlink()`. By closing source it will automatically unlink it for you(and close the destination). Never call `Unlink()` twice.