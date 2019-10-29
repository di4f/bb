package cat
/* Simple module to get output of a few files
  * and put in in one only. */
import(
	"os"
	"io"
	"log"
	"flag"
)
type Catter struct {
	warn *log.Logger
	out *os.File
	bufSiz int
	buf []byte
}

func New(out *os.File, warn *log.Logger, bufSiz int) *Catter {
	c := new(Catter)
	c.out = out
	c.warn = warn
	c.bufSiz = bufSiz
	c.buf = make([]byte, c.bufSiz)
	return c
}

func (c Catter) Cat(files []*os.File) error {
	for _, file := range files {
		for {
			n, err := file.Read(c.buf)
			if n>0 {
				_, err := c.out.Write(c.buf[:n])
				if err != nil {
					return err
				}
			}
			if err == io.EOF {
				break
			}else if err != nil {
				c.warn.Println(err)
			}
		}
	}
	return nil
}

func (c Catter)OpenArr(filePaths []string) []*os.File {
	files := make([]*os.File, len(filePaths))
	for i, filePath := range filePaths {
		var err error
		files[i], err = os.Open(filePath)
		if err != nil {
			c.warn.Println(err)
		}
	}
	return files
}

func Run(args []string) int {
	flagSet := flag.NewFlagSet(args[0], flag.ExitOnError)
	flagSet.Parse(args[1:])
	status := 0
	warn := log.New(os.Stderr, args[0]+": ", 0)
	c := New(os.Stdout, warn, 512)
	files := c.OpenArr(flagSet.Args())
	err := c.Cat(files)
	if err != nil {
		status = 1
	}
	return status
}