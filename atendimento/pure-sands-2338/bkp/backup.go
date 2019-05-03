package bkp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	//"net/url"
	"os"
	"strconv"
	"strings"
)

type Comment struct {
	Author_id int    `json:"author_id"`
	Body      string `json:"body"`
	Public    bool   `json:"public"`
}

type CustomFields struct {
	Id    int    `json:"id"`
	Value string `json:"value"`
}

type CommentSingle struct {
	Comment Comment `json:"comment"`
}

type Ticket struct {
	Ticket struct {
		Id           int            `json:"id"`
		Type         string         `json:"type"`
		Subject      string         `json:"subject"`
		Description  string         `json:"description"`
		Organization int            `json:"organization_id"`
		GroupId      int            `json:"group_id"`
		Requester_id int            `json:"requester_id"`
		Comment      CommentSingle  `json:"comment"`
		Status       string         `json:"status"`
		CustomFields []CustomFields `json:"custom_fields"`
	} `json:"ticket"`
}

/*
func main() {

	fmt.Println("Em funcionamento...")
	http.HandleFunc("/intsugestzd", IntSugestZd)
	http.HandleFunc("/intmultidados", IntMultiDados)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)

	panic_err(err)

}
*/
func IntSugestZd(res http.ResponseWriter, req *http.Request) {

	token := req.FormValue("token")
	usuSoli := req.FormValue("ususoli")
	tipoNotifAre := req.FormValue("notifarea")
	grupoKCSII := req.FormValue("grupokcsii")
	campoNotifArea := req.FormValue("camponotifarea")
	subject := req.FormValue("subject")
	commentSugest := req.FormValue("commentSugest")
	campoTipoSoli := req.FormValue("campotiposoli")
	idTipoSoli := req.FormValue("idtiposoli")

	if token != "sqkv586mB79ESo6gg51mZiP1T4GF70yn" || usuSoli == "" || tipoNotifAre == "" || grupoKCSII == "" || campoNotifArea == "" || campoTipoSoli == "" || idTipoSoli == "" {
		res.Write([]byte("Favor informar todos os campos corretamente."))
		return
	}

	/* trecho de log.
	v := url.Values{}

	v.Add("token", token)
	v.Add("ususoli", usuSoli)
	v.Add("notifarea", tipoNotifAre)
	v.Add("grupokcsii", grupoKCSII)
	v.Add("camponotifarea", campoNotifArea)
	v.Add("subject", subject)
	v.Add("commentSugest", commentSugest)

	doLogRequestBin(v.Encode(), "http://requestb.in/ofgy0qof")*/

	client := &http.Client{}

	var ticket Ticket

	intUsuSoli, err := strconv.Atoi(usuSoli)
	intCampoNotifArea, err := strconv.Atoi(campoNotifArea)
	intGroupId, err := strconv.Atoi(grupoKCSII)
	intCampoTipoSoli, err := strconv.Atoi(campoTipoSoli)

	ticket.Ticket.Subject = "Sugestão - " + subject
	ticket.Ticket.Requester_id = intUsuSoli

	ticket.Ticket.CustomFields = make([]CustomFields, 2)
	ticket.Ticket.CustomFields[0].Id = intCampoNotifArea
	ticket.Ticket.CustomFields[0].Value = tipoNotifAre
	ticket.Ticket.CustomFields[1].Id = intCampoTipoSoli
	ticket.Ticket.CustomFields[1].Value = idTipoSoli

	ticket.Ticket.Type = "task"
	ticket.Ticket.GroupId = intGroupId
	ticket.Ticket.Description = strings.Replace(commentSugest, "#SUGESTAO_ARTIGO", "", -1)

	byteTicket, err := json.Marshal(ticket)

	readerTicket := bytes.NewReader(byteTicket)

	reqPostTicket, err := http.NewRequest("POST", "https://megasistemas.zendesk.com/api/v2/tickets.json", readerTicket)
	reqPostTicket.SetBasicAuth("qualiteam@mega.com.br/token", "pzwsg7ItHBIQqX1wzV66sYjl36awkhCXLgyMcjMx")
	reqPostTicket.Header.Set("Content-Type", "application/json")

	respPostTicket, err := client.Do(reqPostTicket)

	panic_err(err)

	b := new(bytes.Buffer)
	b.ReadFrom(respPostTicket.Body)

	res.Write(b.Bytes())

}

