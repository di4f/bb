package date

import(
	"fmt"
	"time"
	"github.com/omnipunk/cli/mtool"
)

func Run(flagSet *mtool.Flags) {
	flagSet.Parse()
	args := flagSet.Args()

	if len(args) > 0 {
		flagSet.Usage()
	}
	
	date := time.Now()
	
	fmt.Println(date)

}
