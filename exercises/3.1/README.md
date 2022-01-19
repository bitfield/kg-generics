# Exercise 3.1: Stringy beans

Flex your generics muscles a little now, by writing a generic function constrained by `fmt.Stringer`.

Your job here is to write a generic function `StringifyTo[T]` that takes an `io.Writer` and a value of some arbitrary type constrained by `fmt.Stringer`, and prints the value to the writer.

If you're not sure what to do, peek at [Solution 3.1](../../solutions/3.1/stringy.go) for hints.

---

[Index](../../README.md) - [Next](../3.2/)