func panic_err(err error) {
	if err != nil {
		panic(err)
	}
}

func doLogRequestBin(values, url string) {

	fmt.Println(url + "?" + values)

	client := http.Client{}
	reqLog, err := http.NewRequest("POST", url+"?"+values, nil)
	_, err = client.Do(reqLog)

	panic_err(err)
}

func doLogGo(message string) {
	fmt.Println(message)
}

type RetornoSoap struct {
	retorno              []byte
	url                  string
	codSolicitacao       string
	codDivisao           string
	codClienteMultiDados string
	codCliente           string
	NomeCliente          string
	IdTicket             string
	Error                string
}

const xmlSoapTmp = `<soapenv:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:urn="urn:server.Multidados">
	<soapenv:Header/>
	<soapenv:Body>
		<&urn& soapenv:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
			<USUARIO_WS xsi:type="xsd:string">andre.luis</USUARIO_WS>
			<SENHA_WS xsi:type="xsd:string">123mudar</SENHA_WS>
	        &dados&
    	</&urn&>
	</soapenv:Body>
	</soapenv:Envelope>`

func IntMultiDados(res http.ResponseWriter, req *http.Request) {

	baseMulti := req.FormValue("basemulti") // C = Construcao - M = Matriz Servico
	acao := req.FormValue("acao")
	codDivisao := req.FormValue("codivisao")
	codSolicitacao := req.FormValue("codsolicitacao")
	codCliente := req.FormValue("codcliente")
	token := req.FormValue("token")
	idTicket := req.FormValue("idticket")

	var ret RetornoSoap

	//Verifica se os campos foram informados
	if baseMulti == "" || acao == "" || codDivisao == "" || codSolicitacao == "" || token != "w3W18V8a1VF3Uw1E5MfHex49O05dkdPi" || codCliente == "" {

		res.Write([]byte("Favor preencher todos os campos."))
		doLogGo("message=preenche_todos_campos")
		return

	}

	//Verifica base a ser utilizada.
	if (baseMulti != "") && (baseMulti == "C") {

		ret.url = "http://10.0.0.248/vmulti_homologacao/webservices/index.php"

	} else if (baseMulti != "") && (baseMulti == "M") {

		ret.url = "http://10.0.0.248/vmulti_homologacao/webservices/index.php"

	} else {

		res.Write([]byte("Favor informar um código correto para seleção da base na qual deseja inserir a ocorrência."))
		return

	}

	acao = strings.ToLower(acao)

	switch acao {
	case "cadoco":

		doLogGo("message=CADOCO")
		ret.codCliente = codCliente
		ret.codDivisao = codDivisao
		ret.codSolicitacao = codSolicitacao
		ret.IdTicket = idTicket

		doGetClienteCodMultidados(&ret)
		if len(ret.Error) <= 0 {
			doCadOcorrencia(&ret)
		} else {
			res.Write([]byte(ret.Error))
			return
		}

		break
	}

	res.Write(ret.retorno)

}

/*func doRequisicaoGetClientCPFCNPJ(ret *RetornoSoap) {

	strCampos := `<CNPJ_CPF xsi:type="xsd:string">` + ret.cnpjCliente + `</CNPJ_CPF>`

	strSoapAction := "urn:server.Multidados#GetClienteCpfCnpj"
	strUrn := "urn:GetClienteCpfCnpj"
	strXmlSoap := strings.Replace(xmlSoapTmp, "&urn&", strUrn, -1)
	strXmlSoap = fmt.Sprintf(strXmlSoap, strCampos)

	readerSoap := strings.NewReader(strXmlSoap)
	retBytes := soapRequest(strUrn, strSoapAction, "POST", ret.url, readerSoap)

	ret.retorno = retBytes

}*/

type ClienteZendesk struct {
	Organization struct {
		Name               string `json:"name"`
		OrganizationFields struct {
			CodigoCliente string `json:"cdigo_cliente"`
		}
	} `json:"organization"`
	Error string `json:"error"`
}

