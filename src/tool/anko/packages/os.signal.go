package packages

import (
	"os/signal"
	"reflect"

	"github.com/surdeus/goblin/src/tool/anko/env"
)

func init() {
	env.Packages["os/signal"] = map[string]reflect.Value{
		"Notify": reflect.ValueOf(signal.Notify),
		"Stop":   reflect.ValueOf(signal.Stop),
	}
}
