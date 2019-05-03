package agente_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"math/rand"
	. "mega/servico/agente"
	"mega/servico/banco"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"
)

var ts *httptest.Server

var _ = BeforeSuite(func() {
	banco.Conectar()
	r := mux.NewRouter()

	Inicializa(r)

	ts = httptest.NewServer(r)

	rand.Seed(time.Now().UTC().UnixNano())
})

var _ = Describe("Serviço Agente", func() {

	var agenteGravado Agente
	var agenteRetornado Agente

	It("Insert agente", func() {
		nome := fmt.Sprintf("Agente inserido pelo serviço com vários caracteres ãéíóúõ %d", rand.Int())
		agenteJson := fmt.Sprintf(`{"Nome": "%s", "TipoPessoa": "F", "Tipos": [{"Tipo": "C"}]}`, nome)

		res, err := http.Post(ts.URL+"/api/agente", "application/json", strings.NewReader(agenteJson))
		Expect(err).ShouldNot(HaveOccurred())

		err = json.NewDecoder(res.Body).Decode(&agenteGravado)
		Expect(err).ShouldNot(HaveOccurred())

		Expect(agenteGravado.Chave.Codigo).To(BeNumerically(">", 0))
	})

	It("Get agente", func() {
		res, err := http.Get(fmt.Sprintf("%s/api/agente/%d", ts.URL, agenteGravado.Chave.Codigo))
		Expect(err).ShouldNot(HaveOccurred())

		err = json.NewDecoder(res.Body).Decode(&agenteRetornado)
		Expect(err).ShouldNot(HaveOccurred())

		Expect(agenteRetornado.Nome).To(Equal(agenteGravado.Nome))
		Expect(len(agenteRetornado.Tipos)).To(Equal(len(agenteGravado.Tipos)))
		Expect(agenteRetornado.Tipos[0].Tipo).To(Equal(agenteGravado.Tipos[0].Tipo))
	})

	It("Update agente", func() {
		agenteRetornado.Fantasia = fmt.Sprintf("Nome alterado com vários caracteres ãéíóúõ %d", rand.Int())

		agenteJson, err := json.Marshal(agenteRetornado)
		Expect(err).ShouldNot(HaveOccurred())

		client := http.Client{}
		req, _ := http.NewRequest("PUT", fmt.Sprintf("%s/api/agente/%d", ts.URL, agenteGravado.Chave.Codigo), bytes.NewBuffer(agenteJson))
		res, err := client.Do(req)
		Expect(err).ShouldNot(HaveOccurred())

		var retorno Agente
		body, err := ioutil.ReadAll(res.Body)
		Expect(err).ShouldNot(HaveOccurred())
		bodyStr := string(body)
		if strings.Index(bodyStr, "{") > -1 {
			err = json.Unmarshal(body, &retorno)
		} else {
			Expect(bodyStr).To(BeEmpty())
		}
		Expect(err).ShouldNot(HaveOccurred())
		Expect(retorno.Fantasia).To(Equal(strings.ToUpper(agenteRetornado.Fantasia)))

		res, err = http.Get(fmt.Sprintf("%s/api/agente/%d", ts.URL, agenteGravado.Chave.Codigo))
		Expect(err).ShouldNot(HaveOccurred())

		var agenteRetornadoDepoisAlteracao Agente
		err = json.NewDecoder(res.Body).Decode(&agenteRetornadoDepoisAlteracao)
		Expect(err).ShouldNot(HaveOccurred())

		Expect(agenteRetornadoDepoisAlteracao.Fantasia).To(Equal(strings.ToUpper(agenteRetornado.Fantasia)))
	})

})
