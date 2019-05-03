package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/tomazk/envcfg"
	"io"
	"io/ioutil"
	"mega/go-util/erro"
	"net/http"
	"os"
	"strings"
)

type Config struct {
	Diretorio           string `json:"diretorio"`           //"D:\DESENV\git\go\GoBamboo\Arquivos"
	NomeRotinaExecuta   string `json:"nomerotinaexecuta"`   //mgtst.ven_pck_pedidovendatst.p_executactpedido
	NomeRotinaResultado string `json:"nomerotinaresultado"` //mgtst.ven_pck_pedidovendatst.f_ResultadoCTPedidoPipe
	ConnString          string `json:"oracle"`
	Servico             string `json:"servico"` //true - server http || false - executavel
	ServicoAddress      string `json:"servicoaddress"`
	ConfigServico       struct {
		Address string `json:"Address"`
		Check   struct {
			Http     string `json:"http"`
			ID       string `json:"id"`
			Interval string `json:"interval"`
			Name     string `json:"name"`
			Timeout  string `json:"timeout"`
		} `json:"Check"`
		Name string   `json:"Name"`
		Port int      `json:"Port"`
		Tags []string `json:"Tags"`
	} `json:"configServico"`
}

var conf Config

func main() {
	carregarConfiguracao(os.Args, &conf)
	fmt.Println(conf.Servico)
	if strings.ToUpper(conf.Servico) == "FALSE" {
		ExecutaTestesPedido()
	} else {
		if ok := registrarServico(); ok {
			ExecutaServico()
		}
	}
}

func ExecutaServico() {
	fmt.Println("Serviço <gerapedido:5000> em funcionamento...")
	http.HandleFunc("/gerapedido", executaTestesPedidoSrv)
	http.HandleFunc("/teste", executaTestesServico)
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		panic(err)
	}
}

func executaTestesServico(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("ok"))
}

func executaTestesPedidoSrv(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte(ExecutaTestesPedido()))
}

func ExecutaTestesPedido() string {
	var x string
	if conf.Diretorio == "" {
		x = executarTesteSimples()

	} else {
		x = executarTesteArquivo(conf.Diretorio)
	}
	return x
}

func carregarConfiguracao(args []string, resp *Config) {
	msgErro := ""
	tam := len(args)
	if tam == 1 {
		msgErro = fmt.Sprintf("Falha nos parâmetros da aplicação. \n Parametros possíveis: \n   <config> - carregar parametros do arquivo Config.json \n  ou \n  <connectstring> - string com a connect string de conexão com o Oracle (usuario/senha@host:porta/instancia)\n   <nomeRotinaExecucao> - string com o nome da rotina PLSQL de execucao dos testes (mgtst.ven_pck_pedidovendatst.p_executactpedido) \n  <nomeRotinaResultado> - string com o nome da rotina plsql que trata o cursor com as mensagens de retono da execução (mgtst.ven_pck_pedidovendatst.f_ResultadoCTPedidoPipe) \nExemplo: TestePLSQL config \n     ou  TestePLSQL mgtst/megatst@pc_renato:1521/orcl mgtst.ven_pck_pedidovendatst.p_executactpedido mgtst.ven_pck_pedidovendatst.f_ResultadoCTPedidoPipe\n")
	} else {
		if strings.Contains(strings.ToUpper(args[1]), "CONFIG") {
			conf, err := ioutil.ReadFile("config.json")
			if err == nil {
				errj := json.Unmarshal(conf, resp)

				if errj != nil {
					msgErro = fmt.Sprintf("Erro ao carregar as configurações do arquivo Config.json - %s", errj.Error())
				}
			} else {
				msgErro = fmt.Sprintf("Erro ao carregar as configuraçõesdo arquivo Config.json - ", err.Error())
			}
		} else {
			if tam < 4 {
				msgErro = fmt.Sprintf("Falha nos parâmetros da aplicação. \n Parametros possíveis: \n   <config> - carregar parametros do arquivo Config.json \n  ou \n  <connectstring> - string com a connect string de conexão com o Oracle (usuario/senha@host:porta/instancia)\n   <nomeRotinaExecucao> - string com o nome da rotina PLSQL de execucao dos testes (mgtst.ven_pck_pedidovendatst.p_executactpedido) \n  <nomeRotinaResultado> - string com o nome da rotina plsql que trata o cursor com as mensagens de retono da execução (mgtst.ven_pck_pedidovendatst.f_ResultadoCTPedidoPipe) \n <servico> - string 'true' para rodar os testes do pedido como serviço ou 'false' para rodar como executável. \nExemplo: TestePLSQL config \n     ou  TestePLSQL mgtst/megatst@pc_renato:1521/orcl mgtst.ven_pck_pedidovendatst.p_executactpedido mgtst.ven_pck_pedidovendatst.f_ResultadoCTPedidoPipe false\n")
			} else {
				resp.ConnString = args[1]
				resp.NomeRotinaExecuta = args[2]
				resp.NomeRotinaResultado = args[3]
				resp.Servico = args[4]
			}
		}
	}

	if msgErro != "" {
		fmt.Println(os.Stderr, msgErro)
		os.Exit(1)
	}

}

