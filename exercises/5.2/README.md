# Exercise 5.2: Channelling frustration

It seems people have been using channels in Go a little too freely, and a scheme is now proposed to bill users by activity. Accordingly, you'll need to define a generic `Channel` type with suitable instrumentation to track the number of sends and receives.

You have the following tasks to complete:

1. Define the `Channel` type.
2. Define a `NewChannel` function that returns an initialised `Channel` with a specified capacity, ready for use.
3. Define `Send` and `Receive` methods on the `Channel` to allow users to send and receive values of the appropriate type.
4. Define `Sends` and `Receives` methods that will report the number of sends and receives on the channel, for billing purposes.

For example:

```go
c := NewChannel[string](1)
c.Send("hello")
fmt.Println(c.Sends())
// 1
fmt.Println(c.Receive())
// hello
fmt.Println(c.Receives())
// 1
```

It is expected that the introduction of charges will discourage the use of channels altogether, so you don't need to worry about concurrency safety for this exercise. We'll return to that topic later in this book.

See [Solution 5.2](../../solutions/5.2/billable.go) if you'd like something to compare with your answer.

---

[Index](../../README.md) - [Next](../6.1/)
