package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"text/template"
	"time"
)

var horaOcorrencia string
var base_ string

type EntradaMultiDados struct {
	Base             string `json:"basemultidados"`
	Descricao        string `json:"descricao"`
	Divisao          string `json:"codigodivisao"`
	Solicitacao      string `json:"codigosolicitacao"`
	AbertoPor        string `json:"abertopor"`
	CodigoCliente    string `json:"codigocliente"`
	NomeCliente      string `json:"nomecliente"`
	Origem           string `json:"origem"`
	Ticket           string `json:"idticket"`
	Token            string `json:"token"`
	CodigoOperador   string `json:"codigooperador"`
	NomeOperador     string `json:"nomeoperador"`
	EmailCliente     string `json:"emailcliente"`
	Status           string `json:"status"`
	CodigoOcorrencia string `json:"codigoocorrencia"`
	TipoComentario   string `json:"comentariopublico"`
}

type Multidados struct {
	Entrada EntradaMultiDados `json:"entrada"`
}

func doLog(log string, msg string) string {
	fmt.Println(msg)
	log = fmt.Sprintf("%s\n %s", log, msg)
	return log
}

type RetornoSoap struct {
	Retorno              []byte
	Url                  string
	CodSolicitacao       string
	CodDivisao           string
	CodClienteMultiDados string
	CodCliente           string
	NomeCliente          string
	IdTicket             string
	CodOrigem            string
	CodAbertoPor         string
	Descricao            string
	Error                string
	Narrativas           string
	CodOperador          string
	NomeOperador         string
	EmailCliente         string
	Status               string
	CodOcorrencia        string
	TipoComentario       string
}

const usuariowsHomologacao = "andre.luis"
const senhawsHomologacao = "123mudar"

const usuariowsConstrucao = "andre.luis"
const senhawsConstrucao = "123mudar"

const usuariowsMatriz = "debora.toniate"
const senhawsMatriz = "debora12243568"

var usuariows string
var senhaws string

