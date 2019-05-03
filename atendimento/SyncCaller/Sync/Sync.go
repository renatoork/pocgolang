package Sync

import (
	"mega/atendimento/SyncCaller/Caller"
	"mega/atendimento/SyncCaller/LogError"
	"mega/atendimento/SyncCaller/Types"

	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var IntervalSync Types.IntervalSync
var SyncConfig Types.SyncConfig

func ClearConsole() {
	clear := exec.Command("cmd", "/c", "cls")
	clear.Stdout = os.Stdout
	clear.Run()
}

func ConsoleMain() {

	var opt string

	fmt.Println("###########################################")
	fmt.Println("   App: SyncCaller")
	fmt.Println("   Author: Qualiteam [Bernardo P. O.]")
	fmt.Println("")
	fmt.Println("   Rotina que realiza chamadas para ")
	fmt.Println("   sincronize Mega.")
	fmt.Println("\n\n   1 - Tickets\n   2 - Organization\n   3 - User\n   4 - Group\n")
	fmt.Println("   5 - Check last error\n   6 - Clear Console\n   99 - Console Especial\n   custom - Requisição custom")
	fmt.Println("###########################################\n")

	fmt.Print("\nDigite a opção desejada: ")
	_, _ = fmt.Scanln(&opt)

	switch opt {
	case "1":
		fmt.Println("\n..Calling Tickets..\n")
		Caller.CallHttp("ticket", &SyncConfig)
		fmt.Println("..Called..\n")
		getLineString("Pressione qualquer tecla para continuar...")
		ClearConsole()
		ConsoleMain()
	case "2":
		fmt.Println("\n..Calling Organizations..\n")
		Caller.CallHttp("organization", &SyncConfig)
		fmt.Println("..Called..\n")
		getLineString("Pressione qualquer tecla para continuar...")
		ClearConsole()
		ConsoleMain()
	case "3":
		fmt.Println("\n..Calling Users..\n")
		Caller.CallHttp("user", &SyncConfig)
		fmt.Println("..Called..\n")
		getLineString("Pressione qualquer tecla para continuar...")
		ClearConsole()
		ConsoleMain()
	case "4":
		fmt.Println("\n..Calling Groups..\n")
		Caller.CallHttp("group", &SyncConfig)
		fmt.Println("..Called..\n")
		getLineString("Pressione qualquer tecla para continuar...")
		ClearConsole()
		ConsoleMain()
	case "5":
		fmt.Println("\nERRO :: [" + LogError.GetLastError().Log + "]\n")
		getLineString("Pressione qualquer tecla para continuar....")
		ClearConsole()
		ConsoleMain()
	case "6":
		ClearConsole()
		ConsoleMain()
	case "99":
		ConsoleEspecial()
	case "custom":
		var call Types.Caller
		var customPath string

		call.Metodo = "GET"
		customPath = getLineString("Digite o caminho (sem barra, ex: savetickets): ")
		call.Url = SyncConfig.GetCustomUrl() + customPath
		LogError.SetLog(true)

		fmt.Println("\n..Calling Custom..\n")
		Caller.HttpReq(&call)
		fmt.Println("..Called..\n")

		getLineString("Pressione qualquer tecla para continuar....")
		ClearConsole()
		ConsoleMain()
	default:
		fmt.Println("Opção ainda não disponível.")
		ClearConsole()
		ConsoleMain()
	}

}

func ConsoleEspecial() {

	var cmdOpt string

	ClearConsole()
	fmt.Println("###############################################")
	fmt.Println("\tDigite o comando desejado.")
	fmt.Println("###############################################\n\n")

	cmdOpt = getLineString("#> ")

	//cmdOpt = strings.Replace(strings.Replace(strings.ToLower(cmdOpt), "\n", "", -1), "\r", "", -1)

	switch cmdOpt {
	case "exit":
		return
	case "log on":
		LogError.SetLog(true)
		fmt.Println(LogError.CheckLog())
		ConsoleMain()
	case "log off":
		LogError.SetLog(false)
		ConsoleMain()
	case "tickets interval on":
		IntervalSync.CreateInterval()
		Caller.CallTicketsInterval(&IntervalSync, &SyncConfig)
		ConsoleMain()
	case "tickets interval off":
		IntervalSync.StopInterval()
		ConsoleMain()
	case "tickets interval":
		var tempo string
		tempo = getLineString("Digite o tempo: ")
		IntervalSync.SetInterval(tempo)
		ConsoleMain()
	default:
		ConsoleMain()
	}

}

func getLineString(msg string) string {

	var value string

	reader := bufio.NewReader(os.Stdin)

	fmt.Print(msg)
	value, err := reader.ReadString('\n')

	LogError.LogInline(err)

	return strings.Replace(strings.Replace(value, "\n", "", -1), "\r", "", -1)

}

func VariablesInitiate() {

	//Iniciando variável de chamadas com intervalo de tempo
	IntervalSync.DefineIntervalTicket()
	IntervalSync.Interval = 15 //Minutos
	IntervalSync.CreateInterval()

	//Verificando arquivo de configuração
	var path string

	path = "Logs/"

	_, err := os.Stat(path)
	if err != nil {
		panic(errors.New("\n-------------------------------------------------\n   Voce precisa criar a pasta de logs [Logs/]!   \n-------------------------------------------------\n"))
	} else {

		fSyncConfig, err := os.OpenFile("syncconfig.json", os.O_RDONLY, 0666)
		if err != nil {
			panic(errors.New("\n-------------------------------------------------\n   Voce precisa criar o arquivo de configuracao!   \n-------------------------------------------------\n"))
		}

		fSyncStat, err := fSyncConfig.Stat()
		bConfig := make([]byte, fSyncStat.Size())
		n, err := fSyncConfig.Read(bConfig)
		if n > 0 {
			err = json.Unmarshal(bConfig, &SyncConfig)
			if err != nil {
				panic(errors.New("Erro de inicialização [Config]: " + err.Error()))
			}
		} else {
			panic(errors.New("\n-------------------------------------------------\n   Não foi possivel ler arquivo de configuracao!   \n-------------------------------------------------\n"))
		}
	}

}

func IntervalTicketImport() {
	Caller.CallTicketsInterval(&IntervalSync, &SyncConfig)
}

func IntervalOrganizationImport() {
	Caller.CallOrganizationsInterval(&IntervalSync, &SyncConfig)
}

func IntervalUserImport() {
	Caller.CallUsersInterval(&IntervalSync, &SyncConfig)
}
