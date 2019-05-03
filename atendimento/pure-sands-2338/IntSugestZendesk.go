package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	fmt.Println("Em funcionamento...")
	http.HandleFunc("/geraocorrencia", geraOcorrencia)
	http.HandleFunc("/atualizaocorrencia", atualizaOcorrencia)
	http.HandleFunc("/atualizaticket", atualizaTicket)
	http.HandleFunc("/intsugestzd", sugestZendesk)
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		panic(err)
	}
}

func geraOcorrencia(res http.ResponseWriter, req *http.Request) {
	res, req = GeraOcorrenciaMultiDados(res, req)
}

func atualizaOcorrencia(res http.ResponseWriter, req *http.Request) {
	res, req = AtualizaOcorrenciaMultiDados(res, req)
}

func atualizaTicket(res http.ResponseWriter, req *http.Request) {
	res, req = AtualizaTicketZenDesk(res, req)
}

func sugestZendesk(res http.ResponseWriter, req *http.Request) {
	res, req = IntSugestZd(res, req)
}
