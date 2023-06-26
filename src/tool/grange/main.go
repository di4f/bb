package grange
/* Concatenate files in "stdout". */
import(
	"fmt"
	"strconv"
	"github.com/mojosa-software/gomtool/src/mtool"
)

var(
	flagSet *mtool.Flags
	args []string
	blockSize int
	rangeType string
	rangeTypeMap map[string] func() = map[string] func() {
		"int" : IntRange,
		"inteq" : IntRangeEq,
		"byte" : ByteRange,
		"rune" : RuneRange,
	}
)

func IntRange() {
	IntRangeHndl(func(a, b int) bool {
		if a < b {
			return true
		} else {
			return false
		}
	}, func(a, b int) bool {
		if a > b {
			return true
		} else {
			return false
		}
	})	
}

func IntRangeEq () {
	IntRangeHndl(func(a, b int) bool {
		if a <= b {
			return true
		} else {
			return false
		}
	}, func(a, b int) bool {
		if a >= b {
			return true
		} else {
			return false
		}
	})
}

func IntRangeHndl(fn1, fn2 func(a, b int) bool) {
	var (
		err error
		i, a, b int
	)
	step := 1

	if len(args) < 2 {
		flagSet.Usage()
	}

	if len(args) == 3 {
		step, err = strconv.Atoi(args[2])
		if err != nil {
			panic(err)
		}
		if step == 0 {
			flagSet.Usage()
		}
	}

	a, err = strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}

	b, err = strconv.Atoi(args[1])
	if err != nil {
		panic(err)
	}

	if step > 0 {
		if a>b {
			flagSet.Usage()
		}
		for i=a ; fn1(i, b) ; i += step {
			fmt.Printf("%d%s", i, "\n")
		}
	} else if a > b {
		for i=a ; fn2(i, b) ; i += step {
			fmt.Printf("%d%s", i, "\n")
		}
	} else {
		flagSet.Usage()
	}

}

func ByteRange() {
}

func RuneRange() {
}

func Run(flags *mtool.Flags) {	
	flags.StringVar(&rangeType, "t", "int", "range type")
	flags.Parse()
	args = flagSet.Args()
	flagSet = flags
	
	rangeTypeMap[rangeType]()
}
