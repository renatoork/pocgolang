package pedidovenda

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mega/customizacaoservico"
	"mega/go-comparastruct/compara"
	"mega/go-util/dbg"
	"mega/go-util/erro"
	"mega/ms-consul/consul"
	"mega/ms-pedidovenda/tipos"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	_ "github.com/mattn/go-oci8"
	_ "gopkg.in/goracle.v1"
	"gopkg.in/gorp.v1"
	"gopkg.in/rana/ora.v3"

	_ "gopkg.in/goracle.v1"
	_ "gopkg.in/goracle.v1/oracle"
)

var DbMap *gorp.DbMap
var DbMapGO *gorp.DbMap

var SrvCredito *consul.Servico
var ConsulCredito bool

var operacao string
var pdvuk tipos.PedidoUk

type Retorno struct {
	Pedido Pedido  `json:"pedidovenda,omitempty"`
	Erros  []Falha `json:"erros,omitempty"`
}

type Pedido struct {
	Mensagem string              `json:"mensagem,omitempty"`
	Pedido   tipos.PedidoVendaDB `json:"pedidovenda,omitempty"`
}

type Falha struct {
	Mensagem string `json:"erro,omitempty" description:"Descricao do Erro"`
}

var ret Retorno
var retOriginal Retorno

type RetornoCredito struct {
	Credito Credito `json:"credito, omitempty"`
	Erros   []Falha `json:"erros,omitempty"`
}

type Credito struct {
	Situacao string `json:"situacao, omitempty"`
	Mensagem string `json:"mensagem, omitempty"`
}

func erroHTTP(msg string) {
	e := Falha{msg}
	fmt.Println(e)
	ret.Erros = append(ret.Erros, e)
	response.WriteHeader(http.StatusInternalServerError)
}

var ReqJSON []byte
var idPedido int

var (
	response http.ResponseWriter
	request  *http.Request
)

func execPedido(res http.ResponseWriter, req *http.Request, pk []string) Retorno {

	dbg.SetDebug(true)

	response = res
	request = req

	var pdv tipos.PedidoVenda
	var pdvorigem tipos.PedidoVenda
	var ped tipos.Pedido
	ret = Retorno{}
	retOriginal = Retorno{}

	idPedido = 0

	switch req.Method {
	case "POST":
		operacao = "C"
	case "GET":
		operacao = "R"
	case "PUT":
		operacao = "U"
	case "DELETE":
		operacao = "D"
	}

	var (
		resp []byte
		eReq error
	)

	if (operacao == "C") || (operacao == "U") {
		// carregar o json do Pedido.
		resp, eReq = ioutil.ReadAll(req.Body)
		if erro.Trata(eReq) {
			msg := fmt.Sprintf("Falha ao carregar o JSON de entrada: %s", eReq.Error())
			erroHTTP(msg)
			return ret
		}
		ReqJSON = resp
	}

	if operacao == "U" {
		pdvorigem = execGetPedidoInt(res, req, pk)
		if carregaDadosPedido(&ped, resp) {
			pdv = ped.PedidoVenda
			executaServico(&pdvorigem, &pdv)
		}
	} else if operacao == "C" {
		if carregaDadosPedido(&ped, resp) {
			pdv = ped.PedidoVenda
			executaServico(&pdv, nil)
		}
	} else if operacao == "D" {
		pdvorigem = execGetPedidoInt(res, req, pk)
		executaServico(&pdvorigem, nil)
	}

	return ret
}

func execGetPedidoInt(res http.ResponseWriter, req *http.Request, pk []string) tipos.PedidoVenda {
	var pdvret tipos.PedidoVenda

	ret = ExecGetPedido(res, req, pk)
	*(&retOriginal) = *(&ret)
	ret = Retorno{}

	bstruct, errM := json.Marshal(&retOriginal.Pedido.Pedido)
	if erro.Trata(errM) {
		msg := fmt.Sprintf("Falha ao ler os dados do Pedido de Venda Origem: %s", errM.Error())
		erroHTTP(msg)
		return pdvret
	} else {
		errU := json.Unmarshal(bstruct, &pdvret)
		if erro.Trata(errU) {
			msg := fmt.Sprintf("Falha ao ler os dados do Pedido de Venda Origem: %s", errU.Error())
			erroHTTP(msg)
			return pdvret
		}
	}
	return pdvret
}

func carregaDadosPedido(pdv *tipos.Pedido, param []byte) bool {
	if err := json.Unmarshal(param, pdv); erro.Trata(err) {
		msg := fmt.Sprintf("Falha ao interpretar o JSON de entrada: %s", err.Error())
		erroHTTP(msg)
		return false
	}
	dbg.Print("pdv: ", pdv)
	pdvuk.Codigo = pdv.PedidoVenda.Codigo
	pdvuk.Organizacao = pdv.PedidoVenda.Organizacao
	pdvuk.Serie = pdv.PedidoVenda.Serie
	return true
}

func InitDb(drv string) *gorp.DbMap {

	var (
		db  *sql.DB
		err error
	)

	connString := os.Getenv("API_PEDIDOVENDA_ORACLE")

	if drv == "GORACLE" {
		if DbMapGO == nil {
			db, err = sql.Open("goracle", connString)
		} else {
			return DbMapGO
		}
	}

	ora.Cfg().Env.StmtCfg.Rset.SetChar(ora.S)
	ora.Cfg().Env.StmtCfg.Rset.SetChar1(ora.S)
	ora.Cfg().Env.StmtCfg.Rset.SetBlob(ora.OraBin)
	if drv == "RANA" {
		if DbMap == nil {
			db, err = sql.Open("ora", connString)
		} else {
			return DbMap
		}
	}

	if erro.Trata(err) {
		erroHTTP(fmt.Sprintf("Falha ao conectar ao Oracle (%s) - %s - %s", connString, drv, err.Error()))
		return nil
	}
	err = db.Ping()
	if erro.Trata(err) {
		fmt.Println(fmt.Sprintf("Falha ao conectar ao Oracle (%s) - %s - %s", connString, drv, err.Error()))
		erroHTTP(fmt.Sprintf("Falha ao conectar ao Oracle (%s) - %s - %s", connString, drv, err.Error()))
		return nil
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.OracleDialect{}}

	// estrutura de pedido para o get
	dbmap.AddTableWithName(tipos.PedidoVendaDB{}, "VEN_PEDIDOVENDA").SetKeys(false, "ORG_IN_CODIGO", "SER_ST_CODIGO", "PED_IN_CODIGO")
	dbmap.AddTableWithNameAndSchema(tipos.ObsPedidoDB{}, "MGVEN", "VEN_OBSERVACAOPEDIDO").SetKeys(false, "ORG_IN_CODIGO", "SER_ST_CODIGO", "PED_IN_CODIGO", "POB_CH_TIPOOBSERVACAO")
	dbmap.AddTableWithNameAndSchema(tipos.OcorrenciaDB{}, "MGVEN", "VEN_OCORRENCIAPED").SetKeys(false, "ORG_IN_CODIGO", "SER_ST_CODIGO", "PED_IN_CODIGO", "OCP_DT_DATAOCORRENCIA")
	dbmap.AddTableWithNameAndSchema(tipos.ArquivoDB{}, "MGVEN", "VEN_PEDIDOARQUIVO").SetKeys(false, "ORG_IN_CODIGO", "SER_ST_CODIGO", "PED_IN_CODIGO", "PAR_IN_CODIGO")
	dbmap.AddTableWithNameAndSchema(tipos.ItemDB{}, "MGVEN", "VEN_ITEMPEDIDOVENDA").SetKeys(false, "ORG_IN_CODIGO", "SER_ST_CODIGO", "PED_IN_CODIGO") //, "ITP_IN_SEQUENCIA"
	dbmap.AddTableWithNameAndSchema(tipos.ObsItemDB{}, "MGVEN", "VEN_OBSITEMPEDIDO").SetKeys(false, "ORG_IN_CODIGO", "SER_ST_CODIGO", "PED_IN_CODIGO", "ITP_IN_SEQUENCIA", "OIP_CH_TIPOOBSERVACAO")
	dbmap.AddTableWithNameAndSchema(tipos.PedProgEntregaDB{}, "MGVEN", "VEN_PEDPROGENTREGA").SetKeys(false, "ORG_IN_CODIGO", "SER_ST_CODIGO", "PED_IN_CODIGO", "ITP_IN_SEQUENCIA", "IPE_IN_SEQUENCIA")
	dbmap.AddTableWithNameAndSchema(tipos.PedProgEstoqueDB{}, "MGVEN", "VEN_PEDPROGESTOQUE").SetKeys(false, "ORG_IN_CODIGO", "SER_ST_CODIGO", "PED_IN_CODIGO", "ITP_IN_SEQUENCIA", "IPE_IN_SEQUENCIA")
	dbmap.AddTableWithNameAndSchema(tipos.ParcFinPedidoDB{}, "MGVEN", "VEN_PARCFINPEDIDO").SetKeys(false, "ORG_IN_CODIGO", "SER_ST_CODIGO", "PED_IN_CODIGO", "PFP_IN_SEQUENCIA")

	dbmap.AddTableWithName(tipos.PedidoVendaMemDB{}, "TABLE(MGVEN.VEN_PCK_PDVCABECALHO.F_GETMEMPEDIDOVENDA)")
	dbmap.AddTableWithNameAndSchema(tipos.TipoDocumentoDB{}, "MGVEN", "VEN_TIPODOCUMENTO").SetKeys(false, "TPD_TAB_IN_CODIGO", "TPD_PAD_IN_CODIGO", "TPD_IN_CODIGO")

	// estrutura de pedido temporaria para o POST
	dbmap.AddTableWithName(tipos.PedidoVenda{}, "VEN_PEDIDOVENDA_INT")
	dbmap.AddTableWithName(tipos.Observacao{}, "VEN_OBSERVACAOPEDIDO_INT")
	dbmap.AddTableWithName(tipos.Ocorrencia{}, "VEN_OCORRENCIAPED_INT")
	dbmap.AddTableWithName(tipos.Arquivo{}, "VEN_PEDIDOARQUIVO_INT")
	dbmap.AddTableWithName(tipos.Item{}, "VEN_ITEMPEDIDOVENDA_INT")
	dbmap.AddTableWithName(tipos.ObsItem{}, "VEN_OBSITEMPEDIDO_INT")
	dbmap.AddTableWithName(tipos.PedProgEntrega{}, "VEN_PEDPROGENTREGA_INT")
	dbmap.AddTableWithName(tipos.PedProgEstoque{}, "VEN_PEDPROGESTOQUE_INT")
	dbmap.AddTableWithName(tipos.ParcFinPedido{}, "VEN_PARCFINPEDIDO_INT")
	return dbmap

}

