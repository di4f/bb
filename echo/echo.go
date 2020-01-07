/* Simple 'echo' implementation. */
package echo

import (
	"fmt"
	"flag"
)

func Run(args []string) int {
	var nflag bool
	flagSet := flag.NewFlagSet(args[0], flag.ExitOnError)
	flagSet.BoolVar(&nflag, "n", false, "Do not print new line character.")
	flagSet.Parse(args[1:])
	args = flagSet.Args()
	l := len(args) - 1
	for i, s := range args {
		fmt.Print(s)
		if i!=l { fmt.Print(" ") }
	}
	if !nflag {
		fmt.Print("\n")
	}
	return 0
}
