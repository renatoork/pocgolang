package main

import (
	"bytes"
	"encoding/json"
	"net/http"
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

func IntSugestZd(res http.ResponseWriter, req *http.Request) (http.ResponseWriter, *http.Request) {

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
		return res, req
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

	return res, req

}

func panic_err(err error) {
	if err != nil {
		panic(err)
	}
}