func trataPedID(pdv *tipos.PedidoVenda, ped string, seq string) *tipos.PedidoVenda {
	if operacao == "C" {
		pdv.Operacao = "I"
	} else {
		pdv.Operacao = operacao
	}
	pdv.IdPai = ped
	pdv.Id = seq
	if operacao == "C" {
		*pdv.Codigo = 0
	}
	if dbg.GetDebug() {
		os.Remove("pedidovenda.txt")
	}
	Debug_sql_gorp(pdv)
	return pdv
}

func trataObsPedID(obs *tipos.Observacao, ped string, seq string) *tipos.Observacao {
	if operacao == "C" {
		obs.Operacao = "I"
	} else {
		obs.Operacao = operacao
	}
	obs.IdPai = ped
	obs.Id = seq
	obs.PedidoUk = pdvuk
	Debug_sql_gorp(obs)
	return obs
}

func trataArquivoPedID(arq *tipos.Arquivo, ped string, seq string) *tipos.Arquivo {
	if operacao == "C" {
		arq.Operacao = "I"
		arq.Sequencia, _ = strconv.Atoi(seq)
	} else {
		arq.Operacao = operacao
	}
	arq.IdPai = ped
	arq.Id = seq
	arq.PedidoUk = pdvuk
	Debug_sql_gorp(arq)
	return arq
}

func trataOcorrenciaPedID(oco *tipos.Ocorrencia, ped string, seq string) *tipos.Ocorrencia {
	if operacao == "C" {
		oco.Operacao = "I"
	} else {
		oco.Operacao = operacao
	}
	oco.IdPai = ped
	oco.Id = seq
	oco.PedidoUk = pdvuk
	Debug_sql_gorp(oco)
	return oco
}

func trataParcFinPedID(fin *tipos.ParcFinPedido, ped string, seq string) *tipos.ParcFinPedido {
	if operacao == "C" {
		fin.Operacao = "I"
		fin.Sequencia, _ = strconv.Atoi(seq)
	} else {
		fin.Operacao = operacao
	}
	fin.IdPai = ped
	fin.Id = seq
	fin.PedidoUk = pdvuk
	Debug_sql_gorp(fin)
	return fin
}

func trataItemID(ite *tipos.Item, ped string, seq string) *tipos.Item {
	if operacao == "C" {
		ite.Operacao = "I"
		ite.Sequencia, _ = strconv.Atoi(seq)
	} else {
		ite.Operacao = operacao
	}
	ite.IdPai = ped
	ite.Id = seq
	ite.PedidoUk = pdvuk
	Debug_sql_gorp(ite)
	return ite
}

func trataObsItemID(obs *tipos.ObsItem, item string, seq string) *tipos.ObsItem {
	if operacao == "C" {
		obs.Operacao = "I"
		obs.SequenciaItem, _ = strconv.Atoi(item)
	} else {
		obs.Operacao = operacao
	}
	obs.IdPai = item
	obs.Id = seq
	obs.PedidoUk = pdvuk
	Debug_sql_gorp(obs)
	return obs
}

func trataProgID(prog *tipos.PedProgEntrega, item string, seq string) *tipos.PedProgEntrega {
	if operacao == "C" {
		prog.Operacao = "I"
		prog.Sequencia, _ = strconv.Atoi(seq)
		prog.SequenciaItem, _ = strconv.Atoi(item)
	} else {
		prog.Operacao = operacao
	}
	prog.IdPai = item
	prog.Id = seq
	prog.PedidoUk = pdvuk
	Debug_sql_gorp(prog)
	return prog
}

func trataEstqID(estq *tipos.PedProgEstoque, item string, prog string, seq string) *tipos.PedProgEstoque {
	if operacao == "C" {
		estq.Operacao = "I"
		estq.SequenciaItem, _ = strconv.Atoi(item)
		estq.SequenciaProg, _ = strconv.Atoi(prog)
		estq.Sequencia, _ = strconv.Atoi(seq)
	} else {
		estq.Operacao = operacao
	}
	estq.IdPai = item
	estq.Id = seq
	estq.PedidoUk = pdvuk
	Debug_sql_gorp(estq)
	return estq
}

var pdvChan chan interface{}
var erroChan chan interface{}

func strID(i int) string {
	return strconv.Itoa(i + 1)
}

func carregaPdv(pdv tipos.PedidoVenda, tr *gorp.Transaction) bool {

	err := tr.Insert(trataPedID(&pdv, strID(0), strID(0)))
	if erro.Trata(err) {
		erroHTTP(err.Error())
		return false
	}
	for i, _ := range pdv.Observacao {
		err = tr.Insert(trataObsPedID(&pdv.Observacao[i], strID(0), strID(i)))
		if erro.Trata(err) {
			erroHTTP(err.Error())
			return false
		}
	}
	for i, _ := range pdv.Ocorrencia {
		err = tr.Insert(trataOcorrenciaPedID(&pdv.Ocorrencia[i], strID(0), strID(i)))
		if erro.Trata(err) {
			erroHTTP(err.Error())
			return false
		}
	}
	for i, _ := range pdv.Arquivo {
		err = tr.Insert(trataArquivoPedID(&pdv.Arquivo[i], strID(0), strID(i)))
		if erro.Trata(err) {
			erroHTTP(err.Error())
			return false
		}
	}
	for iit, _ := range pdv.Item {
		err = tr.Insert(trataItemID(&pdv.Item[iit], strID(0), strID(iit)))
		if erro.Trata(err) {
			erroHTTP(err.Error())
			return false
		}
		for iob, _ := range pdv.Item[iit].ObsItem {
			err = tr.Insert(trataObsItemID(&pdv.Item[iit].ObsItem[iob], strID(iit), strID(iob)))
			if erro.Trata(err) {
				erroHTTP(err.Error())
				return false
			}
		}
		for iprog, _ := range pdv.Item[iit].PedProgEntrega {
			err = tr.Insert(trataProgID(&pdv.Item[iit].PedProgEntrega[iprog], strID(iit), strID(iprog)))
			if erro.Trata(err) {
				erroHTTP(err.Error())
				return false
			}
			for iestq, _ := range pdv.Item[iit].PedProgEntrega[iprog].PedProgEstoque {
				err = tr.Insert(trataEstqID(&pdv.Item[iit].PedProgEntrega[iprog].PedProgEstoque[iestq], strID(iit), strID(iprog), strID(iestq)))
				if erro.Trata(err) {
					erroHTTP(err.Error())
					return false
				}
			}
		}
	}
	for i, _ := range pdv.Parcela {
		err = tr.Insert(trataParcFinPedID(&pdv.Parcela[i], strID(0), strID(i)))
		if erro.Trata(err) {
			erroHTTP(err.Error())
			return false
		}
	}

	return true
}

