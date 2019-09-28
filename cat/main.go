package main

import(
	"fmt"
	"os"
	"io"
	"log"
	"flag"
)

const(
	bufSiz = 1024
)

func catFile(file *os.File) error {
	buf := make([]byte, bufSiz)
	for {
		n, err := file.Read(buf)
		if n>0 {
			fmt.Printf("%s", buf[:n])
		}
		if err == io.EOF {
			break
		}else if err != nil {
			return err
		}
	}
	return nil
}
func openAndCatFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	catFile(file)

	return nil
}

func main() {
	flag.Parse()
	warn := log.New(os.Stderr, os.Args[0]+": ", 0)
	status := 0
	if flag.NArg()>0 { // Files.
		for _, pathFile := range flag.Args() {
			err := openAndCatFile(pathFile)
			if err != nil {
				warn.Println(err)
			}
		}
	} else { // Stdin.
		err := catFile(os.Stdin)
		if err != nil {
			warn.Println(err)
		}
	}
	os.Exit(status)
}
