package main

import (
	"fmt"
	"mega/atendimento/SyncCaller/Sync"
	"os"
)

func main() {

	//Lendo parametros

	//Iniciando variáveis e configurações do aplicativo.
	Sync.ClearConsole()
	Sync.VariablesInitiate()

	if len(os.Args) > 1 && os.Args[1] == "ticketimportinterval" {
		writeConsole()
		Sync.IntervalTicketImport()
	} else if len(os.Args) > 1 && os.Args[1] == "organizationimportinterval" {
		writeConsole()
		Sync.IntervalOrganizationImport()
	} else if len(os.Args) > 1 && os.Args[1] == "userimportinterval" {
		writeConsole()
		Sync.IntervalOrganizationImport()
	} else {
		Sync.ConsoleMain()
	}

}

func writeConsole() {

	fmt.Println("\n")
	fmt.Println("###########################################")
	fmt.Println("   App: SyncCaller")
	fmt.Println("   Author: Qualiteam [Bernardo P. O.]")
	fmt.Println("")
	fmt.Println("   Rotina que realiza chamadas para ")
	fmt.Println("   sincronização entre Zendesk e Mega.")
	fmt.Println("###########################################")

	fmt.Println("\nRotina em execução...")

}
