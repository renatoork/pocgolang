// @APIVersion 1.0.0
// @SwaggerVersion 2.0
// @APITitle MEGA API PDV MGWEB
// @APIDescription Serviço para consumir a API do PDV a partir dos dados do MGWEB.
// @BasePath /api
package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"mega/apipdvmgweb/mgweb"
	"mega/go-util/dbg"

	"gopkg.in/gorp.v1"

	"mega/go-util/erro"
	"mega/ms-consul/consul"
	"mega/ms-pedidovenda/pedidovenda"
	"mega/ms-pedidovenda/tipos"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	_ "gopkg.in/gorp.v1"
	"gopkg.in/rana/ora.v3"
)

var (
	srv    string
	numCPU int
	db     *sql.DB
	errdb  error
	wg     sync.WaitGroup
	tasks  chan int64
	dbmap  *gorp.DbMap
	pedimp mgweb.Pedido
)

var conf consul.ConfigServico
var servico = flag.String("srv", "http://localhost:8080/api/v0/pedido", "Endereço do serviço de Pedido de Venda. Default: http://localhost:8080/api/v0/pedido")
var cpu = flag.Int("cpu", 0, "Número de CPU paralelas consumindo o serviço de pedido. Default 1 .")
var pedido = flag.Int("pedido", 0, "Número do pedido específico a importar.")

func main() {

	os.Setenv("NLS_LANG", ".UTF8")

	flag.Parse()
	if *cpu == 0 {
		numCPU = 1
	} else {
		numCPU = *cpu
	}

	if *servico == "" {
		srv = "http://localhost:8080/api/v0/pedido"
	} else {
		srv = *servico
	}

	//ler pedidos da base
	connString := os.Getenv("API_PEDIDOVENDA_ORACLE")
	if connString == "" {
		connString = "mgven/megaven@pclopez/orcl" // para teste
	}
	ora.Cfg().Env.StmtCfg.Rset.SetChar(ora.S)
	ora.Cfg().Env.StmtCfg.Rset.SetChar1(ora.S)
	db, errdb = sql.Open("ora", connString)
	if erro.Trata(errdb) {
		panic(fmt.Sprintf("ERRO ao conectar ao Oracle - %s", errdb.Error()))
	}
	errp := db.Ping()
	if erro.Trata(errp) {
		panic(fmt.Sprintf("ERRO ao conectar ao Oracle - %s", errp.Error()))
	}

	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.OracleDialect{}}

	// estrutura de pedido para o get
	dbmap.AddTableWithNameAndSchema(mgweb.PedidoVenda{}, "MGWEB", "VEN_PEDIDOVENDA").SetKeys(false, "PED_IN_SEQUENCIA")
	dbmap.AddTableWithNameAndSchema(mgweb.ObsPedido{}, "MGWEB", "VEN_OBSERVACAOPEDIDO").SetKeys(false, "PED_IN_SEQUENCIA", "POB_CH_TIPOOBSERVACAO")
	dbmap.AddTableWithNameAndSchema(mgweb.Item{}, "MGWEB", "VEN_ITEMPEDIDOVENDA").SetKeys(false, "PED_IN_SEQUENCIA") //, "ITP_IN_SEQUENCIA"
	dbmap.AddTableWithNameAndSchema(mgweb.ObsItem{}, "MGWEB", "VEN_OBSITEMPEDIDO").SetKeys(false, "PED_IN_SEQUENCIA", "ITP_IN_SEQUENCIA", "OIP_CH_TIPOOBSERVACAO")
	dbmap.AddTableWithNameAndSchema(mgweb.PedProgEntrega{}, "MGWEB", "VEN_PEDPROGENTREGA").SetKeys(false, "PED_IN_SEQUENCIA", "ITP_IN_SEQUENCIA", "IPE_IN_SEQUENCIA")
	dbmap.AddTableWithNameAndSchema(mgweb.PedProgEstoque{}, "MGWEB", "VEN_PEDPROGESTOQUE").SetKeys(false, "PED_IN_SEQUENCIA", "ITP_IN_SEQUENCIA", "IPE_IN_SEQUENCIA")
	dbmap.AddTableWithNameAndSchema(mgweb.ParcFinPedido{}, "MGWEB", "VEN_PARCFINPEDIDO").SetKeys(false, "PED_IN_SEQUENCIA", "PFP_IN_SEQUENCIA")

	dbg.SetDebug(true)
	//dbmap.TraceOn("[mgweb] ", log.New(os.Stdout, ":", 0))

	tasks = make(chan int64, 200)

	//for true {
	criarFilaProcessamento(numCPU) //cria a fila de gorotinas para processamento do pedido
	carregarPedidoNaFila()         //cria a fila de gorotinas para processamento do pedido
	//colocar timer
	//}

}

