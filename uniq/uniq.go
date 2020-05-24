/* Yes program implementation. */
package uniq

import(
	"os"
	"fmt"
	"flag"
	"bufio"
	"io"
)

func Run(args []string) int {
	var(
		Uflag bool
	)
	status := 0
	arg0 := args[0]
	args = args[1:]
	flagSet := flag.NewFlagSet(arg0, flag.ExitOnError)
	flagSet.BoolVar(&Uflag, "U", false, "Print every line just one time.")
	flagSet.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s: %s [options] [string]\n", arg0, arg0)
		flagSet.PrintDefaults()
	}
	flagSet.Parse(args)
	args = flagSet.Args()

	r := bufio.NewReader(os.Stdin)
	u := make(map[string]int)
	if Uflag {
		for{
			l, e := r.ReadString('\n')	
			if e==io.EOF {
				break
			}
			_, haskey := u[l]
			if !haskey {
				u[l] = 1
				fmt.Print(l)
			}
		}
	}

	return status
}
