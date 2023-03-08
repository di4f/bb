package ln

import (
	"flag"
	"fmt"
	"os"
)

func Run(args []string) {
	var lflag bool
	flagSet := flag.NewFlagSet(args[0], flag.ExitOnError)
	flagSet.BoolVar(&lflag, "s", false, "make a symbolic link, not a hard one")

	flagSet.Parse(args[1:])
	args = flagSet.Args()

	if len(args) != 2 {
		flagSet.Usage()
		os.Exit(1)
	}

	src := args[0]
	dst := args[1]

	var err error
	if lflag {
		err = os.Symlink(src, dst)
	} else {
		err = os.Link(src, dst)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
