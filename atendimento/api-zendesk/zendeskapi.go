package main

import (
	"fmt"
	//"github.com/gorilla/mux"
	"bytes"
	"encoding/json"
	"html/template"
	"io"
	//"mega/dbg"
	"mega/erro"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
)

type Ticket struct {
	Ticket struct {
		Requester struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"requester"`
		Subject string `json:"subject"`
		Comment struct {
			Body string `json:"body"`
		} `json:"comment"`
	} `json:"ticket"`
}

func main() {
	//TransitarIssueLink("TAG-11")

	//IssueResolution()
	linkIssueTicket()
	/*
		fmt.Println("listening..." + os.Getenv("PORT"))
		http.HandleFunc("/ticket", ticket)
		http.HandleFunc("/ticketLink", ticketLink)
		http.HandleFunc("/issue", issue)
		http.HandleFunc("/components", components)

		err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
		if err != nil {
			panic(err)
		}*/

}

func ticket(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	subject := req.FormValue("subject")
	comment := req.FormValue("comment")

	ticket := Ticket{}
	ticket.Ticket.Requester.Email = "cliente@teste.com"
	ticket.Ticket.Requester.Name = "Cliente"
	ticket.Ticket.Subject = subject
	ticket.Ticket.Comment.Body = comment

	fmt.Println(subject)
	fmt.Println(comment)

	ticketJson, _ := json.MarshalIndent(ticket, "", "  ")
	fmt.Println(string(ticketJson))

	client := &http.Client{}

	reqZen, err := http.NewRequest("POST", "https://megasistemas.zendesk.com/api/v2/tickets.json", bytes.NewBuffer(ticketJson))
	reqZen.SetBasicAuth("andreo@mega.com.br/token", "pzwsg7ItHBIQqX1wzV66sYjl36awkhCXLgyMcjMx")
	reqZen.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(reqZen)
	io.Copy(res, resp.Body)
	fmt.Println(resp.Body, err)
}

type LinkTicket struct {
	AuthToken string `json:"auth_token"`
	Link      Link   `json:"link"`
}

type LinksTicket struct {
	AuthToken string `json:"auth_token"`
	Links     []Link `json:"links"`
}

type Link struct {
	TicketId string `json:"ticket_id"`
	IssueId  string `json:"issue_id"`
}

type IssueExiste struct {
	Issue  int
	Existe bool
}

func linkIssueTicket() {
	lista, err := os.Open("tickets-476.json")
	erro.Trata(err)

	dec := json.NewDecoder(lista)

	var linksTicket LinksTicket
	err = dec.Decode(&linksTicket)
	erro.Trata(err)
	listaIssue := make(map[string]bool)

	chanIssue := make(chan IssueExiste, 10)

	for _, link := range linksTicket.Links {
		listaIssue[link.IssueId] = true
	}
	for issueId := range listaIssue {
		_, err := strconv.Atoi(issueId)
		if err == nil {
			//go func(id int) {
			url := fmt.Sprintf("https://megasistemas.atlassian.net/rest/api/2/issue/%s", issueId)
			fmt.Println(url)
			resp := Request("GET", url, nil)
			if resp.StatusCode == 404 {
				listaIssue[issueId] = false
			}
			//}()
			//fmt.Println(fmt.Sprintf(",%s", issueId))
		}

		/*



			var issues Issues
			sort.Sort(issues.Issues)

			dec := json.NewDecoder(resp.Body)
			err := dec.Decode(&issues)
			erro.Trata(err)

			for _, issue := range issues.Issues {
				if issue.Fields.Status.Name == "Desenv Pronto" {

					for _, version := range issue.Fields.FixVersions {
						if version.Name == "Desenv Pronto" {

					}


				}
			}*/
	}
	for _, link := range linksTicket.Links {
		if !listaIssue[link.IssueId] {
			fmt.Println(link.IssueId, "----", link.TicketId)
		}
	}

}

func ticketLink(res http.ResponseWriter, req *http.Request) {

	acao := req.FormValue("acao")
	issueId := req.FormValue("issueId")
	ticketId := req.FormValue("ticketId")

	if acao == "vinc" {
		res.Write(criaLink(issueId, ticketId))
	} else {
		res.Write(apagaLink(issueId, ticketId))
	}

}

func apagaLink(pIssueId string, pTicketId string) []byte {
	var data LinkTicket

	data.AuthToken = "85882758f0757a03ab0f8c6883281dc3"
	data.Link.IssueId = pIssueId
	data.Link.TicketId = pTicketId

	client := &http.Client{}

	bytePost, err := json.Marshal(data)

	fmt.Printf("%s", string(bytePost))

	readerPost := bytes.NewReader(bytePost)

	reqPOST, err := http.NewRequest("DELETE", "https://jiraplugin.zendesk.com/integrations/jira/account/megasistemas/links/unlink", readerPost)
	reqPOST.Header.Set("Content-Type", "application/json")

	respPost, err := client.Do(reqPOST)

	b := new(bytes.Buffer)
	b.ReadFrom(respPost.Body)

	erro.Trata(err)

	return b.Bytes()
}

func criaLink(pIssueId string, pTicketId string) []byte {
	var data LinkTicket

	data.AuthToken = "85882758f0757a03ab0f8c6883281dc3"
	data.Link.IssueId = pIssueId
	data.Link.TicketId = pTicketId

	client := &http.Client{}

	bytePost, err := json.Marshal(data)

	fmt.Printf("%s", string(bytePost))

	readerPost := bytes.NewReader(bytePost)

	reqPOST, err := http.NewRequest("POST", "https://jiraplugin.zendesk.com/integrations/jira/account/megasistemas/links", readerPost)
	reqPOST.Header.Set("Content-Type", "application/json")

	respPost, err := client.Do(reqPOST)

	b := new(bytes.Buffer)
	b.ReadFrom(respPost.Body)

	erro.Trata(err)

	return b.Bytes()
}

type Issue struct {
	Id     string `json:"id"`
	Key    string `json:"key"`
	Fields struct {
		Project struct {
			Key string `json:"key"`
		} `json:"project"`
		FixVersions []FixVersion `json:"fixVersions,omitempty"`
		Summary     string       `json:"summary"`
		Description string       `json:"description"`
		Issuetype   struct {
			Name string `json:"name"`
		} `json:"issuetype"`
		Components []Component `json:"components,omitempty"`
		Issuelinks []Issuelink `json:"issuelinks,omitempty"`
		Status     struct {
			Name string `json:"name"`
		} `json:"status"`
		Resolution struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"resolution"`
	} `json:"fields"`
}

