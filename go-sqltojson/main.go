package main

import (
	"database/sql"
	"fmt"
	"github.com/elgs/gosqljson"
	_ "github.com/mattn/go-oci8"
	"os"
	"os/exec"
)

func main() {
	db, err := sql.Open("oci8", os.Args[1]) //"mgven:megaven@pc_renato:1521/orcl"
	defer db.Close()
	if err != nil {
		fmt.Println("sql.Open:", err)
	}
	for i, a := range os.Args {
		if i > 1 {
			_, err1 := db.Exec(fmt.Sprintf("insert into %s (REG_ST_ID) values ('1')", a))
			if err1 != nil {
				fmt.Println("sql.insert: ", err1.Error())
			}

			theCase := "lower" // "lower" default, "upper", camel
			m, _ := gosqljson.QueryDbToMapJson(db, theCase, fmt.Sprintf("SELECT * from %s", a))
			j := fmt.Sprintf(`{"%s": %s}`, a, m)
			nome := a + ".json"
			geraArq(nome, j)
			geraStruct(nome)
		}
	}
}

func geraArq(nome string, js string) {

	filename := nome
	f, _ := os.Create(filename)
	f.WriteString(js)

}

func geraStruct(nome string) {
	cmd := exec.Command("JSONGen", nome)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err := cmd.Wait()
	if err != nil {
		fmt.Println(err.Error())
	}

}
