package credito

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mega/go-util/dbg"
	"mega/go-util/erro"
	"mega/ms-pedidovenda/tipos"
	"net/http"
	"os"

	_ "github.com/mattn/go-oci8" // utilizado no gorp.in
	"gopkg.in/gorp.v1"
)

var DbMap *gorp.DbMap

type Retorno struct {
	Credito Credito `json:"credito, omitempty"`
	Erros   []Erro  `json:"erros, omitempty"`
}

type Credito struct {
	Situacao string `json:"situacao, omitempty"`
	Mensagem string `json:"mensagem, omitempty"`
}
type Erro struct {
	Mensagem string `json:"erro, omitempty"`
}

var ret Retorno
var doc tipos.PedidoVendaDB

func erroHTTP(res http.ResponseWriter, msg string) {
	e := Erro{msg}
	ret.Erros = append(ret.Erros, e)
	res.WriteHeader(http.StatusExpectationFailed)
}

func InitDb() *gorp.DbMap {
	connString := os.Getenv("API_ANALISECREDITO_ORACLE")
	fmt.Println(connString)
	db, err := sql.Open("oci8", connString)
	if erro.Trata(err) {
		//erroHTTP(res, fmt.Sprintf("Erro ao conectar ao Oracle (%s)", connString))
		return nil
	}
	err = db.Ping()
	if erro.Trata(err) {
		//erroHTTP(res, fmt.Sprintf("Erro ao conectar ao Oracle (%s)", connString))
		return nil
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.OracleDialect{}}

	return dbmap
}

func trataEntrada(res http.ResponseWriter, req *http.Request) []byte {
	// carregar o json do Pedido.
	resp, eReq := ioutil.ReadAll(req.Body)
	if erro.Trata(eReq) {
		msg := fmt.Sprintf("erro ao tratar o parametro de entrada: %s", eReq.Error())
		erroHTTP(res, msg)
		return []byte("")
	}

	return resp

}

func execAnalise(res http.ResponseWriter, req *http.Request) Retorno {
	ret = Retorno{}

	resp := trataEntrada(res, req)

	if carregaDadosDocumento(&doc, res, resp) {
		executaServico(&doc, res)
	}

	return ret

}

func carregaDadosDocumento(doc *tipos.PedidoVendaDB, res http.ResponseWriter, param []byte) bool {
	fmt.Println(string(param))
	if err := json.Unmarshal(param, &doc); erro.Trata(err) {
		msg := fmt.Sprintf("erro ao interpretar o json do Documento: %s", err.Error())
		erroHTTP(res, msg)
		return false
	}
	return true
}

func setRetorno(situacao string, mensagem string) {
	ret.Credito.Situacao = situacao
	ret.Credito.Mensagem = mensagem
}

func executaServico(doc *tipos.PedidoVendaDB, res http.ResponseWriter) {

	os.Setenv("NLS_LANG", ".AL32UTF8")

	dbg.SetDebug(false)

	// insert tabela temp

	trans, errB := DbMap.Begin()
	if erro.Trata(errB) {
		erroHTTP(res, errB.Error())
		setRetorno("B", errB.Error())
		return
	}

	if dbg.GetDebug() {
		DbMap.TraceOn("[ms-cred]", log.New(os.Stdout, ":", log.Lmicroseconds))
		defer DbMap.TraceOff()
	}
	fmt.Println(doc)
	valCredito, errS := trans.SelectFloat("SELECT CCR_RE_CREDITO FROM MGGLO.GLO_CLIENTE_CREDITO C WHERE AGN_IN_CODIGO = :agn", *doc.Cliente)
	if erro.Trata(errS) {
		valCredito = 0
	}
	valPedido := *doc.Valortotal
	if valPedido > valCredito {
		setRetorno("B", "Pedido Bloqueado por falta de crédito.")
	} else {
		setRetorno("A", "Pedido Aprovado.")
	}

	trans.Commit()
}
