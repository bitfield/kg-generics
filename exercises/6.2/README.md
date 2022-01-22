# Exercise 6.2: Compose yourself

_Composing_ functions means applying them in a chain, so that each function operates on the result of the previous one.

For example, if we had two arbitrary functions `outer` and `inner`, then we could compose them directly by writing:

```go
outer(inner())
```

Your job in this exercise is to write a `Compose` function that will do just this. Given two functions, `outer` and `inner`, and a value of arbitrary type, it should first apply `inner` to the value, then apply `outer` to the result returned by `inner`, then return the result returned by `outer`.

Got all that? Here's an example. Suppose we had a function `double` that returns double its input, and another function `addOne` that returns the result of adding 1 to its input.

And suppose our input value is 1. If we first apply `addOne` to 1, we should get 2. If we then apply `double` to 2, we should get 4. So:

```go
fmt.Println(Compose(double, addOne, 1))
// 4
```

What if we composed the functions in reverse order, so that we applied `double` first, then `addOne` to the result? Well:

```go
fmt.Println(Compose(addOne, double, 1))
// 3
```

If you're having trouble with this, remember that the `inner` function will need to take a parameter of the same type as the value, and its result type is arbitrary. The `outer` function will need to take a parameter of whatever type `inner` returns, but _its_ result type is also arbitrary.

So, all told, there are _three_ type parameters involved here.

If you're still stuck, have a look at one possible way of doing it in [Solution 6.2](../../solutions/6.2/compose.go).

---

[Index](../../README.md) - [Next](../7.1/)
