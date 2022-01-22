# Exercise 9.2: Merging in turn

This task is about writing a `Merge` function that takes any number of maps (all of the same arbitrary type) and merges them together, returning the result.

In other words, the result returned by `Merge` is a map containing all the key-value pairs in all its arguments.

For example:

```go
m1 := map[int]bool{ 1: true }
m2 := map[int]bool{ 2: true }
fmt.Println(Merge(m1, m2)
// map[1:true 2:true]
```

If there are any conflicts (that is, if the maps have any key in common), then the "later" map wins. In other words, the maps are merged in the order that they're passed to `Merge`. For example:

```go
m1 := map[int]bool{ 1: false }
m2 := map[int]bool{ 1: true }
fmt.Println(Merge(m1, m2)
// map[1:true]
```

If you'd like to see one way of solving this, take a peek at [Solution 9.2](../../solutions/9.2/merge.go).

---

[Index](../../README.md)
