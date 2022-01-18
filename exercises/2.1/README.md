# Exercise 2.1: Hello, generics

It's time to write your first generic function in Go!

Take a look at the [`print_test.go`](print_test.go) file in this folder. You'll find the following test:

```go
func TestPrintAnythingTo(t *testing.T) {
	t.Parallel()
	buf := &bytes.Buffer{}
	print.PrintAnythingTo[string](buf, "Hello, world")
	want := "Hello, world\n"
	got := buf.String()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
```

First, run the `go mod tidy` command to download the necessary modules for this test. Once that's completed, you're ready to start the exercise.

## Running the test

Run the test using your editor, or the `go test` command. You'll find right now the test doesn't compile, because the required function doesn't exist:

```
undefined: print.PrintAnythingTo
```

Remember, you'll need at least Go version 1.18 to be able to use generics, including running this test and implementing the function to make it pass. That's because we're using some new syntax that doesn't exist in Go version 1.17 and earlier.

If you try to compile this code with a version of Go that doesn't support generics, you'll get a rather confusing additional error message:

```
type string is not an expression
```

So if you see this, you need to upgrade your Go.

## Passing the test

To make this test compile, you'll need to define a function in the `print` package named `PrintAnythingTo` that takes one parameter of type `io.Writer`, and another value that's of some unspecified type.

To make the test _pass_, that function will need to write the supplied value to the supplied writer. It's up to you how to do this, but you might like to use `fmt.Fprintln`.

The necessary `go.mod` and `print.go` files are already set up for you. All you need to do is edit `print.go` and add the `PrintAnythingTo` function, then run the test again.

The main point of this exercise is to get you up and running writing generics in Go, and to get some practice declaring functions that take type parameters. When the test passes, you're done, so you can move on to the next exercise.

If you get stuck, you can take a look at my suggested answer in [Solution 2.1](../../solutions/2.1/print.go).

---

[Index](../../README.md) - [Next](../2.2/README.md)