package mkdir

import (
	"fmt"
	"os"

	"path/filepath"
	"github.com/di4f/cli/mtool"
)

func Run(flagSet *mtool.Flags) {
	var (
		parentFlag bool
		modeArg    int
	)
	flagSet.BoolVar(&parentFlag, "p", false, "No error if existing, make parent as needed.")
	flagSet.IntVar(&modeArg, "m", 0766, "Set file `mode`.")
	flagSet.Parse()
	args := flagSet.Args()
	if len(args) == 0 {
		flagSet.Usage()
		os.Exit(1)
	}
	mode := os.FileMode(modeArg)
	for _, path := range args {
		var e error
		pth := filepath.FromSlash(path)
		if parentFlag {
			e = os.MkdirAll(pth, mode)
		} else {
			e = os.Mkdir(pth, mode)
		}
		if e != nil {
			fmt.Fprintf(os.Stderr, "%s\n", e)
		}
	}
}
