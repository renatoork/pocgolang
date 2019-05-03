package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	_ "github.com/mattn/go-oci8" // utilizado no gorp.in
	"io"
	_ "io"
	"io/ioutil"
	_ "log"
	"mega/go-comparastruct/compara"
	"mega/go-util/dbg"
	"mega/go-util/erro"
	"mega/ms-consul/consul"
	"mega/ms-pedidovenda/pedidovenda"
	"mega/ms-pedidovenda/tipos"
	"net/http"
	"net/http/httptest"
	"os"
	//"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
)

var srv *consul.Servico
var nomeSrv string

var numCPU int
var orig bool

var cpu = flag.Int("cpu", 0, "Número de CPU paralelas consumindo o serviço de pedido. Default 1 .")
var qtde = flag.Int("qtde", 0, "Quantidade de pedidos a testar. Exemplo: 100 - testa os ultimos 100 pedidos. Default 1.")
var datainicio = flag.String("inicio", "", "Data inicial (dd/mm/yyyy) do período de teste. ")
var datafinal = flag.String("final", "", "Data final (dd/mm/yyyy) do período de teste. ")
var pedido = flag.Int("pedido", 0, "Número do pedido a testar.")

func main() {

	os.Setenv("NLS_LANG", ".AL32UTF8")
	os.Setenv("API_PEDIDOVENDA_HOST", "127.0.0.1")
	os.Setenv("API_PEDIDOVENDA_PORT", "8080")
	os.Setenv("API_PEDIDOVENDA_ORACLE", "MGVEN/MEGAVEN@PC_LEANDROA:1521/WINE")
	os.Setenv("API_ANALISECREDITO_HOST", "127.0.0.1")
	os.Setenv("API_ANALISECREDITO_PORT", "8081")
	os.Setenv("API_ANALISECREDITO_ORACLE", "MGVEN/MEGAVEN@PC_LEANDROA:1521/WINE")

	flag.Parse()
	if *cpu == 0 {
		numCPU = 1
	} else {
		numCPU = *cpu
	}

	srv = getServico()

	if srv == nil || len(srv.Servico) == 0 {
		fmt.Println(`  - ERRO: o serviço "PedidoVenda" não existe ou não esta em funcionamento.`)
	} else {

		var inicio time.Time

		tasks := make(chan []string, 100)
		var wg sync.WaitGroup
		for i := 0; i < numCPU; i++ {
			wg.Add(1)
			go func(cpu int) {
				for pk := range tasks {

					body, sBody := getPedido(pk)
					if body != "" {
						inicio = time.Now()
						_, sRetorno := executaServico(srv, "application/json; param=value", "POST", body, fmt.Sprintf("CPU:%d | %s", cpu, pk))
						fmt.Print("CPU: " + strconv.Itoa(cpu) + " - " + "Executado serviço em: ")
						dbg.Trace(inicio)
						assertPedido(sBody, sRetorno)
					} else {
						fmt.Println(fmt.Sprintf("  CPU: "+strconv.Itoa(cpu)+" - "+"Pedido %d não encontrado\n", pk[2]))
					}
				}
				wg.Done()
			}(i)
		}

		//ler pedidos da base
		connString := os.Getenv("API_PEDIDOVENDA_ORACLE")
		db, errb := sql.Open("oci8", connString)
		if erro.Trata(errb) {
			panic(fmt.Sprintf("ERRO ao conectar ao Oracle - %s", errb.Error()))
		} else {
			var (
				rows *sql.Rows
				err  error
			)
			if *qtde > 0 {
				rows, err = db.Query("SELECT * FROM (SELECT TO_CHAR(ORG_IN_CODIGO), SER_ST_CODIGO, TO_CHAR(PED_IN_CODIGO) PK FROM MGVEN.VEN_PEDIDOVENDA ORDER BY PED_DT_EMISSAO DESC) WHERE rownum <= :qtde", *qtde)
			} else if *pedido > 0 {
				rows, err = db.Query("SELECT TO_CHAR(ORG_IN_CODIGO), SER_ST_CODIGO, TO_CHAR(PED_IN_CODIGO) PK FROM MGVEN.VEN_PEDIDOVENDA WHERE ped_in_codigo = :pedido", *pedido)
			}
			if erro.Trata(err) {
				fmt.Println(fmt.Sprintf("  ERRO ao recuperar os pedidos de venda - %s", err.Error()))
			} else if rows == nil {
				fmt.Println("  Pedido não informado.")
			} else {

				defer rows.Close()
				var (
					org string
					ser string
					cod string
				)
				for rows.Next() {
					err := rows.Scan(&org, &ser, &cod)
					if err != nil {
						panic(fmt.Sprintf("  ERRO ao recuperar os pedidos de venda - %s", err.Error()))
					}
					tasks <- []string{org, ser, cod}
				}
			}
		}
		close(tasks)
		wg.Wait()
	}
}

