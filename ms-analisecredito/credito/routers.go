package credito

import (
	"mega/ms-analisecredito/tipos"
	"mega/ms-consul/consul"
)

type routes []consul.Route

var Routes = routes{
	consul.Route{
		"CheckServicoCredito",
		[]string{"GET"},
		tipos.Versao,
		"/api/%s/credito/check",
		consul.CheckServico,
	},
	consul.Route{
		"Credito",
		[]string{"POST"},
		tipos.Versao,
		"/api/%s/credito",
		GetAnaliseCredito,
	},
}
