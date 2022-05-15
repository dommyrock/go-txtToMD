package errorHandling

import (
	"fmt"
	"log"
	"os"

	"github.com/TwiN/go-color"
	"github.com/dommyrock/txtToMD/library/textUtil"
	"github.com/dommyrock/txtToMD/types"
)

func ExitGracefully(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

func ShouldExit(dict map[string]types.Prefix) bool {
	args := len(os.Args)
	if args < 2 {
		log.Fatal("No file specified")
		return true
	} else if args == 2 && (os.Args[1] == "-options" || os.Args[1] == "-o") {
		println(color.InCyan("Available mappings:"))
		fmt.Print(textUtil.MapKeys(dict))
		return true
	}
	return false
}
