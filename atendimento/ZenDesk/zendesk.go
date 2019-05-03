package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tgulacsi/goracle/oracle"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	//"time"
)

var pathExe string

func main() {
	//pathExe = fmt.Sprintf("%s\\arquivos\\", filepath.Dir(os.Args[0]))
	pathExe = filepath.Dir(os.Args[0])
	if len(os.Args) > 1 {
		if strings.Contains(strings.ToUpper(os.Args[1]), "CARREGATICKETDOZEN") {
			fmt.Println("CARREGATICKETDOZEN")
			buscarTicketOrg()
			return
		}
		if strings.Contains(strings.ToUpper(os.Args[1]), "CRIATICKETSQUESTION") {
			fmt.Println("CRIATICKETSQUESTION")
			buscaTicketTipo("question")
			return
		}
		if strings.Contains(strings.ToUpper(os.Args[1]), "CRIATICKETSINCIDENT") {
			fmt.Println("CRIATICKETSINCIDENT")
			buscaTicketTipo("incident")
			return
		}
		if strings.Contains(strings.ToUpper(os.Args[1]), "CARREGACOMMENTSDOZEN") {
			fmt.Println("CARREGACOMMENTSDOZEN")
			exportaComents()
			return
		}
		if strings.Contains(strings.ToUpper(os.Args[1]), "EXCLUITICKETNOZEN") {
			fmt.Println("EXCLUITICKETNOZEN")
			excluirTickets()
			return
		}
		if strings.Contains(strings.ToUpper(os.Args[1]), "EXCLUIRCRIAQUESTIONINCIDENTE") {
			fmt.Println("EXCLUIRCRIAQUESTIONINCIDENTE")
			excluirTickets()
			buscaTicketTipo("question")
			buscaTicketTipo("incident")
			return
		}
		if strings.Contains(strings.ToUpper(os.Args[1]), "ALTERATICKETATRIBUICAO") {
			fmt.Println("ALTERATICKETATRIBUICAO")
			alteraTicket()
			return
		}
		if strings.Contains(strings.ToUpper(os.Args[1]), "INSERIRANEXO") {
			fmt.Println("INSERIRANEXO")
			insereAnexo()
			return
		}
		if strings.Contains(strings.ToUpper(os.Args[1]), "DELETARANEXO") {
			fmt.Println("DELETARANEXO")
			deleteAnexo(os.Args[2])
			return
		}
		if strings.Contains(strings.ToUpper(os.Args[1]), "VINCULARANEXO") {
			fmt.Println("VINCULARANEXO")
			pId, _ := strconv.Atoi(os.Args[2])
			vincularAnexoTicket(int32(pId), os.Args[3])
			return
		}
		if strings.Contains(strings.ToUpper(os.Args[1]), "TESTEORACLE") {
			fmt.Println("TESTEORACLE")
			testeOracle()
			return
		}

		if strings.Contains(strings.ToUpper(os.Args[1]), "HELP") {
			fmt.Println("Parametros: \n CARREGATICKETDOZEN \n CRIATICKETSQUESTION \n CRIATICKETSINCIDENT \n CARREGACOMMENTSDOZEN \n EXCLUITICKETNOZEN \n EXCLUIRCRIAQUESTIONINCIDENTE \n ALTERATICKETATRIBUICAO \n INSERIRANEXO")
			return
		}

		fmt.Println("Parametro inválido. Execute ZenDesk HELP")

	}
}

type CustomFields struct {
	Id    int32  `json:"id"`
	Value string `json:"value"`
}

type Comment struct {
	Type      string `json:"type"`
	Body      string `json:"body"`
	Public    string `json:"public"`
	Author_id int32  `json:"author_id"`
}

type Ticket struct {
	Subject       string         `json:"subject"`
	Type          string         `json:"type"`
	AssigneeID    int32          `json:"assignee_id"`
	GroupID       string         `json:"group_id"`
	RequesterID   int32          `json:"requester_id"`
	Priority      string         `json:"priority"`
	Status        string         `json:"status"`
	ProblemID     int32          `json:"problem_id"`
	Tags          []string       `json:"tags"`
	Comment       Comment        `json:"comment"`
	Custom_fields []CustomFields `json:"custom_fields"`
}

type Tickets struct {
	Ticket Ticket `json:"ticket"`
}

type TicketC struct {
	Comment Comment `json:"comment"`
}

type TicketCom struct {
	Ticket TicketC `json:"ticket"`
}

