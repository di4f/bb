package mkdir

import(
	"os"
	"flag"
	"log"
)

func Run(args []string) int {
	var (
		parentFlag, verbFlag bool
		 modeArg int
	)
	flagSet := flag.NewFlagSet(args[0], flag.ExitOnError)
	status := 0
	flagSet.BoolVar(&parentFlag, "p", false, "No error if existing, make parent as needed.")
	flagSet.IntVar(&modeArg, "m", 0766, "Set file `mode`.")
	flagSet.BoolVar(&verbFlag, "v", false, "Print a message for each created directory.")
	if len(args) < 2 {
		flagSet.Usage()
	}
	flagSet.Parse(args[1:])
	mode := os.FileMode(modeArg)
	var verb *log.Logger
	if verbFlag {
		verb = log.New(os.Stdout, flagSet.Args()[0]+": ", 0)
	}
	warn := log.New(os.Stderr, flagSet.Args()[0]+": ", 0)
	for _,  path := range flagSet.Args() {
		var err error
		if parentFlag {
			err = os.MkdirAll(path, mode)
		} else {
			err =  os.Mkdir(path, mode)
		}
		if err != nil {
			warn.Println(err)
			status = 1
		} else if verbFlag {
			verb.Printf("Created directory '%s'.", path)
		}
	}
	return status
}