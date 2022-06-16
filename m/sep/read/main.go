package read
/* Plan9 read or something to read into variables. */
import(
	"os"
	"bufio"
	"flag"
	"fmt"
)

var(
	nLines int
)

func Run(args []string) {	
	arg0 := args[0]
	args = args[1:]
	flagSet := flag.NewFlagSet(arg0, flag.ExitOnError)
	flagSet.IntVar(&nLines, "n", 1, "amount of lines")
	flagSet.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [options] [files]\n", arg0)
		flagSet.PrintDefaults()
		os.Exit(1)
	}
	flagSet.Parse(args)
	args = flagSet.Args()

	if nLines <= 0 {
		flagSet.Usage()
	}

	rd := bufio.NewReader(os.Stdin)
	for nLines != 0 {
		line, err := rd.ReadString('\n')
		if err != nil {
			os.Exit(0)
		}
		fmt.Print(line)
		nLines -= 1
	}
}
