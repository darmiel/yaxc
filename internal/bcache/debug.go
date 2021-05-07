package bcache

import (
	"fmt"
	"github.com/darmiel/yaxc/internal/common"
	"github.com/muesli/termenv"
)

const debugEnabled = false

func printDebugSet(key string, value interface{}) {
	if !debugEnabled {
		return
	}
	if b, o := value.(string); o {
		fmt.Println(common.StyleCache(),
			termenv.String("<-").Foreground(common.Profile().Color("#DBAB79")),
			"Set",
			termenv.String(key).Foreground(common.Profile().Color("#A8CC8C")),
			termenv.String("=").Foreground(common.Profile().Color("#DBAB79")),
			common.PrettyLimit(string(b), 48))
	}
}
func printDebugGet(key string, value interface{}) {
	if !debugEnabled {
		return
	}
	if b, o := value.(string); o {
		fmt.Println(common.StyleCache(),
			termenv.String("->").Foreground(common.Profile().Color("#66C2CD")),
			"Get",
			termenv.String(key).Foreground(common.Profile().Color("#A8CC8C")),
			termenv.String("=").Foreground(common.Profile().Color("#DBAB79")),
			common.PrettyLimit(string(b), 48))
	}
}
func printDebugJanitorStart() {
	if !debugEnabled {
		return
	}
	fmt.Println(common.StyleCache(),
		termenv.String("JANITOR").Foreground(common.Profile().Color("#A8CC8C")),
		"Starting ...")
}
func printDebugJanitorDelete(k string) {
	if !debugEnabled {
		return
	}
	fmt.Println(common.StyleCache(),
		termenv.String("JANITOR").Foreground(common.Profile().Color("#A8CC8C")),
		"Deleting", termenv.String(k).Foreground(common.Profile().Color("#A8CC8C")))
}
