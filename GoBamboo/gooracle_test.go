package main

import (
	"fmt"
	"testing"
)

func getConnection(t *testing.T) {
	ConnectString = "MGTST/MEGATST@PC_LOPEZ/ORC3"

	conn := getConn()

	if conn.IsConnected() {
		fmt.Println("ok.")
	} else {
		fmt.Println("não.")
	}
}
