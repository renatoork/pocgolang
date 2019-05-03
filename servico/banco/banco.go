package banco

import (
	"database/sql"
	_ "github.com/mattn/go-oci8"
	"gopkg.in/gorp.v1"
	"log"
	"os"
)

var Dbmap *gorp.DbMap
var Db *sql.DB

func Conectar() {
	os.Setenv("NLS_LANG", ".AL32UTF8")
	var err error
	connect := os.Getenv("SERVICO_CONNECT")
	if connect == "" {
		panic("Favor informar uma connect string na variável de ambiente SERVICO_CONNECT no formato usuario/senha@maquina/instancia")
	}
	Db, err = sql.Open("oci8", connect)
	if err != nil {
		panic(err)
	}

	Dbmap = &gorp.DbMap{Db: Db, Dialect: gorp.OracleDialect{}}
	Dbmap.TraceOn("teste", log.New(os.Stdout, "gorp", log.Lmicroseconds))
}
