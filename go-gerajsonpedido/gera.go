package main

import (
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-oci8"
	"github.com/tgulacsi/goracle/oracle"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"gopkg.in/gorp.v1"
	"io/ioutil"
	"log"
	"mega/go-util/erro"
	tipos "mega/ms-pedidovenda/tipos"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type arquivo struct {
	CasosTeste CasosTeste `json:"casoteste`
}

type CasosTeste struct {
	CasoTeste   CasoTeste         `json:"casoteste"`
	PedidoVenda tipos.PedidoVenda `json:"pedidovenda"`
}

type CasoTeste struct {
	Grupo     string `json:"grupo"`
	Descricao string `json:"descricao"`
	Rotina    string `json:"rotina"`
	Retorno   string `json:"retorno"`
	Tempo     int    `json:"tempo"`
	Gravar    bool   `json:"gravar"`
	Metodo    string `json:"metodo"`
}

var dbmap *gorp.DbMap

func main() {

	runtime.LockOSThread()

	// ler os casos de testes xml
	connString := os.Getenv("ORACLE_MS_PEDIDOVENDA")
	db, err := sql.Open("oci8", connString)
	if erro.Trata(err) {
		log.Println(fmt.Sprintf("erro ao conectar ao Oracle (%s)", connString))
	}
	defer db.Close()

	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.OracleDialect{}}
	defer dbmap.Db.Close()

	dbmap.AddTableWithName(tipos.PedidoVenda{}, "VEN_PEDIDOVENDA_INT")
	dbmap.AddTableWithName(tipos.Observacao{}, "VEN_OBSERVACAOPEDIDO_INT")
	dbmap.AddTableWithName(tipos.Ocorrencia{}, "VEN_OCORRENCIAPED_INT")
	dbmap.AddTableWithName(tipos.Arquivo{}, "VEN_PEDIDOARQUIVO_INT")
	dbmap.AddTableWithName(tipos.Item{}, "VEN_ITEMPEDIDOVENDA_INT")
	dbmap.AddTableWithName(tipos.ObsItem{}, "VEN_OBSITEMPEDIDO_INT")
	dbmap.AddTableWithName(tipos.PedProgEntrega{}, "VEN_PEDPROGENTREGA_INT")
	dbmap.AddTableWithName(tipos.PedProgEstoque{}, "VEN_PEDPROGESTOQUE_INT")
	dbmap.AddTableWithName(tipos.ParcFinPedido{}, "VEN_PARCFINPEDIDO_INT")

	trans, errB := dbmap.Begin()
	if erro.Trata(errB) {
		log.Println(err.Error())
		return
	}
	var cli [2000]int
	rows1, err := db.Query("select distinct cli_in_codigo from mgven.ven_vw_pedidovenda where pro_in_codigo = 72 or pro_in_codigo = 75 order by 1")
	if err != nil {
		log.Fatal("Erro select AGN: ", err)
		trans.Rollback()
		return
	}
	i := 1
	defer rows1.Close()
	for rows1.Next() {
		err = rows1.Scan(&cli[i])
		i = i + 1
	}

	rows, err := db.Query("select ID, GRUPO, DESCRICAO, ROTINA, RETORNO, TEMPO, GRAVAR, PEDIDO PEDIDO from table(mgtst.PCK_PEDIDOVENDATSTJSON.f_selectcasoteste) x where grupo = 'Performance' and id = 1")

	if err != nil {
		log.Fatal("Erro select: ", err)
		trans.Rollback()
		return
	}
	defer rows.Close()
	var (
		tempo   sql.NullString
		gravar  sql.NullString
		id      sql.NullString
		rotina  sql.NullString
		retorno sql.NullString

		pedido string
	)

	// carregar structs
	i = 1
	cont := 1
	for rows.Next() {
		ct := CasoTeste{}
		err = rows.Scan(&id, &ct.Grupo, &ct.Descricao, &rotina, &retorno, &tempo, &gravar, &pedido)

		ct.Tempo, _ = strconv.Atoi(tempo.String)
		ct.Gravar = gravar.String == "S"
		ct.Rotina = rotina.String
		ct.Retorno = retorno.String

		sr := strings.NewReader(ct.Descricao)
		tr := transform.NewReader(sr, charmap.Windows1252.NewDecoder())
		descricao, _ := ioutil.ReadAll(tr)
		ct.Descricao = string(descricao)
		if erro.Trata(err) {
			log.Println("Lendo xml Pedido: ", err.Error())
			continue
		}

		sr = strings.NewReader(pedido)
		tr = transform.NewReader(sr, charmap.Windows1252.NewDecoder())
		pedtxt, _ := ioutil.ReadAll(tr)
		pedtxt1 := string(pedtxt)

		ped := lerPedidoXml(pedtxt1)
		//ped := lerPedidoJson(id.String, ct)
		ped.Representante = ""

		for a := 1; a <= 2000; a++ {
			ped.Cliente = strconv.Itoa(cli[i])
			i = i + 1
			if i > len(cli) || cli[i] == 0 {
				i = 1
			}
			ct.Metodo = ped.Operacao

			// gerar json
			//log.Println(ped.Emissao)
			cts := CasosTeste{ct, ped}
			c, er := json.MarshalIndent(cts, "", "   ")
			if erro.Trata(er) {
				log.Println("Gerando Json: ", er.Error())
				break
				continue
			}

			desc := strings.Split(ct.Descricao, "**: ")
			var d string
			if len(desc) > 1 {
				d = desc[1]
			} else {
				d = desc[0]
			}
			desc = strings.Split(d, ",")
			desc = strings.Split(desc[0], ".")
			desc = strings.Split(desc[0], "(")
			desc1 := strings.Replace(desc[0], "$", "", -1)
			desc1 = strings.Replace(desc1, "/", "", -1)
			desc1 = strings.Replace(desc1, "\\", "", -1)
			//desc1 := strings.Split(ct.Descricao, ",")
			nomearq := fmt.Sprintf("arquivo\\Wine\\%s_%d_%s.json", ct.Grupo, cont /*id.String*/, strings.Replace(desc1, " ", "_", -1))
			criarArquivo(nomearq, c)
			cont = cont + 1
		}
	}
	trans.Commit()
}