func geraResultXML(ret TestSuite) string {

	filename := "Result.xml"
	file, _ := os.Create(filename)

	xmlWriter := io.Writer(file)

	enc := xml.NewEncoder(xmlWriter)
	enc.Indent("  ", "    ")
	if err := enc.Encode(ret); err != nil {
		fmt.Printf("error: %v\n", err)
	}
	return fmt.Sprintf("%s", ret.TestCases)
}

func executarTesteSimples() string {
	return geraResultXML(executarCTPedido(conf))
}

func executarTesteArquivo(caminho string) string {

	return executaTestes(carregarArquivosTeste(caminho))

	//geraResultXML(executarCTPedidoArq())
	// rotina 2
	// carregar arquivo de configuração com: conn ao Oracle, Diretorio com arquivos de teste
	// Carregar arquivos a executar.
	// Para cada arquivo:
	//   carregar parametros
	//   executar teste
	//   Comparar Resultado
	//   logar resultado.
}

func carregarArquivosTeste(caminho string) []string {
	var testes []string
	files, _ := ioutil.ReadDir(caminho)
	for _, f := range files {
		if !f.IsDir() && strings.Contains(f.Name(), ".json") {
			testes = append(testes, fmt.Sprintf("%s\\%s", caminho, f.Name()))
		}
	}
	return testes
}

func executaTestes(arqs []string) string {
	for _, val := range arqs {
		var ct Casosteste

		file, err := ioutil.ReadFile(val)
		if err == nil {
			json.Unmarshal(file, &ct)
		}

		if ct.CasoTeste.Nome != "" {

			//var entrada map[string]interface{}
			//entrada = entrada.(map[string]interface{})
			//fmt.Println("entrada: ", entrada)
			fmt.Println("entrada: ")
			//field := reflect.ValueOf(ct.CasoTeste.Entrada)
			//fmt.Println("Field: ", field)

			//for i, v1 := range field.MapKeys() {
			//	fmt.Println("i: ", i)
			//	fmt.Println("v1: ", v1)
			//}

			//fmt.Printf("Results: %s \n %s\n", val, ct)
		}
	}
	return "ok."
}

type EnvConsul struct {
	CONSUL_HOST string
	CONSUL_PORT string
}

var configconsul EnvConsul

func registrarServico() bool {

	envcfg.Unmarshal(&configconsul)
	if configconsul.CONSUL_HOST == "" || configconsul.CONSUL_PORT == "" {
		fmt.Println("Erro: configure as variáveis de ambiente:\n - CONSUL_HOST: endereco IP do servidor Consul.\n - CONSUL_PORT: porta do servico HTTP do Consul.\n")
		return false
	}

	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%s", configconsul.CONSUL_HOST, configconsul.CONSUL_PORT)
	client, err := api.NewClient(cfg)
	if erro.Trata(err) {
		fmt.Println(fmt.Sprintf("Problema ao registrar o servico de Pedido: %s \n Configuração = %s", err.Error(), cfg))
		return false
	}

	agent := client.Agent()

	reg := &api.AgentServiceRegistration{
		Name:    conf.ConfigServico.Name,
		Tags:    conf.ConfigServico.Tags,
		Port:    conf.ConfigServico.Port,
		Address: conf.ConfigServico.Address,
		Check: &api.AgentServiceCheck{
			TTL:     conf.ConfigServico.Check.Interval,
			Timeout: conf.ConfigServico.Check.Timeout,
			HTTP:    conf.ConfigServico.Check.Http,
		},
	}
	if err = agent.ServiceRegister(reg); erro.Trata(err) {
		fmt.Println(fmt.Sprintf("Problema ao registrar o servico de Pedido: %s \n Configuração = %s", err.Error(), reg))
		return false
	}

	services, err := agent.Services()
	if erro.Trata(err) {
		return false
	}

	if _, ok := services["pedidovenda"]; !ok {
		fmt.Println("servico não iniciado ou não registrado")
		return false
	} else {
		fmt.Println("Servico Registrado: ", services["pedidovenda"])
	}
	return true
}

func descobrirServicoPedido() {

}