func criarFilaProcessamento(numCPU int) {
	var inicio time.Time

	for i := 0; i < numCPU; i++ {
		wg.Add(1)
		go func(cpu int) {
			for pk := range tasks {
				fmt.Println(cpu, ": ---------------------------------------------")
				trans, errt := dbmap.Begin()
				dbg.Print("begin..", nil)
				if erro.Trata(errt) {
					dbg.Print("  - ERRO ao iniciar a transação: "+errt.Error(), nil)
					panic("  - ERRO ao iniciar a transação: " + errt.Error())
				}

				rows, _ := trans.SelectInt("select p.ped_in_sequencia from mgweb.ven_pedidovenda p where p.ped_in_sequencia = :ped_in_sequencia for update skip locked", pk)
				if rows == 0 {
					dbg.Print("  - sem pedido a importar", nil)
				} else {
					//defer rows.Close()

					body := getPedidoMgWeb(trans, pk)
					dbg.Print("  - retorno ["+string(body)+"]", nil)
					if string(body) != "" {
						inicio = time.Now()

						req, err := http.NewRequest("POST", srv, bytes.NewBuffer(body))
						req.Header.Set("Content-Type", "application/json; param=value")
						client := &http.Client{}
						resp, err := client.Do(req)
						if erro.Trata(err) {
							logarErro(trans, pk, "  - ERRO no serviço de Pedido de Venda : ["+err.Error()+"]")
						} else {
							defer resp.Body.Close()
							r, _ := ioutil.ReadAll(resp.Body)

							dbg.Print("  - executou serviço", nil)

							dbg.Print(fmt.Sprintf("  RETORNO: %s", string(r)), nil)

							type ret struct {
								Pedido tipos.PedidoPkDB    `json:"pedido,omitempty"`
								Erros  []pedidovenda.Falha `json:"erros,omitempty"`
							}

							var retorno ret
							errs := json.Unmarshal(r, &retorno)
							if erro.Trata(errs) {
								logarErro(trans, pk, "  - ERRO no serviço de Pedido de Venda : ["+err.Error()+"]")
							} else {
								dbg.Print(fmt.Sprintf("  RETORNO: %v", retorno), nil)
								if len(retorno.Erros) > 0 {
									logarErro(trans, pk, retorno.Erros[0].Mensagem)
								} else {
									logarImp(trans, pk, *retorno.Pedido.Organizacao, *retorno.Pedido.Serie, *retorno.Pedido.Numero)
									dbg.Trace(inicio)
								}
							}
						}
					}
				}
				fmt.Println("---------------------------------------------")
				trans.Commit()
			}
			wg.Done()
		}(i)
	}
}

func logarErro(tx *gorp.Transaction, pedido int64, mens string) {
	dbg.Print("Erro: "+mens, nil)
	tx.Exec("UPDATE MGWEB.VEN_PEDIDOVENDA "+
		"   SET PED_CH_SITUACAO  = 'N', "+
		"       PED_BO_ERRO      = 'S', "+
		"       PED_BO_TEXTOERRO = :PED_BO_TEXTOERRO,"+
		"       PED_CH_STATUSIMP = 'N',"+
		"       PED_CH_ROTINAIMP = 'N' "+
		" WHERE PED_IN_SEQUENCIA = :PED_IN_SEQUENCIA "+
		"   AND PE_PED_IN_CODIGO IS NULL", mens, pedido)
}