func insereTemp(pdv *tipos.PedidoVenda, trans *gorp.Transaction) bool {

	return carregaPdv(*pdv, trans)

}

func Debug_sql_gorp(estru interface{}) {
	if dbg.GetDebug() {
		colunas, valores := customizacaoservico.Debug_estruturaParaColunasValores(estru)
		tabgorp, err := DbMap.TableFor(reflect.Indirect(reflect.ValueOf(estru)).Type(), false)
		if erro.Trata(err) {
			erroHTTP(err.Error())
			return
		}
		sql := customizacaoservico.Debug_geraInsertCustom("MGVEN", tabgorp.TableName, colunas)
		for i, _ := range colunas {
			var valor string
			if valores[i] != nil {
				if reflect.TypeOf(valores[i]).Kind() == reflect.Int ||
					reflect.TypeOf(valores[i]).Kind() == reflect.Int32 ||
					reflect.TypeOf(valores[i]).Kind() == reflect.Int64 ||
					reflect.TypeOf(valores[i]).Kind() == reflect.Float32 ||
					reflect.TypeOf(valores[i]).Kind() == reflect.Float64 {
					valor = fmt.Sprint("", valores[i])
				} else if reflect.TypeOf(valores[i]) == reflect.TypeOf(time.Now()) {
					valor = "to_date('" + valores[i].(time.Time).Format("2006-01-01 15:04:05") + "','RRRR-MM-DD HH24:MI:SS')"
				} else if reflect.TypeOf(valores[i]) == reflect.TypeOf([]uint8{}) {
					valor = "'" + fmt.Sprintf("%x", valores[i].([]uint8)) + "'"
					//fmt.Println("[]byte")

				} else {
					valor = "'" + fmt.Sprint("", valores[i]) + "'"
				}
			} else {
				valor = "null"
			}
			/*if valor != "null" {
				fmt.Println(colunas[i], ":", " Type: ", reflect.TypeOf(valores[i]), " Kind: ", reflect.TypeOf(valores[i]).Kind())
			}*/
			sql = strings.Replace(sql, ":"+colunas[i], valor, 1)
		}

		lidos, _ := ioutil.ReadFile("pedidovenda.txt")
		os.Remove("pedidovenda.txt")
		buff := bytes.NewBuffer(lidos)
		buff.WriteString(sql)
		buff.WriteString(";\n")
		ioutil.WriteFile("pedidovenda.txt", buff.Bytes(), os.ModeAppend)
	}
}

func atualizaTemp(pdvori tipos.PedidoVenda, pdvalt tipos.PedidoVenda, tr *gorp.Transaction) bool {
	operacaotemp := operacao

	operacao = "U"
	err := tr.Insert(trataPedID(&pdvalt, strID(0), strID(0)))
	if erro.Trata(err) {
		erroHTTP(err.Error())
		return false
	}

	operacoes, estruturas, _, _ := compara.CompararSliceStruct(&pdvalt.Observacao, &pdvori.Observacao, []string{"Organizacao", "Serie", "Codigo", "TipoObservacao"}, []string{"I", "U", "D"})
	for i, _ := range operacoes {
		operacao = operacoes[i]
		err = tr.Insert(trataObsPedID(estruturas[i].(*tipos.Observacao), strID(0), strID(i)))
		if erro.Trata(err) {
			erroHTTP(err.Error())
			return false
		}
		if dbg.GetDebug() {
			fmt.Println("")
			fmt.Println("operacao: ", operacoes[i])
			fmt.Println("estrutura: ", estruturas[i])
		}
	}

	operacao = "I"
	for i, v := range pdvalt.Parcela {
		err = tr.Insert(trataParcFinPedID(&v, strID(0), strID(i)))
		if erro.Trata(err) {
			erroHTTP(err.Error())
			return false
		}
	}

	operacoes, estruturas, _, _ = compara.CompararSliceStruct(&pdvalt.Ocorrencia, &pdvori.Ocorrencia, []string{"Organizacao", "Serie", "Codigo", "DataOcorrencia"}, []string{"I", "U", "D"})
	for i, _ := range operacoes {
		operacao = operacoes[i]
		err = tr.Insert(trataOcorrenciaPedID(estruturas[i].(*tipos.Ocorrencia), strID(0), strID(i)))
		if erro.Trata(err) {
			erroHTTP(err.Error())
			return false
		}
	}

	seqarquivo := 0
	for _, v := range pdvalt.Arquivo {
		if v.Sequencia > seqarquivo {
			seqarquivo = v.Sequencia
		}
	}
	for _, v := range pdvori.Arquivo {
		if v.Sequencia > seqarquivo {
			seqarquivo = v.Sequencia
		}
	}
	for i, _ := range pdvalt.Arquivo {
		if pdvalt.Arquivo[i].Sequencia == 0 {
			seqarquivo += 1
			pdvalt.Arquivo[i].Sequencia = seqarquivo
		}
	}
	operacoes, estruturas, _, _ = compara.CompararSliceStruct(&pdvalt.Arquivo, &pdvori.Arquivo, []string{"Organizacao", "Serie", "Codigo", "Sequencia"}, []string{"I", "U", "D", "N"})
	for i, _ := range operacoes {
		if operacoes[i] != "N" {
			operacao = operacoes[i]
			err = tr.Insert(trataArquivoPedID(estruturas[i].(*tipos.Arquivo), strID(0), strID(i)))
			if erro.Trata(err) {
				erroHTTP(err.Error())
				return false
			}
		}
	}

	operacoesitem, estruturasitem, posnew, posold := compara.CompararSliceStruct(&pdvalt.Item, &pdvori.Item, []string{"Organizacao", "Serie", "Codigo", "Sequencia"}, []string{"I", "U", "D"})
	for i, _ := range operacoesitem {
		operacao = operacoesitem[i]
		err = tr.Insert(trataItemID(estruturasitem[i].(*tipos.Item), strID(0), strID(i)))
		if erro.Trata(err) {
			erroHTTP(err.Error())
			return false
		}

		if operacao == "U" || operacao == "I" {
			var itemori *tipos.Item

			if operacao == "U" {
				itemori = &pdvori.Item[posold[i]]

				if dbg.GetDebug() {
					fmt.Println("")
					fmt.Println("operacao: ", operacoesitem[i])
					fmt.Println("posição alt: ", posnew[i], "posição old: ", posold[i])

					fmt.Println("estrutura alt: ", pdvalt.Item[posnew[i]])
					fmt.Println("estrutura old: ", pdvori.Item[posold[i]])

					ok, texto := compara.CompararStruct(&pdvalt.Item[posnew[i]], &pdvori.Item[posold[i]])
					fmt.Println("compara.CompararStruct: ", ok, " : ", texto)
				}

			} else {
				itemori = &tipos.Item{}
			}

			operacoesobs, estruturaobs, _, _ := compara.CompararSliceStruct(&pdvalt.Item[posnew[i]].ObsItem, &itemori.ObsItem, []string{"Organizacao", "Serie", "Codigo", "SequenciaItem", "TipoObservacao"}, []string{"I", "U", "D"})
			for j, _ := range operacoesobs {
				operacao = operacoesobs[j]
				err = tr.Insert(trataObsItemID(estruturaobs[j].(*tipos.ObsItem), strID(i), strID(j)))
				if erro.Trata(err) {
					erroHTTP(err.Error())
					return false
				}
			}

			operacoesprog, estruturaprog, posnewprog, posoldprog := compara.CompararSliceStruct(&pdvalt.Item[posnew[i]].PedProgEntrega, &itemori.PedProgEntrega, []string{"Organizacao", "Serie", "Codigo", "SequenciaItem", "Sequencia"}, []string{"I", "U", "D"})
			for j, _ := range operacoesprog {
				operacao = operacoesprog[j]
				err = tr.Insert(trataProgID(estruturaprog[j].(*tipos.PedProgEntrega), strID(i), strID(j)))
				if erro.Trata(err) {
					erroHTTP(err.Error())
					return false
				}

				if operacao == "U" || operacao == "I" {

					var progori *tipos.PedProgEntrega

					if operacao == "U" {
						progori = &itemori.PedProgEntrega[posoldprog[j]]
					} else {
						progori = &tipos.PedProgEntrega{}
					}

					operacoespest, estruturapest, _, _ := compara.CompararSliceStruct(&pdvalt.Item[posnew[i]].PedProgEntrega[posnewprog[j]].PedProgEstoque, &progori.PedProgEstoque, []string{"Organizacao", "Serie", "Codigo", "SequenciaItem", "SequenciaProg", "Sequencia"}, []string{"I", "U", "D"})
					for k, _ := range operacoespest {
						operacao = operacoespest[k]
						err = tr.Insert(trataEstqID(estruturapest[k].(*tipos.PedProgEstoque), strID(i), strID(j), strID(k)))
						if erro.Trata(err) {
							erroHTTP(err.Error())
							return false
						}
					}
				}
			}
		}
	}

	operacao = operacaotemp

	return true
}

