package auditutil

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func CaminhoReal(caminho string) string {

	retorno, err := exec.Command("subst").Output()
	if err != nil {
		fmt.Println(err)
	}
	linhas := strings.Split(string(retorno), "\r\n")
	regex := regexp.MustCompile(`(\w:\\): => (.+)`)

	for _, linha := range linhas {
		if linha != "" {
			grupos := regex.FindStringSubmatch(linha)
			if grupos[1][0] == filepath.VolumeName(caminho)[0] {
				relativo, _ := filepath.Rel(grupos[1], caminho)
				junto := filepath.Join(grupos[2], relativo)
				return filepath.ToSlash(junto)
			}
		}

	}
	return ""
}

func main() {
	CaminhoReal("S:/fontes/Financ/uFormClasse.pas")
}
