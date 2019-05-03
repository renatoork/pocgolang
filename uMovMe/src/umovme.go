package main

import (
	"./cadastro"
	"./pedido"
	"./util"
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

var arquivo *os.File

func main() {

	var err error

	arquivo, err = os.Create("Log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer arquivo.Close()

	util.WriterLog = bufio.NewWriter(arquivo)

	ret := carregarConfiguracao()

	if ret == "" {

		integrarGrupo := false
		integrarItem := false
		integrarCliente := false
		integrarTipoDoc := false
		integrarPedido := false
		inativarCadastro := false

		for i := 0; i < len(os.Args); i++ {
			if !integrarGrupo {
				integrarGrupo = strings.Contains(strings.ToUpper(os.Args[i]), "GRUPO")
			}
			if !integrarItem {
				integrarItem = strings.Contains(strings.ToUpper(os.Args[i]), "ITEM")
			}
			if !integrarCliente {
				integrarCliente = strings.Contains(strings.ToUpper(os.Args[i]), "CLIENTE")
			}
			if !integrarTipoDoc {
				integrarTipoDoc = strings.Contains(strings.ToUpper(os.Args[i]), "TIPODOC")
			}
			if !integrarPedido {
				integrarPedido = strings.Contains(strings.ToUpper(os.Args[i]), "PEDIDO")
			}
			if !inativarCadastro {
				inativarCadastro = strings.Contains(strings.ToUpper(os.Args[i]), "INATIVAR")
			}
			if !util.VTela {
				util.VTela = strings.Contains(strings.ToUpper(os.Args[i]), "TELA")
			}
			if !util.VLog {
				util.VLog = strings.Contains(strings.ToUpper(os.Args[i]), "LOG")
			}
		}

		if !integrarGrupo && !integrarItem && !integrarCliente && !integrarTipoDoc && !integrarPedido {
			integrarGrupo = true
			integrarItem = true
			integrarCliente = true
			integrarTipoDoc = true
			integrarPedido = true
		}
		util.Key = "https://api.umov.me/CenterWeb/api/11259e5cfe0fffbefcb0c4500648664b70b582" //master.vendasmegateste 123 - 	//util.Key = "https://api.umov.me/CenterWeb/api/11375e7a88fcbcae3d0571d2fb5facc636cf58" //master.umov455 a

		start := time.Now()
		util.AddLog("Inicio Integração: ")

		if integrarGrupo {
			gru := time.Now()
			util.AddLog("  - Grupo: ")
			cadastro.IntegrarSubGrupo(inativarCadastro)
			util.AddLog(fmt.Sprintln("           ", time.Since(gru)))
		}

		if integrarItem {
			it := time.Now()
			util.AddLog("  - Item: ")
			cadastro.IntegrarItem(inativarCadastro)
			util.AddLog(fmt.Sprintln("           ", time.Since(it)))
		}

		if integrarCliente {
			cl := time.Now()
			util.AddLog("  - Cliente: ")
			cadastro.IntegrarClienteUMov(inativarCadastro)
			util.AddLog(fmt.Sprintln("           ", time.Since(cl)))
		}

		if integrarTipoDoc {
			tp := time.Now()
			util.AddLog("  - TipoDocumento: ")
			cadastro.IntegrarTipoDoc(inativarCadastro)
			util.AddLog(fmt.Sprintln("           ", time.Since(tp)))
		}

		if integrarPedido {
			pe := time.Now()
			util.AddLog("  - Pedido de Venda: ")
			pedido.IntegrarPedidoVenda()
			util.AddLog(fmt.Sprintln("           ", time.Since(pe)))

			cl := time.Now()
			util.AddLog("  - Importa Cliente Pedido ERP: ")
			cadastro.IntegrarClienteERP()
			util.AddLog(fmt.Sprintln("           ", time.Since(cl)))

		}

		util.AddLog(fmt.Sprintln("Fim Integração: ", time.Since(start)))
	} else {
		util.AddLog(ret)
	}
	util.WriterLog.Flush()
}

func criarLog(log bool, conteudo string) {
	util.AddLog("criarLog")
	util.AddLog(string(conteudo))
	if log {
		ioutil.WriteFile("Log.txt", bytes.NewBufferString(conteudo).Bytes(), 0644)
	}

}

func carregarConfiguracao() string {
	conf, err := ioutil.ReadFile("config.json")
	if err == nil {
		resp := new(util.Config)
		errj := json.Unmarshal(conf, resp)

		if errj != nil {
			return errj.Error()
		} else {
			util.Key = "https://api.umov.me/CenterWeb/api/" + resp.KeyUMov //var Key = "https://api.umov.me/CenterWeb/api/11259e5cfe0fffbefcb0c4500648664b70b582" //master.vendasmegateste 123
			util.KeyN = resp.KeyUMov                                       //var KeyN = "11259e5cfe0fffbefcb0c4500648664b70b582"
			util.ConnectString = resp.ConnString                           //"megaumov/megaumov@pc_ana:1521/orc4"
			return ""
		}
	} else {
		return err.Error()
	}
}
