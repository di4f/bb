package tac 
/* Concatenate files in "stdout" reversed. */
import(
	"os"
	"io"
	"fmt"
	"bufio"
	"github.com/mojosa-software/gomtool/src/mtool"
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
	r := bufio.NewReader(f)
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

func Run(flagSet *mtool.Flags) {	
	flagSet.Parse()
	args := flagSet.Args()
	if len(args)>0 {
		for _, p := range args {
			e := tac(p)
			if e != nil {
				fmt.Fprintf(os.Stderr, "%s.\n", e)
			}
		}
	} else {
		ftac(os.Stdin)
	}
}