func GeraOcorrenciaMultiDados(res http.ResponseWriter, req *http.Request) (http.ResponseWriter, *http.Request) {

	var log string
	var ret RetornoSoap
	var dados Multidados

	dados.Entrada.Base = req.FormValue("basemulti") // C = Construcao - M = Matriz Servico - H = Homologação
	dados.Entrada.Descricao = req.FormValue("descricao")
	dados.Entrada.Divisao = req.FormValue("codivisao")
	dados.Entrada.Solicitacao = req.FormValue("codsolicitacao")
	dados.Entrada.AbertoPor = req.FormValue("codabertopor")
	dados.Entrada.CodigoCliente = req.FormValue("codigocliente")
	dados.Entrada.NomeCliente = req.FormValue("nomecliente")
	dados.Entrada.EmailCliente = req.FormValue("emailcliente")
	dados.Entrada.Origem = req.FormValue("codorigem")
	dados.Entrada.Ticket = req.FormValue("idticket")
	dados.Entrada.Token = req.FormValue("token")

	dados.Entrada.CodigoOperador = req.FormValue("codigooperador")
	dados.Entrada.NomeOperador = req.FormValue("nomeoperador")

	if dados.Entrada.Base == "" {
		log = doLog(log, "Tag 'basemulti' não informada: 'C' para Construção, 'M' para Matriz e 'H' para base de Homologação.")
	}
	if dados.Entrada.Descricao == "" {
		log = doLog(log, "Tag 'descricao' não informada: Informe a Descrição para identificar qual a descrição da ocorrência.")
	}
	if dados.Entrada.Divisao == "" {
		log = doLog(log, "Tag 'codivisao' não informada: Informe o Codigo da Divisão para identificar qual divisão será responsável pela ocorrência.")
	}
	if dados.Entrada.Solicitacao == "" {
		log = doLog(log, "Tag 'codsolicitacao' não informada: Informe o Código da Solicitação para identificar a solicitação da divisão.")
	}
	if dados.Entrada.NomeCliente == "" && dados.Entrada.CodigoCliente == "" && dados.Entrada.EmailCliente == "" {
		log = doLog(log, "Tag 'nomecliente', 'codigocliente' e 'emailcliente' não informados. Se o email não for informado é obrigatório informar o nome e o codigo do cliente.")
	}
	if dados.Entrada.Origem == "" {
		log = doLog(log, "Tag 'codorigem' não informada: Informe o Código da Origem para identificar a origem da ocorrência.")
	}
	if dados.Entrada.Ticket == "" {
		log = doLog(log, "Tag 'idticket' não informada: Informe o número do ticket para identificar do ticket na ocorrência.")
	}
	if dados.Entrada.Token == "" || dados.Entrada.Token != "w3W18V8a1VF3Uw1E5MfHex49O05dkdPi" {
		log = doLog(log, "Tag 'token' não informada: Informe o token o para validar a aplicação.")
	}
	//res.Write([]byte(log))

	log = doLog(log, setBase(dados.Entrada.Base, &ret))

	res.Write([]byte(log))

	ret.CodDivisao = dados.Entrada.Divisao
	ret.CodSolicitacao = dados.Entrada.Solicitacao
	ret.IdTicket = dados.Entrada.Ticket
	ret.CodCliente = dados.Entrada.CodigoCliente
	ret.NomeCliente = dados.Entrada.NomeCliente
	ret.CodAbertoPor = dados.Entrada.AbertoPor
	ret.CodOrigem = dados.Entrada.Origem
	ret.Descricao = dados.Entrada.Descricao
	ret.EmailCliente = dados.Entrada.EmailCliente

	ret.CodOperador = dados.Entrada.CodigoOperador
	ret.NomeOperador = dados.Entrada.NomeOperador

	if len(ret.Error) <= 0 {
		doCadOcorrencia(&ret)
		res.Write(ret.Retorno)
	} else {
		res.Write([]byte(ret.Error))
	}
	return res, req
}

func AtualizaOcorrenciaMultiDados(res http.ResponseWriter, req *http.Request) (http.ResponseWriter, *http.Request) {

	var log string
	var ret RetornoSoap
	var dados Multidados

	dados.Entrada.Base = req.FormValue("basemulti") // C = Construcao - M = Matriz Servico - H = Homologação
	dados.Entrada.Descricao = req.FormValue("descricao")
	dados.Entrada.Status = req.FormValue("status")
	dados.Entrada.CodigoOcorrencia = req.FormValue("codigoocorrencia")
	dados.Entrada.Token = req.FormValue("token")

	if dados.Entrada.Base == "" {
		log = doLog(log, "Tag 'basemulti' não informada: 'C' para Construção, 'M' para Matriz e 'H' para base de Homologação.")
	}
	if dados.Entrada.Descricao == "" {
		log = doLog(log, "Tag 'descricao' não informada: Informe a Descrição para identificar qual a descrição da ocorrência.")
	}
	if dados.Entrada.Status == "" {
		log = doLog(log, "Tag 'status' não informada: Informe o Status da Narrativa da Ocorrencia.")
	}
	if dados.Entrada.CodigoOcorrencia == "" {
		log = doLog(log, "Tag 'codigoocorrencia' não informada: Informe o codigo da ocorrencia a alterar.")
	}
	if dados.Entrada.Token == "" || dados.Entrada.Token != "w3W18V8a1VF3Uw1E5MfHex49O05dkdPi" {
		log = doLog(log, "Tag 'token' não informada: Informe o token o para validar a aplicação.")
	}

	log = doLog(log, setBase(dados.Entrada.Base, &ret))

	ret.Error = log

	ret.CodOcorrencia = dados.Entrada.CodigoOcorrencia
	ret.Descricao = dados.Entrada.Descricao
	ret.Status = dados.Entrada.Status

	if len(strings.TrimSpace(ret.Error)) <= 0 {
		doAtuOcorrencia(&ret)
		res.Write(ret.Retorno)
	} else {
		res.Write([]byte(ret.Error))
		fmt.Println("erro: ", ret.Error)
	}
	return res, req
}

