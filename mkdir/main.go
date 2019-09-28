package main

import(
	"os"
	"flag"
	"log"
)

func main() {
	var (
		parentFlag, verbFlag bool
		 modeArg int
	)
	status := 0
	flag.BoolVar(&parentFlag, "p", false, "No error if existing, make parent as needed.")
	flag.IntVar(&modeArg, "m", 0766, "Set file `mode`.")
	flag.BoolVar(&verbFlag, "v", false, "Print a message for each created directory.")
	if len(os.Args) < 2 {
		flag.Usage()
	}
	flag.Parse()
	mode := os.FileMode(modeArg)
	var verb *log.Logger
	if verbFlag {
		verb = log.New(os.Stdout, os.Args[0]+": ", 0)
	}
	warn := log.New(os.Stderr, os.Args[0]+": ", 0)
	for _,  path := range flag.Args() {
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
	os.Exit(status)
}