type ImpTicket struct {
	Count        float64     `json:"count"`
	NextPage     interface{} `json:"next_page"`
	PreviousPage interface{} `json:"previous_page"`
	Tickets      []struct {
		AssigneeID     float64 `json:"assignee_id"`
		CollaboratorId float64 `json:"collaborator_id"`
		Subject        string  `json:"subject"`
		ID             float64 `json:"id"`
		OrganizationID float64 `json:"organization_id"`
		ProblemID      float64 `json:"problem_id"`
		RequesterID    float64 `json:"requester_id"`
		Status         string  `json:"status"`
		SubmitterID    float64 `json:"submitter_id"`
		Type           string  `json:"type"`
		Via            struct {
			Channel string `json:"channel"`
		} `json:"via"`
	} `json:"tickets"`
}

type TicketR struct {
	Id int32 `json:"id"`
}

type TicketRet struct {
	Ticket TicketR `json:"ticket"`
}

func getConn() *oracle.Connection {
	var conn *oracle.Connection
	dsn := "mgtsk/megatsk@10.0.0.216:1521/orcl"
	//dsn := "mega/megamega@PC_RENATO:1521/orcl"

	user, passw, sid := oracle.SplitDSN(dsn)
	var err error
	conn, err = oracle.NewConnection(user, passw, sid, true)
	if err != nil {
		fmt.Println(fmt.Sprintf("Erro ao criar a conexão com [%s]: %s", dsn, err))
		return nil
	}
	return conn
}

func testeOracle() {
	conn := getConn()
	defer conn.Close()
	if conn.IsConnected() {
		fmt.Println("Conectado...")
		conn.Close()
	} else {
		fmt.Println("NÃO Conectado...")
	}
}

func getTicketOrg(id int32) {
	client := &http.Client{}
	url := fmt.Sprintf("https://megasistemas.zendesk.com/api/v2/tickets.json?organization_id=%d", id)
	for {
		reqZen, _ := http.NewRequest("GET", url, nil)
		reqZen.SetBasicAuth("qualiteam@mega.com.br/token", "pzwsg7ItHBIQqX1wzV66sYjl36awkhCXLgyMcjMx")
		reqZen.Header.Set("Content-Type", "application/json")
		resp, ErrC := client.Do(reqZen)
		if ErrC != nil {
			fmt.Println(fmt.Sprintf("Erro ao carregar os tickets da organizacao: %s", ErrC.Error()))
			break
		}

		body, _ := ioutil.ReadAll(resp.Body)
		var tkt ImpTicket
		json.Unmarshal(body, &tkt)

		inserirTicketOrg(tkt)

		if tkt.NextPage == nil {
			break
		}
		url = tkt.NextPage.(string)
	}

}

func buscarTicketOrg() {

	db := getConn()
	defer conn.Close()

	cur := db.NewCursor()
	defer cur.Close()

	cmd := "select UNO_IN_ID_ZENDESK ID from TSK_VW_TICKETSORGIMP"
	if err := cur.Execute(cmd, nil, nil); err != nil {
		fmt.Println(fmt.Sprintf(`Erro ao executar "%s": %s`, cmd, err))
	} else {
		rows, errF := cur.FetchAll()
		if errF != nil {
			fmt.Println(fmt.Sprintf("Erro no Fetch: %s", errF))
		} else {
			for _, row := range rows {
				fmt.Println("org: ", row[0].(int32))
				getTicketOrg(row[0].(int32))

			}
		}
	}
}

func inserirTicketOrg(tkt ImpTicket) {

	var cmd string

	db := getConn()
	defer conn.Close()
	if db.IsConnected() {

		cur := db.NewCursor()
		defer cur.Close()
		fmt.Println("  Qtde Tickets: ", tkt.Count)
		for _, f := range tkt.Tickets {

			cmd = "INSERT INTO SYSTEM.DEPARA_TICKET_TAREFA (TICKET_ID,TICKET_CHANEL,TICKET_DESCRIPTION,TICKET_STATUS,TICKET_REQUESTER_ID,TICKET_ASSIGNEE,TICKET_ORGANIZATION,TICKET_PROBLEM_ID, TICKET_TYPE) " +
				"values(:ID,:CHANEL,:DESCRIPTION,:STATUS,:REQUESTER_ID,:ASSIGNEE,:ORGANIZATION,:PROBLEM_ID, :TICKET_TYPE)"

			param := make([]interface{}, 9)
			param[0] = f.ID
			param[1] = f.Via.Channel
			param[2] = strings.Replace(f.Subject, "log", "xl2p1", -1)
			param[3] = f.Status
			param[4] = f.RequesterID
			param[5] = f.AssigneeID
			param[6] = f.OrganizationID
			param[7] = f.ProblemID
			param[8] = f.Type

			if err := cur.Execute(cmd, param, nil); err != nil {
				fmt.Println("cannot insert into DEPARA_TICKET_TAREFA : %s", err)
			}
		}
	}
}

