package main

import (
	"mega/atendimento/DashBoardSync/Zendesk"

	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {

	var port string
	port = ""

	enviromentPrepare()
	clearScreen()

	// Set variável de ambiente NLS_LANG com codificação UTF8, pois dados retornados das entidades tem a codificação de UTF8.
	os.Setenv("NLS_LANG", "American_America.UTF8")

	port = os.Getenv("PORT")

	if port == "" {
		port = "8899"
	}

	http.HandleFunc("/saveticketlinkissue", SaveTicketLinkIssue)

	http.HandleFunc("/savetickets", SaveTickets)
	http.HandleFunc("/saveorganizations", SaveOrganizations)
	http.HandleFunc("/saveusers", SaveUsers)
	http.HandleFunc("/savegroups", SaveGroups)
	http.ListenAndServe(":"+port, nil)

}

func clearScreen() {

	var runIt *exec.Cmd

	runIt = exec.Command("cmd", "/c", "cls")
	runIt.Stdout = os.Stdout
	_ = runIt.Run()

	fmt.Println("\n")
	fmt.Println("###########################################")
	fmt.Println("   App: DashboardSync")
	fmt.Println("   Author: Qualiteam [Bernardo P. O.]")
	fmt.Println(" ")
	fmt.Println("   Rotina para sincornização de entidades")
	fmt.Println("   do Zendesk para Mega.")
	fmt.Println("###########################################")
	fmt.Println("\nRotina em execução...")

}

func SaveTicketLinkIssue(rw http.ResponseWriter, req *http.Request) {

	Zendesk.SetTicketLinkIssue()
	rw.Write([]byte("Fim da rotina"))

}

func SaveTickets(rw http.ResponseWriter, req *http.Request) {

	var tempoInicio string
	var epochTime time.Time
	var once bool //Rodar apenas 1 vez.

	tempoInicio = req.FormValue("tempoinicio")

	if req.FormValue("apenasum") != "" && req.FormValue("apenasum") == "true" {
		once = true
	} else {
		once = false
	}

	fmt.Println(tempoInicio)

	//Validando parâmetros
	if tempoInicio != "" {
		epochTime, _ = time.Parse("2006-02-01", tempoInicio)
		fmt.Println(strconv.FormatInt(epochTime.Unix(), 10))
	} else {
		timestamp := Zendesk.GetLastDataTimeTicket()
		i, _ := strconv.ParseInt(timestamp, 10, 64)
		epochTime = time.Unix(i, 0)
	}

	Zendesk.GetIncExportTicket(strconv.FormatInt(epochTime.Unix(), 10), once)

	clearScreen()

	rw.Write([]byte("Fim da rotina"))

}

func SaveOrganizations(rw http.ResponseWriter, req *http.Request) {

	var tempoInicio string
	var epochTime time.Time
	var once bool //Rodar apenas 1 vez.

	tempoInicio = req.FormValue("tempoinicio")

	if req.FormValue("apenasum") != "" && req.FormValue("apenasum") == "true" {
		once = true
	} else {
		once = false
	}

	fmt.Println(tempoInicio)

	//Validando parâmetros
	if tempoInicio != "" {
		epochTime, _ = time.Parse("2006-02-01", tempoInicio)
		fmt.Println(strconv.FormatInt(epochTime.Unix(), 10))
	} else {
		timestamp := Zendesk.GetLastDataTimeOrganization()
		i, _ := strconv.ParseInt(timestamp, 10, 64)
		epochTime = time.Unix(i, 0)
	}

	Zendesk.GetIncExportOrganization(strconv.FormatInt(epochTime.Unix(), 10), once)

	clearScreen()

	rw.Write([]byte("Fim da rotina"))

}

func SaveUsers(rw http.ResponseWriter, req *http.Request) {

	var tempoInicio string
	var epochTime time.Time
	var once bool //Rodar apenas 1 vez.

	tempoInicio = req.FormValue("tempoinicio")

	if req.FormValue("apenasum") != "" && req.FormValue("apenasum") == "true" {
		once = true
	} else {
		once = false
	}

	fmt.Println(tempoInicio)

	//Validando parâmetros
	if tempoInicio != "" {
		epochTime, _ = time.Parse("2006-02-01", tempoInicio)
		fmt.Println(strconv.FormatInt(epochTime.Unix(), 10))
	} else {
		timestamp := Zendesk.GetLastDataTimeUser()
		i, _ := strconv.ParseInt(timestamp, 10, 64)
		epochTime = time.Unix(i, 0)
	}

	Zendesk.GetIncExportUser(strconv.FormatInt(epochTime.Unix(), 10), once)

	clearScreen()

	rw.Write([]byte("Fim da rotina"))

}

func SaveGroups(rw http.ResponseWriter, req *http.Request) {

	Zendesk.GetIncExportGroup("", false)

	clearScreen()

	rw.Write([]byte("Fim da rotina"))

}

func enviromentPrepare() {

	var pathLog string

	pathLog = "Logs/"

	_, err := os.Stat(pathLog)

	if err != nil {
		err = os.Mkdir("Logs", 0666)
		if err != nil {
			fmt.Println("Erro de inicialização: " + err.Error())
		}
	}

}
