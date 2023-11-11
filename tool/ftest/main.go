package ftest

import (
	"os"
	"io"
	"bufio"
	"fmt"
	"github.com/omnipunk/cli/mtool"
	"path/filepath"
)

var (
	flags [lastFlag]bool
	flagList []int
	flagMap = map[int] func(string) bool{
		fileFlag : IsFile,
		dirFlag : IsDir,
		writFlag : IsWritable,
		readFlag : IsReadable,
	}
)

const(
	fileFlag = iota
	readFlag
	writFlag
	execFlag
	dirFlag
	sizeFlag
	lastFlag
)

func IsFile(p string) bool {
	st, err := os.Stat(p)
	if err != nil {
		return false
	}
	mode := st.Mode()
	return mode.IsRegular()
}

func IsDir(p string) bool {
	st, err := os.Stat(p)
	if err != nil {
		return false
	}
	return st.IsDir()
}

func IsWritable(p string) bool {
	f, err := os.OpenFile(p, os.O_WRONLY, 0)
	if err != nil {
		return false
	}
	defer f.Close()

	return true
}

func IsReadable(p string) bool {
	f, err := os.OpenFile(p, os.O_RDONLY, 0)
	if err != nil {
		return false
	}
	defer f.Close()

	return true
}

func checkFile(p string) bool {
	p = filepath.FromSlash(p)
	for _, v := range flagList {
		if !flagMap[v](p){
			return false
		}
	}

	return true
}

func Run(flagSet *mtool.Flags) {
	flagSet.BoolVar(&flags[fileFlag], "f", false, "is file")
	flagSet.BoolVar(&flags[dirFlag], "d", false, "is directory")
	flagSet.BoolVar(&flags[writFlag], "w", false, "is writable")
	flagSet.BoolVar(&flags[readFlag], "r", false, "is readable")

	flagSet.Parse()
	args := flagSet.Args()

	if len(args) != 0 {
		flagSet.Usage()
	}

	for i, v := range flags {
		if v {
			break;
		} else if i == (len(flags) - 1) {
			flagSet.Usage()
		}
	}

	for i, v := range flags {
		if v {
			flagList = append(flagList, i)
		}
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
