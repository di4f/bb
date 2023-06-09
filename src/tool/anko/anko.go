// +build !appengine

package anko

import (
	"bufio"
	//"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/surdeus/goblin/src/tool/anko/core"
	"github.com/surdeus/goblin/src/tool/anko/env"
	_ "github.com/surdeus/goblin/src/tool/anko/packages"
	"github.com/surdeus/goblin/src/tool/anko/parser"
	"github.com/surdeus/goblin/src/tool/anko/vm"
	"github.com/surdeus/gomtool/src/mtool"
	"os/exec"
)

type Cmd struct {
	input io.ReadCloser
	*exec.Cmd
	output io.WriteCloser
}

const version = "0.1.8"

var (
	flagExecute string
	file        string
	args        []string
	e           *env.Env
	flag *mtool.Flags
)

func Run(flagSet *mtool.Flags) {
	var exitCode int
	
	flag = flagSet

	parseFlags()
	setupEnv()
	exitCode = runNonInteractive()
	//if flagExecute != "" || flag.NArg() > 0 {
	//} 

	os.Exit(exitCode)
}

func parseFlags() {
	flagVersion := flag.Bool("v", false, "prints out the version and then exits")
	flag.StringVar(&flagExecute, "e", "", "execute the Anko code")
	flag.Parse()

	if *flagVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	if flagExecute != "" || flag.NArg() < 1 {
		args = flag.Args()
		return
	}

	file = flag.Arg(0)
	args = flag.Args()[1:]
}

func setupEnv() {
	e = env.NewEnv()
	cmd := func(args ...string) *Cmd {
		if len(args) < 1 {
			panic("too few arguments")
		}
		
		cmd := exec.Command(args[0], args[1:]...)
		
		return &Cmd{
			Cmd: cmd,
			input: os.Stdin,
			output: os.Stdout,
		}
	}
	e.Define("args", args)
	e.Define("cmd", cmd)
	e.Define("rcmd", func(args ...string) bool {
		rcmd := cmd(args...)
		
		input := rcmd.input
		output := rcmd.output
		
		stdin, err := rcmd.StdinPipe()
		if err != nil {
			panic(err)
		}
		
		stdout, err := rcmd.StdoutPipe()
		if err != nil {
			panic(err)
		}
		
		err = rcmd.Start()
		if err != nil {
			panic(err)
		}
		go func(){
			_, err = io.Copy(stdin, input)
			if err != nil {
				panic(err)
			}
		}()
		go func() {
			_, err = io.Copy(output, stdout)
			if err != nil {
				panic(err)
			}
		}()
	
		err = rcmd.Wait()
		if err != nil {
			fmt.Println(err)
			return false
		}
		
		return true
	})
	core.Import(e)
}

func runNonInteractive() int {
	var source string
	if flagExecute != "" {
		source = flagExecute
	} else {
		var (
			sourceBytes []byte
			err error
		)
		if file == "#stdin" {
			sourceBytes, err = io.ReadAll(os.Stdin)
		} else {
			sourceBytes, err = ioutil.ReadFile(file)
		}
		if err != nil {
			fmt.Println("ReadFile error:", err)
			return 2
		}
		source = string(sourceBytes)
	}

	_, err := vm.Execute(e, nil, source)
	if err != nil {
		fmt.Println("Execute error:", err)
		return 4
	}

	return 0
}

func runInteractive() int {
	var following bool
	var source string
	scanner := bufio.NewScanner(os.Stdin)

	parser.EnableErrorVerbose()

	for {
		if following {
			source += "\n"
			fmt.Print("  ")
		} else {
			fmt.Print("> ")
		}

		if !scanner.Scan() {
			break
		}
		source += scanner.Text()
		if source == "" {
			continue
		}
		if source == "quit()" {
			break
		}

		stmts, err := parser.ParseSrc(source)

		if e, ok := err.(*parser.Error); ok {
			es := e.Error()
			if strings.HasPrefix(es, "syntax error: unexpected") {
				if strings.HasPrefix(es, "syntax error: unexpected $end,") {
					following = true
					continue
				}
			} else {
				if e.Pos.Column == len(source) && !e.Fatal {
					fmt.Fprintln(os.Stderr, e)
					following = true
					continue
				}
				if e.Error() == "unexpected EOF" {
					following = true
					continue
				}
			}
		}

		following = false
		source = ""
		var v interface{}

		if err == nil {
			v, err = vm.Run(e, nil, stmts)
		}
		if err != nil {
			if e, ok := err.(*vm.Error); ok {
				fmt.Fprintf(os.Stderr, "%d:%d %s\n", e.Pos.Line, e.Pos.Column, err)
			} else if e, ok := err.(*parser.Error); ok {
				fmt.Fprintf(os.Stderr, "%d:%d %s\n", e.Pos.Line, e.Pos.Column, err)
			} else {
				fmt.Fprintln(os.Stderr, err)
			}
			continue
		}

		fmt.Printf("%#v\n", v)
	}

	if err := scanner.Err(); err != nil {
		if err != io.EOF {
			fmt.Fprintln(os.Stderr, "ReadString error:", err)
			return 12
		}
	}

	return 0
}
