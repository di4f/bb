/* Simple 'echo' implementation. */
package echo

import (
	"fmt"
	"flag"
)

var(
	del string
	eol = "\n"
)

func Run(args []string) {
	var nflag bool
	flagSet := flag.NewFlagSet(args[0], flag.ExitOnError)
	flagSet.BoolVar(&nflag, "n", false, "Do not print new line character.")
	flagSet.StringVar(&del, "d", " ", "Delimiter of arguments")
	flagSet.Parse(args[1:])
	args = flagSet.Args()
	l := len(args) - 1
	for i, s := range args {
		fmt.Print(s)
		if i!=l { fmt.Print(del) }
	}
	if !nflag {
		fmt.Print(eol)
	}

}
