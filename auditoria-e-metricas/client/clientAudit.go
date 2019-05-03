package main

import (
  "fmt"
	"os"
  "io/ioutil"
  "../auditutil"
  "path/filepath"
  "net/http"
  "encoding/json"
  "bytes"
)


func main() {

  if len(os.Args) > 1 {

    caminhoArquivo  := os.Args[1]
    arq, err := ioutil.ReadFile(caminhoArquivo)

    if err != nil {
      fmt.Printf("Erro [%s] ao ler o arquivo [%s]", err, caminhoArquivo )
    } else {
      var arquivo auditutil.Arquivo
      arquivo.Nome = filepath.Base(caminhoArquivo)
      arquivo.ConteudoAnterior = "Anteriorrrrrrr"
      arquivo.ConteudoAtual = string(arq)

      jsonArq, err := json.MarshalIndent(arquivo, "" , "  ")
      if err != nil {
        fmt.Println("Error [%s] ao empacotar o arquivo.", err)
      } else {
        resposta, err := http.Post("http://localhost:8080/arquivo", "application/json", bytes.NewBuffer(jsonArq))
        if err != nil{
          fmt.Println(err) 
        } else {
          fmt.Println(resposta) 
        }

      }
          
    }
  }

}