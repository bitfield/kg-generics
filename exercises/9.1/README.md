# Exercise 9.1: Contain your excitement

Your job here is to write a generic function `ContainsFunc` that operates on slices of arbitrary type. Given a function argument, `ContainsFunc` should return true if the function returns true for any element in the slice, or false otherwise.

For example:

```go
s := []rune{'a', 'B', 'c'}
fmt.Println(ContainsFunc(s, unicode.IsUpper))
// true
```

See [Solution 9.1](../../solutions/9.1/contains.go) if you'd like something to compare with your answer.

---

[Index](../../README.md) - [Next](../9.2/)