func AtualizaTicketZenDesk(res http.ResponseWriter, req *http.Request) (http.ResponseWriter, *http.Request) {

	var log string
	var ret RetornoSoap
	var dados Multidados

	dados.Entrada.Descricao = req.FormValue("descricao")
	dados.Entrada.Status = req.FormValue("status")
	dados.Entrada.Ticket = req.FormValue("idticket")
	dados.Entrada.TipoComentario = req.FormValue("comentariopublico")
	dados.Entrada.Token = req.FormValue("token")

	if dados.Entrada.Descricao == "" {
		log = doLog(log, "Tag 'descricao' não informada: Informe a Descrição para o comentário do ticket.")
	}
	if dados.Entrada.Status == "" {
		log = doLog(log, "Tag 'status' não informada: Informe o Status do Ticket.")
	}
	if dados.Entrada.Ticket == "" {
		log = doLog(log, "Tag 'idticket' não informada: Informe o codigo do ticket.")
	}
	if dados.Entrada.TipoComentario == "" {
		log = doLog(log, "Tag 'comentariopublico' não informada: `true` para Comentário publico e `false` para privado.")
	}
	if dados.Entrada.Token == "" || dados.Entrada.Token != "w3W18V8a1VF3Uw1E5MfHex49O05dkdPi" {
		log = doLog(log, "Tag 'token' não informada: Informe o token o para validar a aplicação.")
	}

	ret.Error = log

	ret.IdTicket = dados.Entrada.Ticket
	ret.Descricao = dados.Entrada.Descricao
	ret.Status = dados.Entrada.Status
	ret.TipoComentario = dados.Entrada.TipoComentario

	if len(ret.Error) <= 0 {
		doAtuTicket(&ret)
		res.Write(ret.Retorno)
	} else {
		res.Write([]byte(ret.Error))
	}
	return res, req
}

func setBase(base string, ret *RetornoSoap) string {
	//Verifica base a ser utilizada.
	base_ = base // controle variavel global.
	switch base {
	case "C":
		ret.Url = "http://200.155.13.163/mega_construcao/webservices/index.php"
		usuariows = usuariowsConstrucao
		senhaws = senhawsConstrucao
	case "M":
		ret.Url = "http://200.159.62.51:8888/vmulti_matriz/webservices/index.php"
		usuariows = usuariowsMatriz
		senhaws = senhawsMatriz
	case "H":
		ret.Url = "http://189.108.175.204/vmulti_homologacao/webservices/index.php"
		usuariows = usuariowsHomologacao
		senhaws = senhawsHomologacao
	case "L":
		ret.Url = "http://10.0.0.248/vmulti_homologacao/webservices/index.php"
		usuariows = usuariowsHomologacao
		senhaws = senhawsHomologacao
	default:
		return "Tag 'basemulti' informada errada. Valores possíveis: 'C' para Construção, 'M' para Matriz e 'H' para Homologação."
	}
	return ""
}

const xmlSoapTmp = `<soapenv:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:urn="urn:server.Multidados">
	<soapenv:Header/>
	<soapenv:Body>
		<{{.Urn}} soapenv:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
	        {{.Dados}}
    	</{{.Urn}}>
	</soapenv:Body>
	</soapenv:Envelope>`

type HeadSoap struct {
	Urn    string
	Dados  string
	Action string
}

const xmlsoap = `<{{.Tag}}>{{.Valor}}</{{.Tag}}>`

type DadosSoap struct {
	Servicedesk struct {
		Dados []struct {
			Tag   string `json:"tag"`
			Valor string `json:"valor"`
		} `json:"dados"`
	} `json:"servicedesk"`
}

