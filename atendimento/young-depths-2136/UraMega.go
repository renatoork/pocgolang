package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	Ticket struct {
		Id             int `json:"id"`
		RequesterId    int `json:"requester_id"`
		OrganizationId int `json:"organization_id"`
		AssigneeId     int `json:"assignee_id"`
		GroupId        int `json:"group_id"`
		CustomFields   []struct {
			Id    int    `json:"id"`
			Value string `json:"value"`
		} `json:"custom_fields"`
	} `json:"ticket"`
	Error string `json:"error"`
}

type Retorno struct {
	Ramal  int    `json:"ramal"`
	Grupo  int    `json:"grupo"`
	Status string `json:"status"`
}

type Organization struct {
	Organization struct {
		Id int `json:"id"`
	} `json: "organization"`
	OrganizationFields struct {
		Status string `json:"status"`
	} `json:"organization_fields"`
	Error string `json:"error"`
}

type Results struct {
	Id                 int    `json:"id"`
	Name               string `json:"name"`
	OrganizationFields struct {
		Status     string `json:"status"`
		CanalAtend string `json:"canal_de_atendimento"`
		Segmento   string `json:"segmento"`
	} `json:"organization_fields"`
}

type ResultsUser struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	UserFields struct {
		Cargo  string `json:"cargo"`
		Status string `json:"status"`
	} `json:"organization_fields"`
}

type Search struct {
	Results []Results
}

type SearchUser struct {
	Results []ResultsUser
}

type User struct {
	User struct {
		Id             int    `json:"id"`
		Phone          string `json:"phone"`
		OrganizationId int    `json:"organization_id"`
	} `json:"user`
	Error string `json:"error"`
}

type GruposMega struct {
	Id   int
	Name string
}

type arGruposMega []GruposMega

func main() {

	fmt.Printf("%s", "Em funcionamento...")
	http.HandleFunc("/ura", URA)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)

	panic_err(err)

}

