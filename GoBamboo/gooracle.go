package main

import (
	"bytes"
	"fmt"
	"github.com/tgulacsi/goracle/oracle"
	"os"
)

var conn *oracle.Connection
var dsn string

func getConn(connectString string) *oracle.Connection {

	dsn = connectString

	if conn.IsConnected() {
		return conn
	}

	if dsn == "" {
		fmt.Println(os.Stderr, "Problema com a connect string tratada no parâmetro 1.")
		os.Exit(1)
		//return conn
	}

	user, passw, sid := oracle.SplitDSN(dsn)
	var err error
	conn, err = oracle.NewConnection(user, passw, sid, false)
	if err != nil {
		fmt.Println(os.Stderr, fmt.Sprintf("Erro ao criar a conexão em %s: %s", dsn, err))
		os.Exit(1)
		//return conn
	}
	if err = conn.Connect(0, false); err != nil {
		fmt.Println(os.Stderr, fmt.Sprintf("Erro ao conectar em %s: %s", dsn, err))
		os.Exit(1)
	}
	return conn
}

func executarCTPedido(conf Config) TestSuite {
	var buffer bytes.Buffer
	var tc TestCase

	suite := TestSuite{}

	db := getConn(conf.ConnString)
	if db.IsConnected() {

		cur := db.NewCursor()
		defer cur.Close()

		cmd := fmt.Sprintf("begin %s; end;", conf.NomeRotinaExecuta) //mgtst.ven_pck_pedidovendatst.p_executactpedido
		if err := cur.Execute(cmd, nil, nil); err != nil {
			m1 := fmt.Sprintf(`Erro ao executar "%s": %s`, cmd, err)
			buffer.WriteString(m1)
			fmt.Println(os.Stderr, m1)
			os.Exit(1)
		} else {
			cmd = fmt.Sprintf("select grupo, nome, falha, erro from table(%s)", conf.NomeRotinaResultado) //mgtst.ven_pck_pedidovendatst.f_ResultadoCTPedidoPipe
			if err := cur.Execute(cmd, nil, nil); err != nil {
				m2 := fmt.Sprintf(`Erro ao executar "%s": %s`, cmd, err)
				buffer.WriteString(m2)
				fmt.Println(os.Stderr, m2)
				os.Exit(1)
			} else {
				rows, errF := cur.FetchAll()
				if errF != nil {
					m3 := fmt.Sprintf("Erro ao tratar o fetch do resultado: %s", errF)
					buffer.WriteString(m3)
					fmt.Println(os.Stderr, m3)
					os.Exit(1)
				} else {
					for _, row := range rows {
						tc.Name = row[1].(string)
						tc.ClassName = row[0].(string)
						tc.TestFailure = ""
						if row[3].(string) == "S" {
							tc.TestFailure = row[2].(string)
						}
						suite.TestCases = append(suite.TestCases, tc)
						buffer.WriteString(row[1].(string))
					}
				}
			}
		}
	} else {
		buffer.WriteString(fmt.Sprintf("Erro ao conectar em %s", conf.ConnString))
		m4 := fmt.Sprintf("Erro ao conectar em %s", conf.ConnString)
		fmt.Println(os.Stderr, m4)
		os.Exit(1)
	}
	return suite
}
