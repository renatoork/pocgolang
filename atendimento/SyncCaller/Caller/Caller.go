package Caller

import (
	"mega/atendimento/SyncCaller/LogError"
	"mega/atendimento/SyncCaller/Timer"
	"mega/atendimento/SyncCaller/Types"

	"bytes"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func CallHttp(tipo string, cfg *Types.SyncConfig) {

	var call Types.Caller

	call.Metodo = "GET"
	call.Url = cfg.GetUrl(tipo) //Busca a URL de chamada da configuração

	LogError.SetLog(true)
	HttpReq(&call)
	if call.Erro != "" {
		fmt.Println(call.Erro)
	}

}

func CallTicketsInterval(intTick *Types.IntervalSync, cfg *Types.SyncConfig) {

	var waitIt sync.WaitGroup

	waitIt.Add(1)

	go func() {

		//First Run
		Timer.TimeStart()
		fmt.Println("\n..Calling Tickets..")
		CallHttp("ticket", cfg)
		fmt.Println("Executed: \t\n", Timer.TimeEnd())

		for _ = range intTick.SyncInterval.C {
			Timer.TimeStart()
			fmt.Println("\n..Calling Tickets..")
			CallHttp("ticket", cfg)
			fmt.Println("Executed: \t\n", Timer.TimeEnd())
		}

		waitIt.Done()

	}()

	waitIt.Wait()

}

func CallOrganizationsInterval(intOrg *Types.IntervalSync, cfg *Types.SyncConfig) {

	var waitIt sync.WaitGroup

	waitIt.Add(1)

	go func() {

		//First Run
		Timer.TimeStart()
		fmt.Println("\n..Calling Organization..")
		CallHttp("organization", cfg)
		fmt.Println("Executed: \t\n", Timer.TimeEnd())

		for _ = range intOrg.SyncInterval.C {
			Timer.TimeStart()
			fmt.Println("\n..Calling Organization..")
			CallHttp("organization", cfg)
			fmt.Println("Executed: \t\n", Timer.TimeEnd())
		}

		waitIt.Done()

	}()

	waitIt.Wait()

}

func CallUsersInterval(intUser *Types.IntervalSync, cfg *Types.SyncConfig) {

	var waitIt sync.WaitGroup

	waitIt.Add(1)

	go func() {

		//First Run
		Timer.TimeStart()
		fmt.Println("\n..Calling User..")
		CallHttp("user", cfg)
		fmt.Println("Executed: \t\n", Timer.TimeEnd())

		for _ = range intUser.SyncInterval.C {
			Timer.TimeStart()
			fmt.Println("\n..Calling User..")
			CallHttp("user", cfg)
			fmt.Println("Executed: \t\n", Timer.TimeEnd())
		}

		waitIt.Done()

	}()

	waitIt.Wait()

}

func CallGroupsInterval(intUser *Types.IntervalSync, cfg *Types.SyncConfig) {

	var waitIt sync.WaitGroup

	waitIt.Add(1)

	go func() {

		//First Run
		Timer.TimeStart()
		fmt.Println("\n..Calling User..")
		CallHttp("group", cfg)
		fmt.Println("Executed: \t\n", Timer.TimeEnd())

		for _ = range intUser.SyncInterval.C {
			Timer.TimeStart()
			fmt.Println("\n..Calling User..")
			CallHttp("group", cfg)
			fmt.Println("Executed: \t\n", Timer.TimeEnd())
		}

		waitIt.Done()

	}()

	waitIt.Wait()

}

func HttpReq(call *Types.Caller) {

	client := &http.Client{}

	reqUrl, err := http.NewRequest(call.Metodo, call.Url, nil)
	LogError.LogInline(err)

	respUrl, err := client.Do(reqUrl)
	LogError.LogInline(err)

	if err != nil {
		call.Erro = "Erro ao executar requisição!\nERRO:: " + err.Error()
		return
	}

	b := new(bytes.Buffer)
	b.ReadFrom(respUrl.Body)

	if LogError.CheckLog() {
		var recLog Types.RecLog

		recLog = Types.RecLog{}

		fileName := time.Now()
		strFileName := fmt.Sprintf("%d_%d_%d__%d-%d-%d.log", fileName.Day(), int(fileName.Month()), fileName.Year(), fileName.Hour(), fileName.Minute(), fileName.Second())

		recLog.FileName = strFileName
		recLog.Log = b.String()

		LogError.LogFile(&recLog)
	}

	//gravaTxt("retorno bytes", b.String())

	call.Retorno = b.Bytes()

}

func gravaTxt(name, txt string) {
	var recLog Types.RecLog

	recLog = Types.RecLog{}

	recLog.FileName = name
	recLog.Log = txt

	LogError.LogFile(&recLog)
}