func buscaTicketTipo(tipo string) {

	db := getConn()
	defer conn.Close()

	cur := db.NewCursor()
	defer cur.Close()

	cmd := fmt.Sprintf("select * from mgtsk.Ticket where type = '%s'", tipo)
	if err := cur.Execute(cmd, nil, nil); err != nil {
		fmt.Println(fmt.Sprintf(`Erro ao executar "%s": %s`, cmd, err))
	} else {
		for {

			row, _ := cur.FetchOne()
			if row == nil {
				break
			}
			var tkt Tickets
			var tk Ticket
			if row[0] != nil {
				tk.Subject = row[0].(string)
			}
			if row[1] != nil {
				tk.Type = row[1].(string)
			}
			if row[2] != nil {
				tk.AssigneeID = row[2].(int32)
			}
			if row[3] != nil {
				tk.GroupID = row[3].(string)
			}
			if row[4] != nil {
				tk.RequesterID = row[4].(int32)
			}
			if row[5] != nil {
				tk.Priority = row[5].(string)
			}
			if row[6] != nil {
				tk.Status = row[6].(string)
			}
			if row[7] != nil {
				tk.ProblemID = row[7].(int32)
			}
			if row[8] != nil {
				tk.Tags = []string{row[8].(string)}
			}

			if row[10] != nil {
				tk.Comment.Author_id = row[10].(int32)
			}
			if row[9] != nil {
				var ret1 *oracle.ExternalLobVar
				ret1, _ = row[9].(*oracle.ExternalLobVar)
				aux, _ := ret1.ReadAll()
				tk.Comment.Body = string(aux)
				tk.Comment.Public = "true"
				tk.Comment.Type = "Comment"
			}
			var c CustomFields
			c.Id = 22642794
			if row[11] != nil {
				c.Value = row[11].(string)
			}
			tk.Custom_fields = append(tk.Custom_fields, c)
			c.Id = 22861460
			c.Value = row[12].(string)
			tk.Custom_fields = append(tk.Custom_fields, c)
			c.Id = 22639154
			c.Value = row[13].(string)
			tk.Custom_fields = append(tk.Custom_fields, c)
			c.Id = 22801000
			c.Value = row[14].(string)
			tk.Custom_fields = append(tk.Custom_fields, c)
			c.Id = 22648794
			c.Value = row[15].(string)
			tk.Custom_fields = append(tk.Custom_fields, c)

			tkt.Ticket = tk

			if tipo == "incident" {
				ver := " "
				if row[17] != nil {
					ver = row[17].(string)
				}
				var loc int32
				if row[18] != nil {
					loc = row[18].(int32)
				}
				var taridpm int32
				if row[19] != nil {
					taridpm = row[19].(int32)
				}
				tkt.Ticket.ProblemID = criaTicketZenProblem(tkt, ver, loc, taridpm)
			}

			tkJson, _ := json.Marshal(tkt)

			id := criaTicketZen(tkJson)

			atualizaTextoTarefa(row[16].(int32), id)

			marcaDeparaTicketTarefa(row[16].(int32), id, tkt.Ticket.ProblemID)

		}
	}

}

func marcaDeparaTicketTarefa(tar int32, id int32, idProblem int32) {
	var cmd string

	db := getConn()
	defer conn.Close()
	if db.IsConnected() {

		cur := db.NewCursor()
		defer cur.Close()
		cmd = fmt.Sprintf("UPDATE system.PRE_IMPORTACAO set IMPORTADO = '%s', TICKET_IN_ID = %d, PROBLEM_ID = %d where Tar_In_Id=%d", "S", id, idProblem, tar)

		if err := cur.Execute(cmd, nil, nil); err != nil {
			fmt.Println("cannot update into PRE_IMPORTACAO : %s", err)
		}
	}
}

