package auditutil

import (
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

func ConteudoOriginal(caminho string) string {

	pastaAtual, nomeArquivo := path.Split(filepath.ToSlash(caminho))

	CheckError(os.Chdir(pastaAtual))

	//Pesquisa TopLevel
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	topLevel, err := cmd.Output()
	CheckError(err)

	//Pesquisa Conteúdo do arquivo
	cmd = exec.Command("git", "cat-file", "blob", "HEAD:"+substr(pastaAtual, len(topLevel), 255)+nomeArquivo)
	HEAD, err := cmd.Output()
	CheckError(err)
	return string(HEAD)
}
