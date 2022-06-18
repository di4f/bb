package ftest

import (
	"os"
	"io"
	"bufio"
	"fmt"
	"flag"
)

var (
	flags [lastFlag]bool
	flagList []int
	flagMap = map[int] func(os.FileInfo) bool{
		fileFlag : IsFile,
		dirFlag : IsDir,
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

func FileInfo(p string) (os.FileInfo, error) {
	s, err := os.Stat(p)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func IsFile(fi os.FileInfo) bool {
	mode := fi.Mode()
	return mode.IsRegular()
}

func IsDir(fi os.FileInfo) bool {
	return fi.IsDir()
}

func checkFile(p string) bool {
	fi, err := FileInfo(p)
	if err != nil {
		return false
	}

	for _, v := range flagList {
		if !flagMap[v](fi){
			return false
		}
	}

	return true
}

func Run(args []string) {
	arg0 := args[0]
	args = args[1:]
	flagSet := flag.NewFlagSet(arg0, flag.ExitOnError)
	flagSet.Usage = func() {
		flagSet.Usage()
		os.Exit(1)
	}

	flagSet.BoolVar(&flags[fileFlag], "f", false, "is file")
	flagSet.BoolVar(&flags[dirFlag], "d", false, "is directory")

	flagSet.Parse(args)
	args = flagSet.Args()

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