func logarImp(tx *gorp.Transaction, pedido int64, org int, ser string, ped int) {
	dbg.Print(fmt.Sprintf("  - Pedido %d importado", pedido), nil)
	tx.Exec("UPDATE MGWEB.VEN_PEDIDOVENDA "+
		"   SET PED_CH_SITUACAO  = 'I', "+
		"       PED_BO_ERRO      = 'N', "+
		"       PED_BO_TEXTOERRO = null,"+
		"       PED_CH_STATUSIMP = 'S',"+
		"       PED_CH_ROTINAIMP = 'G', "+
		"       PE_ORG_TAB_IN_CODIGO = 53, "+
		"       PE_ORG_PAD_IN_CODIGO = 1, "+
		"       PE_ORG_IN_CODIGO = :ORG, "+
		"       PE_ORG_TAU_ST_CODIGO = 'G', "+
		"       PE_SER_ST_CODIGO = :SER, "+
		"       PE_PED_IN_CODIGO = :PED "+
		" WHERE PED_IN_SEQUENCIA = :PED_IN_SEQUENCIA "+
		"   AND PE_PED_IN_CODIGO IS NULL", org, ser, ped, pedido)
}

const qry = "select v.ped_in_sequencia " +
	"  from mgweb.ven_pedidovenda v " +
	" where v.ped_ch_situacao = 'N' " +
	"   and v.ped_ch_statusimp = 'N' " +
	"   and v.pe_ped_in_codigo is Null " +
	//"   and nvl(v.ped_bo_erro,'N') = 'N' " +
	" order by v.ped_in_prioridade, v.ped_in_sequencia "

func carregarPedidoNaFila() {

	rows, err := dbmap.Db.Query(qry)

	if erro.Trata(err) {
		dbg.Print(fmt.Sprintf("  - ERRO ao recuperar os dados do MGWEB: %s", err.Error()), nil)
	} else if rows == nil {
		dbg.Print("  - Sem pedidos a inportar.", nil)
	} else {
		defer rows.Close()

		var ped_in_sequencia float64

		for rows.Next() {

			err := rows.Scan(&ped_in_sequencia)
			if err != nil {
				dbg.Print(fmt.Sprintf("  - ERRO ao recuperar os dados do MGWEB - %s", err.Error()), nil)
				panic(fmt.Sprintf("  - ERRO ao recuperar os dados do MGWEB - %s", err.Error()))
			}

			tasks <- int64(ped_in_sequencia)

			break
		}
	}

	close(tasks)
	wg.Wait()

}

