package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tealeg/xlsx"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	Ticket struct {
		Id           int    `json:"id"`
		Type         string `json:"type"`
		Subject      string `json:"subject"`
		Description  string `json:"description"`
		Organization int    `json:"organization_id"`
		GroupId      int    `json:"group_id"`
		Requester_id int    `json:"requester_id"`
		Requester    struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"requester"`
		Comment      CommentSingle  `json:"comment"`
		Status       string         `json:"status"`
		CustomFields []CustomFields `json:"custom_fields"`
	} `json:"ticket"`
}

type Group struct {
	Group struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"group"`
}

type Attachment struct {
	FileName    string `json:"file_name"`
	ContentType string `json:"content_type"`
}

type CustomFields struct {
	Id    int    `json:"id"`
	Value string `json:"value"`
}

type Comment struct {
	Author_id   int          `json:"author_id"`
	Body        string       `json:"body"`
	Public      bool         `json:"public"`
	Attachments []Attachment `json:"attachments"`
}

type Comments struct {
	Comments []Comment `json:"comments"`
}

type CommentSingle struct {
	Comment Comment `json:"comment"`
}

type UsersToCreate struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Verified bool   `json:"verified"`
}

type Users struct {
	Users []UsersToCreate `json:"users"`
}

type User struct {
	User UsersToCreate `json:"user"`
}

type Search struct {
	Results []UsersToCreate `json:"results"`
}

type Orgs struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	TipoConta string `json:"tipoConta"`
}

type SearchOrg struct {
	Results []struct {
		Id   int    `json:"id"`
		Name string `json:"name"`

		OrgFields struct {
			TipoConta string `json:"tipo_da_conta"`
		} `json:"organization_fields"`
	} `json:"results"`

	NextPage string `json:"next_page"`
	Count    int    `json:"count"`
}

type ArrayOrgs struct {
	Orgs   []Orgs `json:"orgs"`
	Filled bool
}

type TicketComment struct {
	Ticket struct {
		Comment Comment `json:"comment"`
	} `json:"ticket"`
}

type CommentTicket struct {
	Comments []struct {
		CommentSimples string `json:"body"`
		Author         int    `json:"author_id"`
	} `json:"comments"`
}

