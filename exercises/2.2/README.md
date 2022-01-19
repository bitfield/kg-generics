# Exercise 2.2: Group therapy

This time, you'll need to define a generic slice type `Group[E]` to pass the following test:

```go
func TestGroup(t *testing.T) {
	t.Parallel()
	got := group.Group[string]{}
	got = append(got, "hello")
	got = append(got, "world")
	want := group.Group[string]{"hello", "world"}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
```

If you get stuck, have a look at [Solution 2.2](../../solutions/2.2/group.go).

---

[Index](../../README.md) - [Next](../2.3/)