func doGetClienteCodMultidados(ret *RetornoSoap) {

	doLogGo("CODCLIENTE:" + ret.codCliente)

	reqCliente, _ := http.NewRequest("GET", "https://megasistemas.zendesk.com/api/v2/organizations/"+ret.codCliente+".json", nil)
	reqCliente.SetBasicAuth("andreo@mega.com.br/token", "pzwsg7ItHBIQqX1wzV66sYjl36awkhCXLgyMcjMx")
	reqCliente.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	respCliente, _ := client.Do(reqCliente)

	b := new(bytes.Buffer)
	b.ReadFrom(respCliente.Body)

	var cZdsk ClienteZendesk
	_ = json.Unmarshal(b.Bytes(), &cZdsk)

	ret.codClienteMultiDados = cZdsk.Organization.OrganizationFields.CodigoCliente
	ret.NomeCliente = cZdsk.Organization.Name
	ret.Error = cZdsk.Error

	doLogGo("CodClient=" + b.String())
}

func doCadOcorrencia(ret *RetornoSoap) {

	//<COD_ORIGEM xsi:type="xsd:string"></COD_ORIGEM>

	strCampos := `  <CNPJ_EMPRESA xsi:type="xsd:string"></CNPJ_EMPRESA>
					<COD_DIVISAO xsi:type="xsd:string">` + ret.codDivisao + `</COD_DIVISAO>
					<COD_SOLICITACAO xsi:type="xsd:string">` + ret.codSolicitacao + `</COD_SOLICITACAO>
					<COD_PROJETO xsi:type="xsd:string"></COD_PROJETO>
					<DESCRICAO xsi:type="xsd:string"></DESCRICAO>
					<DESCRICAO_SOLUCAO xsi:type="xsd:string"></DESCRICAO_SOLUCAO>
					<DATA_ABERTURA xsi:type="xsd:string"></DATA_ABERTURA>
					<DATA_FECHAMENTO xsi:type="xsd:string"></DATA_FECHAMENTO>
					<COD_ABERTO_POR xsi:type="xsd:string">ANDREL</COD_ABERTO_POR>
					<COD_ABERTO_PARA xsi:type="xsd:string">ANDREL</COD_ABERTO_PARA>
					<COD_CLIENTE xsi:type="xsd:string">` + ret.codClienteMultiDados + `</COD_CLIENTE>
					<NOME_CLIENTE xsi:type="xsd:string">` + ret.NomeCliente + `</NOME_CLIENTE>
					<CNPJ_CLIENTE xsi:type="xsd:string"></CNPJ_CLIENTE>
					<DT_NASC_CLIENTE xsi:type="xsd:string"></DT_NASC_CLIENTE>
					<END_CLIENTE xsi:type="xsd:string"></END_CLIENTE>
					<BAIRRO_CLIENTE xsi:type="xsd:string"></BAIRRO_CLIENTE>
					<COMPLEMENTO_CLIENTE xsi:type="xsd:string"></COMPLEMENTO_CLIENTE>
					<CEP_CLIENTE xsi:type="xsd:string"></CEP_CLIENTE>
					<CIDADE_CLIENTE xsi:type="xsd:string"></CIDADE_CLIENTE>
					<ESTADO_CLIENTE xsi:type="xsd:string"></ESTADO_CLIENTE>
					<NUM_END_CLIENTE xsi:type="xsd:string"></NUM_END_CLIENTE>
					<DDD_RES_CLIENTE xsi:type="xsd:string"></DDD_RES_CLIENTE>
					<TEL_RES_CLIENTE xsi:type="xsd:string"></TEL_RES_CLIENTE>
					<DDD_CEL_CLIENTE xsi:type="xsd:string"></DDD_CEL_CLIENTE>
					<TEL_CEL_CLIENTE xsi:type="xsd:string"></TEL_CEL_CLIENTE>
					<DDD_COM_CLIENTE xsi:type="xsd:string"></DDD_COM_CLIENTE>
					<TEL_COM_CLIENTE xsi:type="xsd:string"></TEL_COM_CLIENTE>
					<DDD_OUTROS_CLIENTE xsi:type="xsd:string"></DDD_OUTROS_CLIENTE>
					<TEL_OUTROS_CLIENTE xsi:type="xsd:string"></TEL_OUTROS_CLIENTE>
					<EMAIL_CLIENTE xsi:type="xsd:string"></EMAIL_CLIENTE>
					<SEXO_CLIENTE xsi:type="xsd:string"></SEXO_CLIENTE>
					<EST_CIVIL_CLIENTE xsi:type="xsd:string"></EST_CIVIL_CLIENTE>
					<QTDE_DEPENDENTES_CLIENTE xsi:type="xsd:string"></QTDE_DEPENDENTES_CLIENTE>
					<NOME_MAE_CLIENTE xsi:type="xsd:string"></NOME_MAE_CLIENTE>
					<ESTADO_NASC_CLIENTE xsi:type="xsd:string"></ESTADO_NASC_CLIENTE>
					<RG_CLIENTE xsi:type="xsd:string"></RG_CLIENTE>
					<DT_EMISSAO_CLIENTE xsi:type="xsd:string"></DT_EMISSAO_CLIENTE>
					<ORGAO_EMISSOR_CLIENTE xsi:type="xsd:string"></ORGAO_EMISSOR_CLIENTE>
					<ESTADO_EMISSOR_CLIENTE xsi:type="xsd:string"></ESTADO_EMISSOR_CLIENTE>
					<BANCOS_CLIENTES xsi:type="urn:BD_Clientes" soapenc:arrayType="urn:BD_ClientesStruct[]"/>
					<NOME_CONTATO xsi:type="xsd:string"></NOME_CONTATO>
					<ATIVO_RECEPTIVO xsi:type="xsd:string">A</ATIVO_RECEPTIVO>
					<COD_CAMPANHA xsi:type="xsd:string"></COD_CAMPANHA>
					<COD_ORIGEM xsi:type="xsd:string"></COD_ORIGEM>
					<CAMPOS_VARIAVEIS_CLIENTES xsi:type="urn:CA_Clientes" soapenc:arrayType="urn:CA_ClientesStruct[]"/>
					<CAMPOS_VARIAVEIS_OCORRENCIA xsi:type="urn:CA_Clientes" soapenc:arrayType="urn:CA_ClientesStruct[]"/>
					<TELEFONES_CLIENTES_FULL xsi:type="urn:CA_Clientes" soapenc:arrayType="urn:CA_ClientesStruct[]"/>
					<CODIGO_AUXILIAR xsi:type="xsd:string"></CODIGO_AUXILIAR>
					<DDD_RES2_CLIENTE xsi:type="xsd:string"></DDD_RES2_CLIENTE>
					<TEL_RES2_CLIENTE xsi:type="xsd:string"></TEL_RES2_CLIENTE>
					<DDD_CEL2_CLIENTE xsi:type="xsd:string"></DDD_CEL2_CLIENTE>
					<TEL_CEL2_CLIENTE xsi:type="xsd:string"></TEL_CEL2_CLIENTE>
					<DDD_COM2_CLIENTE xsi:type="xsd:string"></DDD_COM2_CLIENTE>
					<TEL_COM2_CLIENTE xsi:type="xsd:string"></TEL_COM2_CLIENTE>
					<DDD_OUTROS2_CLIENTE xsi:type="xsd:string"></DDD_OUTROS2_CLIENTE>
					<TEL_OUTROS2_CLIENTE xsi:type="xsd:string"></TEL_OUTROS2_CLIENTE>
					<DDD_RES3_CLIENTE xsi:type="xsd:string"></DDD_RES3_CLIENTE>
					<TEL_RES3_CLIENTE xsi:type="xsd:string"></TEL_RES3_CLIENTE>
					<DDD_CEL3_CLIENTE xsi:type="xsd:string"></DDD_CEL3_CLIENTE>
					<TEL_CEL3_CLIENTE xsi:type="xsd:string"></TEL_CEL3_CLIENTE>
					<DDD_COM3_CLIENTE xsi:type="xsd:string"></DDD_COM3_CLIENTE>
					<TEL_COM3_CLIENTE xsi:type="xsd:string"></TEL_COM3_CLIENTE>
					<DDD_OUTROS3_CLIENTE xsi:type="xsd:string"></DDD_OUTROS3_CLIENTE>
					<TEL_OUTROS3_CLIENTE xsi:type="xsd:string"></TEL_OUTROS3_CLIENTE>
					<DDD_RES4_CLIENTE xsi:type="xsd:string"></DDD_RES4_CLIENTE>
					<TEL_RES4_CLIENTE xsi:type="xsd:string"></TEL_RES4_CLIENTE>
					<DDD_CEL4_CLIENTE xsi:type="xsd:string"></DDD_CEL4_CLIENTE>
					<TEL_CEL4_CLIENTE xsi:type="xsd:string"></TEL_CEL4_CLIENTE>
					<DDD_COM4_CLIENTE xsi:type="xsd:string"></DDD_COM4_CLIENTE>
					<TEL_COM4_CLIENTE xsi:type="xsd:string"></TEL_COM4_CLIENTE>
					<DDD_OUTROS4_CLIENTE xsi:type="xsd:string"></DDD_OUTROS4_CLIENTE>
					<TEL_OUTROS4_CLIENTE xsi:type="xsd:string"></TEL_OUTROS4_CLIENTE>
					<DDD_TEL_RES_CLIENTE xsi:type="xsd:string"></DDD_TEL_RES_CLIENTE>
					<DDD_TEL_CEL_CLIENTE xsi:type="xsd:string"></DDD_TEL_CEL_CLIENTE>
					<DDD_TEL_COM_CLIENTE xsi:type="xsd:string"></DDD_TEL_COM_CLIENTE>
					<DDD_TEL_OUTROS_CLIENTE xsi:type="xsd:string"></DDD_TEL_OUTROS_CLIENTE>
					<DDD_TEL_RES2_CLIENTE xsi:type="xsd:string"></DDD_TEL_RES2_CLIENTE>
					<DDD_TEL_CEL2_CLIENTE xsi:type="xsd:string"></DDD_TEL_CEL2_CLIENTE>
					<DDD_TEL_COM2_CLIENTE xsi:type="xsd:string"></DDD_TEL_COM2_CLIENTE>
					<DDD_TEL_OUTROS2_CLIENTE xsi:type="xsd:string"></DDD_TEL_OUTROS2_CLIENTE>
					<IDOCORRENCIA xsi:type="xsd:string"></IDOCORRENCIA>
					<CODIGO_OC xsi:type="xsd:string">` + ret.IdTicket + `</CODIGO_OC>
					<CODIGO_OP_ENCAMINHADO xsi:type="xsd:string"></CODIGO_OP_ENCAMINHADO>
					<IDMONITORIA xsi:type="xsd:string"></IDMONITORIA>
					<NOME_DO_ARQUIVO_MONITORIA xsi:type="xsd:string"></NOME_DO_ARQUIVO_MONITORIA>
					<DH_ARQUIVO_GERADO xsi:type="xsd:string"></DH_ARQUIVO_GERADO>
					<COD_STATUS xsi:type="xsd:string"></COD_STATUS>`

	strSoapAction := "urn:server.Multidados#DadosServicedesk"
	strUrn := "urn:DadosServicedesk"
	strXmlSoap := strings.Replace(xmlSoapTmp, "&urn&", strUrn, -1)
	strXmlSoap = strings.Replace(strXmlSoap, "&dados&", strCampos, -1)
	//strXmlSoap = fmt.Sprintf(strXmlSoap, strCampos)

	file, _ := os.Create("teste.xml")
	defer file.Close()

	file.WriteString(strXmlSoap)

	readerSoap := strings.NewReader(strXmlSoap)
	retBytes := soapRequest(strUrn, strSoapAction, "POST", ret.url, readerSoap)

	ret.retorno = retBytes

}

func soapRequest(urn, soapAction, metodo, url string, readerXml *strings.Reader) []byte {

	client := &http.Client{}

	reqSoap, err := http.NewRequest(metodo, url, readerXml)
	reqSoap.Header.Set("SOAPAction", soapAction)
	reqSoap.Header.Add("Content-Type", "text/xml;charset=UTF-8")

	respSoap, err := client.Do(reqSoap)

	b := new(bytes.Buffer)
	b.ReadFrom(respSoap.Body)

	if err != nil {
		panic(err)
	}

	return b.Bytes()

}
