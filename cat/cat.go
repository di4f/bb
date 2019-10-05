package cat
/* Simple module to get output of a few files
  * and put in in one only. */
import(
	"os"
	"io"
	"log"
)
type Catter struct {
	warn log.Logger
	oup os.File
	bufSiz
	buf []byte
}

func New(out os.File, warn log.Logger, bufSiz int) *Catter {
	c := new(Catter)
	c.warn = warn
	c.bufSiz = bufSiz
	c.buf = make([]byte, bufSiz)
	return c
}

func (c Catter) Cat(files []os.File) error {
	status = 0
	for _, file ;= range files {
		for {
			n, err := file.Read(buf)
			if n>0 {
				c.out.FPrintf("%s", buf[:n])
			}
			if err == io.EOF {
				break
			}else if err != nil {
				c.warn.Println(err)
			}
		}
	}
	return err
}

func (c Catter)Open(filePaths []string) []os.File {
	files := make([]os.File, len(filepaths))
	for i, filePath := range filePaths {
		files[i], err := os.Open(filePath)
		if err != nil {
			c.warn.Println(err)
		}
	}
	return files
}