func atualizaTextoTarefa(tar int32, id int32) {
	db := getConn()
	defer conn.Close()

	//fmt.Println("TAR: ", tar)
	cur := db.NewCursor()
	defer cur.Close()
	cmd := fmt.Sprintf("Select Replace(Txx.Tte_Bl_Texto, Chr(10), ' \n ') Tte_Bl_Texto  "+
		"      ,Txx.Tte_Bo_Visualiza  "+
		"	      ,nvl(u.Usu_In_Id_Zendesk, 631566480) Usu_In_Id_Zendesk "+
		"	      , Case "+
		"	          When Uno.Pai_Uno_In_Id = 3 Then "+
		"	           'M' "+
		"	          Else "+
		"	           'C' "+
		"	        End Tipo_Usu "+
		"	  From Mgtsk.Tsk_Tarefatexto       Txx "+
		"	      ,Mgtsk.Tsk_Usuario           u "+
		"	      ,Mgtsk.Tsk_Unidadeorg_Cmpesp Uno "+
		"	 Where Txx.Tar_In_Id =  %d"+
		"	   And Txx.Tex_In_Id Not In (1) "+
		"	   And Txx.Usu_In_Inclusao = u.Usu_In_Id "+
		"	   And u.Uno_In_Id = Uno.Uno_In_Id "+
		"	   And Length(Txx.Tte_Bl_Texto) > 10 "+
		"	 Order By Txx.Tte_Dt_Inclusao asc ", tar)

	if err := cur.Execute(cmd, nil, nil); err != nil {
		fmt.Println(fmt.Sprintf(`Erro ao executar "%s": %s`, cmd, err))
	} else {
		for {
			row, _ := cur.FetchOne()
			if row == nil {
				break
			}
			var com TicketC
			var tk TicketCom

			com.Comment.Author_id = row[2].(int32)
			var ret1 *oracle.ExternalLobVar
			ret1, _ = row[0].(*oracle.ExternalLobVar)
			aux, _ := ret1.ReadAll()
			com.Comment.Body = string(aux)
			if row[1].(string) == "C" {
				com.Comment.Public = "false"
			} else {
				com.Comment.Public = "true"
			}
			com.Comment.Type = "Comment"

			tk.Ticket = com

			tkJson, _ := json.Marshal(tk)
			//fmt.Println(string(tkJson))
			criaCommentTicketZen(tkJson, id)

		}
	}
}

func criaTicketZen(tk []byte) int32 {
	client := &http.Client{}
	url := "https://megasistemas.zendesk.com/api/v2/tickets.json"
	reqZen, _ := http.NewRequest("POST", url, strings.NewReader(string(tk)))
	reqZen.SetBasicAuth("qualiteam@mega.com.br/token", "pzwsg7ItHBIQqX1wzV66sYjl36awkhCXLgyMcjMx")
	reqZen.Header.Set("Content-Type", "application/json")
	resp, _ := client.Do(reqZen)

	body, _ := ioutil.ReadAll(resp.Body)
	var tkr TicketRet
	json.Unmarshal(body, &tkr)
	fmt.Println(fmt.Sprintf("ID: %d ", tkr.Ticket.Id))

	return tkr.Ticket.Id
}

func criaTicketZenProblem(tkt Tickets, ver string, loc int32, taridpm int32) int32 {
	tkt.Ticket.Type = "problem"
	if tkt.Ticket.AssigneeID == 0 {
		val, _ := strconv.ParseInt(tkt.Ticket.GroupID, 0, 32)
		tkt.Ticket.AssigneeID = int32(val)
	}
	tkt.Ticket.RequesterID = tkt.Ticket.AssigneeID

	var cmd string

	var texto string
	var idJira int32

	db := getConn()
	defer conn.Close()
	if db.IsConnected() {

		cur := db.NewCursor()
		defer cur.Close()

		dbl := "GCS"
		dbl1 := "GCS"
		if loc == 15 {
			dbl = "TO_CYCLUS_CURITIBA"
			dbl1 = "DB_INT"
		}

		cmd = fmt.Sprintf("Select * From (Select a.tte_bl_texto , t.tar_in_id, i.id From Mgtsk.Tsk_Vw_Tarefatexto@%s a, Mgtsk.Tsk_Tarefa@%s t, mgtsk.issues@%s i Where a.Tex_In_Tipo = 2 And t.Tpt_In_Id = 2 And t.Tar_In_Id = a.Tar_In_Id and i.PM (+) = t.tar_in_id and t.Tar_In_Id = %d)", dbl, dbl, dbl1, taridpm)

		if err := cur.Execute(cmd, nil, nil); err != nil {
			fmt.Println(fmt.Sprintf("Erro ao buscar o texto da PM : %s", err.Error()))
		} else {
			row, _ := cur.FetchOne()
			if row != nil {
				if row[0] != nil {
					texto = row[0].(string)
				}
				if row[2] != nil {
					idJira = row[2].(int32)
				}
			}
		}
	}

	tkt.Ticket.Comment.Body = fmt.Sprintf("------------------------------ \n Versão: %s \n  Cod_PM: %d \n ------------------------------ \n \n %s", ver, taridpm, texto)

	tkJson, _ := json.Marshal(tkt)

	id := criaTicketZen(tkJson)

	atualizaPMTicketHeroku(id, idJira)

	return id
}

func atualizaPMTicketHeroku(id int32, idJira int32) {
	if idJira != 0 {

		client := &http.Client{}
		url := fmt.Sprintf("https://limitless-bayou-7525.herokuapp.com/vinculaissue?token=bKWhAffE6qpG8x3ZvS4gGRU3v4LLsc9g&acao=vinc&issue=%d&ticket=%d", idJira, id)
		fmt.Println(url)
		heroku, _ := http.NewRequest("GET", url, nil)
		resp, err := client.Do(heroku)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
		if err != nil {
			fmt.Println(fmt.Sprintf("Erro ao vincular a issue %d na pm %d no Heroku", id, idJira))
		}
	}
}

