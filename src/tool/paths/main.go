package paths

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/surdeus/goblin/src/pathx"
)

var (
	part     string
	handlers = map[string]func(string) string{
		"base": path.Base,
		"ext":  path.Ext,
		"dir":  path.Dir,
		"all":  func(v string) string { return v },
	}
	handler   func(string) string
	r         bool
	fromReal  bool
	ec        bool
	noPartErr = errors.New("no such part")
)

func handlePath(p string) {
	if fromReal {
		p = pathx.FromReal(p).String()
	}
	if handler != nil {
		p = handler(p)
	}

	var toPrint string
	if r {
		toPrint = pathx.From(p).Real()
	} else {
		toPrint = p
	}

	if ec {
		toPrint = strings.ReplaceAll(toPrint, "\\", "\\\\")
	}

	fmt.Println(toPrint)
}

func Run(args []string) {
	arg0 := args[0]
	args = args[1:]
	flags := flag.NewFlagSet(arg0, flag.ExitOnError)
	flags.StringVar(&part, "p", "all", "part of path you want to print")
	flags.BoolVar(&r, "r", false, "print real OS dependent paths")
	flags.BoolVar(&fromReal, "fr", false, "take input paths as real ones")
	flags.BoolVar(&ec, "ec", false, "escape characters (mostly for '\\' char in Git bash")

	flags.Parse(args)
	args = flags.Args()
	lhandler, ok := handlers[part]
	if !ok {
		log.Fatal(noPartErr)
	}
	handler = lhandler

	for _, p := range args {
		handlePath(p)
	}

	rd := bufio.NewScanner(os.Stdin)
	for rd.Scan() {
		handlePath(rd.Text())
	}
}