func URA(res http.ResponseWriter, req *http.Request) {

	token := req.FormValue("token")
	ura := req.FormValue("ura")

	emailApi := "qualiteam@mega.com.br"

	fmt.Println("URA: " + ura + " --- token: " + token)

	if token != "R1C2BogK3KKBxx1N2CoyLU1qTts4Wl67" || ura == "" {
		res.Write([]byte("Favor informar todos os parâmetros."))
		return
	}

	client := &http.Client{}

	var arGMega []GruposMega

	//Popula grupos MEGA
	arGMega = PopulaGruposMega(arGMega, "Resolvedores", 41)
	arGMega = PopulaGruposMega(arGMega, "Gestão de Mudanças", 2)
	arGMega = PopulaGruposMega(arGMega, "Desbravadores", 31)
	arGMega = PopulaGruposMega(arGMega, "Conhecedores Infra", 47)

	/*Esses itens deixarão de existir*/
	arGMega = PopulaGruposMega(arGMega, "Conhecedores Incorporação - Construção", 51)
	arGMega = PopulaGruposMega(arGMega, "Conhecedores Engenharia - Construção", 6)

	arGMega = PopulaGruposMega(arGMega, "Conhecedores - Construção Itu", 7)
	arGMega = PopulaGruposMega(arGMega, "Conhecedores - RH", 81)
	arGMega = PopulaGruposMega(arGMega, "Conhecedores - Manufatura", 30)
	arGMega = PopulaGruposMega(arGMega, "Conhecedores - Logística", 40)
	arGMega = PopulaGruposMega(arGMega, "Administraçao", 8602)

	tipo := func() string {
		if strings.Contains(ura, "#") {
			ura = strings.Replace(ura, "#", "", -1)
			return "C"
		} else {
			return "T"
		}
	}()

	fmt.Println("tipo: " + tipo)

	//codGrupo := 0
	var retorno Retorno

	if tipo == "T" {
		reqTicket, err := http.NewRequest("GET", "https://megasistemas.zendesk.com/api/v2/tickets/"+ura+".json", nil)
		reqTicket.SetBasicAuth(emailApi+"/token", "pzwsg7ItHBIQqX1wzV66sYjl36awkhCXLgyMcjMx")
		reqTicket.Header.Set("Content-Type", "application/json")

		respTicket, err := client.Do(reqTicket)

		b := new(bytes.Buffer)
		b.ReadFrom(respTicket.Body)

		var ticket Ticket
		var user User

		err = json.Unmarshal(b.Bytes(), &ticket)

		panic_err(err)

		if ticket.Error != "" {
			res.Write(MarshalObj(0, 0, "", retorno))
			return
		} else {

			if ticket.Ticket.AssigneeId != 0 {

				reqUser, err := http.NewRequest("GET", "https://megasistemas.zendesk.com/api/v2/users/"+strconv.Itoa(ticket.Ticket.AssigneeId)+".json", nil)
				reqUser.SetBasicAuth(emailApi+"/token", "pzwsg7ItHBIQqX1wzV66sYjl36awkhCXLgyMcjMx")
				reqUser.Header.Set("Content-Type", "application/json")

				respUser, err := client.Do(reqUser)

				b := new(bytes.Buffer)
				b.ReadFrom(respUser.Body)

				err = json.Unmarshal(b.Bytes(), &user)

				if user.Error != "" {
					res.Write(MarshalObj(0, 0, "", retorno))
					return
				}

				if user.User.Phone == "" || user.User.Phone == "0" {
					res.Write(MarshalObj(arGMega[EvalCodGrupo(ticket.Ticket.GroupId, "")].Id, 0, "S", retorno))
					return
				}

				ramalUser, err := strconv.Atoi(user.User.Phone)

				panic_err(err)

				res.Write(MarshalObj(0, ramalUser, "S", retorno))
				return
			} else if ticket.Ticket.GroupId != 0 {

				ramalDesbrav, err := strconv.Atoi(achaDesbravGrupo(ticket.Ticket.GroupId, emailApi))

				if err == nil {
					res.Write(MarshalObj(ramalDesbrav, 0, "S", retorno))
				} else {
					res.Write(MarshalObj(arGMega[EvalCodGrupo(ticket.Ticket.GroupId, "")].Id, 0, "S", retorno))
				}

				return

			} else {

				res.Write(MarshalObj(0, 0, "", retorno))
				return

			}

		}

	} else {

		var org Search

		fmt.Println("URA: " + ura)

		reqOrg, err := http.NewRequest("GET", "https://megasistemas.zendesk.com/api/v2/search.json?query=type:organization%20cdigo_cliente:"+ura, nil)
		reqOrg.SetBasicAuth(emailApi+"/token", "pzwsg7ItHBIQqX1wzV66sYjl36awkhCXLgyMcjMx")
		reqOrg.Header.Set("Content-Type", "application/json")

		respOrg, err := client.Do(reqOrg)

		b := new(bytes.Buffer)
		b.ReadFrom(respOrg.Body)

		err = json.Unmarshal(b.Bytes(), &org)

		panic_err(err)

		if len(org.Results) > 0 {

			fmt.Println("Contrato: " + org.Results[0].OrganizationFields.Status)
			fmt.Println("len STATUS: " + strconv.Itoa(len(org.Results[0].OrganizationFields.Status)))

			if org.Results[0].OrganizationFields.Status != "contrato_ativo" && len(org.Results[0].OrganizationFields.Status) > 0 {
				res.Write(MarshalObj(arGMega[EvalCodGrupo(0, "A")].Id, 0, "N", retorno))
				return
			}

			// EXCECOES - ATENDIMENTO
			var userSearch SearchUser
			if org.Results[0].OrganizationFields.CanalAtend == "atend_mega_sa_itu" && org.Results[0].OrganizationFields.Segmento == "segmento_manufatura" {

				reqUsuDesbrav, err := http.NewRequest("GET", "https://megasistemas.zendesk.com/api/v2/search.json?query=type:user%20group:21449070%20cargo:cargo_analista_desbravador", nil)
				reqUsuDesbrav.SetBasicAuth(emailApi+"/token", "pzwsg7ItHBIQqX1wzV66sYjl36awkhCXLgyMcjMx")
				reqUsuDesbrav.Header.Set("Content-Type", "application/json")

				respUsuDesbrav, err := client.Do(reqUsuDesbrav)

				b := new(bytes.Buffer)
				b.ReadFrom(respUsuDesbrav.Body)

				json.Unmarshal(b.Bytes(), &userSearch)

				panic_err(err)

				if len(userSearch.Results) > 0 {

					ramalDesbrav, err := strconv.Atoi(userSearch.Results[0].Phone)

					panic_err(err)

					res.Write(MarshalObj(0, ramalDesbrav, "S", retorno))
				}
				return

			} else if org.Results[0].OrganizationFields.CanalAtend == "atend_mega_sa_itu" && org.Results[0].OrganizationFields.Segmento == "segmento_construcao" {

				ramalDesbrav, err := strconv.Atoi(achaDesbravGrupo(21252930, emailApi))

				if err == nil {
					res.Write(MarshalObj(0, ramalDesbrav, "S", retorno))
				}

				return

			} else if org.Results[0].OrganizationFields.CanalAtend == "atend_informaction" {

				ramalDesbrav, err := strconv.Atoi(achaDesbravGrupo(21538024, emailApi))

				if err == nil {
					res.Write(MarshalObj(0, ramalDesbrav, "S", retorno))
				}

				return

			} else {

				//res.Write(MarshalObj(0, 8670, "S", retorno))

				ramalDesbrav, err := strconv.Atoi(achaDesbravGrupo(21252940, emailApi))

				if err == nil {
					res.Write(MarshalObj(0, ramalDesbrav, "S", retorno))
				}

				return

			}
			//FIM DE EXCECAO

			res.Write(MarshalObj(arGMega[EvalCodGrupo(0, "D")].Id, 0, "S", retorno))
			return

		} else {
			res.Write(MarshalObj(0, 0, "", retorno))
			return
		}
	}

}