func criaCommentTicketZen(tk []byte, id int32) {
	client := &http.Client{}
	url := fmt.Sprintf("https://megasistemas.zendesk.com/api/v2/tickets/%d.json", id)
	reqZen, _ := http.NewRequest("PUT", url, strings.NewReader(string(tk)))
	reqZen.SetBasicAuth("qualiteam@mega.com.br/token", "pzwsg7ItHBIQqX1wzV66sYjl36awkhCXLgyMcjMx")
	reqZen.Header.Set("Content-Type", "application/json")
	client.Do(reqZen)
}

func exportaComents() {
	db := getConn()
	defer conn.Close()

	cur := db.NewCursor()
	defer cur.Close()

	cmd := "select UNO_IN_ID_ZENDESK ID from TSK_VW_TICKETSORGIMP"
	if err := cur.Execute(cmd, nil, nil); err != nil {
		fmt.Println(fmt.Sprintf(`Erro ao executar "%s": %s`, cmd, err))
	} else {
		rows, errF := cur.FetchAll()
		if errF != nil {
			fmt.Println(fmt.Sprintf("Erro no Fetch: %s", errF))
		} else {
			for _, row := range rows {
				fmt.Println("org: ", row[0].(int32))
				getCommentTicketOrg(row[0].(int32))
			}
		}
	}
}

func getCommentTicketOrg(id int32) {
	client := &http.Client{}
	url := fmt.Sprintf("https://megasistemas.zendesk.com/api/v2/tickets.json?organization_id=%d", id)
	for {
		reqZen, _ := http.NewRequest("GET", url, nil)
		reqZen.SetBasicAuth("qualiteam@mega.com.br/token", "pzwsg7ItHBIQqX1wzV66sYjl36awkhCXLgyMcjMx")
		reqZen.Header.Set("Content-Type", "application/json")
		resp, _ := client.Do(reqZen)

		body, _ := ioutil.ReadAll(resp.Body)
		var tkt ImpTicket
		json.Unmarshal(body, &tkt)

		gerarCommentTicketOrg(tkt)

		if tkt.NextPage == nil {
			break
		}
		url = tkt.NextPage.(string)
	}

}

type OrgComment struct {
	NextPage interface{} `json:"next_page"`
	Comments []struct {
		AuthorID  float64 `json:"author_id"`
		Body      string  `json:"body"`
		CreatedAt string  `json:"created_at"`
		Public    bool    `json:"public"`
		Via       struct {
			Channel string `json:"channel"`
		} `json:"via"`
	} `json:"comments"`
}

func gerarCommentTicketOrg(tkt ImpTicket) {

	var cmd string

	db := getConn()
	defer conn.Close()

	if db.IsConnected() {

		cur := db.NewCursor()
		defer cur.Close()
		fmt.Println("  Qtde Tickets: ", tkt.Count)
		for _, f := range tkt.Tickets {

			client := &http.Client{}
			url := fmt.Sprintf("https://megasistemas.zendesk.com/api/v2/tickets/%d/comments.json", int32(f.ID))

			for {
				reqZen, errC := http.NewRequest("GET", url, nil)
				if errC != nil {
					fmt.Println(errC.Error())
					break
				}
				reqZen.SetBasicAuth("qualiteam@mega.com.br/token", "pzwsg7ItHBIQqX1wzV66sYjl36awkhCXLgyMcjMx")
				reqZen.Header.Set("Content-Type", "application/json")
				resp, _ := client.Do(reqZen)

				body, _ := ioutil.ReadAll(resp.Body)
				var oc OrgComment
				json.Unmarshal(body, &oc)

				fmt.Println(fmt.Sprintf("    %d - Qtde Comments: %d", int32(f.ID), len(oc.Comments)))
				for _, i := range oc.Comments {

					cmd = "INSERT INTO SYSTEM.ORGANIZACAOCOMMENTS (TICKET_ID, TICKET_ORGANIZATION, COMMENT_BODY, COMMENT_VIA, COMMENT_AUTHORID) " +
						"values(:TICKET_ID, :TICKET_ORGANIZATION, :COMMENT_BODY, :COMMENT_VIA, :COMMENT_AUTHORID)"

					param := make([]interface{}, 5)
					param[0] = f.ID
					param[1] = f.OrganizationID
					param[2] = i.Body
					param[3] = i.Via.Channel
					param[4] = i.AuthorID
					if err := cur.Execute(cmd, param, nil); err != nil {
						fmt.Println("cannot insert into ORGANIZACAOCOMMENTS : %s", err)
					}
				}
				if oc.NextPage == nil {
					break
				}
				url = oc.NextPage.(string)
			}
		}
	}
}

