/* Simple 'echo' implementation. */
package echo

import (
	"fmt"
	"github.com/vultras/cli/mtool"
)

var(
	del string
	eol = "\n"
)

func Run(flagSet *mtool.Flags) {
	var nflag bool
	
	flagSet.BoolVar(&nflag, "n", false, "Do not print new line character.")
	flagSet.StringVar(&del, "d", " ", "Delimiter of arguments")
	
	flagSet.Parse()
	args := flagSet.Args()
	
	l := len(args) - 1
	for i, s := range args {
		fmt.Print(s)
		if i!=l { fmt.Print(del) }
	}
	if !nflag {
		fmt.Print(eol)
	}

}
