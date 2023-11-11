package wc

import (
	"os"
	"io"
	"bufio"
	"fmt"
	"unicode"
	"strconv"
	"github.com/omnipunk/cli/mtool"
)

var (
	vals [lastFlag]int
	flags [lastFlag]bool
	flagList []int
)

const (
	wordsFlag = iota
	runesFlag
	charsFlag
	linesFlag
	lastFlag
)

func rdRune(rd *bufio.Reader) (rune, int) {
	r, siz, err := rd.ReadRune()
	if err == io.EOF {
		finish()
	} else if err != nil {
		panic(err)
	}
	return r, siz
}

func finish() {
	str := ""
	for i, v := range flagList {
		str += strconv.Itoa(vals[v])
		if i != ( len(flagList)-1 ) {
			str += "\t"
		}
	}
	fmt.Print(str)
	os.Exit(0)
}

func Run(flagSet *mtool.Flags) {
	var (
		r rune
		siz int
	)
	flagSet.BoolVar(&flags[charsFlag], "c", false, "print amount of chars(bytes)")
	flagSet.BoolVar(&flags[runesFlag], "r", false, "print amount of runes in UTF stream")
	flagSet.BoolVar(&flags[wordsFlag], "w", false, "print amount of words")
	flagSet.BoolVar(&flags[linesFlag], "l", false, "print amount of lines")

	flagSet.Parse()

	for i, v := range flags {
		if v {
			break;
		} else if i == (len(flags) - 1){
			flagSet.Usage()
			os.Exit(1)
		}
	}

	for i, v := range flags {
		if v {
			flagList = append(flagList, i)
		}
	}

	rd := bufio.NewReader(os.Stdin)
	inWord := false
	for {
		r, siz = rdRune(rd)
		if inWord {
			if unicode.IsSpace(r) {
				inWord = false
			}
		} else {
			if !unicode.IsSpace(r) {
				vals[wordsFlag]++
				inWord = true
			}
		}
		vals[runesFlag]++
		vals[charsFlag] += siz
		if r == rune('\n') {
			vals[linesFlag]++
		}
	}
	finish()
}
