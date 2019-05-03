package main

import (
	"fmt"
	"log"
	"mega/ms-analisecredito/credito"
	"mega/ms-analisecredito/tipos"
	"mega/ms-consul/consul"
	"net/http"
	"strconv"
)

var conf consul.ConfigServico

func main() {

	conf := carregaConfiguracao()
	if consul.RegistrarServico(conf) {

		//carregaServicoDelphi()

		rt := consul.NewRouter(credito.Routes)

		credito.DbMap = credito.InitDb()
		credito.SrvDelphi = consul.GetServico("credito/json")

		porta := strconv.Itoa(conf.Port)
		fmt.Println(fmt.Sprintf("Serviço em funcionamento na porta %s.", porta))

		log.Fatal(http.ListenAndServe(":"+porta, rt))
	}
}

func carregaConfiguracao() consul.ConfigServico {
	var c consul.ConfigServico
	c.Name = "api/" + tipos.Versao + "/credito"
	c.Address = "127.0.0.1"
	c.Tags = []string{tipos.Versao}
	c.Port = consul.RandomInt(2000, 5000)
	c.Check.ID = "cred1"
	c.Check.Http = fmt.Sprintf("http://%s:%d/api/%s/credito/check", c.Address, c.Port, tipos.Versao)
	c.Check.Interval = "10s"
	c.Check.Name = "Analise de Crédito"
	c.Check.Timeout = "2s"
	return c
}

func carregaServicoDelphi() {
	var c consul.ConfigServico
	c.Name = "api/" + tipos.Versao + "/creditodelphi"
	c.Address = "PC_ALAIR"
	c.Tags = []string{tipos.Versao}
	c.Port = 7099
	c.Check.ID = "credDelphi"
	c.Check.Http = fmt.Sprintf("http://%s:%d/api/%s/creditodelphi", c.Address, c.Port, tipos.Versao)
	c.Check.Interval = "10s"
	c.Check.Name = "Analise de Crédito Delphi"
	c.Check.Timeout = "2s"

	consul.RegistrarServico(conf)

}
