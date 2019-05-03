package main

import (
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

func main() {

	os.Setenv("NLS_LANG", ".AL32UTF8")

	conf := carregaConfiguracao()
	if consul.RegistrarServico(conf) {

		rt := consul.NewRouter(pedidovenda.Routes)

		porta := strconv.Itoa(conf.Port)
		fmt.Println(fmt.Sprintf("Serviço em funcionamento na porta %s.", porta))

		pedidovenda.DbMap = pedidovenda.InitDb()
		pedidovenda.SrvCredito = consul.GetServico("api/v0/credito")

		if pedidovenda.DbMap == nil {
			log.Fatal("Erro ao conectar ao Oracle. \nVerifique a configuração da variável de ambiente [ORACLE_MS_PEDIDOVENDA]")
		}
		log.Fatal(http.ListenAndServe(":"+porta, rt))
	}
}

func carregaConfiguracao() consul.ConfigServico {
	var c consul.ConfigServico
	c.Name = "api/" + tipos.Versao + "/pedido"
	c.Address = "127.0.0.1"
	c.Tags = []string{tipos.Versao}
	c.Port = consul.RandomInt(2000, 5000)
	c.Check.ID = "pdv1"
	c.Check.Http = fmt.Sprintf("http://%s:%d/api/%s/check", c.Address, c.Port, tipos.Versao)
	c.Check.Interval = "10s"
	c.Check.Name = "Pedido de Venda"
	c.Check.Timeout = "2s"
	return c
}
