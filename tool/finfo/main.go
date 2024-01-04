package finfo

import (
	"os"
	"io"
	"bufio"
	"fmt"
	"github.com/vultras/cli/mtool"
	"path/filepath"
)

var (
	flags [lastFlag] bool
	flagVals []Flag
)

type Flag uint8

const (
	SizeFlag Flag = 1 << iota
	ModeFlag
	ModTimeFlag
	LastFlag
)

func checkFile(p string) bool {
	p = filepath.FromSlash(p)
	for _, v := range flagList {
		if !flagMap[v](p){
			return false
		}
	}

	return true
}

func Fileinfo(flags []Flag, p string) string {
	p = filepath.FromSlash(p)
	ret := p
	for _, flag := range flags {
		ret 
	}
}

func Run(flagSet *mtool.Flags) {
	flagSet.BoolVar(
		&flags[sizeFlag],
		false,
		"print size of files",
	)

	flagSet.BoolVar(
		&flags[modeFlag],
		false,
		"print mode of files",
	)
	flagSet.BoolVar(
		&flags[modTimeFlag],
		false,
		"print last modification time in standard Golang format",
	)
	flagSet.Parse()
	args := flagSet.Args()

	for _, p := range args {
		checkFile(p)
	}

	rd := bufio.NewReader(os.Stdin)
	for {
		s, err := rd.ReadString('\n')
		if err == io.EOF {
			os.Exit(0)
		} else if err != nil {
			panic(err)
		}

		if s[len(s)-1] == '\n' {
			s = s[:len(s)-1]
		}

		if checkFile(s) {
			fmt.Println(s)
		}
	}
}
