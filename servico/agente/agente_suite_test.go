package agente_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestAgente(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Agente Suite")
}
