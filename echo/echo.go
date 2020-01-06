/* Simple 'echo' implementation. */
package echo

import (
	"fmt"
	"flag"
)

func Run(args []string) int {
	flagSet := flag.NewFlagSet(args[0], flag.ExitOnError)
	flagSet.Parse(args[1:])
	args = args[1:]
	status := 0
	l := len(args) - 1
	for i, s := range args {
		fmt.Print(s)
		if i!=l { fmt.Print(" ") }
	}
	fmt.Print("\n")
	return status
}
