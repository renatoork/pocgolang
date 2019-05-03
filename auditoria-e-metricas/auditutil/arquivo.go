package auditutil

import (
	"errors"
	"fmt"
	"io/ioutil"
)

type Arquivo struct {
	Nome             string
	ConteudoAnterior string
	ConteudoAtual    string
}

func LerArquivoTeste(arquivo string) (string, error) {
	conteudo, erro := ioutil.ReadFile(arquivo)

	if erro != nil {
		return "", errors.New(fmt.Sprintf("Erro ao ler o arquivo [%s]", arquivo))
	}

	return string(conteudo), nil
}
