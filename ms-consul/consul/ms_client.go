package consul

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hashicorp/consul/api"
	"github.com/tomazk/envcfg"
	"html"
	"io/ioutil"
	"log"
	"math/rand"
	"mega/go-util/erro"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var config EnvConsul

type ConfigServico struct {
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
}

type EnvConsul struct {
	CONSUL_HOST string
	CONSUL_PORT string
}

func DefaultURL() *url.URL {
	return &url.URL{
		Scheme:   "http",
		Opaque:   "",
		Host:     "",
		Path:     "",
		RawQuery: "",
		Fragment: "",
	}
}

type Servico struct {
	Servico []Srv `json:"servico"`
}
type Srv struct {
	Address        string   `json:"Address"`
	Node           string   `json:"Node"`
	ServiceAddress string   `json:"ServiceAddress"`
	ServiceID      string   `json:"ServiceID"`
	ServiceName    string   `json:"ServiceName"`
	ServicePort    int      `json:"ServicePort"`
	ServiceTags    []string `json:"ServiceTags"`
	ServicePath    string   `json:"ServicePath"`
}

const pathConsul = "/v1/catalog/service"

func RegistrarServico(conf ConfigServico) bool {

	var configconsul EnvConsul
	envcfg.Unmarshal(&configconsul)

	if configconsul.CONSUL_HOST == "" || configconsul.CONSUL_PORT == "" {
		fmt.Println("Erro: configure as variáveis de ambiente:\n - CONSUL_HOST: endereco IP do servidor Consul.\n - CONSUL_PORT: porta do servico HTTP do Consul.\n")
		return false
	}

	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%s", configconsul.CONSUL_HOST, configconsul.CONSUL_PORT)
	client, err := api.NewClient(cfg)
	if erro.Trata(err) {
		fmt.Println(fmt.Sprintf("Problema ao registrar o servico %s: %s \n Configuração = %s", conf.Name, err.Error(), cfg))
		return false
	}

	agent := client.Agent()

	reg := &api.AgentServiceRegistration{
		Name:    conf.Name,
		Tags:    conf.Tags,
		Port:    conf.Port,
		Address: conf.Address,
		Check: &api.AgentServiceCheck{
			TTL:     conf.Check.Interval,
			Timeout: conf.Check.Timeout,
			HTTP:    conf.Check.Http,
		},
	}
	if err = agent.ServiceRegister(reg); erro.Trata(err) {
		fmt.Println(fmt.Sprintf("Problema ao registrar o servico %s: %s \n Configuração = %s", conf.Name, err.Error(), reg))
		return false
	}

	services, err := agent.Services()
	if erro.Trata(err) {
		return false
	}

	if _, ok := services[conf.Name]; !ok {
		fmt.Println(fmt.Sprintf("Servico %s não iniciado ou não registrado.", conf.Name))
		return false
	} else {
		fmt.Println(fmt.Sprintf("Servico %s registrado.", conf.Name))
	}
	return true
}

func RandomInt(min, max int) int {

	//return 5001

	randInt := func(min int, max int) int {
		return min + rand.Intn(max-min)
	}

	rand.Seed(time.Now().Unix())
	return randInt(min, max)
}

func GeraURLGet(scheme, host, port, path, servico string, parametros [][]string) string {

	u := DefaultURL()
	u.Scheme = scheme
	u.Host = fmt.Sprintf("%s:%s", host, port)
	u.Path = fmt.Sprintf("%s/%s", path, servico)
	if len(parametros) > 0 {
		q := u.Query()
		for _, v := range parametros {
			q.Set(v[0], v[1])
		}
		u.RawQuery = q.Encode()
	}
	return fmt.Sprint(u)
}

func GeraURLPost(scheme, host, port, path, servico string, parametros [][]string) (string, url.Values) {
	u := DefaultURL()
	u.Scheme = scheme
	u.Host = fmt.Sprintf("%s:%s", host, port)
	if path != "" {
		u.Path = path
	}
	if servico != "" {
		if path != "" {
			u.Path = u.Path + "/"
		}
		u.Path = u.Path + servico
	}
	u.Path = fmt.Sprintf("%s/%s", path, servico)
	par := url.Values{}
	if len(parametros) > 0 {
		for _, v := range parametros {
			par.Set(v[0], v[1])
		}
	}
	return fmt.Sprint(u), par
}