type Component struct {
	Name string `json:"name"`
}

type FixVersion struct {
	Name string `json:"name"`
}

type Issuelink struct {
	Type struct {
		Name string `json:"name"`
	} `json:"type"`
	OutwardIssue Issue `json:"outwardIssue"`
}

type Retorno struct {
	Id  int    `json:"id"`
	Key string `json:"key"`
}

type Transition struct {
	Transition struct {
		Id int `json:"id"`
	} `json:"transition"`
}

type Resolution struct {
	Key    string `json:"key"`
	Fields struct {
		Resolution struct {
			Id   string `json:"id"`
			Name string `json:"name,omitempty"`
		} `json:"resolution"`
	} `json:"fields"`
}

type Issues struct {
	Issues IssueArray `json:"issues"`
}

type IssueArray []Issue

func (s IssueArray) Len() int {
	return len(s)
}
func (s IssueArray) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s IssueArray) Less(i, j int) bool {
	return s[i].Fields.Components[0].Name < s[j].Fields.Components[0].Name
}

func Request(metodo string, url string, body io.Reader) *http.Response {
	//fmt.Println(metodo, url)
	var client http.Client
	req, err := http.NewRequest(metodo, url, body)
	erro.Trata(err)
	req.SetBasicAuth("remoto", "remoto")
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	erro.Trata(err)
	return resp
}

func issue(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	project := req.FormValue("project")
	issuetype := req.FormValue("issuetype")
	summary := req.FormValue("summary")
	description := req.FormValue("description")
	component := req.FormValue("component")

	statusid, err := strconv.Atoi(req.FormValue("statusid"))
	erro.Trata(err)
	//101 Banco de ideias
	//241 PNF
	//221 Projetos
	fmt.Println("Status id")
	fmt.Println(statusid)

	issue := Issue{}
	issue.Fields.Project.Key = project
	issue.Fields.Issuetype.Name = issuetype
	issue.Fields.Summary = summary
	issue.Fields.Description = description
	if component != "" {
		issue.Fields.Components = make([]Component, 1)
		issue.Fields.Components[0].Name = component
	}

	issueJson, _ := json.MarshalIndent(issue, "", "  ")
	fmt.Println(string(issueJson))

	resp := Request("POST", "https://megasistemas.atlassian.net/rest/api/2/issue.json", bytes.NewBuffer(issueJson))

	var retorno Retorno
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&retorno)

	fmt.Println("Issue Id")
	fmt.Println(retorno.Id)
	fmt.Println("Issue Key")
	fmt.Println(retorno.Key)

	TransitarIssue(statusid, retorno.Key)
}