func postSoap(ret *RetornoSoap, dj DadosSoap, head HeadSoap) []byte {
	var erroM error
	tempXmlSoap, erroM := template.New("soap").Parse(xmlsoap)
	if erroM != nil {
		fmt.Println("ERRO Template Parse xmlsoap: ", erroM.Error())
		return nil
	}

	headXmlSoap, erroH := template.New("headsoap").Parse(xmlSoapTmp)
	if erroH != nil {
		fmt.Println("ERRO Template Parse xmlSoapTmp: ", erroH.Error())
		return nil
	}

	buf := new(bytes.Buffer)
	bufsoap := new(bytes.Buffer)

	if tempXmlSoap != nil {
		for _, d := range dj.Servicedesk.Dados {
			erro := tempXmlSoap.Execute(buf, d)
			if erro != nil {
				fmt.Println("ERRO ao executar o template tempXmlSoap:", erro.Error())
			}
		}

		head.Dados = buf.String()

		erro := headXmlSoap.Execute(bufsoap, head)
		if erro != nil {
			fmt.Println("ERRO ao executar o template xmlSoapTmp:", erro.Error())
		}

		readerSoap := strings.NewReader(bufsoap.String())
		retBytes := soapRequest(head.Urn, head.Action, "POST", ret.Url, readerSoap)

		//fmt.Println(fmt.Sprintf("URL: %s \n SOAP: \n%s", ret.Url, bufsoap.String()))

		return retBytes
	}

	return nil

}

func doCadOcorrencia(ret *RetornoSoap) {

	dj := getDadosSoapJson(ret)

	var head HeadSoap
	head.Urn = "urn:DadosServicedesk"
	head.Action = "urn:server.Multidados#DadosServicedesk"

	t := time.Now()
	t = t.Add(-4 * time.Hour)
	horaOcorrencia = t.Format("2006-01-02 15:04:05")

	resp := postSoap(ret, dj, head)

	if !strings.Contains(string(resp), " (e:") && !strings.Contains(string(resp), "erro") {
		codOc := doCadNarrativaOc(ret)
		if codOc != "" {
			atualizaCustomField(ret, codOc)
			ret.Retorno = []byte(fmt.Sprintf("Ocorrência Gerada: %s ", codOc))
		} else {
			fmt.Println("Erro ao gravar a ocorrencia... ")
		}
	} else {
		ret.Retorno = resp
		if strings.Contains(string(resp), "<html>") {
			fmt.Println("ERRO não indentificado pelo WebService ", ret.Url)
		} else {
			fmt.Println("ERRO: ", string(resp))
		}
	}
}

func doAtuOcorrencia(ret *RetornoSoap) {
	dj := getDadosSoapEditOCJson(ret)
	//fmt.Println(dj)
	var head HeadSoap
	head.Urn = "urn:EditOC"
	head.Action = "urn:server.Multidados#EditOC"

	resp := postSoap(ret, dj, head)
	//fmt.Println(string(resp))
	if !strings.Contains(string(resp), " (e:") && !strings.Contains(string(resp), "erro") && !strings.Contains(string(resp), ":false") {
		ret.Retorno = []byte(fmt.Sprintf("Ocorrência Alterada: %s ", ret.CodOcorrencia))
	} else {
		ret.Retorno = resp
		if strings.Contains(string(resp), "<html>") {
			fmt.Println("ERRO não indentificado pelo WebService ", ret.Url)
		} else {
			fmt.Println("ERRO: ", string(resp))
		}
	}
}

type ZComment struct {
	Type      string `json:"type"`
	Body      string `json:"body"`
	Public    string `json:"public"`
	Author_id int32  `json:"author_id"`
}

type ZTicketC struct {
	Comment ZComment `json:"comment"`
	Status  string   `json:"status"`
}

type ZTicketCom struct {
	Ticket ZTicketC `json:"ticket"`
}

