// +build go1.10

package packages

import (
	"reflect"
	"strings"

	"github.com/surdeus/goblin/src/tool/run/env"
)

func stringsGo110() {
	env.PackageTypes["strings"] = map[string]reflect.Type{
		"Builder": reflect.TypeOf(strings.Builder{}),
	}
}
