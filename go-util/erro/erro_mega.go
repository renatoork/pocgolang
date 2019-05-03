package erro

import (
	"bytes"
	"fmt"
	"mega/go-util/dbg"
	"path/filepath"
	"runtime"
)

func Trata(err error) bool {
	if err != nil {
		_, arquivo, linha, _ := runtime.Caller(1)
		fmt.Println(fmt.Sprintf("%s:%d - %s", filepath.Base(arquivo), linha, err))
		if dbg.GetDebug() {
			trace := make([]byte, 16384)
			runtime.Stack(trace, true)
			fmt.Printf("%s", string(bytes.Trim(trace, "\x00")))
		}
		return true
	} else {
		return false
	}
}