func doAtuTicket(ret *RetornoSoap) {

	var com ZTicketC
	var tk ZTicketCom

	com.Comment.Body = ret.Descricao
	com.Comment.Public = ret.TipoComentario
	com.Comment.Type = "Comment"
	com.Status = ret.Status

	tk.Ticket = com

	tkJson, _ := json.Marshal(tk)

	client := &http.Client{}
	url := fmt.Sprintf("https://megasistemas.zendesk.com/api/v2/tickets/%s.json", ret.IdTicket)
	reqZen, err1 := http.NewRequest("PUT", url, strings.NewReader(string(tkJson)))
	if err1 != nil {
		msgerr := fmt.Sprintf("Erro ao editar o ticket %s - %s \n %s", ret.IdTicket, err1.Error(), string(tkJson))
		ret.Retorno = []byte(msgerr)
		fmt.Println(msgerr)
	}
	reqZen.SetBasicAuth("qualiteam@mega.com.br/token", "pzwsg7ItHBIQqX1wzV66sYjl36awkhCXLgyMcjMx")
	reqZen.Header.Set("Content-Type", "application/json")
	respS, errR := client.Do(reqZen)

	resp, _ := ioutil.ReadAll(respS.Body)
	//fmt.Println(string(resp))

	if errR != nil {
		msgerr := fmt.Sprintf("Erro ao editar o ticket %s - %s ", ret.IdTicket, errR.Error())
		ret.Retorno = []byte(msgerr)
		fmt.Println(msgerr)
	} else {
		if !strings.Contains(string(resp), "error") {
			ret.Retorno = []byte(fmt.Sprintf("Ticket Alterado: %s ", ret.IdTicket))
		} else {
			msgerr := fmt.Sprintf("Erro ao editar o ticket %s - %s ", ret.IdTicket, string(resp))
			ret.Retorno = []byte(msgerr)
			fmt.Println(msgerr)
		}
	}

}

type Envelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    *TBody
}

type TBody struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	Servico *TServico
}

type TServico struct {
	XMLName xml.Name `xml:"urn:server.Multidados ConsultarOcorrenciasResponse"`
	Retorno string   `xml:"return"`
}

type ConsultaOC struct {
	Dados []struct {
		CodigoTicket string `json:"Código da ocorrência"`
		CodigoOC     string `json:"N.º"`
	} `json:"dados"`
}

func getCodigoOC(ret *RetornoSoap) string {

	dj := getDadosCodigoOCJson(ret)

	var head HeadSoap
	head.Urn = "urn:ConsultarOcorrencias"
	head.Action = "urn:server.Multidados#ConsultarOcorrencias"

	body := postSoap(ret, dj, head)
	if body != nil {
		var retorno Envelope
		err := xml.Unmarshal(body, &retorno)
		if err == nil {
			auxJson := []byte(fmt.Sprintf("%s%s%s", `{"dados":`, retorno.Body.Servico.Retorno, "}"))
			var aux ConsultaOC
			json.Unmarshal(auxJson, &aux)
			for _, v := range aux.Dados {
				if ret.IdTicket == v.CodigoTicket {
					return strings.TrimSpace(v.CodigoOC)
				}
			}
		} else {
			fmt.Println(fmt.Sprintf("ERRO ao busca o Código da Ocorrencia: %s\n%s", err.Error(), string(body)))
		}
	}
	return ""
}

func doCadNarrativaOc(ret *RetornoSoap) string {

	narr := doGetComentsZenDesk(ret)

	codigoOC := getCodigoOC(ret)

	if len(narr) > 0 && codigoOC != "" {

		ret.Narrativas = narr

		dj := getDadosSoapJsonNarr(codigoOC, ret)

		var head HeadSoap
		head.Urn = "urn:DadosNarrativasOC"
		head.Action = "urn:server.Multidados#DadosNarrativasOC"

		retorno := postSoap(ret, dj, head)
		fmt.Println(fmt.Sprintf("Retorno Atualiza Narrativas: %s", retorno))

		if strings.Contains(string(retorno), " (e:") {
			ret.Retorno = []byte(fmt.Sprintf("%s\n%s", ret.Retorno, retorno))
		}
	}
	return codigoOC
}