func excluirTickets() {
	db := getConn()
	defer conn.Close()

	cur := db.NewCursor()
	defer cur.Close()

	cmd := "SELECT * FROM SYSTEM.EXCLUIR_TICKETS where nvl(excluido,'N') = 'N'"
	if err := cur.Execute(cmd, nil, nil); err != nil {
		fmt.Println(fmt.Sprintf(`Erro ao executar "%s": %s`, cmd, err))
	} else {
		rows, errF := cur.FetchAll()
		if errF != nil {
			fmt.Println(fmt.Sprintf("Erro no Fetch: %s", errF))
		} else {
			fmt.Println(fmt.Sprintf("Excluindo %d Ticket... ", len(rows)))
			for _, row := range rows {
				fmt.Println("Ticket: ", row[0].(int32))
				excluirTicket(row[0].(int32))
			}
		}
	}
}

func excluirTicket(id int32) {
	client := &http.Client{}
	url := fmt.Sprintf("https://megasistemas.zendesk.com/api/v2/tickets/%d.json", id)
	reqZen, _ := http.NewRequest("DELETE", url, nil)
	reqZen.SetBasicAuth("qualiteam@mega.com.br/token", "pzwsg7ItHBIQqX1wzV66sYjl36awkhCXLgyMcjMx")
	reqZen.Header.Set("Content-Type", "application/json")
	//fmt.Println(url)
	_, errD := client.Do(reqZen)
	if errD != nil {
		fmt.Println(fmt.Sprintf("Erro ao excluir Ticket %d: %s", id, errD.Error()))
	} else {
		db := getConn()
		defer conn.Close()
		cur := db.NewCursor()
		defer cur.Close()
		cmd := fmt.Sprintf("update SYSTEM.EXCLUIR_TICKETS set excluido = '%s' where ticket_in_id = %d", "S", id)
		if err := cur.Execute(cmd, nil, nil); err != nil {
			fmt.Println(fmt.Sprintf("Erro ao excluir o Ticket %d", id))
		} else {
			fmt.Println(fmt.Sprintf("Excluido Ticket %d", id))
		}
	}
}

type TicketAltGru struct {
	Ticket struct {
		AssigneeID string `json:"assignee_id"`
		GroupID    string `json:"group_id"`
	} `json:"ticket"`
}

type TicketAltAss struct {
	Ticket struct {
		AssigneeID string `json:"assignee_id"`
	} `json:"ticket"`
}

func alteraTicket() {
	db := getConn()
	defer conn.Close()

	cur := db.NewCursor()
	defer cur.Close()

	cmd := "select p.Ticket_In_Id, p.Ticket_Assignee_Id, p.ticket_group_id from system.Pre_Importacao p "
	if err := cur.Execute(cmd, nil, nil); err != nil {
		fmt.Println(fmt.Sprintf(`Erro ao executar "%s": %s`, cmd, err))
	} else {
		rows, errF := cur.FetchAll()
		if errF != nil {
			fmt.Println(fmt.Sprintf("Erro no Fetch: %s", errF))
		} else {
			var tk int32
			var ass int32
			var gru string
			tot := len(rows)
			fmt.Println(fmt.Sprintf("Alterar %d Tickets ", tot))
			for i, row := range rows {
				tk = 0
				if row[0] != nil {
					tk = row[0].(int32)
				}
				ass = 0
				if row[1] != nil {
					ass = row[1].(int32)
				}
				gru = "0"
				if row[2] != nil {
					gru = row[2].(string)
				}
				fmt.Println(fmt.Sprintf("Ticket %d/%d: %d (%d - %d - %s)", i+1, tot, tk, tk, ass, gru))
				if tk != 0 {
					if ass == 0 && !strings.EqualFold(gru, "0") {
						alteraTicketZenGru(tk, gru)
					} else {
						if ass != 0 {
							alteraTicketZenAss(tk, strconv.Itoa(int(ass)))
						}
					}
				}
			}
		}
	}
}

func alteraTicketZenGru(tk int32, gru string) {
	var alt TicketAltGru
	alt.Ticket.AssigneeID = ""
	alt.Ticket.GroupID = gru

	tkJson, _ := json.Marshal(alt)

	executaAlteracao(tk, tkJson)
}

func alteraTicketZenAss(tk int32, ass string) {
	var alt TicketAltAss
	alt.Ticket.AssigneeID = ass

	tkJson, _ := json.Marshal(alt)

	executaAlteracao(tk, tkJson)
}

func executaAlteracao(id int32, tk []byte) {
	client := &http.Client{}
	url := fmt.Sprintf("https://megasistemas.zendesk.com/api/v2/tickets/%d.json", id)
	reqZen, _ := http.NewRequest("PUT", url, strings.NewReader(string(tk)))
	reqZen.SetBasicAuth("qualiteam@mega.com.br/token", "pzwsg7ItHBIQqX1wzV66sYjl36awkhCXLgyMcjMx")
	reqZen.Header.Set("Content-Type", "application/json")
	_, errD := client.Do(reqZen)
	if errD != nil {
		fmt.Println(fmt.Sprintf("    Erro ao alterar o Ticket %d: %s", id, errD.Error()))
	}
}

