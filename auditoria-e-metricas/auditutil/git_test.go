package auditutil

import (
	"fmt"
	"testing"
)

func TestGit(t *testing.T) {
	/*
		A variável "caminho" recebe o caminho real do path.
		Ela será usada na função "ConteudoOriginal" em conjunto com o GIT e deve
		conter o caminho escrito da mesma forma que o S.O., pois, é CaseSensitive
	*/

	//caminho := "c:/temp/testegit/teste.txt"
	//caminho := "D:/Git/MegaEmpresarial/MegaEmpresarial_401/Fontes/FinCPA/uDataModuleCPagar.pas"
	caminho := CaminhoReal("S:/Fontes/FinCPA/uDataModuleCPagar.pas")

	conteudo := ConteudoOriginal(caminho)
	fmt.Println("GIT_TEST = " + conteudo)
}