type CommentsTicket struct {
	Comments []struct {
		Body   string `json:"body"`
		Public bool   `json:"public"`
	} `json:"comments"`
	Count        float64     `json:"count"`
	NextPage     interface{} `json:"next_page"`
	PreviousPage interface{} `json:"previous_page"`
}

func doGetComentsZenDesk(ret *RetornoSoap) string {

	var narrativas string
	client := &http.Client{}
	url := "https://megasistemas.zendesk.com/api/v2/tickets/" + ret.IdTicket + "/comments.json"

	for {
		reqCliente, _ := http.NewRequest("GET", url, nil)
		reqCliente.SetBasicAuth("qualiteam@mega.com.br/token", "pzwsg7ItHBIQqX1wzV66sYjl36awkhCXLgyMcjMx")
		reqCliente.Header.Set("Content-Type", "application/json")

		respCliente, _ := client.Do(reqCliente)

		body, _ := ioutil.ReadAll(respCliente.Body)

		var comments CommentsTicket
		err := json.Unmarshal(body, &comments)
		if err == nil {
			if len(comments.Comments) > 0 {
				for _, v := range comments.Comments {
					if v.Public {
						narrativas = fmt.Sprintf("%s \n %s", narrativas, v.Body)
					}
				}
			}
		} else {
			fmt.Println("Erro buscar Narrativas: ", err.Error())
		}
		if comments.NextPage == nil {
			break
		}
		url = comments.NextPage.(string)
	}

	return narrativas
}

type CustomField struct {
	Id    int32  `json:"id"`
	Value string `json:"value"`
}

type CustomTicket struct {
	Custom_field []CustomField `json:"custom_fields"`
}

type CustomTickets struct {
	Ticket CustomTicket `json:"ticket"`
}

func atualizaCustomField(ret *RetornoSoap, idOcorrencia string) {
	var tkt CustomTickets
	var tk CustomTicket
	var c CustomField
	c.Id = 22719144
	c.Value = idOcorrencia
	tk.Custom_field = append(tk.Custom_field, c)
	tkt.Ticket = tk
	tkJson, _ := json.Marshal(tkt)

	client := &http.Client{}
	url := fmt.Sprintf("https://megasistemas.zendesk.com/api/v2/tickets/%s.json", ret.IdTicket)
	//url := fmt.Sprintf("https://megasistemas1411735609.zendesk.com/api/v2/tickets/%s.json", ret.IdTicket)
	reqZen, _ := http.NewRequest("PUT", url, strings.NewReader(string(tkJson)))
	reqZen.SetBasicAuth("qualiteam@mega.com.br/token", "pzwsg7ItHBIQqX1wzV66sYjl36awkhCXLgyMcjMx")
	reqZen.Header.Set("Content-Type", "application/json")
	resp, _ := client.Do(reqZen)
	_, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(fmt.Sprintf("Erro ao atualizar a ocorrencia %s no ticket %s : %s", idOcorrencia, ret.IdTicket, err.Error()))
	}
}

func soapRequest(urn, soapAction, metodo, url string, readerXml *strings.Reader) []byte {

	client := &http.Client{}

	reqSoap, err := http.NewRequest(metodo, url, readerXml)
	reqSoap.Header.Set("SOAPAction", soapAction)
	reqSoap.Header.Add("Content-Type", "text/xml;charset=UTF-8")

	respSoap, err := client.Do(reqSoap)

	if err == nil {
		b := new(bytes.Buffer)
		b.ReadFrom(respSoap.Body)
		if err != nil {
			panic(err)
		}
		return b.Bytes()
	} else {
		fmt.Println(fmt.Sprintf("ERRO ao executar o serviço: %s\n", err.Error()))
		return nil
	}

}

