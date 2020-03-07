package tac 
/* Concatenate files in "stdout" reversed. */
import(
	"os"
	"io"
	"flag"
	"fmt"
	"bufio"
)

func reverse(a []string) chan string {
	ret := make(chan string)
	go func() {
		l := len(a)
		for i, _ := range a {
			ret <- a[l-1-i]
		}
		close(ret)
	}()
	return ret
}

func tac(p string) error {
	f, e := os.Open(p)
	if e != nil {
		return e
	}
	defer f.Close()
	ftac(f)
	return nil
}


func ftac(f *os.File) error {
	r := bufio.NewReader(os.Stdin)
	var lines []string
	for {
		line, e := r.ReadString('\n')
		if e == io.EOF {
			break
		}
		lines  = append(lines, line)
	}
	for l := range reverse(lines) {
		fmt.Print(l)
	}
	return nil
}

func Run(args []string) int {	
	status := 0
	arg0 := args[0]
	args = args[1:]
	flagSet := flag.NewFlagSet(arg0, flag.ExitOnError)
	flagSet.Parse(args)
	args = flagSet.Args()
	if len(args)>0 {
		for _, p := range args {
			e := tac(p)
			if e != nil {
				fmt.Fprintf(os.Stderr, "%s: %s.\n", arg0, e)
				status = 1
			}
		}
	} else {
		ftac(os.Stdin)
	}
	return status
}
