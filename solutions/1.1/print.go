package print

import (
	"fmt"
	"io"
)

func PrintAnythingTo[T any](w io.Writer, p T) {
	fmt.Fprintln(w, p)
}