func getPedidoMgWeb(tx *gorp.Transaction, pk int64) []byte {
	var mens []byte
	dbg.Print(fmt.Sprintf("  pk [%d]", pk), nil)

	ped, err := tx.Get(mgweb.PedidoVenda{}, pk)
	if err != nil {
		dbg.Print(fmt.Sprintf("  - ERRO ao recuperar os dados do pedido: %s - %s", fmt.Sprintf("  pk [%d]", pk), err.Error()), nil)
		return []byte("")
	} else {
		if ped != nil {
			dbg.Print(fmt.Sprintf("  tem ped: %v ", ped), nil)

			pedimp.PedidoVenda = *ped.(*mgweb.PedidoVenda)
			codigo := 0
			pedimp.PedidoVenda.PED_IN_CODIGO = &codigo
			pedimp.PedidoVenda.ORG_IN_CODIGO = getOrganizacao(pedimp.PedidoVenda.FIL_IN_CODIGO)
			pedimp.PedidoVenda.GRU_IN_CODIGO = getUsuario(pedimp.PedidoVenda.GRU_IN_CODIGO)
			pedimp.PedidoVenda.SER_ST_CODIGO = getSerie(tx, *pedimp.PedidoVenda.SER_ST_CODIGO, *pedimp.PedidoVenda.TPD_IN_CODIGO)

			var err1 error

			if tpd.AltParc == "S" {
				var parcfin []mgweb.ParcFinPedido
				_, err1 = tx.Select(&parcfin, "select PED_IN_SEQUENCIA,PFP_IN_SEQUENCIA,PFP_ST_DOCUMENTO,PFP_ST_PARCELA,PFP_DT_VENCTO,PFP_RE_VALORMOE,PFP_RE_PERC,PFP_HCOB_IN_SEQUENCIA from mgweb.Ven_parcfinpedido where PED_IN_SEQUENCIA = :pk", pk)
				if erro.Trata(err1) {
					dbg.Print("erro: "+err1.Error(), nil)
				}
				pedimp.PedidoVenda.Parcela = parcfin
			}

			var obsped []mgweb.ObsPedido
			_, err1 = tx.Select(&obsped, "select POB_CH_TIPOOBSERVACAO,PED_IN_SEQUENCIA,POB_ST_OBSERVACAO from mgweb.Ven_Observacaopedido where PED_IN_SEQUENCIA = :pk", pk)
			if erro.Trata(err1) {
				dbg.Print("erro: "+err1.Error(), nil)
			}
			pedimp.PedidoVenda.Observacao = obsped

			var itp []mgweb.Item
			_, err1 = tx.Select(&itp, "select PED_IN_SEQUENCIA,ITP_ST_COMPLEMENTO,ITP_IN_SEQUENCIA,APL_IN_CODIGO,PRO_ST_ALTERNATIVO,COS_IN_CODIGO,PROJ_IDE_ST_CODIGO,PROJ_IN_REDUZIDO,TPR_IN_CODIGO,TPP_IN_CODIGO,CUS_IDE_ST_CODIGO,CUS_IN_REDUZIDO,ITP_CH_FRETEPCONTA,PRO_IN_CODIGO,UIN_IN_CODIGO,UIN_DT_INCLUSAO,FMT_ST_CODIGO,EMB_IN_CODIGO,ALM_IN_CODIGO,UNI_ST_UNIDADE,LOC_IN_CODIGO,ITP_ST_DESCRICAO,EMB_UNI_ST_UNIDADE,EMB_FMT_ST_CODIGO,ITP_RE_QUANTIDADE,ITP_RE_QTDEAUX,ITP_RE_VALORUNITARIO,ITP_RE_LARGURA,ITP_RE_COMPRIMENTO,ITP_RE_VALORMERCADORIA,ITP_RE_MEDIDAREAL,ITP_RE_VALORMERCEMPREG,ITP_RE_VALORMAOOBRA,ITP_RE_VALORTOTAL,ITP_RE_FRETE,ITP_ST_PEDIDOCLIENTE,ITP_ST_CODPROCLI,ITP_RE_PERCDESCONTO,ITP_RE_VALORDESCONTO,ITP_RE_PERCACRESCIMO,ITP_RE_VALORACRESCIMO,TPC_ST_CLASSE,ITP_CH_STATUSIMP,ITP_RE_VALORCAUCAO,ITP_RE_PERCCAUCAO,ITP_RE_ALIQIPI,ITP_RE_BASEIPI,ITP_RE_ICMSRETIDO,ITP_RE_BASESUBTRIB,ITP_RE_ALIQICMS,ITP_RE_VALORICMS,ITP_RE_BASEICMS,ITP_RE_PRECOTABELA,CFOP_ST_EXTENSO,ITP_RE_SEGURO,ITP_RE_DESPACESSORIA,ITP_RE_PERCREDIPI,ITP_RE_VALORIPI,ITP_RE_PERCDESCDIG,ITP_RE_VALDESCUNITARIO,ITP_RE_PERDESCUNITARIO,ITP_BO_ESPECIAL,CPS_IN_CODIGO from mgweb.ven_itempedidovenda where PED_IN_SEQUENCIA = :pk", pk)
			if erro.Trata(err1) {
				dbg.Print("erro: "+err1.Error(), nil)
			}
			pedimp.PedidoVenda.Item = itp

			for i, v := range pedimp.PedidoVenda.Item {

				var obsitem []mgweb.ObsItem
				_, err1 := tx.Select(&obsitem, "select PED_IN_SEQUENCIA,ITP_IN_SEQUENCIA,OIP_CH_TIPOOBSERVACAO,ITO_ST_OBSERVACAO from mgweb.Ven_Obsitempedido where PED_IN_SEQUENCIA = :pk and ITP_IN_SEQUENCIA = :it", pk, v.ITP_IN_SEQUENCIA)
				if erro.Trata(err1) {
					dbg.Print("erro: "+err1.Error(), nil)
				}
				pedimp.PedidoVenda.Item[i].ObsItem = obsitem

				var pedprog []mgweb.PedProgEntrega
				_, err1 = tx.Select(&pedprog, "select PED_IN_SEQUENCIA,ITP_IN_SEQUENCIA,IPE_IN_SEQUENCIA,UNI_ST_UNIDADE,CLI_ST_CODIGO,ENA_IN_CODIGO,CLI_ST_TIPOCODIGO,CLI_TAU_ST_CODIGO,IPE_DT_DATAENTREGA,IPE_CH_TIPOENTREGA,IPE_CH_ENTREGAPARCIAL,IPE_ST_NUMEROORDEM,IPE_RE_QUANTIDADE,IPE_DT_DATAEMISSAO,UIN_IN_CODIGO,UIN_DT_INCLUSAO,IPE_CH_TIPODATA,IPE_DT_DATAORIGINAL,IPE_DT_DATAPLANEJADA,FMT_ST_CODIGO,EMB_IN_CODIGO from mgweb.Ven_Pedprogentrega where PED_IN_SEQUENCIA = :pk and ITP_IN_SEQUENCIA = :it", pk, v.ITP_IN_SEQUENCIA)
				if erro.Trata(err1) {
					dbg.Print("erro: "+err1.Error(), nil)
				}
				pedimp.PedidoVenda.Item[i].PedProgEntrega = pedprog

				for c, e := range pedimp.PedidoVenda.Item[i].PedProgEntrega {
					var progestq []mgweb.PedProgEstoque
					_, err1 := tx.Select(&progestq, "select PED_IN_SEQUENCIA,ITP_IN_SEQUENCIA,IPE_IN_SEQUENCIA,IPPE_IN_SEQUENCIA,ALM_IN_CODIGO,LOC_IN_CODIGO,NAT_ST_CODIGO,MVS_ST_REFERENCIA,MVS_ST_LOTEFORNE,MVS_DT_ENTRADA,MVS_DT_VALIDADE,LMS_RE_QUANTIDADE From mgweb.Ven_Pedprogestoque where PED_IN_SEQUENCIA = :pk and ITP_IN_SEQUENCIA = :it and IPE_IN_SEQUENCIA = :ipe", pk, v.ITP_IN_SEQUENCIA, e.IPE_IN_SEQUENCIA)
					if erro.Trata(err1) {
						dbg.Print("erro: "+err1.Error(), nil)
					}
					pedimp.PedidoVenda.Item[i].PedProgEntrega[c].PedProgEstoque = progestq
				}
			}
			mens, _ = json.Marshal(pedimp)
		}
	}
	return mens
}

