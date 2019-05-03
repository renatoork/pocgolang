package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	_ "io"
	"io/ioutil"
	"mega/go-util/dbg"
	"mega/go-util/erro"
	"mega/ms-consul/consul"
	"mega/ms-pedidovenda/pedidovenda"
	"mega/ms-pedidovenda/tipos"
	"net/http"
	"os"
	//"runtime/pprof"
	"strconv"
	"strings"
	"sync"
	"time"
)

var srv *consul.Servico
var nomeSrv string

var numCPU int
var orig bool

func main() {

	//f, err := os.Create("teste.pprof")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//pprof.StartCPUProfile(f)
	//defer pprof.StopCPUProfile()

	var errX error
	numCPU = 1
	if len(os.Args) > 1 {
		numCPU, errX = strconv.Atoi(os.Args[1])
		if errX != nil {
			numCPU = 1
		}
	}
	srv = consul.GetServico("api/v0/pedido")
	if srv == nil || len(srv.Servico) == 0 {
		fmt.Println(`  - ERRO: o serviço "PedidoVenda" não existe ou não esta em funcionamento.`)
	} else {
		nomeSrv = srv.Servico[0].ServiceName
		execS := "nenhum"
		if len(os.Args) > 2 {
			execS = os.Args[2]
		}
		orig = false
		if len(os.Args) > 3 && strings.ToLower(os.Args[3]) == "arq" {
			orig = true
		}

		if execS == "nenhum" {
			fmt.Println("\nGO REST - Padrão: ")
			setPedido("")
		}

		if execS == "todos" || execS == "semcredito" {
			fmt.Println("\nGO REST- Sem Analise Credito: ")
			setPedido("semcredito")
		}
		if execS == "todos" || execS == "comcreditopl" {
			fmt.Println("\nGO REST- Com Analise Credito via PL: ")
			setPedido("comcreditopl")
		}
		if execS == "todos" || execS == "comcreditosrv" {
			fmt.Println("\nGO REST- Com Analise Credito via Servico PL ")
			setPedido("comcreditosrv")
		}
		if execS == "todos" || execS == "comcreditodelphi" {
			fmt.Println("\nGO REST- Com Analise Credito via Servico Delphi ")
			setPedido("comcreditodelphi")
		}
		if execS == "todos" || execS == "comcreditogo" {
			fmt.Println("\nGO REST- Com Analise Credito via Servico GO ")
			setPedido("comcreditogo")
		}

	}
}

func setPedido(tiposervico string) {
	dir := "..\\gerajsonpedido\\arquivo"
	if orig {
		dir = "..\\gerajsonpedido\\arquivo\\Wine"
	}

	if tiposervico != "" {
		srv.Servico[0].ServiceName = fmt.Sprintf("%s/%s", nomeSrv, tiposervico)
	}

	var inicio time.Time

	tasks := make(chan string, 100)
	var wg sync.WaitGroup
	for i := 0; i < numCPU; i++ {
		wg.Add(1)
		go func(cpu int) {
			for nomearq := range tasks {
				arq := fmt.Sprintf("%s\\%s", dir, nomearq)
				body := lerArquivo(arq)
				inicio = time.Now()
				executaServico(srv, "application/json; param=value", "POST", body, fmt.Sprintf("CPU:%d | %s", cpu, arq[1:25]))
				dbg.Trace(inicio)
			}
			wg.Done()
		}(i)
	}

	readerDir, _ := ioutil.ReadDir(dir)
	for _, fileInfo := range readerDir {
		if !fileInfo.IsDir() {
			tasks <- fileInfo.Name()
		}
	}

	close(tasks)
	wg.Wait()

}

func executaServico(srv *consul.Servico, content string, metodo string, body string, msg string) error {
	os.Setenv("NLS_LANG", ".AL32UTF8")

	url, _ := consul.GeraURLPost("http",
		srv.Servico[0].ServiceAddress,
		strconv.Itoa(srv.Servico[0].ServicePort),
		srv.Servico[0].ServicePath,
		srv.Servico[0].ServiceName,
		[][]string{})

	req, err := http.NewRequest(metodo, url, bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", content)

	client := &http.Client{}
	resp, err := client.Do(req)
	if erro.Trata(err) {
		fmt.Println(msg, err.Error())
		return err
	}
	defer resp.Body.Close()

	fmt.Print(msg, resp.Status, ";")
	r, _ := ioutil.ReadAll(resp.Body)

	var ret pedidovenda.Retorno
	json.Unmarshal(r, &ret)
	//fmt.Println(" - ", ret.Erros)
	if len(ret.Erros) > 0 {
		return errors.New("erro")
	}

	return nil

}

func getUrl(nomeSrv string) string {
	return fmt.Sprintf("/api/%s/%s", tipos.Versao, nomeSrv)
}

func lerArquivo(nomeArq string) string {
	arq, _ := ioutil.ReadFile(fmt.Sprintf("../testes/%s", nomeArq))
	return string(arq)
}