func gravarTokenImportado(tan int32, token string) bool {
	var cmd string

	db := getConn()
	defer conn.Close()
	if db.IsConnected() {

		cur := db.NewCursor()
		defer cur.Close()
		cmd = fmt.Sprintf("UPDATE MGTSK.PRE_IMPORTACAOANEXOS set ANEXOIMPORTADO = '%s', TOKEN = '%s' where TaN_In_Id=%d", "S", token, tan)
		fmt.Println(cmd)
		if err := cur.Execute(cmd, nil, nil); err != nil {
			fmt.Println("Erro ao atualizar os dados do anexo na PRE_IMPORTACAO : %s", err.Error())
			return false
		}
		fmt.Println("  Marcado registro: ", tan)
		return true
	} else {
		return false
	}
}

func insereAnexo() {
	db := getConn()
	defer conn.Close()

	cur := db.NewCursor()
	defer cur.Close()

	cmd := "select PA.Tan_In_Id, pa.Tar_In_Id, pa.Ticket_In_Id, pa.Tan_St_Nome, pa.Gru_St_Email, pa.Tan_Dt_Inclusao, TA.Tan_Bl_Anexo" +
		"  from PRE_IMPORTACAOANEXOS PA, Tsk_Tarefaanexo Ta " +
		" where PA.Tan_In_Id = TA.Tan_In_Id  " +
		"   and PA.Anexoimportado = 'N' order by 1 "

	if err := cur.Execute(cmd, nil, nil); err != nil {
		fmt.Println(fmt.Sprintf(`Erro ao executar "%s": %s`, cmd, err))
	} else {

		for {
			var tanId int32
			var tarId int32
			var ticketId int32
			var nomeArq string
			var conteudoArq []byte
			//var emailUsu string
			//var data time.Time
			row, _ := cur.FetchOne()
			if row == nil {
				break
			}
			if row[1] != nil {
				tarId = row[1].(int32)
				//fmt.Println(fmt.Sprintf("Tarefa: %d", tarId))
			}
			if row[2] != nil {
				ticketId = row[2].(int32)
				//fmt.Println(fmt.Sprintf("Ticket: %d", ticketId))
			}
			if row[3] != nil {
				nomeArq = row[3].(string)
				//fmt.Println(fmt.Sprintf("NomeArq: %s", nomeArq))
			}
			if row[6] != nil {
				var ret1 *oracle.ExternalLobVar
				ret1, _ = row[6].(*oracle.ExternalLobVar)
				conteudoArq, _ = ret1.ReadAll()
			}

			if row[0] != nil {
				tanId = row[0].(int32)
				//fmt.Println(fmt.Sprintf("tanId: %s", tanId))
			}

			/*          if row[4] != nil {
							emailUsu = row[4].(string)
							fmt.Println(fmt.Sprintf("email: %s", emailUsu))
						}
			/*
			/*
						if row[5] != nil {
							data = row[5].(time.Time)
							fmt.Println(fmt.Sprintf("Data: %s", data))
						}
			*/

			fmt.Println(fmt.Sprintf("Ticket %d - Tarefa %d - anexo %d - %s", tarId, ticketId, tanId, nomeArq))

			var token string
			var ret Attachments
			if ticketId != 0 {
				ret, token = uploadAnexo(nomeArq, conteudoArq)
				if ret.ID == 0 {
					fmt.Println(fmt.Sprintf("  Upload do arquivo %s inválido.", nomeArq))
				} else {
					fmt.Println("  TOKEN: ", token)
					if vincularAnexoTicket(ticketId, token) {
						if !gravarTokenImportado(tanId, token) {
							deleteAnexo(token)
						}
					}
				}
			} else {
				fmt.Println(fmt.Sprintf("  Tarefa %d esta sem ticket [0] vinculado.", tarId))
			}
		}
	}
}

type Attachments struct {
	ContentType string `json:"content_type"`
	ContentURL  string `json:"content_url"`
	FileName    string `json:"file_name"`
	ID          int32  `json:"id"`
	Size        int32  `json:"size"`
	URL         string `json:"url"`
}

type TicketAttac struct {
	Comment CommentAttac `json:"comment"`
}

type TicketsAttac struct {
	Ticket TicketAttac `json:"ticket"`
}

type CommentAttac struct {
	Body    string   `json:"body"`
	Uploads []string `json:"uploads"`
}