func apagaTemp(pdv tipos.PedidoVenda, tr *gorp.Transaction) bool {

	err := tr.Insert(trataPedID(&pdv, strID(0), strID(0)))
	if erro.Trata(err) {
		erroHTTP(err.Error())
		return false
	}

	return true
}

func executaServico(pdv *tipos.PedidoVenda, pdvalt *tipos.PedidoVenda) {

	var err error

	trans, errB := DbMap.Begin()

	if erro.Trata(errB) {
		erroHTTP(errB.Error())
		return
	}

	defer func() {
		if len(ret.Erros) > 0 {
			err = trans.Rollback()
			if erro.Trata(err) {
				erroHTTP(err.Error())
			}
			return
		}
		err = trans.Commit()
		if erro.Trata(err) {
			erroHTTP(err.Error())
		}
	}()

	if dbg.GetDebug() {
		DbMap.TraceOn("[ms-pdv]", log.New(os.Stdout, ":", log.Lmicroseconds))
		DbMapGO.TraceOn("[ms-pdv]", log.New(os.Stdout, ":", log.Lmicroseconds))
		defer DbMap.TraceOff()
		defer DbMapGO.TraceOff()
	}

	if operacao == "C" {
		if !insereTemp(pdv, trans) {
			trans.Rollback()
		} else {
			executaServicoPedido(pdv, trans)
		}
	} else if operacao == "U" {
		if !atualizaTemp(*pdv, *pdvalt, trans) {
			trans.Rollback()
		} else {
			executaServicoPedido(pdvalt, trans)
		}
	} else if operacao == "D" {
		if !apagaTemp(*pdv, trans) {
			trans.Rollback()
		} else {
			executaServicoPedido(pdv, trans)
		}
	}
}

func ExecGetPedido(res http.ResponseWriter, req *http.Request, pk []string) Retorno {
	ret = Retorno{}
	if len(pk) < 3 {
		ret.Pedido.Mensagem = "Pedido de Venda não encontrado. \n A estrutura correta da chamada da API é /pedido/{organização|serie|numero}. Exemplo: /pedido/3|1|778541."
	} else {

		trans, errt := DbMap.Begin()
		if erro.Trata(errt) {
			ret.Pedido.Mensagem = fmt.Sprintf("Erro ao procurar o Pedido de Venda. \n %s", errt.Error())
		}
		org, _ := strconv.Atoi(pk[0])
		cod, _ := strconv.Atoi(pk[2])
		pdk := tipos.PedidoUk{&org, &cod, pk[1]}
		execGet(trans, pdk)
		getCustomizacoes(req, &ret.Pedido.Pedido, trans)
		trans.Commit()

	}
	return ret

}

func getPedidoPK(trans *gorp.Transaction, pk tipos.PedidoUk) {

	ped, err := trans.Get(tipos.PedidoVendaDB{}, *pk.Organizacao, pk.Serie, *pk.Codigo)

	if erro.Trata(err) {
		ret.Pedido.Mensagem = "Erro ao disponibilizar os dados de identificação do documento."
	} else {
		ret.Pedido.Mensagem = ""
		if ped == nil {
			ret.Pedido.Mensagem = "Pedido de Venda não encontrado."
		} else {
			ret.Pedido.Pedido = *ped.(*tipos.PedidoVendaDB)
			if err != nil {
				ret.Pedido.Mensagem = "Erro ao disponibilizar os dados de identificação do documento."
				fmt.Println(err.Error())
			}
		}
	}
}

func getSelect(col interface{}, where string) string {

	etype := reflect.TypeOf(col)
	t, err := DbMap.TableFor(etype, false)
	if err != nil {
		fmt.Println(reflect.Value{}, "\n", err)
	}

	var colunas string
	for _, v := range t.Columns {
		if v.ColumnName != "-" {
			if colunas == "" {
				colunas = v.ColumnName
			} else {
				colunas = fmt.Sprintf("%s,%s", colunas, v.ColumnName)
			}
		}
	}
	sql := fmt.Sprintf("%s %s %s %s", "SELECT", colunas, "FROM", t.TableName)
	if where != "" {
		sql = fmt.Sprintf("%s %s %s", sql, "WHERE", where)
	}
	return sql
}

func getItemPK(trans *gorp.Transaction, ped tipos.PedidoVendaDB) {

	if *ped.Numero != 0 {

		sql := getSelect(tipos.ItemDB{}, "ORG_TAB_IN_CODIGO = 53 AND ORG_PAD_IN_CODIGO = 1 AND ORG_IN_CODIGO = :org AND SER_ST_CODIGO = :ser AND PED_IN_CODIGO = :ped")

		var itp []tipos.ItemDB
		_, err := trans.Select(&itp, sql, ped.Organizacao, ped.Serie, ped.Numero)

		if erro.Trata(err) {
			ret.Pedido.Mensagem = fmt.Sprintf("Erro ao disponibilizar os dados de identificação dos itens do documento: %s", err.Error())
		} else {
			ret.Pedido.Mensagem = ""
			if itp == nil {
				ret.Pedido.Mensagem = "Itens do Pedido de Venda não encontrado."
			} else {
				for _, reg := range itp {
					ret.Pedido.Pedido.Item = append(ret.Pedido.Pedido.Item, reg)
				}
			}
		}
	}
}

func getObsPedidoPK(trans *gorp.Transaction, ped tipos.PedidoVendaDB) {

	if *ped.Numero != 0 {
		sql := getSelect(tipos.ObsPedidoDB{}, "ORG_TAB_IN_CODIGO = 53 AND ORG_PAD_IN_CODIGO = 1 AND ORG_IN_CODIGO = :org AND SER_ST_CODIGO = :ser AND PED_IN_CODIGO = :ped")

		var obs []tipos.ObsPedidoDB

		rows, err := DbMap.Db.Query(sql, ped.Organizacao, ped.Serie, ped.Numero)

		if erro.Trata(err) {
			ret.Pedido.Mensagem = "Erro ao disponibilizar as observações do documento. [" + err.Error() + "]"
		} else {
			ret.Pedido.Mensagem = ""
			var (
				orgtab int
				orgpad int
				org    int
				orgtau string
				serie  string
				numero int
				tipo   *string
				texto  ora.Lob
			)
			defer rows.Close()
			for rows.Next() {
				err := rows.Scan(&orgtab, &orgpad, &org, &orgtau,
					&serie, &numero, &tipo, &texto)
				if err != nil {
					log.Fatal(err)
				}
				gots, _ := texto.Bytes()
				aux := tipos.ObsPedidoDB{tipos.PedidoPkDB{&orgtab, &orgpad, &org, &orgtau, &serie, &numero}, tipo, string(gots), customizacaoservico.Customizacao{}}
				obs = append(obs, aux)
			}
			ret.Pedido.Pedido.Observacao = obs
		}

	}
}

