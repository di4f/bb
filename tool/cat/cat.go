package cat
/* Concatenate files in "stdout". */
import(
	"os"
	"io"
	"fmt"
	"github.com/reklesio/mtool"
)

var(
	blockSize int
)

func Cat(p string) error {
	f, e := os.Open(p)
	if e != nil {
		return e
	}
	defer f.Close()
	fcat(f)
	return nil
}

func fcat(f *os.File) error {
	b := make([]byte, blockSize)
	for {
		n, e := f.Read(b)
		if n>0 {
			fmt.Print(string(b[:n]))
		}
		if e == io.EOF {
			break
		}else if e != nil {
			return e
		}
	}
	return nil
}

func Run(flags *mtool.Flags) {	
	flags.IntVar(&blockSize, "bs", 512, "block size")
	flags.Parse()
	args := flags.Args()
	if len(args)>0 {
		for _, p := range args {
			e := Cat(p)
			if e != nil {
				fmt.Fprintf(os.Stderr, "%s.\n", e)
			}
		}
	} else {
		fcat(os.Stdin)
	}
}