func vincularAnexoTicket(id int32, token string) bool {

	var tkts TicketsAttac
	var tkt TicketAttac
	var com CommentAttac

	com.Body = "Anexos..."
	com.Uploads = append(com.Uploads, token)
	tkt.Comment = com
	tkts.Ticket = tkt

	tkJson, _ := json.Marshal(tkts)
	//fmt.Println("Json envio anexo: ", string(tkJson))

	client := &http.Client{}
	url := fmt.Sprintf("https://megasistemas.zendesk.com/api/v2/tickets/%d.json", id)
	reqZen, _ := http.NewRequest("PUT", url, strings.NewReader(string(tkJson)))
	reqZen.SetBasicAuth("qualiteam@mega.com.br/token", "pzwsg7ItHBIQqX1wzV66sYjl36awkhCXLgyMcjMx")
	reqZen.Header.Set("Content-Type", "application/json")
	resp, errR := client.Do(reqZen)
	if errR != nil {
		fmt.Println(fmt.Sprintf("  Erro ao vincular o anexo %s ao ticket %d", token, id))
		deleteAnexo(token)
		return false
	}

	var b string
	if resp.StatusCode > 250 {
		body, _ := ioutil.ReadAll(resp.Body)
		b = string(body)
	} else {
		b = "ok"
	}
	fmt.Println(fmt.Sprintf("  Resposta Atualização Anexo: %d - %s", resp.StatusCode, b))
	return true
}

type AttachmentResp struct {
	ContentType string `json:"content_type"`
	ContentURL  string `json:"content_url"`
	FileName    string `json:"file_name"`
	ID          int32  `json:"id"`
	Size        int32  `json:"size"`
	URL         string `json:"url"`
}

type UploadResp struct {
	Attachment AttachmentResp `json:"attachment"`
	Token      string         `json:"token"`
}

type AnexoResp struct {
	Upload UploadResp `json:"upload"`
}

func uploadAnexo(nomeArq string, conteudo []byte) (Attachments, string) {

	//fmt.Println(nomeArq)
	var ata Attachments
	ata.ID = 0

	client := &http.Client{}
	url := fmt.Sprintf("https://megasistemas.zendesk.com/api/v2/uploads.json?filename=%s", strings.Replace(nomeArq, " ", "", -1))
	//fmt.Println(url)
	reqZen, err1 := http.NewRequest("POST", url, bytes.NewBuffer(conteudo))
	if err1 != nil {
		fmt.Println(fmt.Sprintf("  Erro ao realizar o upload do arquivo [%s]", err1.Error()))
		return ata, ""
	}
	reqZen.SetBasicAuth("qualiteam@mega.com.br/token", "pzwsg7ItHBIQqX1wzV66sYjl36awkhCXLgyMcjMx")
	reqZen.Header.Set("Content-Type", "application/binary")
	resp, err := client.Do(reqZen)
	if err != nil {
		fmt.Println(fmt.Sprintf("  Erro ao realizar o upload do arquivo [%s]", nomeArq))
		return ata, ""
	}
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("BODY: ", string(body))
	var upl AnexoResp
	json.Unmarshal(body, &upl)

	//fmt.Println(fmt.Sprintf("Token: %s - ID: %d ", upl.Upload.Token, upl.Upload.Attachment.ID))

	ata.FileName = upl.Upload.Attachment.FileName
	ata.ContentType = "application/unknown" //upl.Upload.Attachment.ContentType
	ata.ContentURL = upl.Upload.Attachment.ContentURL
	ata.ID = int32(upl.Upload.Attachment.ID)
	ata.Size = int32(upl.Upload.Attachment.Size)
	ata.URL = upl.Upload.Attachment.URL

	return ata, upl.Upload.Token
}

func deleteAnexo(token string) {

	client := &http.Client{}
	url := fmt.Sprintf("https://megasistemas.zendesk.com/api/v2/uploads/%s.json", token)
	fmt.Println("  Deletando anexo do ticket ", url)
	//fmt.Println(url)
	reqZen, err1 := http.NewRequest("DELETE", url, nil)
	if err1 != nil {
		fmt.Println(fmt.Sprintf("Erro ao deletar o arquivo [%s]", err1.Error()))
	}
	reqZen.SetBasicAuth("qualiteam@mega.com.br/token", "pzwsg7ItHBIQqX1wzV66sYjl36awkhCXLgyMcjMx")
	reqZen.Header.Set("Content-Type", "application/binary")
	resp, err := client.Do(reqZen)
	if err != nil {
		fmt.Println(fmt.Sprintf("  Erro ao realizar o delete do arquivo [%s]", token))
	}
	if resp.StatusCode == 201 || resp.StatusCode == 200 {
		fmt.Println(resp.StatusCode, "  Delete do arquivo OK")
	} else {
		fmt.Println(resp.StatusCode, "  Delete do arquivo com erro.")
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	}
}
