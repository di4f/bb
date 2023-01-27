package path

import (
	"path/filepath"
	"fmt"
)

func Run(args []string) {
	path := filepath.Join(args[1:]...)
	fmt.Print(path)
}