func panic_err(err error) {

	if err != nil {

		panic(err)

	}

}

func MarshalObj(grupo, ramal int, status string, ret Retorno) []byte {

	ret.Grupo = grupo
	ret.Ramal = ramal
	ret.Status = status

	byteRetorno, err := json.Marshal(ret)

	panic_err(err)

	return byteRetorno

}

func EvalCodGrupo(codGrupo int, grupo string) int {

	retorno := 0

	fmt.Println("EvalCodGrupo / Grupo: " + grupo)

	if grupo != "" {

		switch {
		case grupo == "D":
			return 2
		case grupo == "A":
			return 10
		}

	} else {

		switch {
		case codGrupo == 21148614:
			return 0
		case codGrupo == 21252950:
			return 1
		case codGrupo == 21229524:
			return 2
		case codGrupo == 21148624:
			return 3
		case codGrupo == 21506730:
			return 4
		case codGrupo == 21094684:
			return 5
		case codGrupo == 21252930:
			return 6
		case codGrupo == 21148604:
			return 7
		case codGrupo == 21449070:
			return 8
		case codGrupo == 21252940:
			return 9
		}

	}

	return retorno
}

func PopulaGruposMega(g []GruposMega, nome string, codigo int) []GruposMega {

	var arTemp GruposMega

	arTemp.Id = codigo
	arTemp.Name = nome

	g = append(g, arTemp)

	return g

}

func achaDesbravGrupo(groupId int, emailApi string) string {

	strGroupId := strconv.Itoa(groupId)

	client := http.Client{}

	reqUsuDesbrav, err := http.NewRequest("GET", "https://megasistemas.zendesk.com/api/v2/search.json?query=type:user%20group:"+strGroupId+"%20cargo:cargo_analista_desbravador", nil)
	reqUsuDesbrav.SetBasicAuth(emailApi+"/token", "pzwsg7ItHBIQqX1wzV66sYjl36awkhCXLgyMcjMx")
	reqUsuDesbrav.Header.Set("Content-Type", "application/json")

	respUsuDesbrav, err := client.Do(reqUsuDesbrav)

	b := new(bytes.Buffer)
	b.ReadFrom(respUsuDesbrav.Body)

	var userSearch SearchUser

	json.Unmarshal(b.Bytes(), &userSearch)

	panic_err(err)

	if len(userSearch.Results) > 0 {
		return userSearch.Results[0].Phone
	}

	return ""

}

func extraiValorCustomField(tick Ticket, valueCustomField int) string {

	for i := 0; i < len(tick.Ticket.CustomFields); i++ {
		if tick.Ticket.CustomFields[i].Id == valueCustomField {
			return tick.Ticket.CustomFields[i].Value
		}
	}

	return ""

}
