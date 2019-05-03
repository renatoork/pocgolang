// @APIVersion 1.0.0
// @SwaggerVersion 2.0
// @APITitle MEGA API PV
// @APIDescription Serviço para incluir Pedido de Venda no Mega ERP.
// @BasePath /api/v0
package main

import (
	"flag"
	"fmt"
	"log"
	"mega/ms-consul/consul"
	"mega/ms-pedidovenda/pedidovenda"
	"mega/ms-pedidovenda/tipos"
	"net/http"
	"os"
	"strconv"
)

var conf consul.ConfigServico
var port = flag.Int("porta", 0, "Número de uma porta fixa para publicar o serviço. Se não informado o serviço define uma porta.")
var regConsul = flag.Bool("consul", false, "Define se o serviço se registra no Consul para descoberta pelo cliente. Se true, o Consul deve estar em execução para o serviço funcionar. O Defult é true.")

func main() {

	os.Setenv("NLS_LANG", ".UTF8")

	flag.Parse()
	if *port == 0 {
		*port, _ = strconv.Atoi(os.Getenv("API_PEDIDOVENDA_PORT"))
		if *port == 0 {
			*port = consul.RandomInt(2000, 5000)
		}
	}
	pedidovenda.ConsulCredito = *regConsul

	conf = carregaConfiguracao()
	if regServico(conf) {

		rt := consul.NewRouter(pedidovenda.Routes) //carrega as rotas do serviço.

		porta := strconv.Itoa(conf.Port) //encontra uma porta para publicar o serviço.

		pedidovenda.DbMap = pedidovenda.InitDb("RANA")      //cria a conexão única para o serviço.
		pedidovenda.DbMapGO = pedidovenda.InitDb("GORACLE") //cria a conexão única para o serviço.

		if pedidovenda.DbMap == nil {
			log.Fatal("Erro ao conectar ao Oracle. \nVerifique a configuração da variável de ambiente [API_PEDIDOVENDA_ORACLE]")
		}

		pedidovenda.SrvCredito = pedidovenda.GetCredito() //encontra o serviço publicado no consul.

		fs := http.FileServer(http.Dir("swagger/"))
		http.Handle("/docs/", http.StripPrefix("/docs/", fs))
		http.Handle("/api/", rt)

		fmt.Println(fmt.Sprintf("Serviço em funcionamento na porta %s.", porta))

		log.Fatal(http.ListenAndServe(":"+porta, nil))
	}
}

func regServico(conf consul.ConfigServico) bool {
	if pedidovenda.ConsulCredito {
		return consul.RegistrarServico(conf)
	}
	return true
}

func carregaConfiguracao() consul.ConfigServico {
	address := os.Getenv("API_PEDIDOVENDA_HOST")
	if address == "" {
		address = "127.0.0.1"
	}
	var c consul.ConfigServico
	c.Name = "api/" + tipos.Versao + "/pedido"
	c.Address = address
	c.Tags = []string{tipos.Versao}
	c.Port = *port
	c.Check.ID = "pdv1"
	c.Check.Http = fmt.Sprintf("http://%s:%d/api/%s/check", c.Address, c.Port, tipos.Versao)
	c.Check.Interval = "10s"
	c.Check.Name = "Pedido de Venda"
	c.Check.Timeout = "2s"
	return c
}
