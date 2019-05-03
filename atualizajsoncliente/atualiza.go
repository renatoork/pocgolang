package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
)

type Cliente struct {
	Cliente []int64
}

func main() {

	connString := os.Getenv("ORACLE_MS_PEDIDOVENDA") //"mgven/megaven@pc_lopez:1521/orc3"
	db, err := sql.Open("oci8", connString)

	db, err := sqlx.Connect("postgres", "user=foo dbname=bar sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	// carrega clientes struct
	agn := cliente.Cliente{}
	db.Select(&agn, "select distinct cli_in_codigo from mgven.ven_pedidovenda order by 1")
	fmt.println(agn[0], agn[1])
	// le arquivos de diretórios
	// carrega arquivo struct
	// altera cliente
	// salva arquivo

}
