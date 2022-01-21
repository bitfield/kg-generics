# Exercise 7.1: Stack overflow

A [stack](https://en.wikipedia.org/wiki/Stack_(abstract_data_type)) is an abstract data type which stores items in a last-in-first-out way. The only thing you can do with a stack is "push" a new item onto it, or "pop" the top item off it.

Your task here is to define a generic `Stack` type that holds an ordered collection of values of as broad a range of types as possible.

You'll need to write a `Push` method that can append any number of items to the stack, in order, and a `Pop` method that retrieves (and removes) the last item from the stack.

`Pop` should also return a second `ok` value indicating whether an item was actually popped. If the stack is empty, then `Pop` should return `false` for this second value, but `true` otherwise.

Also, you should provide a `Len` method that returns the number of items in the stack.

For example:

```go
s := Stack[int]{}
fmt.Println(s.Pop())
// 0 false
s.Push(5, 6)
fmt.Println(s.Len())
// 2
v, ok := s.Pop()
fmt.Println(v)
// 6
fmt.Println(ok)
// true
```

If you'd like some hints, check out [Solution 7.1](../../solutions/7.1/stack.go) for one possible way to solve this.

---

[Index](../../README.md)
