/* Yes program implementation. */
package yes

import (
	"os"
	"fmt"
	"strings"
	"vultras.su/core/bb/input"
	"vultras.su/core/cli/mtool"
)

var (
	nArg int
)

func yes(s string) {
	if nArg < 0 {
		for {
			fmt.Print(s)
		}
	} else {
		for i := 0; i < nArg; i += 1 {
			fmt.Print(s)
		}
	}
}

func Run(flagSet *mtool.Flags) {
	var (
		stdinFlag bool
		nFlag     bool
		s         string
	)
	flagSet.BoolVar(&stdinFlag, "s", false, "Read string from stdin.")
	flagSet.BoolVar(&nFlag, "n", false, "Do not add net line character.")
	flagSet.IntVar(&nArg, "N", -1, "Repeat input N times. Negative value means infinite cycle.")

	flagSet.Parse()
	args := flagSet.Args()

	if stdinFlag {
		in, _ := input.ReadAllRaw(os.Stdin)
		s = string(in)
	} else {
		if len(args) > 0 {
			s = strings.Join(args, " ")
		} else {
			s = "y"
		}
	}

	if !nFlag {
		s += "\n"
	}

	yes(s)

}
