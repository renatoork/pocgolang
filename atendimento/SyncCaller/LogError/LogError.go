package LogError

import (
	"mega/atendimento/SyncCaller/Types"

	"fmt"
	"os"
	"runtime"
	"strconv"
)

var doLog bool
var lLog Types.LastError

func LogInline(err error) {

	var rTimePc uintptr
	var File string
	var lineError int

	lLog = Types.LastError{}

	rTimePc, File, lineError, _ = runtime.Caller(1)

	if err != nil {
		log := "\n" + runtime.FuncForPC(rTimePc).Name() + "\n\t" + File + "\n\t" + strconv.Itoa(lineError) + "\n\t" + err.Error() + "\n\n"
		lLog.Log = log
		fmt.Println(log)
	}
}

func LogFile(logParam *Types.RecLog) {
	file, _ := os.OpenFile("Logs/"+logParam.FileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
	defer file.Close()
	file.WriteString(logParam.Log)
}

func CheckLog() bool {
	return doLog
}

func SetLog(value bool) {
	doLog = value
}

func GetLastError() Types.LastError {
	return lLog
}
