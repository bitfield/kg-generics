# Exercise 5.1: Empty promises

Your job is to create a `Sequence` type that can hold multiple values, and an `Empty` method on it that will report whether or not the sequence is empty.

For example:

```go
s := Sequence[string]{"a", "b", "c"}
fmt.Println(s.Empty())
// false
```

See [Solution 5.1](../../solutions/5.1/empty.go) if you'd like something to compare with your answer.

---

[Index](../../README.md) - [Next](../6.1/)
