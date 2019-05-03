package dbg

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

func SetDebug(debug bool) {
	os.Setenv("GO-DEBUG", strconv.FormatBool(debug))
}

func GetDebug() bool {
	debugStr := os.Getenv("GO-DEBUG")
	debug, _ := strconv.ParseBool(debugStr)
	return debug
}

func Print(mensagem string, valor interface{}) {
	if GetDebug() {
		fmt.Println(mensagem)
		valorJson, _ := json.MarshalIndent(valor, "", "  ")
		if string(valorJson) != "null" {
			fmt.Println(string(valorJson))
		}
	}
}

func Trace(startTime time.Time) {
	endTime := time.Now()
	fmt.Printf("%.2fs\n", endTime.Sub(startTime).Seconds())
}
