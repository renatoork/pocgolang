package main

import (
	"bytes"
	"fmt"
	"github.com/tgulacsi/goracle/oracle"
	"os"
)

type Config struct {
	ConnString string `json: "connectstring"`
	Servico    string `json: "servico"`
	Tipo       string `json: "tipo"`
}

func main() {
	c := &Config{
		ConnString: "mgven/megaven@cacador",
		Servico:    "mgven.VEN_PCK_PEDIDOVENDADEF",
		Tipo:       "DEF",
	}

	jpdv := getJson(c)
	fmt.Println(jpdv)

}

var conn *oracle.Connection
var buffer bytes.Buffer

func getConn(connectString string) *oracle.Connection {

	dsn := connectString

	if conn.IsConnected() {
		return conn
	}

	if dsn == "" {
		fmt.Println(os.Stderr, "Dados de conexão não definidos.")
		os.Exit(1)
	}

	user, passw, sid := oracle.SplitDSN(dsn)
	var err error
	conn, err = oracle.NewConnection(user, passw, sid, false)
	if err != nil {
		fmt.Println(os.Stderr, fmt.Sprintf("Erro ao criar a conexão em %s: %s", dsn, err))
		os.Exit(1)
	}
	if err = conn.Connect(0, false); err != nil {
		fmt.Println(os.Stderr, fmt.Sprintf("Erro ao conectar em %s: %s", dsn, err))
		os.Exit(1)
	}
	return conn
}

func getJson(conf *Config) string {

	var x map[string]interface{}
	x = make(map[string]interface{}, 20)
	db := getConn(conf.ConnString)
	if db.IsConnected() {

		cur := db.NewCursor()
		defer cur.Close()

		cmd := `select level, tag, pai, owner, tabela ` +
			`  from table(mgven.Ven_Pck_Pedidovendadef.F_TABELAS) ` +
			` start with pai is null ` +
			` connect by prior tag=pai `

		if err := cur.Execute(cmd, nil, nil); err != nil {
			fmt.Println(fmt.Sprintf(`Erro ao executar "%s": %s`, cmd, err))
			os.Exit(1)
		} else {
			rows, errF := cur.FetchAll()
			if errF != nil {
				fmt.Println(fmt.Sprintf("Erro ao tratar o fetch do resultado: %s", errF))
				os.Exit(1)
			} else {
				var (
					tabela string
					owner  string
					tag    string
					tagold string
				)
				for _, row := range rows {
					tabela = row[4].(string)
					owner = row[3].(string)
					tagold = tag
					if row[2] != nil {
						tag = row[2].(string)
					}
					//buffer.WriteString(fmt.Sprintf(`{"%s": `, tag))
					cmdtab := `select t.COLUMN_NAME , t.DATA_TYPE ` +
						`  from sys.All_Tab_Columns t, table(mgven.ven_pck_pedidovendadef.F_FIELDS) d ` +
						` where t.OWNER = :1 ` +
						`   and t.TABLE_NAME = :2 ` +
						`   and t.COLUMN_NAME = d.NOME ` +
						`   and d.tag = :3 ` +
						` order by t.COLUMN_ID `
					if err := cur.Execute(cmdtab, []interface{}{owner, tabela, tag}, nil); err != nil {
						fmt.Println(fmt.Sprintf(`Erro ao executar "%s": %s`, cmd, err))
						os.Exit(1)
					} else {
						rows, errF := cur.FetchAll()
						if errF != nil {
							fmt.Println(fmt.Sprintf("Erro ao tratar o fetch do resultado: %s", errF))
							os.Exit(1)
						} else {
							x[tag] = rows
							if tagold != "" && tagold != tag {
								x[tagold] = fmt.Sprintf("%s%s", x[tagold], x[tag])
							}
						}
					}
				}
				buffer.WriteString(`}`)
			}
		}

	} else {
		buffer.WriteString(fmt.Sprintf("Erro ao conectar em %s", conf.ConnString))
		m4 := fmt.Sprintf("Erro ao conectar em %s", conf.ConnString)
		fmt.Println(os.Stderr, m4)
		os.Exit(1)
	}
	fmt.Println(x)
	return buffer.String()
}
