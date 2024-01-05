package urlprs

import (
	"os"
	"fmt"
	"net"
	"net/url"
	"vultras.su/core/cli/mtool"
)

var nilStr = "_"

func printNil(s string) {
	if len(s) != 0 {
		fmt.Println(s)
	} else {
		fmt.Println(nilStr)
	}
}
func Run(flags *mtool.Flags) {
	nilStr = "_"

	flags.Parse()
	args := flags.Args()
	if len(args) < 1 { flags.Usage() }
	s := args[0]

	u, err := url.Parse(s)
	if err != nil {
		os.Exit(1)
	}

	printNil(u.Scheme)
	printNil(u.User.Username())
	p, _ := u.User.Password()
	printNil(p)
	if host, port, err := net.SplitHostPort(u.Host) ; err!=nil {
		printNil(u.Host)
		fmt.Println(nilStr)
	} else {
		printNil(host)
		printNil(port)
	}
	printNil(u.Path)
	printNil(u.Fragment)
	printNil(u.RawQuery)

	/*fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])*/
}

