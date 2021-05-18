package basename
/* Print file path without extension. */

import(
	"fmt"
	"flag"
	"os"
	"path"
)

var(
	arg0 string
	args []string
	slash =  '/'
)

func Base(p string) string {
	return path.Base(p)
}

func Run(argv []string) int {
	status := 0
	arg0 = argv[0]
	args = argv[1:]
	flagSet := flag.NewFlagSet(arg0, flag.ExitOnError)
	flagSet.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [files]\n", arg0, arg0)
		flagSet.PrintDefaults()
	}
	flagSet.Parse(args)
	args = flagSet.Args()

	lasti := len(args) - 1
	for i, v := range args {
		fmt.Printf("%s", Base(v))
		if i != lasti {
			fmt.Println("")
		}
	}

	return status
}
