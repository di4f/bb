package whoami

import(
	"os"
	"os/user"
	"flag"
	"fmt"
	"log"
)

func Run(args []string) {
	arg0 := args[0]
	args = args[1:]
	flagSet := flag.NewFlagSet(arg0, flag.ExitOnError)
	flagSet.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s: %s\n", arg0, arg0)
		flagSet.PrintDefaults()
	}
	flagSet.Parse(args)

	if len(flagSet.Args())>0 {
		flagSet.Usage()
		os.Exit(1)
	}
	
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(u.Username)
}