var tpd mgweb.Tipodoc

func getSerie(tx *gorp.Transaction, ser string, tipoDoc int) *string {
	serie := ser
	tpd = mgweb.Tipodoc{}
	rows, err1 := dbmap.Db.Query("SELECT TPD_SER_ST_CODIGO, TPD_BO_ALTERAPARCELAS FROM MGVEN.VEN_TIPODOCUMENTO WHERE TPD_TAB_IN_CODIGO = 155 AND TPD_PAD_IN_CODIGO = MGGLO.PCK_MEGA.Achapadraodatabela(:FIL,155,'',Sysdate) AND TPD_IN_CODIGO = :cod", pedimp.PedidoVenda.FIL_IN_CODIGO, tipoDoc)
	if erro.Trata(err1) {
		fmt.Println("Erro tipo documento: ", err1.Error())
	} else {
		for rows.Next() {
			err := rows.Scan(&tpd.Serie, &tpd.AltParc)
			if err != nil {
				log.Fatal(err)
			}
		}
		if strings.TrimSpace(serie) == "" {
			serie = tpd.Serie
		}
		fmt.Println(tpd, serie)

	}
	return &serie
}

func getUsuario(usu *int) *int {
	usuario := 1
	if usu != nil && *usu != 0 {
		usuario = *usu
	}
	return &usuario
}

func getOrganizacao(fil *int) *int {
	rows, err := dbmap.Db.Query("Select pai_org_in_Codigo From mgglo.glo_vw_organizacao Where org_in_Codigo = :fil", fil)
	var org int
	if err != nil {
		org = 0
	} else {

		for rows.Next() {
			err := rows.Scan(&org)
			if err != nil {
				org = 0
			}
		}
	}
	return &org

}