func getArquivoPedidoPK(trans *gorp.Transaction, ped tipos.PedidoVendaDB) {

	if *ped.Numero != 0 {
		sql := getSelect(tipos.ArquivoDB{}, "ORG_TAB_IN_CODIGO = 53 AND ORG_PAD_IN_CODIGO = 1 AND ORG_IN_CODIGO = :org AND SER_ST_CODIGO = :ser AND PED_IN_CODIGO = :ped")

		var arq []tipos.ArquivoDB

		//rows, err := DbMap.Db.Query(sql, ped.Organizacao, ped.Serie, ped.Numero)
		stmt, err := DbMap.Db.Prepare(sql)
		if erro.Trata(err) {
			ret.Pedido.Mensagem = "Erro ao disponibilizar os arquivos do documento. [" + err.Error() + "]"
		} else {
			ret.Pedido.Mensagem = ""

			rows, err := stmt.Query(ped.Organizacao, ped.Serie, ped.Numero)
			if erro.Trata(err) {
				ret.Pedido.Mensagem = "Erro ao disponibilizar os arquivos do documento. [" + err.Error() + "]"
				return
			}

			var (
				orgtab         int
				orgpad         int
				org            int
				orgtau         string
				serie          string
				numero         int
				usuario        *int
				sequencia      *int
				nomearquivo    *string
				resumoconteudo *string
				conteudo       ora.Lob
				datainclusaoO  *time.Time
			)
			defer rows.Close()
			for rows.Next() {
				err := rows.Scan(&orgtab, &orgpad, &org, &orgtau,
					&serie, &numero, &usuario, &sequencia, &nomearquivo,
					&resumoconteudo, &conteudo, &datainclusaoO)
				if err != nil {
					log.Fatal(err)
				}
				barq, _ := conteudo.Bytes()
				aux := tipos.ArquivoDB{tipos.PedidoPkDB{&orgtab, &orgpad, &org, &orgtau, &serie, &numero}, usuario, sequencia, nomearquivo,
					resumoconteudo, barq, datainclusaoO, customizacaoservico.Customizacao{}}
				arq = append(arq, aux)
			}
			ret.Pedido.Pedido.Arquivo = arq
		}

	}
}

func getOcorrenciaPK(trans *gorp.Transaction, ped tipos.PedidoVendaDB) {

	if *ped.Numero != 0 {

		sql := getSelect(tipos.OcorrenciaDB{}, "ORG_TAB_IN_CODIGO = 53 AND ORG_PAD_IN_CODIGO = 1 AND ORG_IN_CODIGO = :org AND SER_ST_CODIGO = :ser AND PED_IN_CODIGO = :ped")

		var oco []tipos.OcorrenciaDB

		rows, err := DbMap.Db.Query(sql, ped.Organizacao, ped.Serie, ped.Numero)

		if erro.Trata(err) {
			ret.Pedido.Mensagem = "Erro ao disponibilizar as ocorrências do documento. [" + err.Error() + "]"
		} else {
			ret.Pedido.Mensagem = ""
			var (
				orgtab int
				orgpad int
				org    int
				orgtau string
				serie  string
				numero int
				data   *time.Time
				codigo *int
				texto  string
			)
			defer rows.Close()
			for rows.Next() {
				err := rows.Scan(&orgtab, &orgpad, &org, &orgtau,
					&serie, &numero, &data, &codigo, &texto)
				if err != nil {
					log.Fatal(err)
				}
				aux := tipos.OcorrenciaDB{tipos.PedidoPkDB{&orgtab, &orgpad, &org, &orgtau, &serie, &numero}, data, codigo, texto, customizacaoservico.Customizacao{}}
				oco = append(oco, aux)
			}
			ret.Pedido.Pedido.Ocorrencia = oco
		}

	}
}

func getParcFinPedidoPK(trans *gorp.Transaction, ped tipos.PedidoVendaDB) {

	if *ped.Numero != 0 {

		sql := getSelect(tipos.ParcFinPedidoDB{}, "ORG_TAB_IN_CODIGO = 53 AND ORG_PAD_IN_CODIGO = 1 AND ORG_IN_CODIGO = :org AND SER_ST_CODIGO = :ser AND PED_IN_CODIGO = :ped ")

		var parc []tipos.ParcFinPedidoDB

		_, err := trans.Select(&parc, sql, ped.Organizacao, ped.Serie, ped.Numero)

		if erro.Trata(err) {
			ret.Pedido.Mensagem = "Erro ao disponibilizar os dados das parcelas financeiras do documento. [" + err.Error() + "]"
		} else {
			ret.Pedido.Mensagem = ""
			for _, reg := range parc {
				ret.Pedido.Pedido.Parcelas = append(ret.Pedido.Pedido.Parcelas, reg)
			}
		}
	}
}
func getObsItemPK(trans *gorp.Transaction, ped tipos.PedidoVendaDB) {

	if *ped.Numero != 0 {

		sql := getSelect(tipos.ObsItemDB{}, "ORG_TAB_IN_CODIGO = 53 AND ORG_PAD_IN_CODIGO = 1 AND ORG_IN_CODIGO = :org AND SER_ST_CODIGO = :ser AND PED_IN_CODIGO = :ped and ITP_IN_SEQUENCIA = :itp ")

		var obs []tipos.ObsItemDB

		for i, v := range ret.Pedido.Pedido.Item {

			rows, err := DbMap.Db.Query(sql, ped.Organizacao, ped.Serie, ped.Numero, v.Sequencia)

			if erro.Trata(err) {
				ret.Pedido.Mensagem = "Erro ao disponibilizar as observações do item. [" + err.Error() + "]"
			} else {
				ret.Pedido.Mensagem = ""
				var (
					orgtab int
					orgpad int
					org    int
					orgtau string
					serie  string
					numero int
					item   *int
					tipo   *string
					texto  ora.Lob
				)
				defer rows.Close()
				for rows.Next() {
					err := rows.Scan(&orgtab, &orgpad, &org, &orgtau,
						&serie, &numero, &item, &tipo, &texto)
					if err != nil {
						log.Fatal(err)
					}
					gots, _ := texto.Bytes()
					aux := tipos.ObsItemDB{tipos.PedidoPkDB{&orgtab, &orgpad, &org, &orgtau, &serie, &numero}, item, tipo, string(gots), customizacaoservico.Customizacao{}}
					obs = append(obs, aux)
				}
				ret.Pedido.Pedido.Item[i].Observacao = obs
			}
		}
	}
}

func getPedProgEntregaPK(trans *gorp.Transaction, ped tipos.PedidoVendaDB) {

	if *ped.Numero != 0 {

		sql := getSelect(tipos.PedProgEntregaDB{}, "ORG_TAB_IN_CODIGO = 53 AND ORG_PAD_IN_CODIGO = 1 AND ORG_IN_CODIGO = :org AND SER_ST_CODIGO = :ser AND PED_IN_CODIGO = :ped and ITP_IN_SEQUENCIA = :itp ")

		for i, v := range ret.Pedido.Pedido.Item {

			var prog []tipos.PedProgEntregaDB

			_, err := trans.Select(&prog, sql, ped.Organizacao, ped.Serie, ped.Numero, v.Sequencia)

			if erro.Trata(err) {
				ret.Pedido.Mensagem = "Erro ao disponibilizar os dados da programação de entrega dos itens do documento. [" + err.Error() + "]"
			} else {
				ret.Pedido.Mensagem = ""
				for _, reg := range prog {
					ret.Pedido.Pedido.Item[i].ProgEntrega = append(ret.Pedido.Pedido.Item[i].ProgEntrega, reg)
				}
			}
		}
	}
}

