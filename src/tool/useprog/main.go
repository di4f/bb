package useprog

import (
	"os"
	"os/exec"
	"fmt"
	"flag"
)

func Run(args []string) {
	arg0 := args[0]
	args = args[1:]
	flagSet := flag.NewFlagSet(arg0, flag.ExitOnError)
	flagSet.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s <prog1> [prog2, ...prog3]\n", arg0)
		flagSet.PrintDefaults()
		os.Exit(1)
	}

	flagSet.Parse(args)
	args = flagSet.Args()

	if len(args) == 0 {
		flagSet.Usage()
	}

	for _, v := range args {
		_, err := exec.LookPath(v)
		if err != nil {
			continue
		}
		fmt.Printf("%s", v)
		break
	}
}