type Authors struct {
	Authors []struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"authors"`
}

type TicketLista struct {
	Ticket []struct {
		Id            int    `json:"id"`
		Subject       string `json:"subject"`
		Requester_id  int    `json:"requester_id"`
		Requester     string `json:"requester"`
		Assignee      string `json:"assignee"`
		UpdatedAt     string `json:"updated_at"`
		Modulo        string `json:"modulo"`
		status        string `json:"status"`
		ProblemId     int    `json:"problem_id"`
		DataFormatada string `json:"data_formatada"`
	} `json:"tickets"`
}

func main() {

	fmt.Println("Em funcionamento...")
	http.HandleFunc("/intzendesk", TransfTickets)
	http.HandleFunc("/getjson", getJson)
	http.HandleFunc("/savexcel", saveExcel)
	http.HandleFunc("/savejson", saveJson)
	http.HandleFunc("/teste", teste)
	http.HandleFunc("/getAllOrgs", getAllOrgs)
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		panic(err)
	}

}

func getAllOrgs(res http.ResponseWriter, r *http.Request) {

	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Add("Cache-Control", "no-cache")

	url := r.FormValue("url")
	metodo := r.FormValue("metodo")
	query := r.FormValue("query")

	page := 0

	client := &http.Client{}
	var arOrgs ArrayOrgs
	var sOrg SearchOrg

	arOrgs.Filled = false
	index := 0

	for {

		page++

		//fmt.Println(page)
		sOrg = SearchOrg{}

		reqOrgs, _ := http.NewRequest(metodo, url+"?query="+query+"&page="+strconv.Itoa(page), nil)
		reqOrgs.SetBasicAuth("qualiteam@mega.com.br/token", "pzwsg7ItHBIQqX1wzV66sYjl36awkhCXLgyMcjMx")
		reqOrgs.Header.Set("Content-Type", "application/json")
		reqOrgs.Header.Add("Cache-Control", "no-cache")

		respOrgs, _ := client.Do(reqOrgs)

		b := new(bytes.Buffer)
		b.ReadFrom(respOrgs.Body)

		json.Unmarshal(b.Bytes(), &sOrg)

		//fmt.Println("PAGE:: " + sOrg.NextPage)

		if arOrgs.Filled == false {
			arOrgs.Orgs = make([]Orgs, sOrg.Count)
			arOrgs.Filled = true
		}

		if len(sOrg.Results) > 0 {

			for i := 0; i < len(sOrg.Results); i++ {

				arOrgs.Orgs[index].Id = sOrg.Results[i].Id
				arOrgs.Orgs[index].Name = sOrg.Results[i].Name
				arOrgs.Orgs[index].TipoConta = sOrg.Results[i].OrgFields.TipoConta

				index++

			}
		}

		if sOrg.NextPage == "" {
			break
		}

	}

	retorno, _ := json.Marshal(arOrgs)

	res.Write(retorno)

}

type ErroRetorno struct {
	Error struct {
		Message string `json:"message"`
		Title   string `json:"title"`
	} `json:"error"`
}

func logErro(msg string, res http.ResponseWriter) {
	fmt.Println(msg)
	res.Write([]byte(msg))
}

func TransfTickets(res http.ResponseWriter, req *http.Request) {

	client := &http.Client{}

	id_ticket := req.FormValue("id")
	requesterName := req.FormValue("reqname")
	requesterEmail := req.FormValue("reqemail")
	token := req.FormValue("token")
	satisfactionString := req.FormValue("satisfaction")
	score := req.FormValue("score")
	subject := req.FormValue("subject")
	groupId := req.FormValue("group_id")

	// token megsistemas
	//pzwsg7ItHBIQqX1wzV66sYjl36awkhCXLgyMcjMx

	// token megasistemasqualiteam
	// EYS57O8XtM4l1qkRzBq255UOCHz9Ul9jLmesnIdj

	erroPar := false

	i, err := strconv.Atoi(id_ticket)
	if id_ticket == "" || i <= 0 || err != nil {
		logErro(fmt.Sprintf("ERRO: Parametro id não informado ou invalido [%s].", err.Error()), res)
		erroPar = true
	}
	if groupId == "" {
		logErro("ERRO: Parametro group_id não informado.", res)
		erroPar = true

	}
	if subject == "" {
		logErro("ERRO: Parametro subject não informado.", res)
		erroPar = true

	}
	if requesterName == "" {
		logErro("ERRO: Parametro reqname não informado.", res)
		erroPar = true

	}
	if requesterEmail == "" {
		logErro("ERRO: Parametro reqemail não informado.", res)
		erroPar = true

	}
	if score == "" {
		logErro("ERRO: Parametro score não informado.", res)
		erroPar = true

	}
	if token == "" || token != "910ffU42VVQEXzTjVK8W3VMcj854AA5J" {
		logErro("ERRO: Parametro token não informado ou não esta correto.", res)
		erroPar = true
	}

	if erroPar {
		return
	}

	var ticket Ticket
	var group Group

	//Recuperando nome do grupo
	reqGroup, err := http.NewRequest("GET", "https://megasistemas.zendesk.com/api/v2/groups/"+groupId+".json", nil)
	reqGroup.SetBasicAuth("qualiteam@mega.com.br/token", "pzwsg7ItHBIQqX1wzV66sYjl36awkhCXLgyMcjMx")
	reqGroup.Header.Set("Content-Type", "application/json")

	respGroup, err := client.Do(reqGroup)

	b := new(bytes.Buffer)
	b.ReadFrom(respGroup.Body)

	err = json.Unmarshal(b.Bytes(), &group)

	//Atribui dados do requester
	ticket.Ticket.Requester.Name = requesterName
	ticket.Ticket.Requester.Email = requesterEmail
	ticket.Ticket.Subject = subject
	ticket.Ticket.Description = "**Ticket:** " + id_ticket + " \n **Tipo da avaliaçao:** " + score + " \n\n **Avaliação:** \n " + satisfactionString
	ticket.Ticket.CustomFields = make([]CustomFields, 1)
	ticket.Ticket.CustomFields[0].Id = 22930575
	ticket.Ticket.CustomFields[0].Value = strings.ToLower(strings.Replace(group.Group.Name, " ", "_", -1))

	byteTicket, err := json.Marshal(ticket)
	readerTicket := bytes.NewReader(byteTicket)

	reqPostTicket, err := http.NewRequest("POST", "https://megasistemasqualiteam.zendesk.com/api/v2/tickets.json", readerTicket)
	reqPostTicket.SetBasicAuth("betania.oliveira@mega.com.br/token", "EYS57O8XtM4l1qkRzBq255UOCHz9Ul9jLmesnIdj")
	reqPostTicket.Header.Set("Content-Type", "application/json")

	respPostTicket, err := client.Do(reqPostTicket)

	if err != nil {
		fmt.Println(fmt.Sprintf("ERRO:  fatal - Não Gerado Ticket: %s - erro: %s\n", id_ticket, err.Error()))
		// log
		panic_err(err)
	}

	bTicket := new(bytes.Buffer)
	bTicket.ReadFrom(respPostTicket.Body)

	var erroRetornoTicket ErroRetorno
	err = json.Unmarshal(bTicket.Bytes(), &erroRetornoTicket)

	if erroRetornoTicket.Error.Title == "" {
		err = json.Unmarshal(bTicket.Bytes(), &ticket)

		if err != err {
			fmt.Println(fmt.Sprintln("ERRO: fatal: %s \n", err.Error()))
			panic_err(err)
		}

		fmt.Printf(fmt.Sprintf("Gerado ticket: %s\n", string(byteTicket)))

		res.Write(bTicket.Bytes())

	} else {
		fmt.Println(fmt.Sprintf("ERRO: Não Gerado Ticket: %s\n     - erro: %s \n             %s", id_ticket, erroRetornoTicket.Error.Title, erroRetornoTicket.Error.Message))
	}

}

func VerifyUserQualiTeam(email string, client *http.Client) int {

	reqGetUser, err := http.NewRequest("GET", "https://megasistemasqualiteam.zendesk.com/api/v2/search.json?query=type:user%20email:"+email, nil)
	reqGetUser.SetBasicAuth("betania.oliveira@mega.com.br/token", "EYS57O8XtM4l1qkRzBq255UOCHz9Ul9jLmesnIdj")
	reqGetUser.Header.Set("Content-Type", "application/json")

	respUser, err := client.Do(reqGetUser)

	b := new(bytes.Buffer)
	b.ReadFrom(respUser.Body)

	var results Search

	err = json.Unmarshal(b.Bytes(), &results)

	panic_err(err)

	if len(results.Results) > 0 {
		return results.Results[0].Id
	} else {
		return 0
	}
}

func getJson(res http.ResponseWriter, r *http.Request) {

	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Add("Cache-Control", "no-cache")

	urlGet := r.FormValue("url")
	metodo := r.FormValue("metodo")
	per_page := r.FormValue("perpage")
	query := r.FormValue("query")
	page := r.FormValue("page")

	if urlGet != "" {

		client := http.Client{}

		Params := url.Values{}

		if per_page != "" {
			Params.Add("per_page", per_page)
		}

		if query != "" {
			Params.Add("query", query)
		}

		if page != "" {
			Params.Add("page", page)
		}

		urlX := urlGet + "?" + Params.Encode()
		reqTriggers, err := http.NewRequest(metodo, urlX, nil)
		reqTriggers.SetBasicAuth("qualiteam@mega.com.br/token", "pzwsg7ItHBIQqX1wzV66sYjl36awkhCXLgyMcjMx")
		reqTriggers.Header.Set("Content-Type", "application/json")
		reqTriggers.Header.Add("Cache-Control", "no-cache")

		resp, err := client.Do(reqTriggers)

		if err != nil {
			fmt.Println(fmt.Sprintf("Erro: %s\n      %s\n", err.Error(), urlX))
			panic(err)
		}

		b := new(bytes.Buffer)
		b.ReadFrom(resp.Body)

		res.Write(b.Bytes())

	} else {

		res.Write([]byte("Parametros devem ser informados!"))

	}

}

/*func saveExcel(res http.ResponseWriter, r *http.Request) {

	res.Header().Set("Access-Control-Allow-Origin", "*")

	vIdTicket := r.FormValue("idticket")
	vAuthors := r.FormValue("authors")
	vStringName := "../tmp/ListaPendencia.xlsx"

	if vIdTicket != "" {

		var objJson CommentTicket
		var objAuthors Authors
		var fileXls *xlsx.File
		var sheet *xlsx.Sheet
		var row *xlsx.Row
		var cellUser *xlsx.Cell
		var cellComentario *xlsx.Cell
		var styleCellBg *xlsx.Style
		//var styleCellFont *xlsx.Style
		var arrayAuthors map[int]string

		client := http.Client{}

		reqTicketComment, err := http.NewRequest("GET", "https://megasistemas.zendesk.com/api/v2/tickets/"+vIdTicket+"/comments.json", nil)
		reqTicketComment.SetBasicAuth("qualiteam@mega.com.br/token", "pzwsg7ItHBIQqX1wzV66sYjl36awkhCXLgyMcjMx")
		reqTicketComment.Header.Set("Content-Type", "application/json")
		reqTicketComment.Header.Add("Cache-Control", "no-cache")

		respTicketComment, err := client.Do(reqTicketComment)

		b := new(bytes.Buffer)
		b.ReadFrom(respTicketComment.Body)

		err = json.Unmarshal(b.Bytes(), &objJson)
		err = json.Unmarshal([]byte(vAuthors), &objAuthors)

		fileXls = xlsx.NewFile()
		sheet = fileXls.AddSheet("ListaDePendencias")

		arrayAuthors = make(map[int]string)

		for i := 0; i < len(objAuthors.Authors); i++ {
			arrayAuthors[objAuthors.Authors[i].Id] = objAuthors.Authors[i].Name
		}

		styleCellBg = xlsx.NewStyle()
		styleCellBg.Font = *xlsx.NewFont(12, "Helvetica")
		styleCellBg.Fill = *xlsx.NewFill("solid", "", "333")
		styleCellBg.ApplyFont = true
		styleCellBg.ApplyFill = true

		row = sheet.AddRow()
		cellUser = row.AddCell()
		cellComentario = row.AddCell()

		cellUser.SetStyle(*styleCellBg)

		cellUser.Value = "Ticket Id"
		cellComentario.Value = "Comentário"

		for i := 0; i < len(objJson.Comments); i++ {

			row = sheet.AddRow()
			cellUser = row.AddCell()
			cellComentario = row.AddCell()

			/*if i%2 == 0 {

				styleCellBg = cellUser.GetStyle()
				styleCellBg.Fill.BgColor = "#333"

				cellUser.SetStyle(styleCellBg)
				cellComentario.SetStyle(styleCellBg)

			} * /

			fmt.Println(arrayAuthors[objJson.Comments[i].Author])

			cellUser.Value = arrayAuthors[objJson.Comments[i].Author]
			cellComentario.Value = objJson.Comments[i].CommentSimples

		}

		err = fileXls.Save(vStringName)
		fileDownload, err := os.Open(vStringName)
		statDownload, err := fileDownload.Stat()

		byteDownload := make([]byte, statDownload.Size())

		_, err = fileDownload.Read(byteDownload)

		bReader := bytes.NewReader(byteDownload)

		panic_err(err)

		res.Header().Add("Content-Disposition", "attachment; filename="+"ListaDePendencias.xlsx")
		res.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		res.Header().Add("Content-Length", strconv.FormatInt(statDownload.Size(), 10))

		io.Copy(res, bReader)

	}

}*/

func saveExcel(res http.ResponseWriter, r *http.Request) {

	res.Header().Set("Access-Control-Allow-Origin", "*")

	//vStringName := "ListaTickets.json"
	vStringName := "../tmp/ListaTickets.json"

	vXlsxArq := "../tmp/ListaPendencia.xlsx"
	var objJson TicketLista
	var fileXls *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cellId *xlsx.Cell
	var cellAssunto *xlsx.Cell
	var cellOrigem *xlsx.Cell
	var cellSolicitante *xlsx.Cell
	var cellAtribuido *xlsx.Cell
	var cellUltAtividade *xlsx.Cell
	var cellStatus *xlsx.Cell
	var cellIdProblema *xlsx.Cell

	fileDownload, err := os.Open(vStringName)
	defer fileDownload.Close()

	if err == nil {

		statFile, err := fileDownload.Stat()
		stringByte := make([]byte, statFile.Size())
		_, err = fileDownload.Read(stringByte)

		err = json.Unmarshal(stringByte, &objJson)

		fileXls = xlsx.NewFile()
		sheet = fileXls.AddSheet("ListaDePendencias")

		row = sheet.AddRow()
		cellId = row.AddCell()
		cellAssunto = row.AddCell()
		cellAtribuido = row.AddCell()
		cellSolicitante = row.AddCell()
		cellOrigem = row.AddCell()
		cellUltAtividade = row.AddCell()
		cellStatus = row.AddCell()
		cellIdProblema = row.AddCell()

		cellId.Value = "Id"
		cellAssunto.Value = "Assunto"
		cellAtribuido.Value = "Atribuído"
		cellSolicitante.Value = "Solicitante"
		cellOrigem.Value = "Origem"
		cellUltAtividade.Value = "Ult. Atividade"
		cellStatus.Value = "Status"
		cellIdProblema.Value = "Id Problema"

		for i := 0; i < len(objJson.Ticket); i++ {

			row = sheet.AddRow()

			cellId = row.AddCell()
			cellAssunto = row.AddCell()
			cellAtribuido = row.AddCell()
			cellSolicitante = row.AddCell()
			cellOrigem = row.AddCell()
			cellUltAtividade = row.AddCell()
			cellStatus = row.AddCell()
			cellIdProblema = row.AddCell()

			cellId.Value = strconv.Itoa(objJson.Ticket[i].Id)
			cellAssunto.Value = objJson.Ticket[i].Subject
			cellAtribuido.Value = objJson.Ticket[i].Assignee
			cellSolicitante.Value = objJson.Ticket[i].Requester
			cellOrigem.Value = objJson.Ticket[i].Modulo
			cellUltAtividade.Value = objJson.Ticket[i].DataFormatada
			cellStatus.Value = objJson.Ticket[i].status
			cellIdProblema.Value = strconv.Itoa(objJson.Ticket[i].ProblemId)

		}

		err = fileXls.Save(vXlsxArq)
		fileDownload, err := os.Open(vXlsxArq)
		defer fileDownload.Close()
		statDownload, err := fileDownload.Stat()

		byteDownload := make([]byte, statDownload.Size())

		fileDownload.Read(byteDownload)

		bReader := bytes.NewReader(byteDownload)

		res.Header().Add("Content-Disposition", "attachment; filename="+"ListaDePendencias.xlsx")
		res.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		res.Header().Add("Content-Length", strconv.FormatInt(statDownload.Size(), 10))

		panic_err(err)

		io.Copy(res, bReader)

	} else {

		res.Write([]byte(`{ "status": "Não existe arquivo JSON." }`))

	}
}

func saveJson(res http.ResponseWriter, r *http.Request) {

	res.Header().Set("Access-Control-Allow-Origin", "*")

	vStringName := "../tmp/ListaTickets.json"
	//vStringName := "ListaTickets.json"

	vJson := r.FormValue("json")

	if vJson != "" {

		fileJson, err := os.Create(vStringName)
		defer fileJson.Close()

		fileJson.WriteString(vJson)

		panic_err(err)

		res.Write([]byte(`{ "status": "Arquivo criado" }`))
		return

	}

	res.Write([]byte(`{ "status": "Houve erro" }`))
	return

}

func panic_err(err error) {

	if err != nil {
		panic(err)
	}

}

func teste(res http.ResponseWriter, req *http.Request) {

	reqBinLog, _ := http.NewRequest("POST", "http://requestb.in/1mg4ak91?cod="+req.FormValue("cod"), nil)

	client := &http.Client{}

	resp, _ := client.Do(reqBinLog)

	b := new(bytes.Buffer)
	b.ReadFrom(resp.Body)

	res.Write(b.Bytes())

}
