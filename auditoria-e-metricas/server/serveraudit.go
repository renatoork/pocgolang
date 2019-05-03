package main

import (
	"../auditutil"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"reflect"
)

func main() {

	http.HandleFunc("/arquivo", arquivoHandler)
	http.HandleFunc("/resultado", resultadoHandler)
	fmt.Println("Rodando na porta 8080")
	http.ListenAndServe(":8080", nil)

}

func gravaArquivoTemp(pastaTempSistema string, arquivo auditutil.Arquivo, tipo string) {
	var caminhoCompleto = path.Join(pastaTempSistema, tipo)
	err := os.Mkdir(caminhoCompleto, 0700)
	fmt.Println(err)
	var conteudo string
	if tipo == "anterior" {
		conteudo = arquivo.ConteudoAnterior
	} else {
		conteudo = arquivo.ConteudoAtual
	}

	writer, err := os.Create(path.Join(caminhoCompleto, arquivo.Nome))
	fmt.Println(err)
	defer writer.Close()
	writer.WriteString(conteudo)
}

func arquivoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		dec := json.NewDecoder(r.Body)
		var arquivo auditutil.Arquivo
		err := dec.Decode(&arquivo)
		fmt.Println(err)
		pastaTempSistema, err := ioutil.TempDir("", "audit_")
		gravaArquivoTemp(pastaTempSistema, arquivo, "atual")
		gravaArquivoTemp(pastaTempSistema, arquivo, "anterior")
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(auditutil.AuditaArquivo)
	}
}

func resultadoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		r.ParseForm()
		id := r.FormValue("id")

		resultado, err := auditutil.RetornaResultadoAuditoria(id)

		fmt.Println(resultado, err)

		if err == nil {

			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(&resultado)
		} else {
			fmt.Println(reflect.TypeOf(err))
			switch err.(type) {
			case *auditutil.InvalidIdError:
				w.WriteHeader(400)
				json.NewEncoder(w).Encode(err)
			case *auditutil.IdNotFoundError:
				w.WriteHeader(404)
				json.NewEncoder(w).Encode(err)
			}
		}

		//resultado, err := auditutil.RetornaResultado(id)

		/*if bson.IsObjectIdHex(id) {
			objectId := bson.ObjectIdHex(id)

			var resultados auditutil.Resultados

			err := dbAudit.Find(bson.M{"_id": objectId}).One(&resultados)

			if err != nil {
				w.WriteHeader(404)
				json.NewEncoder(w).Encode("Registro não encontrado")
			}

			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(&resultados)
		} else {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode("Chave inválida. Favor informar uma chave válida no parâmetro id.")
		}*/

	}
}
