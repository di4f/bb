package mkdir

import(
	"fmt"
	"os"
	"flag"
)

func Run(args []string) int {
	status := 0
	arg0 := args[0]
	args = args[1:]
	var (
		parentFlag bool
		 modeArg int
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
	}
	mode := os.FileMode(modeArg)
	for _,  path := range args {
		var e error
		if parentFlag {
			e = os.MkdirAll(path, mode)
		} else {
			e =  os.Mkdir(path, mode)
		}
		if e != nil {
			fmt.Fprintf(os.Stderr, "%s: %s.\n", arg0, e)
			status = 1
		}
	}
	return status
}
