# Exercise 3.4: Greater love

You've been given the following (incomplete) function:

```go
func IsGreater[T /* Your constraint goes here! */ }](x, y T) bool {
	return x.Greater(y)
}
```

This takes two values of some arbitrary type, and compares them by calling the `Greater` method on the first value, passing it the second value.

The tests exercise this function by calling it with two values of a defined type `MyInt`, which has the required `Greater` method. So to make these tests pass, you'll need to write an appropriate type constraint for `IsGreater`.

Can you see what to do?

You can check your answer against [Solution 3.4](https://github.com/bitfield/kg-generics/blob/main/solutions/3.4/greater.go).

---

[Index](../../README.md) - [Next](../4.1/)