func criarArquivo(nomearq string, conteudo []byte) {
	// criar arquivo
	err := ioutil.WriteFile(nomearq, conteudo, 0644)
	if erro.Trata(err) {
		log.Println(err.Error())
	} else {
		log.Println(fmt.Sprintf("Arquivo %s gerado.", nomearq))
	}
}

func lerPedidoXml(pedido string) tipos.PedidoVenda {
	ped := tipos.PedidoVenda{}
	//log.Println(pedido)
	err := xml.Unmarshal([]byte(pedido), &ped)
	if erro.Trata(err) {
		log.Println("Erro Unmarshal: ", err.Error())
		return tipos.PedidoVenda{}
	}
	//log.Println(ped)
	ped.DataEmissao = time.Time(ped.Emissao)

	return ped
}

var conn *oracle.Connection
var dsn string

func getConn() *oracle.Connection {

	connString := os.Getenv("ORACLE_MS_PEDIDOVENDA")
	dsn = connString

	if conn.IsConnected() {
		return conn
	}

	if !(dsn != "") {
		fmt.Println("Impossível conexão sem o host/sid")
		return conn
	}
	user, passw, sid := oracle.SplitDSN(dsn)
	var err error
	conn, err = oracle.NewConnection(user, passw, sid, false)
	if err != nil {
		fmt.Println(fmt.Sprintf("error creating connection to %s: %s", dsn, err))
	}
	if err = conn.Connect(0, false); err != nil {
		fmt.Println(fmt.Sprintf("error connecting: %s", err))
	}
	return conn
}

func lerPedidoJson(id string, ct CasoTeste) tipos.PedidoVenda {
	ped := tipos.PedidoVenda{}

	dbx := getConn()
	if dbx.IsConnected() {

		cur := dbx.NewCursor()
		defer cur.Close()

		cmd := `MGTST.ven_pck_pedidovendatst.f_ExecutaCTPedidoJson`
		var listaArgumento []interface{}
		var nomeArgumento map[string]interface{}
		listaArgumento = append(listaArgumento, ct.Grupo)
		listaArgumento = append(listaArgumento, id)

		if ret, err := cur.CallFunc(cmd, *oracle.ClobVarType, listaArgumento, nomeArgumento); err != nil {
			fmt.Println(fmt.Sprintf(`Erro ao executar "%s": %s`, cmd, err))
		} else {

			// close the underlying cursor, see whether it invalidates the LOB handle
			cur.Close()
			//var ret1 *oracle.ExternalLobVar
			log.Println("ret= ", ret)
			clob, ass := ret.(*oracle.ExternalLobVar)
			log.Println("clob= ", clob, ass, err)
			got, err := clob.ReadAll()
			log.Println("got= ", got, err)
		}
	}
	return ped
}

func GetRowJSon(row *sqlx.Row) string {

	cont := 0
	res := make([]interface{}, 1000)

	columns, _ := row.Columns()
	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for i, _ := range columns {
		valuePtrs[i] = &values[i]
	}
	row.Scan(valuePtrs...)
	lin := make([]interface{}, count)
	for i, _ := range columns {
		var v interface{}
		val := values[i]
		b, ok := val.([]byte)
		if ok {
			v = string(b)
		} else {
			v = val
		}
		lin[i] = v
	}
	res[cont] = lin
	cont = cont + 1

	ret := make([]interface{}, cont)
	for i, v := range res {
		if i >= cont {
			break
		}
		ret[i] = v
	}

	b, err := json.Marshal(ret)
	if err != nil {
		fmt.Println("error:", err)
	}
	return string(b)
}