type DadosClienteEmail struct {
	Dados []struct {
		CodigoCliente string `json:"codigo_cliente"`
		CodigoContato string `json:"codigo_contato"`
		EmailContato  string `json:"email_contato"`
		NomeCliente   string `json:"nome_cliente"`
		NomeContato   string `json:"nome_contato"`
		Telefones     string `json:"telefones"`
	} `json:"dados"`
}

func getDadosClienteEmail(ret *RetornoSoap) *RetornoSoap {

	client := http.Client{}

	var end string

	switch base_ {
	case "C":
		end = "200.155.13.163/mega_construcao3"
	case "M":
		end = "200.159.62.51:8888/vmulti_matriz"
	case "H":
		end = "189.108.175.204/vmulti_homologacao"
	default:
		end = ""
	}

	url := fmt.Sprintf("http://%s/webservices/crontab/contato_clientes.php?email=%s", end, ret.EmailCliente)
	reqLog, err := http.NewRequest("POST", url, nil)
	resp, err1 := client.Do(reqLog)
	if err1 == nil {
		body, _ := ioutil.ReadAll(resp.Body)
		var dadosEmail DadosClienteEmail
		aux := fmt.Sprintf("%s%s%s", `{"dados":`, string(body), "}")
		err = json.Unmarshal([]byte(aux), &dadosEmail)
		if err == nil && len(dadosEmail.Dados) > 0 {
			ret.CodCliente = dadosEmail.Dados[0].CodigoCliente
			if ret.CodAbertoPor == "" {
				ret.CodAbertoPor = dadosEmail.Dados[0].CodigoContato
			}
			ret.NomeCliente = dadosEmail.Dados[0].NomeCliente
		}
	}
	return ret
}

func getDadosSoapJson(ret *RetornoSoap) DadosSoap {

	var dad DadosSoap
	conf, err := ioutil.ReadFile("dadosoap_.json")
	if err == nil {
		err = json.Unmarshal(conf, &dad)
		if err != nil {
			fmt.Println(fmt.Sprintf("ERRO ao gerar DadosSoapJson: %s\n%s", string(conf), err.Error()))
		}
		if len(dad.Servicedesk.Dados) > 0 {
			if ret.EmailCliente != "" {
				ret = getDadosClienteEmail(ret)
			}
			dad.Servicedesk.Dados[searchSoap(dad, "USUARIO_WS")].Valor = usuariows
			dad.Servicedesk.Dados[searchSoap(dad, "SENHA_WS")].Valor = senhaws
			dad.Servicedesk.Dados[searchSoap(dad, "COD_DIVISAO")].Valor = ret.CodDivisao
			dad.Servicedesk.Dados[searchSoap(dad, "COD_SOLICITACAO")].Valor = ret.CodSolicitacao
			dad.Servicedesk.Dados[searchSoap(dad, "COD_CLIENTE")].Valor = ret.CodCliente
			dad.Servicedesk.Dados[searchSoap(dad, "NOME_CLIENTE")].Valor = ret.NomeCliente
			dad.Servicedesk.Dados[searchSoap(dad, "CODIGO_OC")].Valor = ret.IdTicket
			dad.Servicedesk.Dados[searchSoap(dad, "COD_ABERTO_POR")].Valor = ret.CodAbertoPor
			dad.Servicedesk.Dados[searchSoap(dad, "COD_ORIGEM")].Valor = ret.CodOrigem
			dad.Servicedesk.Dados[searchSoap(dad, "DESCRICAO")].Valor = ret.Descricao
		}
	}
	return dad
}