func TransitarIssue(codTransition int, key string) {

	transition := Transition{}
	transition.Transition.Id = codTransition

	issueTransitionJson, err := json.MarshalIndent(transition, "", "  ")
	fmt.Println(err)
	fmt.Println(string(issueTransitionJson))

	url := fmt.Sprintf("https://megasistemas.atlassian.net/rest/api/2/issue/%s/transitions", key)

	resp := Request("POST", url, bytes.NewBuffer(issueTransitionJson))
	fmt.Println(resp.Body, err)
}

func TransitarIssueLink(key string) {
	url := fmt.Sprintf("https://megasistemas.atlassian.net/rest/api/2/issue/%s", key)
	resp := Request("GET", url, nil)

	var issue Issue
	dec := json.NewDecoder(resp.Body)
	err := dec.Decode(&issue)
	erro.Trata(err)

	for _, r := range issue.Fields.Issuelinks {
		fmt.Println(r.OutwardIssue.Key)
		fmt.Println(r.OutwardIssue.Fields.Status)

	}

}

type Project struct {
	Name       string      `json:"name"`
	Key        string      `json:"key"`
	Components []Component `json:"components,omitempty"`
}

type RetornoComponent struct {
	Name       string   `json:"name"`
	Components []string `json:"components,omitempty"`
}

type RetornoComponents []RetornoComponent

func (s RetornoComponents) Len() int {
	return len(s)
}
func (s RetornoComponents) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s RetornoComponents) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}

type LinhaTable struct {
	Project    string
	Components string
}

func components(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	client := &http.Client{}

	url := "https://megasistemas.atlassian.net/rest/api/2/project"

	reqJira, err := http.NewRequest("GET", url, nil)
	reqJira.SetBasicAuth("remoto", "remoto")
	reqJira.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(reqJira)
	fmt.Println(err)

	var projects []Project

	dec := json.NewDecoder(resp.Body)
	dec.Decode(&projects)

	chanProject := make(chan Project, len(projects))

	for _, project := range projects {

		go RetornaProject(url, project.Key, chanProject)

	}
	fmt.Println("terminou de disparar os projetos")

	retorno := make(RetornoComponents, 0, len(projects))

	for _ = range projects {
		var project Project
		project = <-chanProject
		retornoComponent := RetornoComponent{}
		retornoComponent.Name = project.Name
		retornoComponent.Components = make([]string, len(project.Components))
		for i, component := range project.Components {
			retornoComponent.Components[i] = component.Name
		}
		retorno = append(retorno, retornoComponent)
	}

	sort.Sort(retorno)

	res.Header().Set("Content-Type", "text/html")

	var linhas []LinhaTable

	for _, project := range retorno {
		for _, component := range project.Components {
			linhas = append(linhas, LinhaTable{project.Name, component})
		}

	}
	res.Write([]byte("\n</table>\n</body>\n</html>"))
	t, _ := template.ParseFiles("Component.html")
	t.Execute(res, linhas)

	fmt.Println("terminou de processar os projetos")
	close(chanProject)
}

func RetornaProject(url string, key string, ret chan Project) {
	reqJira, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", url, key), nil)
	reqJira.SetBasicAuth("remoto", "remoto")
	reqJira.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(reqJira)
	if err != nil {
		fmt.Println(err)
	}

	var project Project

	dec := json.NewDecoder(resp.Body)
	dec.Decode(&project)
	ret <- project
}

func IssueResolution() {

	filterResolution := url.QueryEscape(`filter = "Issues prontas - status Unresolved"`)

	query := fmt.Sprintf(`https://megasistemas.atlassian.net/rest/api/2/search?jql=%s&maxResults=1000`, filterResolution)

	resp := Request("GET", query, nil)

	var issues Issues
	sort.Sort(issues.Issues)

	dec := json.NewDecoder(resp.Body)
	err := dec.Decode(&issues)
	erro.Trata(err)

	for _, issue := range issues.Issues {
		AlterarResolution("6", issue.Key)
	}

}

func AlterarResolution(codResolution string, key string) {

	resolution := Resolution{}
	resolution.Key = key
	resolution.Fields.Resolution.Id = codResolution

	issueJson, _ := json.MarshalIndent(resolution, "", "  ")
	fmt.Println(string(issueJson))

	url := fmt.Sprintf("https://megasistemas.atlassian.net/rest/api/2/issue/%s", key)

	resp := Request("PUT", url, bytes.NewBuffer(issueJson))
	fmt.Println(resp.Body)
}
