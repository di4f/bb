package mergelbl

import(
	"os"
	"fmt"
	"bufio"
	"log"
	"github.com/surdeus/gomtool/src/mtool"
)

func Run(flagSet *mtool.Flags) {
	var(
		e error
		buf string
	)
	rsep := '\n'
	wsep := "\n"
	//del := ""
	flagSet.Parse()
	args := flagSet.Args()

	files := make([]*os.File, len(args))
	for i, v := range args {
		files[i], e = os.Open(v)
		if e!=nil {
			log.Fatal(e)
		}
	}

	rds := make([]*bufio.Reader, len(files))
	for i, v := range files {
		rds[i] = bufio.NewReader(v)
	}

	for{
		s := ""
		for _, r := range rds{
			buf, e = r.ReadString(byte(rsep))
			if e!=nil {
				os.Exit(1)
			}
			s += buf[:len(buf)-1]
		}
		fmt.Printf("%s%s", s, wsep)
	}
}

