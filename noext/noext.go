package noext
/* Print file path without extension. */

import(
	"fmt"
	"flag"
	"os"
)

var(
	arg0 string
	args []string
	dot = '.'
	slash =  '.'
)

func NoExt(p string) string {
	l := len(p)-1
	i := l
	for ; rune(p[i])!=dot ; i -= 1 {
		if rune(p[i]) == slash || i==0 {
			return p
		}
	}
	return p[:i]
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
	fmt.Printf("%s", NoExt(args[0]))
	return status
}
