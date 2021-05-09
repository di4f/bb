package mergelbl

import(
	"os"
	"fmt"
	"bufio"
	"flag"
	"log"
)

func
Run(args []string) int {
	var(
		e error
		buf string
	)
	rsep := '\n'
	wsep := "\n"
	//del := ""
	status := 0
	arg0 := args[0]
	args = args[1:]
	flagSet := flag.NewFlagSet(arg0, flag.ExitOnError)
	flagSet.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s <file1> <file2> .. [fileN]\n", arg0)
		flagSet.PrintDefaults()
	}
	flagSet.Parse(args)
	args = flagSet.Args()

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
				return 0
			}
			s += buf[:len(buf)-1]
		}
		fmt.Printf("%s%s", s, wsep)
	}
	return status
}

