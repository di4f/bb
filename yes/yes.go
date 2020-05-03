/* Yes program implementation. */
package yes

import(
	"os"
	"fmt"
	"flag"
	"strings"
	"github.com/jienfak/goblin/input"
)
var(
	nArg int
)

func yes(s string){
	if nArg<0{
		for{
			fmt.Print(s)
		}
	} else {
		for i:=0 ; i<nArg ; i+=1 {
			fmt.Print(s)
		}
	}
}

func Run(args []string) int {
	var(
		stdinFlag bool
		nFlag bool
		s string
	)
	status := 0
	arg0 := args[0]
	args = args[1:]
	flagSet := flag.NewFlagSet(arg0, flag.ExitOnError)
	flagSet.BoolVar(&stdinFlag, "s", false, "Read string from stdin.")
	flagSet.BoolVar(&nFlag, "n", false, "Do not add net line character.")
	flagSet.IntVar(&nArg, "N", -1, "Repeat input N times. Negative value means infinite cycle.")
	flagSet.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s: %s [options] [string]\n", arg0, arg0)
		flagSet.PrintDefaults()
	}
	flagSet.Parse(args)
	args = flagSet.Args()

	if stdinFlag {
		in, _ := input.ReadAllRaw(os.Stdin)
		s = string(in)
	} else {
		if len(args)>0 {
			s = strings.Join(args, " ")
		} else {
			s = "y"
		}
	}

	if !nFlag {
		s += "\n"
	}

	yes(s)

	return status
}
