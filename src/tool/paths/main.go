package paths

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	//"github.com/mojosa-software/goblin/src/pathx"
	"path/filepath"
	"github.com/mojosa-software/gomtool/src/mtool"
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
		p = filepath.ToSlash(p)
	}
	if handler != nil {
		p = handler(p)
	}

	var toPrint string
	if r {
		toPrint = filepath.FromSlash(p)
	} else {
		toPrint = p
	}

	if ec {
		toPrint = strings.ReplaceAll(toPrint, "\\", "\\\\")
	}

	fmt.Println(toPrint)
}

func Run(flags *mtool.Flags) {
	flags.StringVar(&part, "p", "all", "part of path you want to print")
	flags.BoolVar(&r, "r", false, "print real OS dependent paths")
	flags.BoolVar(&fromReal, "fr", false, "take input paths as real ones")
	flags.BoolVar(&ec, "ec", false, "escape characters (mostly for '\\' char in Git bash")

	flags.Parse()
	args := flags.Args()
	
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
