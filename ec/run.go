package ec

import(
	"fmt"
	"io"
	"os"
	"bufio"
	"log"
)

const(
	esc = '\\'
)

var(
	CharMap = map[string]string{
		"n":"\n",
		"b":"\b",
		"r":"\r",
	}
)

func
handleEscChar(rd *bufio.Reader) error {
	r, _, e := rd.ReadRune()
	if e != nil {
		return e
	}
	fmt.Print(CharMap[string(r)])
	return nil
}

func
Run(args []string) int {
	rd := bufio.NewReader(os.Stdin)
	for {
		r, _, e := rd.ReadRune()
		if e == io.EOF {
			break;
		} else if e != nil {
			log.Fatal(e)
		}

		
		if r == esc {
			handleEscChar(rd)
		} else {
			fmt.Print(string(r))
		}
	}
	return 0
}
