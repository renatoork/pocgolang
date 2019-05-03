package main

import (
	//	"./util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	//"os"
)

func main() {
	/* para cada json do diretorio
	        entender o serviço
			executar o setup
			casar o parametro de entrada com a DEF
			executar o serviço
			comparar o resultado
			executar teardown
	*/

	ct := "D:/DESENV/git/go/TesteServicoPDV/CTs/CT_1.json"
	if len(ct) > 1 {

		caminhoArquivo := ct
		arq, err := ioutil.ReadFile(caminhoArquivo)

		if err != nil {
			fmt.Printf("Erro [%s] ao ler o arquivo [%s]", err, caminhoArquivo)
		} else {
			//trataStruct(arq)
			trataMap(arq)
			//util.GetSession("mgven/megaven@PC_ana:1521/orc4").MustExec(insert_PedidoVenda, Pedidovenda)
		}
	}
}

func trataStruct(arq []byte) {
	var casoTesteStruct CasoTeste
	err := json.Unmarshal([]byte(arq), &casoTesteStruct)
	if err != nil {
		fmt.Println("Error [%s] ao desempacotar o arquivo.", err)
	}
	fmt.Println(casoTesteStruct)
}

func trataMap(arq []byte) {
	var casoTesteMap map[string]interface{}
	err := json.Unmarshal([]byte(arq), &casoTesteMap)
	if err != nil {
		fmt.Println("Error [%s] ao desempacotar o arquivo.", err)
	}
	fmt.Println(casoTesteMap)

}
