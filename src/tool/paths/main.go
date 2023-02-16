package paths

import (
	"fmt"
	"flag"
	"bufio"
	"os"
	"github.com/surdeus/goblin/src/pathx"
)

var (
)

func handlePath(p string) {
	pth := pathx.From(p)
	fmt.Println(pth.Real())
}

func Run(args []string) {
	arg0 := args[0]
	args = args[1:]
	flags := flag.NewFlagSet(arg0, flag.ExitOnError)
	flags.Parse(args)
	args = flags.Args()

	for _, p := range args {
		handlePath(p)
	}

	rd := bufio.NewScanner(os.Stdin)
	for rd.Scan() {
		handlePath(rd.Text())
	}
}