func getPedProgEstoquePK(trans *gorp.Transaction, ped tipos.PedidoVendaDB) {

	if *ped.Numero != 0 {

		sql := getSelect(tipos.PedProgEstoqueDB{}, "ORG_TAB_IN_CODIGO = 53 AND ORG_PAD_IN_CODIGO = 1 AND ORG_IN_CODIGO = :org AND SER_ST_CODIGO = :ser AND PED_IN_CODIGO = :ped and ITP_IN_SEQUENCIA = :itp and IPE_IN_SEQUENCIA = :prog")

		for i, v := range ret.Pedido.Pedido.Item {

			for x, p := range v.ProgEntrega {

				var estq []tipos.PedProgEstoqueDB

				_, err := trans.Select(&estq, sql, ped.Organizacao, ped.Serie, ped.Numero, v.Sequencia, p.Sequencia)

				if erro.Trata(err) {
					ret.Pedido.Mensagem = "Erro ao disponibilizar os dados de estoque da programação de entrega dos itens do documento. [" + err.Error() + "]"
				} else {
					ret.Pedido.Mensagem = ""
					for _, reg := range estq {
						ret.Pedido.Pedido.Item[i].ProgEntrega[x].ProgEstoque = append(ret.Pedido.Pedido.Item[i].ProgEntrega[x].ProgEstoque, reg)
					}
				}
			}
		}
	}
}

func execGet(trans *gorp.Transaction, pk tipos.PedidoUk) Retorno {

	getPedidoPK(trans, pk)
	if !reflect.DeepEqual(ret.Pedido.Pedido, tipos.PedidoVendaDB{}) {
		getObsPedidoPK(trans, ret.Pedido.Pedido)
		getArquivoPedidoPK(trans, ret.Pedido.Pedido)
		getOcorrenciaPK(trans, ret.Pedido.Pedido)
		getParcFinPedidoPK(trans, ret.Pedido.Pedido)
		getItemPK(trans, ret.Pedido.Pedido)
		getObsItemPK(trans, ret.Pedido.Pedido)
		getPedProgEntregaPK(trans, ret.Pedido.Pedido)
		getPedProgEstoquePK(trans, ret.Pedido.Pedido)
	}

	return ret
}

func processaDocumento(trans *gorp.Transaction) bool {

	cmd := "declare pCursor sys_refcursor; begin mgven.ven_pck_pedidovendarn.p_ProcessaDocumentoServico(pCursor); end;"
	/*
		cmd := "begin " +
			"  MGVEN.VEN_PCK_PEDIDOVENDARN.vPedidoVendaPK := null; " +
			"  MGVEN.VEN_PCK_PEDIDOVENDARN.P_ApagaDadosTodosTypeTab; " +
			"  MGVEN.VEN_PCK_PEDIDOVENDARN.P_InicializaVariaveis; " +
			"  MGVEN.VEN_PCK_PEDIDOVENDARN.P_ProcessaDocumento; " +
			"exception " +
			"  when MGVEN.VEN_PCK_PEDIDOVENDARN.vErroTratado then " +
			"    MGVEN.VEN_PCK_PEDIDOVENDARN.P_ApagaDadosTodosTypeTab; " +
			"    raise_application_error(MGVEN.VEN_PCK_PEDIDOVENDARN.vCodeErro ,MGGLO.GLO_PCK_EXCECAO.P_GetErro); " +
			"  when others then " +
			"    MGVEN.VEN_PCK_PEDIDOVENDARN.P_ApagaDadosTodosTypeTab; " +
			"    raise_application_error(MGVEN.VEN_PCK_PEDIDOVENDARN.vCodeErro ,MGGLO.GLO_PCK_EXCECAO.P_GetErro||chr(13)||sqlerrm); " +
			"end;"
	*/
	_, err := trans.Exec(cmd)
	if erro.Trata(err) {
		erroHTTP(fmt.Sprintf("Erro ao processar o documento: %s", err.Error()))
		return false
	}
	return true
}

func gravaDocumento(trans *gorp.Transaction) bool {
	cmd := "declare pCursor sys_refcursor; begin mgven.ven_pck_pedidovendarn.p_GravaDocumentoServico(pCursor); end;"
	/*
		cmd := "begin " +
			"  MGVEN.VEN_PCK_PEDIDOVENDARN.P_AtualizaAntesGravacao; " +
			"  MGVEN.VEN_PCK_PEDIDOVENDARN.P_GravaDocumento; " +
			"  MGVEN.VEN_PCK_PEDIDOVENDARN.P_AtualizaAposGravacao; " +
			"  MGVEN.VEN_PCK_PEDIDOVENDARN.P_ValidaAposGravacao; " +
			"  MGVEN.VEN_PCK_PEDIDOVENDARN.P_ValidaAposGravacao; " +
			"  MGVEN.VEN_PCK_PEDIDOVENDARN.P_ApagaDadosTodosTypeTab; " +
			"exception " +
			"  when MGVEN.VEN_PCK_PEDIDOVENDARN.vErroTratado then " +
			"    MGVEN.VEN_PCK_PEDIDOVENDARN.P_ApagaDadosTodosTypeTab; " +
			"    raise_application_error(MGVEN.VEN_PCK_PEDIDOVENDARN.vCodeErro ,MGGLO.GLO_PCK_EXCECAO.P_GetErro); " +
			"  when others then " +
			"    MGVEN.VEN_PCK_PEDIDOVENDARN.P_ApagaDadosTodosTypeTab; " +
			"    raise_application_error(MGVEN.VEN_PCK_PEDIDOVENDARN.vCodeErro ,MGGLO.GLO_PCK_EXCECAO.P_GetErro||chr(13)||sqlerrm); " +
			"end;"
	*/
	_, err := trans.Exec(cmd)
	if erro.Trata(err) {
		erroHTTP(fmt.Sprintf("Erro ao gravar o documento: %s", err.Error()))
		return false
	}
	return true
}

func getMemItemDB(trans *gorp.Transaction) []tipos.ItemDB {
	var pedMem []tipos.ItemDB
	var pedaux []interface{}

	sql := getSelect(tipos.ItemDB{}, "")
	sql = strings.Replace(sql, "FROM VEN_ITEMPEDIDOVENDA", "FROM table(MGVEN.VEN_PCK_PDVITEM.F_CarregaItemPipe)", -1)
	pedaux, err := trans.Select(&pedMem, sql)
	if erro.Trata(err) {
		erroHTTP(fmt.Sprintf("Erro ao carregar os dados em memória da package de Item de Pedido de Venda. %s", err.Error()))
		fmt.Println("Erro ao carregar os dados em memória da package de Pedido de Venda.")
	} else {
		for _, reg := range pedaux {
			aux := *reg.(*tipos.ItemDB)
			pedMem = append(pedMem, aux)
		}

	}
	return pedMem
}

func getMemPedidoVendaDB(trans *gorp.Transaction) tipos.PedidoVendaDB {
	var pedMem tipos.PedidoVendaDB
	var pedaux []interface{}

	sql := getSelect(tipos.PedidoVendaDB{}, "")
	sql = strings.Replace(sql, "FROM VEN_PEDIDOVENDA", "FROM table(MGVEN.VEN_PCK_PDVCABECALHO.F_GetMemPedidoVenda)", -1)
	pedaux, err := trans.Select(&pedMem, sql)
	if erro.Trata(err) {
		erroHTTP(fmt.Sprintf("Erro ao carregar os dados em memória da package de Pedido de Venda. %s", err.Error()))
		fmt.Println("Erro ao carregar os dados em memória da package de Pedido de Venda.")
	} else {
		pedMem = *pedaux[0].(*tipos.PedidoVendaDB)
		pedMem.Item = getMemItemDB(trans)
	}
	return pedMem
}

