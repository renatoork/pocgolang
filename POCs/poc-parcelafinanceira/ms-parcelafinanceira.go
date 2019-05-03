package main

import (
	"fmt"
	"log"
	"mega/ms-consul/consul"
	"mega/ms-parcelafinanceira/parcela"
	"mega/ms-parcelafinanceira/tipos"
	"net/http"
	"strconv"
)

var conf consul.ConfigServico

func main() {

	conf := carregaConfiguracao()
	if consul.RegistrarServico(conf) {

		rt := consul.NewRouter(parcela.Routes)

		porta := strconv.Itoa(conf.Port)
		fmt.Println(fmt.Sprintf("Serviço em funcionamento na porta %s.", porta))

		log.Fatal(http.ListenAndServe(":"+porta, rt))
	}
}

func carregaConfiguracao() consul.ConfigServico {
	var c consul.ConfigServico
	c.Name = "api/" + tipos.Versao + "/parcela"
	c.Address = "127.0.0.1"
	c.Tags = []string{tipos.Versao}
	c.Port = consul.RandomInt(2000, 5000)
	c.Check.ID = "parc1"
	c.Check.Http = fmt.Sprintf("http://%s:%d/api/%s/parcela/check", c.Address, c.Port, tipos.Versao)
	c.Check.Interval = "10s"
	c.Check.Name = "Parcelas Financeiras"
	c.Check.Timeout = "2s"
	return c
}
