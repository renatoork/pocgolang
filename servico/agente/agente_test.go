package agente_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "mega/servico/agente"
	"strings"
)

var _ = Describe("Agente", func() {

	It("Alimenta corretamente fantasia com base no nome", func() {
		nome := "Esse é o nome do agente com vários caracteres e acentos"
		agente := Agente{}
		agente.Nome = nome

		AlimentaFantasia(&agente)

		Expect(agente.Fantasia).To(Equal(strings.ToUpper(nome)))
	})

	It("Deve transformar o fantasia em maiúsculas se estiver alimentado", func() {
		fantasia := "Esse é o nome do agente com vários caracteres e acentos"
		agente := Agente{}
		agente.Fantasia = fantasia

		AlimentaFantasia(&agente)

		Expect(agente.Fantasia).To(Equal(strings.ToUpper(fantasia)))
	})

	It("Não deve deixar inserir agente sem tipo", func() {

		agente := Agente{}
		err := ValidaTiposAgente(&agente)
		Expect(err).Should(HaveOccurred())
	})
})