func executaServicoPedido(pdv *tipos.PedidoVenda, trans *gorp.Transaction) {

	if processaDocumento(trans) {
		if operacao != "D" {
			pedMem := getMemPedidoVendaDB(trans)
			if execCredito(trans, pedMem) {
				if gravaDocumento(trans) {
					var pk []tipos.PedidoUk
					trans.Select(&pk, "select ORG_IN_CODIGO, SER_ST_CODIGO, PED_IN_CODIGO from table(MGVEN.VEN_PCK_PEDIDOVENDARN.F_GetPKPedidoVenda)")
					fmt.Println(pk)
					ret = execGet(trans, pk[0])
					copiaCustomizacoes(pdv, &ret.Pedido.Pedido)
					if len(ret.Erros) == 0 {
						var err error
						if operacao == "C" {
							err = gravaCustomizacoes(ret.Pedido.Pedido, trans)
						} else if operacao == "U" {
							err = atualizaCustomizacoes(ret.Pedido.Pedido, retOriginal.Pedido.Pedido, trans)
						}
						if erro.Trata(err) {
							ret.Pedido = Pedido{}
							erroHTTP(fmt.Sprintf("Erro ao gravar as customizações. %s", err.Error()))
						}
					}
				}
			}
		} else {
			if processaDocumento(trans) {
				if gravaDocumento(trans) {
					*(&ret) = *(&retOriginal)
				}
			}

		}
	}
}

func getTipoDocumento(pdv tipos.PedidoVendaDB, trans *gorp.Transaction) tipos.TipoDocumentoDB {

	tpd, err := trans.Get(tipos.TipoDocumentoDB{}, *pdv.TPD_TAB_IN_CODIGO, *pdv.TipoDocumentoPad, *pdv.TipoDocumento)

	tpdret := tipos.TipoDocumentoDB{}

	if erro.Trata(err) || tpd == nil {
		ret.Pedido.Mensagem = ret.Pedido.Mensagem + "\nTipo de Documento do Pedido não encontrado."
	} else {
		tpdret = *tpd.(*tipos.TipoDocumentoDB)
	}
	return tpdret
}

func execCredito(trans *gorp.Transaction, ped tipos.PedidoVendaDB) bool {

	var retCred RetornoCredito

	retCred.Credito.Situacao = "B"
	retCred.Credito.Mensagem = ""

	tpd := getTipoDocumento(ped, trans)

	if tpd.Codigo == nil {
		retCred.Credito.Mensagem = fmt.Sprintf("ERRO: Tipo de Documento %s não encontrado. Verifique as configurações do pedido de venda.", ped)
	} else {
		if *tpd.SujeitoAprovacao == "N" {
			retCred.Credito.Situacao = "P"
			retCred.Credito.Mensagem = ""
		} else {
			body, _ := json.Marshal(ped)

			if SrvCredito == nil || len(SrvCredito.Servico) == 0 {
				SrvCredito = GetCredito()
				if SrvCredito == nil || len(SrvCredito.Servico) == 0 {
					retCred.Credito.Mensagem = "ERRO: o serviço de Analise de Credito não existe ou não esta em funcionamento."
				}
			} else {
				url, _ := consul.GeraURLPost("http",
					SrvCredito.Servico[0].ServiceAddress,
					strconv.Itoa(SrvCredito.Servico[0].ServicePort),
					SrvCredito.Servico[0].ServicePath,
					SrvCredito.Servico[0].ServiceName,
					[][]string{})

				req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
				req.Header.Set("Content-Type", "application/json; param=value")

				client := &http.Client{}
				resp, err := client.Do(req)
				if erro.Trata(err) {
					retCred.Credito.Mensagem = "Erro na análise de crédito: [" + err.Error() + "]"
					return true
				}
				defer resp.Body.Close()
				r, _ := ioutil.ReadAll(resp.Body)

				errc := json.Unmarshal(r, &retCred)

				if errc != nil {
					retCred.Credito.Situacao = "B"
					retCred.Credito.Mensagem = "Erro na análise de crédito: [" + errc.Error() + "]"
				}
				if len(retCred.Erros) > 0 {
					retCred.Credito.Situacao = "B"
					retCred.Credito.Mensagem = retCred.Credito.Mensagem + " " + retCred.Erros[0].Mensagem
				}
			}
		}
	}
	_, err := trans.Exec("begin MGVEN.VEN_PCK_PDVSITUACAO.P_ProcessaSetSituacao(:sit, :men); end;", retCred.Credito.Situacao, retCred.Credito.Mensagem)
	if erro.Trata(err) {
		erroHTTP(fmt.Sprintf("Erro ao processar a analise de credito : %s", err.Error()))
		return false
	}
	return true

}

func getPedidoJson(res http.ResponseWriter, req *http.Request) string {
	dat, err := ioutil.ReadFile("swagger\\data\\pedido\\index.json")
	if !erro.Trata(err) {
		return string(dat)
	} else {
		erroHTTP(fmt.Sprintf("Erro para disponibilizar o json do pedido: %s", err.Error()))
	}
	return ""
}

func GetCredito() *consul.Servico { //encontra o serviço publicado no consul.
	pathServico := os.Getenv("API_ANALISECREDITO_PATH")
	if pathServico == "" {
		pathServico = "avi/" + tipos.Versao
	}
	nomeServico := os.Getenv("API_ANALISECREDITO_NOME")
	if nomeServico == "" {
		nomeServico = "credito"
	}
	hostCredito := os.Getenv("API_ANALISECREDITO_HOST")
	if hostCredito == "" {
		hostCredito = "127.0.0.1"
	}
	portCredito := os.Getenv("API_ANALISECREDITO_PORT")
	if portCredito == "" {
		portCredito = "8081"
	}
	if ConsulCredito {
		return consul.GetServico(pathServico + "/" + nomeServico)
	} else {
		var srv consul.Servico
		if hostCredito != "" && portCredito != "" {
			var s consul.Srv
			s.ServiceAddress = hostCredito
			s.ServicePort, _ = strconv.Atoi(portCredito)
			s.ServicePath = pathServico
			s.ServiceName = nomeServico
			srv.Servico = append(srv.Servico, s)
			return &srv
		} else {
			return nil
		}
	}
}

func getCustomizacoes(req *http.Request, pdv *tipos.PedidoVendaDB, trans *gorp.Transaction) {
	tabelascustom := customizacaoservico.CarregaTabelasCustomDoHeader(req)

	customizacaoservico.CarregaCustomizacao("MGVEN", "VEN_PEDIDOVENDA", pdv, tabelascustom, &pdv.Customizacoes.Tabela, trans)
	customizacaoservico.CarregaCustomizacao("MGVEN", "VEN_OBSERVACAOPEDIDO", &pdv.Observacao, tabelascustom, &pdv.Observacao, trans)
	customizacaoservico.CarregaCustomizacao("MGVEN", "VEN_OCORRENCIAPED", &pdv.Ocorrencia, tabelascustom, &pdv.Ocorrencia, trans)
	customizacaoservico.CarregaCustomizacao("MGVEN", "VEN_PEDIDOARQUIVO", &pdv.Arquivo, tabelascustom, &pdv.Arquivo, trans)
	customizacaoservico.CarregaCustomizacao("MGVEN", "VEN_PARCFINPEDIDO", &pdv.Parcelas, tabelascustom, &pdv.Parcelas, trans)

	customizacaoservico.CarregaCustomizacao("MGVEN", "VEN_ITEMPEDIDOVENDA", &pdv.Item, tabelascustom, &pdv.Item, trans)
	for i, _ := range pdv.Item {
		customizacaoservico.CarregaCustomizacao("MGVEN", "VEN_OBSITEMPEDIDO", &pdv.Item[i].Observacao, tabelascustom, &pdv.Item[i].Observacao, trans)

		customizacaoservico.CarregaCustomizacao("MGVEN", "VEN_PEDPROGENTREGA", &pdv.Item[i].ProgEntrega, tabelascustom, &pdv.Item[i].ProgEntrega, trans)
		for k, _ := range pdv.Item[i].ProgEntrega {
			customizacaoservico.CarregaCustomizacao("MGVEN", "VEN_PEDPROGESTOQUE", &pdv.Item[i].ProgEntrega[k].ProgEstoque, tabelascustom, &pdv.Item[i].ProgEntrega[k].ProgEstoque, trans)
		}
	}
}

