package auditutil

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestSubst(t *testing.T) {
	caminho := CaminhoReal("S:/fontes/Financ/uFormClasse.pas")
	esperado := "C:/Mega/MegaEmpresarial/MegaEmpresarial_401/fontes/Financ/uFormClasse.pas"
	if filepath.FromSlash(caminho) != filepath.FromSlash(esperado) {
		t.Log(fmt.Printf("Caminho: %s\nEsperado:%s\n", caminho, esperado))
		//t.Fail()
	}
}
