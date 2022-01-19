# Exercise 3.3: A first approximation

Here you're provided with a function `IsPositive`, which determines whether a given value is greater than zero, and a set of accompanying tests.

The `IsPositive` function is constrained by some interface `Intish`, and the tests call it with values of the following type:

```go
type MyInt int
```

Your task here is to define the `Intish` interface. A method set won't work, because `int` has no methods. On the other hand, the type literal `int` won't work either, because `MyInt` is not `int`.

What could you use instead?

If you'd like some clues, have a look at [Solution 3.3](../../solutions/3.3/intish.go).

---

[Index](../../README.md) - [Next](../3.4/README.md)
