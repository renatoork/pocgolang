package auditutil

import (
	"fmt"
)

func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
