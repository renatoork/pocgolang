// @SubApi Inclusão de Pedido de Venda API [/api/v0] [/pedido]
// @BasePath /api/v0
package pedidovenda

import (
	"encoding/json"
	"fmt"
	"mega/go-util/erro"
	"mega/ms-consul/consul"
	"mega/ms-pedidovenda/tipos"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func varPed(vars map[string]string) string {
	return fmt.Sprintf("Pedido: %s", vars["pedido"])
}

func varItem(vars map[string]string) string {
	return fmt.Sprintf("%s - Item: %s", varPed(vars), vars["item"])
}

func varItemObs(vars map[string]string) string {
	return fmt.Sprintf("%s - Obs: %s", varItem(vars), vars["obs"])
}

func varItemProg(vars map[string]string) string {
	return fmt.Sprintf("%s - Prog: %s", varItem(vars), vars["prog"])
}

func varItemProgEstoque(vars map[string]string) string {
	return fmt.Sprintf("%s - Estoque: %s", varItem(vars), vars["estq"])
}

func GetPedidoJson(res http.ResponseWriter, req *http.Request) {
	ret := getPedidoJson(res, req)
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.Header().Add("Access-Control-Allow-Origin", "*")
	res.Header().Add("Access-Control-Allow-Methods", "GET, OPTIONS")
	res.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	res.Write([]byte(ret))
}

// @Title SetPedido
// @Description Inclui um pedido de venda.
// @Accept  json
// @Param   pedidovenda    body    tipos.Pedido    true    "Pedido de Venda"
// @Success 200 {object} Retorno "Pedido inserido."
// @Failure 400 {object} Retorno "Serviço fora de operação."
// @Failure 500 {object} Retorno "Erro interno."
// @Resource /pedido
// @Router /pedido [post]
func SetPedido(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if origin := req.Header.Get("Origin"); origin != "" {
		res.Header().Set("Access-Control-Allow-Origin", origin)
	}
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type")
	res.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	if req.Method == "OPTIONS" {
		return
	}

	type retj struct {
		Pedido tipos.PedidoPkDB `json:"pedido,omitempty"`
		Erro   []Falha          `json:"erros,omitempty"`
	}

	var aux retj

	ret := execPedido(res, req, nil)
	aux.Erro = ret.Erros
	aux.Pedido = ret.Pedido.Pedido.PedidoPkDB

	mens, _ := json.Marshal(aux)
	res.Write(mens)
}

// @Title GetPedido
// @Description Retorna um pedido de venda. O código do pedido é composto por {organização|serie|numero}. Exemplo: /pedido/3|1|778541
// @Accept  json
// @Success 200 {object} Retorno "Pedido retornado."
// @Failure 400 {object} Retorno "Serviço fora de operação."
// @Failure 500 {object} Retorno "Erro interno."
// @Resource /pedido
// @Router /pedido/{pedido} [get]
func GetPedido(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	pk := strings.Split(vars["pedido"], "|")
	ret := ExecGetPedido(res, req, pk)
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.Header().Add("Access-Control-Allow-Origin", "*")
	res.Header().Add("Access-Control-Allow-Methods", "GET")
	res.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	mens, err := json.Marshal(ret)
	if erro.Trata(err) {
		res.Write([]byte(fmt.Sprintf("Erro ao formatar a saída dos dados do pedido de venda \n%s", err.Error())))
	}
	res.Write(mens)

}
func putPedido(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if origin := req.Header.Get("Origin"); origin != "" {
		res.Header().Set("Access-Control-Allow-Origin", origin)
	}
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type")
	res.Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	if req.Method == "OPTIONS" {
		return
	}

	type retj struct {
		Pedido tipos.PedidoPkDB `json:"pedido,omitempty"`
		Erro   []Falha          `json:"erros,omitempty"`
	}

	var aux retj

	vars := mux.Vars(req)
	pk := strings.Split(vars["pedido"], "|")
	ret := execPedido(res, req, pk)
	aux.Erro = ret.Erros
	aux.Pedido = ret.Pedido.Pedido.PedidoPkDB

	//mens, _ := json.Marshal(pdvexterno)
	mens, _ := json.Marshal(aux)
	res.Write(mens)
}

func delPedido(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if origin := req.Header.Get("Origin"); origin != "" {
		res.Header().Set("Access-Control-Allow-Origin", origin)
	}
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type")
	res.Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	if req.Method == "OPTIONS" {
		return
	}

	type retj struct {
		Pedido tipos.PedidoPkDB `json:"pedido,omitempty"`
		Erro   []Falha          `json:"erros,omitempty"`
	}

	var aux retj

	vars := mux.Vars(req)
	pk := strings.Split(vars["pedido"], "|")
	ret := execPedido(res, req, pk)
	aux.Erro = ret.Erros
	aux.Pedido = ret.Pedido.Pedido.PedidoPkDB

	//mens, _ := json.Marshal(pdvexterno)
	mens, _ := json.Marshal(aux)
	res.Write(mens)
}

func setObs(res http.ResponseWriter, req *http.Request) {
	consul.EmConstrucao(res, req, varPed(mux.Vars(req)))
}
func getObs(res http.ResponseWriter, req *http.Request) {
	consul.EmConstrucao(res, req, varPed(mux.Vars(req)))
}
func putObs(res http.ResponseWriter, req *http.Request) {
	consul.EmConstrucao(res, req, varPed(mux.Vars(req)))
}
func delObs(res http.ResponseWriter, req *http.Request) {
	consul.EmConstrucao(res, req, varPed(mux.Vars(req)))
}

func setItems(res http.ResponseWriter, req *http.Request) {
	consul.EmConstrucao(res, req, varPed(mux.Vars(req)))
}
func getItems(res http.ResponseWriter, req *http.Request) {
	consul.EmConstrucao(res, req, varPed(mux.Vars(req)))
}
func putItem(res http.ResponseWriter, req *http.Request) {
	consul.EmConstrucao(res, req, varItem(mux.Vars(req)))
}
func delItem(res http.ResponseWriter, req *http.Request) {
	consul.EmConstrucao(res, req, varItem(mux.Vars(req)))
}

func setItemObs(res http.ResponseWriter, req *http.Request) {
	consul.EmConstrucao(res, req, varItem(mux.Vars(req)))
}
func getItemObs(res http.ResponseWriter, req *http.Request) {
	consul.EmConstrucao(res, req, varItem(mux.Vars(req)))
}
func putItemObs(res http.ResponseWriter, req *http.Request) {
	consul.EmConstrucao(res, req, varItemObs(mux.Vars(req)))
}
func delItemObs(res http.ResponseWriter, req *http.Request) {
	consul.EmConstrucao(res, req, varItemObs(mux.Vars(req)))
}

func setItemProg(res http.ResponseWriter, req *http.Request) {
	consul.EmConstrucao(res, req, varItemProg(mux.Vars(req)))
}
func getItemProg(res http.ResponseWriter, req *http.Request) {
	consul.EmConstrucao(res, req, varItemProg(mux.Vars(req)))
}
func putItemProg(res http.ResponseWriter, req *http.Request) {
	consul.EmConstrucao(res, req, varItemProg(mux.Vars(req)))
}
func delItemProg(res http.ResponseWriter, req *http.Request) {
	consul.EmConstrucao(res, req, varItemProg(mux.Vars(req)))
}

func setItemProgEstoque(res http.ResponseWriter, req *http.Request) {
	consul.EmConstrucao(res, req, varItemProgEstoque(mux.Vars(req)))
}
func getItemProgEstoque(res http.ResponseWriter, req *http.Request) {
	consul.EmConstrucao(res, req, varItemProgEstoque(mux.Vars(req)))
}
func putItemProgEstoque(res http.ResponseWriter, req *http.Request) {
	consul.EmConstrucao(res, req, varItemProgEstoque(mux.Vars(req)))
}
func delItemProgEstoque(res http.ResponseWriter, req *http.Request) {
	consul.EmConstrucao(res, req, varItemProgEstoque(mux.Vars(req)))
}
