package stringy

import (
	"fmt"
	"io"
)

func StringifyTo[T fmt.Stringer](w io.Writer, p T) {
	fmt.Fprintln(w, p.String())
}
