package whoami

import(
	"os"
	"os/user"
	"fmt"
	"log"
	"github.com/vultras/cli/mtool"
)

func Run(flagSet *mtool.Flags) {
	flagSet.Parse()

	if len(flagSet.Args())>0 {
		flagSet.Usage()
		os.Exit(1)
	}
	
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(u.Username)
}