func getDadosSoapJsonNarr(codigooc string, ret *RetornoSoap) DadosSoap {

	var dad DadosSoap
	conf, err := ioutil.ReadFile("dadosoapnar_.json")
	if err == nil {
		err = json.Unmarshal(conf, &dad)
		if err != nil {
			fmt.Println(fmt.Sprintf("ERRO ao gerar DadosSoapJsonNarr: %s\n%s", string(conf), err.Error()))
		}
		if len(dad.Servicedesk.Dados) > 0 {
			t := time.Now()
			dad.Servicedesk.Dados[searchSoap(dad, "USUARIO_WS")].Valor = usuariows
			dad.Servicedesk.Dados[searchSoap(dad, "SENHA_WS")].Valor = senhaws
			dad.Servicedesk.Dados[searchSoap(dad, "N_OCORRENCIA")].Valor = codigooc

			reg, err := regexp.Compile("[^A-Za-z0-9]+")
			if err == nil {
				ret.Narrativas = reg.ReplaceAllString(ret.Narrativas, " ")
			}

			dad.Servicedesk.Dados[searchSoap(dad, "DESCRICAO")].Valor = ret.Narrativas
			dad.Servicedesk.Dados[searchSoap(dad, "COD_OPERADOR")].Valor = ret.CodOperador
			dad.Servicedesk.Dados[searchSoap(dad, "NOME_OPERADOR")].Valor = ret.NomeOperador
			dad.Servicedesk.Dados[searchSoap(dad, "DATA_HORA_ATENDIMENTO")].Valor = t.Format("2006-01-02 15:04:05")
		}
	}
	return dad
}

func getDadosCodigoOCJson(ret *RetornoSoap) DadosSoap {

	var dad DadosSoap
	conf, err := ioutil.ReadFile("dadosoc_.json")
	if err == nil {
		err = json.Unmarshal(conf, &dad)
		if err != nil {
			fmt.Println(fmt.Sprintf("ERRO ao gerar getDadosCodigoOCJson: %s\n%s", string(conf), err.Error()))
		}
		if len(dad.Servicedesk.Dados) > 0 {
			t := time.Now()
			dad.Servicedesk.Dados[searchSoap(dad, "USUARIO_WS")].Valor = usuariows
			dad.Servicedesk.Dados[searchSoap(dad, "SENHA_WS")].Valor = senhaws
			dad.Servicedesk.Dados[searchSoap(dad, "TIPO_DATA")].Valor = "data_abertura"
			dad.Servicedesk.Dados[searchSoap(dad, "DATA_INI")].Valor = horaOcorrencia
			dad.Servicedesk.Dados[searchSoap(dad, "DATA_FIM")].Valor = t.Format("2006-01-02 15:04:05")
			dad.Servicedesk.Dados[searchSoap(dad, "RETORNO")].Valor = "json"
			dad.Servicedesk.Dados[searchSoap(dad, "CAMPOS")].Valor = "numero,codigo_oc"
		}
	}
	return dad
}

func getDadosSoapEditOCJson(ret *RetornoSoap) DadosSoap {

	var dad DadosSoap
	conf, err := ioutil.ReadFile("dadosoapeditoc_.json")
	if err == nil {
		err = json.Unmarshal(conf, &dad)
		if err != nil {
			fmt.Println(fmt.Sprintf("ERRO ao gerar DadosSoapEditOCJson: %s\n%s", string(conf), err.Error()))
		}
		if len(dad.Servicedesk.Dados) > 0 {
			params := [][]string{{"numero", ret.CodOcorrencia}, {"status", ret.Status}, {"responsavel", "ANDREL"}, {"descricao", ret.Descricao}}
			par := url.Values{}
			for _, v := range params {
				par.Set(v[0], v[1])
			}
			dad.Servicedesk.Dados[searchSoap(dad, "USUARIO_WS")].Valor = usuariows
			dad.Servicedesk.Dados[searchSoap(dad, "SENHA_WS")].Valor = senhaws
			dad.Servicedesk.Dados[searchSoap(dad, "PARAMS")].Valor = strings.Replace(par.Encode(), "&", "&amp;", -1)
			//"lt": "<",
			//	"gt": ">",
			//	"apos": "'",
			//	"quot": `"`,

		}
	}
	return dad
}

func searchSoap(dados DadosSoap, valor string) int {
	for i, v := range dados.Servicedesk.Dados {
		if v.Tag == valor {
			return i
		}
	}
	return 0
}
