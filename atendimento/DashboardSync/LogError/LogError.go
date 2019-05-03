package LogError

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

type LogErrorText struct {
	textError string
	fileName  string
}

var panicAttack bool = true
var errorText LogErrorText

func SetTextLogError(text string, fileName string) {

	errorText = LogErrorText{}
	errorText.textError = text
	errorText.fileName = fileName

}

func DontPanicAttack() {

	panicAttack = false

}

func Log(fnc string, err error) {

	if err != nil {
		pc, file, line, _ := runtime.Caller(1)
		fmt.Println("\nFunction: " + fnc + "\nErro: " + err.Error() + "\nLineError: " + strconv.Itoa(line) + "\nFileName: " + file + "\nFuncPc: " + runtime.FuncForPC(pc).Name() + "\n")

		if errorText.fileName != "" {
			recOnTextFile()
		}

		if panicAttack {
			panic(err)
			panicAttack = true
		}
	}

}

func recOnTextFile() {

	fileC, _ := os.Create("Logs/" + errorText.fileName + ".txt")
	defer fileC.Close()

	fileC.WriteString(errorText.textError)

	errorText = LogErrorText{}

}
