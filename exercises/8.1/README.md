# Exercise 8.1: Channelling frustration

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

You'll need to ensure that your `Channel` is concurrency-safe, so make sure your solution passes the tests when the race detector is enabled:

**`go test -race`**

See [Solution 8.1](../../solutions/8.1/billable.go) if you'd like something to compare with your answer.

---

[Index](../../README.md) - [Next](../9.1/)
