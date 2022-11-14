package in

import (
	"os"
	"io"
	"bufio"
	"fmt"
	"flag"
)

func Run(args []string) {
	var (
		print bool
		not bool
	)
	arg0 := args[0]
	args = args[1:]
	flagSet := flag.NewFlagSet(arg0, flag.ExitOnError)
	flagSet.BoolVar(&print, "p", false, "print matching lines")
	flagSet.BoolVar(&not, "n", false, "find not matching lines")
	flagSet.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s <str1> [str2, ...str3]\n", arg0)
		flagSet.PrintDefaults()
		os.Exit(1)
	}

	flagSet.Parse(args)
	args = flagSet.Args()

	if len(args) == 0 {
		//flagSet.Usage()
		if !not {
			os.Exit(1)
		}
	}

	mp := make(map[string] int)
	for _, v := range args {
		mp[v] = 0
	}

	status := 1
	r := bufio.NewReader(os.Stdin)

	if not {
		status = 0
	}
	for {
		l, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}

		if len(l) != 0 && l[len(l)-1] == '\n' {
			l = l[:len(l)-1]
		}

		_, ok := mp[l]
		if not {
			if ok {
				status = 1
				ok = false
			} else {
				ok = true
			}
		} else {
			if ok {
				status = 0
			}
		}

		if print && ok {
			fmt.Println(l)
		}
	}
	os.Exit(status)
}