func mockServico(metodo string, url string, contentType string, body io.Reader) (*httptest.ResponseRecorder, *http.Request) {
	r, _ := http.NewRequest(metodo, url, body)
	w := httptest.NewRecorder()
	r.Header.Set("Content-Type", contentType)
	w.Code = 400
	return w, r
}

func getPedido(pk []string) (string, tipos.PedidoVendaDB) {
	w, r := mockServico("GET", getUrl("pedido"), "", nil)
	pedidovenda.DbMap = pedidovenda.InitDb("RANA") //cria a conexão única para o serviço.

	ret := pedidovenda.ExecGetPedido(w, r, pk)

	mensDB, _ := json.Marshal(ret.Pedido)

	var ret_ tipos.Pedido
	json.Unmarshal(mensDB, &ret_)

	mens, _ := json.Marshal(ret_)

	return string(mens), ret.Pedido.Pedido
}

type Pedido struct {
	PedidoVenda tipos.PedidoUk `json:"pedido"`
}

func executaServico(srv *consul.Servico, content string, metodo string, body string, msg string) (string, tipos.PedidoVendaDB) {
	os.Setenv("NLS_LANG", ".AL32UTF8")

	jBody := strings.NewReader(body)

	w, r := mockServico("POST", getUrl("pedido"), "application/json; param=value", jBody)

	pedidovenda.SetPedido(w, r)

	rBody, err := ioutil.ReadAll(w.Body)
	if erro.Trata(err) {
		fmt.Println(err.Error())
		return "", tipos.PedidoVendaDB{}
	}
	var pk Pedido
	json.Unmarshal(rBody, &pk)

	if pk.PedidoVenda.Codigo != nil {
		newbody, ret_ := getPedido([]string{strconv.Itoa(*pk.PedidoVenda.Organizacao), pk.PedidoVenda.Serie, strconv.Itoa(*pk.PedidoVenda.Codigo)})
		return newbody, ret_
	} else {
		nome := "erro" + time.Now().String() + ".txt"
		fl, _ := os.Create(nome)
		fl.WriteString(string(rBody))
		fl.Close()
		fmt.Println(fmt.Sprintf("  Erro ao gerar o pedido. Arquivo %s gerado.", nome))

		return "", tipos.PedidoVendaDB{}
	}
}

func assertPedido(jGet tipos.PedidoVendaDB, jRet tipos.PedidoVendaDB) {
	var cont string
	igual, diferencas := compara.ComparaStruct(&jGet, &jRet)
	if igual {
		fmt.Println(fmt.Sprintf("  Pedidos: %d == %d", *jGet.Numero, *jRet.Numero))
	} else {
		if jGet.Numero != nil && jRet.Numero != nil {
			cont = fmt.Sprintf("  Pedidos: %d != %d \n%s", *jGet.Numero, *jRet.Numero, diferencas)
			nome := fmt.Sprintf("%d_%d.txt", *jGet.Numero, *jRet.Numero)
			fl, errA := os.Create(nome)
			fl.WriteString(cont)
			fl.Close()
			if erro.Trata(errA) {
				fmt.Println(errA.Error())
			}
			fmt.Println(fmt.Sprintf("  Existem diferenças. Arquivo %s gerado.", nome))
		} else {
			cont = fmt.Sprintf("  Erro ao gerar o pedido: %d \n", *jGet.Numero)
		}
	}
}

func getUrl(nomeSrv string) string {
	return fmt.Sprintf("/api/%s/%s", tipos.Versao, nomeSrv)
}

func getServico() *consul.Servico {
	pathServico := os.Getenv("API_PEDIDOVENDA_PATH")
	if pathServico == "" {
		pathServico = "avi/" + tipos.Versao
	}
	nomeServico := os.Getenv("API_PEDIDOVENDA_NOME")
	if nomeServico == "" {
		nomeServico = "pedido"
	}
	hostServico := os.Getenv("API_PEDIDOVENDA_HOST")
	if hostServico == "" {
		hostServico = "127.0.0.1"
	}
	portServico := os.Getenv("API_PEDIDOVENDA_PORT")
	if portServico == "" {
		portServico = "8080"
	}
	var srv consul.Servico
	if hostServico != "" && portServico != "" {
		var s consul.Srv
		s.ServiceAddress = hostServico
		s.ServicePort, _ = strconv.Atoi(portServico)
		s.ServicePath = pathServico
		s.ServiceName = nomeServico
		srv.Servico = append(srv.Servico, s)
		return &srv
	} else {
		return nil
	}
}
