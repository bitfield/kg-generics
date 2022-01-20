# Exercise 4.2: Duplicate keys

The task here is to determine if a given slice contains duplicated elements. For example, if a slice of numbers contains any given number more than once, then the slices contains duplicates. The duplicate elements need not be consecutive.

To do this, you'll need to write a function `Dupes`, with an appropriate type constraint, that returns `true` if the given slice contains duplicates, or `false` otherwise.

For example:

```go
fmt.Println(Dupes([]int{1, 2, 3}]
// false
fmt.Println(Dupes([]int{1, 2, 1}]
// true
```

See [Solution 4.2](../../solutions/4.2/product.go) if you'd like something to compare with your answer.

---

[Index](../../README.md) - [Next](../5.1/)
