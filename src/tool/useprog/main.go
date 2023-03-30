package useprog

import (
	"os"
	"os/exec"
	"fmt"
	"github.com/surdeus/gomtool/src/mtool"
	"path/filepath"
)

func Run(flagSet *mtool.Flags) {
	var nVar bool
	flagSet.BoolVar(&nVar, "n", false, "print only name, not path")
	flagSet.Parse()
	args := flagSet.Args()

	if len(args) < 1 {
		flagSet.Usage()
	}

	for _, v := range args {
		pth, err := exec.LookPath(v)
		if err != nil {
			continue
		}
		var toPrint string
		if nVar {
			toPrint = v
		} else {
			toPrint = filepath.ToSlash(pth)
		}
		fmt.Printf("%s", toPrint)
		os.Exit(0)
	}
	os.Exit(1)
}
