# Exercise 6.1: Func to funky

For this exercise, you'll be building a simple _dynamic dispatch_ system in Go. What's required here is a way of storing a number of different functions by name, and then applying the one you want.

Your job is to create a `FuncMap` type which can store a number of named functions. That is, it's a map of `string` to some arbitrary `func` type. For example:

```go
fm := FuncMap[int, int]{
	"double": func(i int) int {
		return i * 2
	},
	"addOne": func(i int) int {
		return i + 1
	},
}
```

You'll also need to implement an `Apply` method on `FuncMap` which will apply the specified function to some value and return the result. For example:

```go
fmt.Println(fm.Apply("double", 2))
// 4
```

If it stops being fun, take a look at one possible answer in [Solution 6.1](../../solutions/6.1/funcmap.go).

---

[Index](../../README.md) - [Next](../6.2/)
