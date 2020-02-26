package sort

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"io"
	"flag"
)

func readLines() []string {

	r := bufio.NewReader(os.Stdin)
	a := make([]string, 0)
	for {
		l, e := r.ReadString('\n')
		if e==io.EOF {
			break
		}
		a = append(a, l)
	}
	return a
}

func Run(args []string) int {
	flagSet := flag.NewFlagSet(args[0], flag.ExitOnError)	
	flagSet.Parse(args[1:])
	status := 0

	lines := readLines()
	sort.Strings(lines)
	for _, l := range lines {
		fmt.Print(l)
	}
	return status
}
