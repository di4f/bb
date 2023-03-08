package mkdir

import (
	"flag"
	"fmt"
	"os"

	"github.com/surdeus/goblin/src/pathx"
)

func Run(args []string) {
	arg0 := args[0]
	args = args[1:]
	var (
		parentFlag bool
		modeArg    int
	)
	flagSet := flag.NewFlagSet(arg0, flag.ExitOnError)
	flagSet.BoolVar(&parentFlag, "p", false, "No error if existing, make parent as needed.")
	flagSet.IntVar(&modeArg, "m", 0766, "Set file `mode`.")
	flagSet.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s: %s [options] [files]\n", arg0, arg0)
		flagSet.PrintDefaults()
	}
	flagSet.Parse(args)
	args = flagSet.Args()
	if len(args) == 0 {
		flagSet.Usage()
		os.Exit(1)
	}
	mode := os.FileMode(modeArg)
	for _, path := range args {
		var e error
		pth := pathx.From(path).Real()
		if parentFlag {
			e = os.MkdirAll(pth, mode)
		} else {
			e = os.Mkdir(pth, mode)
		}
		if e != nil {
			fmt.Fprintf(os.Stderr, "%s: %s.\n", arg0, e)
		}
	}
}
