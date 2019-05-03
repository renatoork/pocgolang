package pedidovenda

import (
	"mega/ms-consul/consul"
	"mega/ms-pedidovenda/tipos"
)

type routes []consul.Route

const servuri = "/api/%s/pedido"

var Routes = routes{
	consul.Route{
		"CheckServicoPedido",
		[]string{"GET"},
		tipos.Versao,
		servuri + "/check",
		consul.CheckServico,
	},
	consul.Route{
		"GetPedidoJson",
		[]string{"GET"},
		tipos.Versao,
		servuri + "/pedido.json",
		GetPedidoJson,
	},
	consul.Route{
		"SetPedido",
		[]string{"POST", "OPTIONS"},
		tipos.Versao,
		"/api/%s/pedido",
		SetPedido,
	},
	consul.Route{
		"GetPedido",
		[]string{"GET"},
		tipos.Versao,
		"/api/%s/pedido/{pedido}", //"/api/%s/pedido/{pedido:[0-9]+}",
		GetPedido,
	},
	consul.Route{
		"PutPedido",
		[]string{"PUT"},
		tipos.Versao,
		"/api/%s/pedido/{pedido}", //"/api/%s/pedido/{pedido:[0-9]+}",
		putPedido,
	},
	consul.Route{
		"DelPedido",
		[]string{"DELETE"},
		tipos.Versao,
		"/api/%s/pedido/{pedido}", //"/api/%s/pedido/{pedido:[0-9]+}",
		delPedido,
	},
	consul.Route{
		"SetObsPedido",
		[]string{"POST"},
		tipos.Versao,
		"/api/%s/pedido/{pedido:[0-9]+}/obs",
		setObs,
	},
	consul.Route{
		"GetObsPedido",
		[]string{"GET"},
		tipos.Versao,
		"/api/%s/pedido/{pedido:[0-9]+}/obs",
		getObs,
	},
	consul.Route{
		"PutObsPedido",
		[]string{"POST"},
		tipos.Versao,
		"/api/%s/pedido/{pedido:[0-9]+}/obs/{tipo: [a-z]}",
		putObs,
	},
	consul.Route{
		"DelObsPedido",
		[]string{"DELETE"},
		tipos.Versao,
		"/api/%s/pedido/{pedido:[0-9]+}/obs/{tipo: [a-z]}",
		delObs,
	},

	consul.Route{
		"SetItemPedido",
		[]string{"POST"},
		tipos.Versao,
		"/api/%s/pedido/{pedido:[0-9]+}/items",
		setItems,
	},
	consul.Route{
		"GetItemsPedido",
		[]string{"GET"},
		tipos.Versao,
		"/api/%s/pedido/{pedido:[0-9]+}/items",
		getItems,
	},

	consul.Route{
		"PutItemPedido",
		[]string{"POST"},
		tipos.Versao,
		"/api/%s/pedido/{pedido:[0-9]+}/item/{item: [0-9]+}",
		putItem,
	},

	consul.Route{
		"DelItemPedido",
		[]string{"DELETE"},
		tipos.Versao,
		"/api/%s/pedido/{pedido:[0-9]+}/item/{item: [0-9]+}",
		delItem,
	},
	consul.Route{
		"SetItemObsPedido",
		[]string{"POST"},
		tipos.Versao,
		"/api/%s/pedido/{pedido:[0-9]+}/item/{item: [0-9]+}/obs",
		setItemObs,
	},
	consul.Route{
		"GetItemObsPedido",
		[]string{"GET"},
		tipos.Versao,
		"/api/%s/pedido/{pedido:[0-9]+}/item/{item: [0-9]+}/obs",
		getItemObs,
	},
	consul.Route{
		"PutItemObsPedido",
		[]string{"POST"},
		tipos.Versao,
		"/api/%s/pedido/{pedido:[0-9]+}/item/{item: [0-9]+}/obs/{tipo: [a-z]+}",
		putItemObs,
	},
	consul.Route{
		"DelItemObsPedido",
		[]string{"DELETE"},
		tipos.Versao,
		"/api/%s/pedido/{pedido:[0-9]+}/item/{item: [0-9]+}/obs/{tipo: [a-z]+}",
		delItemObs,
	},
	consul.Route{
		"SetItemProgPedido",
		[]string{"POST"},
		tipos.Versao,
		"/api/%s/pedido/{pedido:[0-9]+}/item/{item: [0-9]+}/prog",
		setItemProg,
	},
	consul.Route{
		"GetItemProgPedido",
		[]string{"GET"},
		tipos.Versao,
		"/api/%s/pedido/{pedido:[0-9]+}/item/{item: [0-9]+}/prog",
		getItemProg,
	},
	consul.Route{
		"PutItemProgPedido",
		[]string{"POST"},
		tipos.Versao,
		"/api/%s/pedido/{pedido:[0-9]+}/item/{item: [0-9]+}/prog/{prog: [0-9]+}",
		putItemProg,
	},
	consul.Route{
		"DelItemProgPedido",
		[]string{"DELETE"},
		tipos.Versao,
		"/api/%s/pedido/{pedido:[0-9]+}/item/{item: [0-9]+}/prog/{prog: [0-9]+}",
		delItemProg,
	},
	consul.Route{
		"SetItemProgEstoquePedido",
		[]string{"POST"},
		tipos.Versao,
		"/api/%s/pedido/{pedido:[0-9]+}/item/{item: [0-9]+}/prog/{prog: [0-9]+}/estoque",
		setItemProgEstoque,
	},
	consul.Route{
		"GetItemProgEstoquePedido",
		[]string{"GET"},
		tipos.Versao,
		"/api/%s/pedido/{pedido:[0-9]+}/item/{item: [0-9]+}/prog/{prog: [0-9]+}/estoque",
		getItemProgEstoque,
	},
	consul.Route{
		"PutItemProgEstoquePedido",
		[]string{"POST"},
		tipos.Versao,
		"/api/%s/pedido/{pedido:[0-9]+}/item/{item: [0-9]+}/prog/{prog: [0-9]+}/estoque/{estq: [0-9]+}",
		putItemProgEstoque,
	},
	consul.Route{
		"DelItemProgEstoquePedido",
		[]string{"DELETE"},
		tipos.Versao,
		"/api/%s/pedido/{pedido:[0-9]+}/item/{item: [0-9]+}/prog/{prog: [0-9]+}/estoque/{estq: [0-9]+}",
		delItemProgEstoque,
	},
}
