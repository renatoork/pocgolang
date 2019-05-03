package util

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-oci8"
	"os"
)

var session *sqlx.DB

type Estrutura struct {
	Servico string
}

func checkErrorDb(err error) {
	if err != nil {
		panic(fmt.Sprintf("DB ERROR: %s", err))
	}
}

func GetSession(conn string) *sqlx.DB {
	os.Setenv("NLS_LANG", ".WE8PC850")
	var dns string
	if session == nil {
		var err error
		session, err = sqlx.Open("oci8", conn)
		checkErrorDb(err)
	}
	fmt.Println("Conectado em ", dns)
	return session
}

func CloseSession() {
	session.Close()
	session = nil
	fmt.Println("Conexão fechada.")
}

func geraEntrada(servico string, conn string) {
	db := GetSession(conn)
	defer CloseSession()
	cmd := "Select * From (Select LEVEL LEV, T.TAG, T.PAI, T.TABELA From Table(" + servico + "DEF.F_TABELAS()) T Start With PAI Is Null Connect By Prior TAG = PAI) T ,(Select TAG, NOME, VALOR From Table(" + servico + "DEF.F_FIELDS())) F Where T.TAG = F.TAG  ORDER BY lev"
	est := Estrutura{}
	db.Select(&est, cmd)
}
