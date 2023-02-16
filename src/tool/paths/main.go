package paths

import (
	"path/filepath"
	"fmt"
	"flag"
	"bufio"
	"os"
)

var (
)

func handlePath(p string) {
	fin := filepath.FromSlash(p)
	fmt.Println(fin)
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

