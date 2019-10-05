package main

import(
	"fmt"
	"os"
	"io"
	"log"
	"flag"
	"cat"
)

func main() {
	flag.Parse()
	status := 0
	c = cat.New(os.Stdout, warn, 512)
	os.Exit(status)
}