func copiaCustomizacoes(pdvint *tipos.PedidoVenda, pdv *tipos.PedidoVendaDB) {
	pdv.Customizacoes = pdvint.Customizacoes

	for i, _ := range pdv.Item {

		for _, v := range pdvint.Item {
			if *pdv.Item[i].Sequencia == v.Sequencia {
				pdv.Item[i].Customizacoes = v.Customizacoes

				for k, _ := range pdv.Item[i].Observacao {
					for _, e := range v.ObsItem {
						if *pdv.Item[i].Observacao[k].TipoobservacaoO == e.TipoObservacao {
							pdv.Item[i].Observacao[k].Customizacoes = e.Customizacoes
							break
						}
					}
				}

				for k, _ := range pdv.Item[i].ProgEntrega {
					for _, e := range v.PedProgEntrega {
						if *pdv.Item[i].ProgEntrega[k].Sequencia == e.Sequencia {
							pdv.Item[i].ProgEntrega[k].Customizacoes = e.Customizacoes

							for j, _ := range pdv.Item[i].ProgEntrega[k].ProgEstoque {
								for _, u := range e.PedProgEstoque {
									if *pdv.Item[i].ProgEntrega[k].ProgEstoque[j].PPE_IN_SEQUENCIA == u.Sequencia {
										pdv.Item[i].ProgEntrega[k].ProgEstoque[j].Customizacoes = u.Customizacoes
										break
									}

								}

							}
							break
						}
					}

				}
				break
			}
		}
	}

	for i, _ := range pdv.Observacao {
		for _, v := range pdvint.Observacao {
			if *pdv.Observacao[i].Tipoobservacao == v.TipoObservacao {
				pdv.Observacao[i].Customizacoes = v.Customizacoes
				break
			}
		}
	}

	for i, _ := range pdv.Ocorrencia {
		for _, v := range pdvint.Ocorrencia {
			if *pdv.Ocorrencia[i].DataOcorrencia == *v.DataOcorrencia {
				pdv.Ocorrencia[i].Customizacoes = v.Customizacoes
				break
			}
		}
	}

	for i, _ := range pdv.Arquivo {
		for _, v := range pdvint.Arquivo {
			if *pdv.Arquivo[i].Sequencia == v.Sequencia {
				pdv.Arquivo[i].Customizacoes = v.Customizacoes
				break
			}
		}
	}

	for i, _ := range pdv.Parcelas {
		for _, v := range pdvint.Parcela {
			if *pdv.Parcelas[i].Sequencia == v.Sequencia {
				pdv.Parcelas[i].Customizacoes = v.Customizacoes
				break
			}
		}
	}
}

func gravaCustomizacoes(pdv tipos.PedidoVendaDB, trans *gorp.Transaction) error {

	err := customizacaoservico.InsereCustomizacao("MGVEN", "VEN_PEDIDOVENDA", &pdv, trans)
	if err != nil {
		return err
	}

	err = customizacaoservico.InsereCustomizacao("MGVEN", "VEN_OBSERVACAOPEDIDO", &pdv.Observacao, trans)
	if err != nil {
		return err
	}

	err = customizacaoservico.InsereCustomizacao("MGVEN", "VEN_OCORRENCIAPED", &pdv.Ocorrencia, trans)
	if err != nil {
		return err
	}

	err = customizacaoservico.InsereCustomizacao("MGVEN", "VEN_PEDIDOARQUIVO", &pdv.Arquivo, trans)
	if err != nil {
		return err
	}

	err = customizacaoservico.InsereCustomizacao("MGVEN", "VEN_PARCFINPEDIDO", &pdv.Parcelas, trans)
	if err != nil {
		return err
	}

	err = customizacaoservico.InsereCustomizacao("MGVEN", "VEN_ITEMPEDIDOVENDA", &pdv.Item, trans)
	if err != nil {
		return err
	}

	for _, v := range pdv.Item {
		err := customizacaoservico.InsereCustomizacao("MGVEN", "VEN_OBSITEMPEDIDO", &v.Observacao, trans)
		if err != nil {
			return err
		}
		err = customizacaoservico.InsereCustomizacao("MGVEN", "VEN_PEDPROGENTREGA", &v.ProgEntrega, trans)
		if err != nil {
			return err
		}

		for _, e := range v.ProgEntrega {
			err := customizacaoservico.InsereCustomizacao("MGVEN", "VEN_PEDPROGESTOQUE", &e.ProgEstoque, trans)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func atualizaCustomizacoes(pdvalt tipos.PedidoVendaDB, pdvori tipos.PedidoVendaDB, trans *gorp.Transaction) error {

	err := customizacaoservico.AtualizaCustomizacao("MGVEN", "VEN_PEDIDOVENDA", &pdvori, &pdvalt, trans)
	if err != nil {
		return err
	}

	err = customizacaoservico.AtualizaCustomizacao("MGVEN", "VEN_OBSERVACAOPEDIDO", &pdvori.Observacao, &pdvalt.Observacao, trans)
	if err != nil {
		return err
	}

	err = customizacaoservico.AtualizaCustomizacao("MGVEN", "VEN_OCORRENCIAPED", &pdvori.Ocorrencia, &pdvalt.Ocorrencia, trans)
	if err != nil {
		return err
	}

	err = customizacaoservico.AtualizaCustomizacao("MGVEN", "VEN_PEDIDOARQUIVO", &pdvori.Arquivo, &pdvalt.Arquivo, trans)
	if err != nil {
		return err
	}

	err = customizacaoservico.AtualizaCustomizacao("MGVEN", "VEN_PARCFINPEDIDO", &pdvori.Parcelas, &pdvalt.Parcelas, trans)
	if err != nil {
		return err
	}

	err = customizacaoservico.AtualizaCustomizacao("MGVEN", "VEN_ITEMPEDIDOVENDA", &pdvori.Item, &pdvalt.Item, trans)
	if err != nil {
		return err
	}

	operacoesitem, _, posnew, posold := compara.CompararSliceStruct(&pdvalt.Item, &pdvori.Item, []string{"Organizacao", "Serie", "Numero", "Sequencia"}, []string{"I", "U", "N"})
	for i, _ := range operacoesitem {
		if operacoesitem[i] == "N" || operacoesitem[i] == "U" || operacoesitem[i] == "I" {
			var itemori *tipos.ItemDB
			var itemalt *tipos.ItemDB

			if operacoesitem[i] == "N" || operacoesitem[i] == "U" {
				itemori = &pdvori.Item[posold[i]]
				itemalt = &pdvalt.Item[posnew[i]]
			} else if operacoesitem[i] == "I" {
				itemori = &tipos.ItemDB{}
				itemalt = &pdvalt.Item[posnew[i]]
			} /*else operacoesitem[i]=="D" {
			itemori = &pdvori.Item[posold[i]]
			itemalt = nil
			}*/

			err := customizacaoservico.AtualizaCustomizacao("MGVEN", "VEN_OBSITEMPEDIDO", &itemori.Observacao, &itemalt.Observacao, trans)
			if err != nil {
				return err
			}

			err = customizacaoservico.AtualizaCustomizacao("MGVEN", "VEN_PEDPROGENTREGA", &itemori.ProgEntrega, &itemalt.ProgEntrega, trans)
			if err != nil {
				return err
			}
			operacoesprog, _, posnewprog, posoldprog := compara.CompararSliceStruct(&itemalt.ProgEntrega, &itemori.ProgEntrega, []string{"Organizacao", "Serie", "Numero", "Sequenciaitem", "Sequencia"}, []string{"I", "U", "N"})
			for j, _ := range operacoesprog {

				if operacoesprog[j] == "N" || operacoesprog[j] == "U" || operacoesprog[j] == "I" {

					var progori *tipos.PedProgEntregaDB
					var progalt *tipos.PedProgEntregaDB

					if operacoesprog[j] == "N" || operacoesprog[j] == "U" {
						progori = &itemori.ProgEntrega[posoldprog[j]]
						progalt = &itemalt.ProgEntrega[posnewprog[j]]
					} else if operacoesprog[j] == "I" {
						progori = &tipos.PedProgEntregaDB{}
						progalt = &itemalt.ProgEntrega[posnewprog[j]]
					} /*else operacoesprog[j]=="D" {
					progori = &itemori.ProgEntrega[posoldprog[j]]
					progalt = &tipos.PedProgEntregaDB{}
					}*/
					err := customizacaoservico.AtualizaCustomizacao("MGVEN", "VEN_PEDPROGESTOQUE", &progori.ProgEstoque, &progalt.ProgEstoque, trans)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}
