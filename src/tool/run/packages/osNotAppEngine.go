// +build !appengine

package packages

import (
	"os"
	"reflect"

	"github.com/surdeus/goblin/src/tool/run/env"
)

func osNotAppEngine() {
	env.Packages["os"]["Getppid"] = reflect.ValueOf(os.Getppid)
}
