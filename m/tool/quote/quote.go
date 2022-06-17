package quote
/* Quote quotes string if it contains white space character. */

import(
	"os"
	"io"
	"flag"
	"fmt"
	"unicode"
	"bufio"
)

func HasWhiteSpace(s string) bool {
	for _, r := range s {
		if(unicode.IsSpace(r)){
			return true
		}
	}
	return false
}

func Run(args []string) {	
	arg0 := args[0]
	args = args[1:]
	flagSet := flag.NewFlagSet(arg0, flag.ExitOnError)
	flagSet.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s: %s\n", arg0, arg0)
		flagSet.PrintDefaults()
	}
	flagSet.Parse(args)
	args = flagSet.Args()

	r := bufio.NewReader(os.Stdin)
	for{
		l, e := r.ReadString('\n')	
		if e==io.EOF {
			break
		}
		last := len(l) - 1
		if l[last] == '\n' {
			l = l[:last]	
		}
		if HasWhiteSpace(l) {
			fmt.Printf("'%s'\n", l)
		}else {
			fmt.Println(l)
		}
	}
}