func GetServico(servico string) *Servico {

	envcfg.Unmarshal(&config)
	if config.CONSUL_HOST == "" || config.CONSUL_PORT == "" {
		fmt.Println("Erro: configure as variáveis de ambiente:\n - CONSUL_HOST: endereco IP do servidor Consul.\n - CONSUL_PORT: porta do servico HTTP do Consul.\n")
		return nil
	}

	url := GeraURLGet("http", config.CONSUL_HOST, config.CONSUL_PORT, pathConsul, servico, nil)
	response, err := http.Get(url)

	var srv Servico

	if !erro.Trata(err) {
		defer response.Body.Close()
		contents, errr := ioutil.ReadAll(response.Body)
		contentssrv := fmt.Sprintf(`{"servico": %s}`, string(contents))
		//fmt.Println(contentssrv)
		if !erro.Trata(errr) {
			err = json.Unmarshal([]byte(contentssrv), &srv)
			if erro.Trata(err) {
				fmt.Println(fmt.Sprintf("Erro: configure as variáveis de ambiente:\n - CONSUL_HOST: endereco IP do servidor Consul.\n - CONSUL_PORT: porta do servico HTTP do Consul.\n %s\n", err.Error()))
			}
		}
	} else {
		return nil
	}
	return &srv
}

var contentType string = "application/x-www-form-urlencoded"

func SetContentType(cont string) {
	contentType = cont
}

func ExecutaServico(srv *Servico, serviceparam [][]string) (string, string) {
	url, dado := GeraURLPost("http",
		srv.Servico[0].ServiceAddress,
		strconv.Itoa(srv.Servico[0].ServicePort),
		srv.Servico[0].ServicePath,
		srv.Servico[0].ServiceName,
		serviceparam)

	response, err := http.Post(url, contentType, strings.NewReader(dado.Encode()))

	var (
		resp    []byte
		codErro string
		menErro string
	)

	if erro.Trata(err) {
		codErro = "500"
		menErro = err.Error()
	} else {
		defer response.Body.Close()
		resp, err = ioutil.ReadAll(response.Body)
		codErro = strconv.Itoa(response.StatusCode)
		menErro = string(resp)
		if erro.Trata(err) {
			codErro = "500"
			menErro = err.Error()
		}
	}
	return trataRetorno(menErro, codErro)
}

func trataRetorno(resposta string, err string) (string, string) {
	mens := fmt.Sprintf("Cod %s", err)
	switch err {
	case "200":
		mens = mens + " - sucesso."
	case "201":
		mens = mens + " - salvo."
	case "400":
		mens = mens + " - request mal realizado."
	case "401":
		mens = mens + " - não autorizado."
	default:
		mens = mens + " - Erro não tratado."
	}
	return mens, resposta
}

type Route struct {
	Name        string
	Method      []string
	Version     string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func NewRouter(routes []Route) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		router.
			Methods(route.Method...).
			Path(fmt.Sprintf(route.Pattern, route.Version)).
			Name(route.Name).
			Handler(handler)

	}
	return router
}

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

// API check de servico
func CheckServico(res http.ResponseWriter, req *http.Request) {
	msg := fmt.Sprintf("Serviço em funcionamento, %q", html.EscapeString(req.URL.Path))
	res.Header().Set("Content-Type", "text/html; charset=UTF-8")
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(msg))
	//fmt.Println(msg)
}

// controle em Construcao.
func EmConstrucao(res http.ResponseWriter, req *http.Request, vars string) {
	msg := fmt.Sprintf("Funcionalidade em construcao, %q %s", html.EscapeString(req.URL.Path), vars)
	res.Header().Set("Content-Type", "text/html; charset=UTF-8")
	res.WriteHeader(http.StatusNotImplemented)
	res.Write([]byte(msg))
	fmt.Println(msg)
}
