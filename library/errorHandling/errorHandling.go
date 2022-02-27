package errorHandling

import (
	"fmt"
	"os"
)

func ExitGracefully(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}
