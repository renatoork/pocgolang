package main

import (
	"flag"
	"fmt"
	"log"
	"mega/ms-analisecredito/credito"
	"mega/ms-analisecredito/tipos"
	"mega/ms-consul/consul"
	"net/http"
	"os"
	"strconv"
)

var conf consul.ConfigServico
var port = flag.Int("porta", 0, "Número de uma porta fixa para publicar o serviço. Se não informado o serviço define uma porta.")
var regConsul = flag.Bool("consul", false, "Define se o serviço se registra no Consul para descoberta pelo cliente. Se true, o Consul deve estar em execução para o serviço funcionar. O Defult é true.")

func main() {
	flag.Parse()
	if *port == 0 {
		*port, _ = strconv.Atoi(os.Getenv("API_ANALISECREDITO_PORT"))
		if *port == 0 {
			*port = consul.RandomInt(2000, 5000)
		}
	}

	conf = carregaConfiguracao()
	if regServico(conf) {

		rt := consul.NewRouter(credito.Routes)

		credito.DbMap = credito.InitDb()

		porta := strconv.Itoa(conf.Port)
		fmt.Println(fmt.Sprintf("Serviço em funcionamento na porta %s.", porta))

		log.Fatal(http.ListenAndServe(":"+porta, rt))
	}
}

func carregaConfiguracao() consul.ConfigServico {
	address := os.Getenv("API_ANALISECREDITO_HOST")
	if address == "" {
		address = "127.0.0.1"
	}
	var c consul.ConfigServico
	c.Name = "api/" + tipos.Versao + "/credito"
	c.Address = address
	c.Tags = []string{tipos.Versao}
	c.Port = *port
	c.Check.ID = "cred1"
	c.Check.Http = fmt.Sprintf("http://%s:%d/api/%s/credito/check", c.Address, c.Port, tipos.Versao)
	c.Check.Interval = "10s"
	c.Check.Name = "Analise de Crédito"
	c.Check.Timeout = "2s"
	return c
}

func regServico(conf consul.ConfigServico) bool {
	if *regConsul {
		return consul.RegistrarServico(conf)
	}
	return true
}
