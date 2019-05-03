package serverutil

import (
	"fmt"
	"regexp"
	"strings"
)

func validaArquivo(arquivo string) string {
	//fonte := separarDeclaracaoComponente(arquivo)
	//componente := separarComponente(fonte)
	audit := validaNomenclatura()
	return audit
}

func separarDeclaracaoComponenteSlice(arquivo string) []string {
	re := regexp.MustCompile("(?s)class(.+?)private")
	return re.FindAllString(arquivo, -1)
}

func separarDeclaracaoComponente(arquivo string) string {
	re := regexp.MustCompile("(?s)class(.+?)private")
	return strings.Join(re.FindAllString(arquivo, -1), " ")
}

func separarComponente(arquivo string) [][]string {

	// remover comentario
	reComent := regexp.MustCompile("(?s)[/*|{|(*|//^](.*?)[*/$|}|*)]")
	arquivoComp := reComent.ReplaceAllString(arquivo, "")

	// separar declaracao componentes
	regComp := regexp.MustCompile("(.+?): (.+?);")
	decl := regComp.FindAllStringSubmatch(arquivoComp, -1)

	var obj [][]string
	for _, x := range decl {
		if strings.Contains(x[0], "procedure") || strings.Contains(x[0], "function") {
			break
		}
		x[0] = strings.Trim(x[0], " ")
		x[1] = strings.Trim(x[1], " ")
		x[2] = strings.Trim(x[2], " ")
		obj = append(obj, x)

	}
	return obj
}

func validaNomenclatura() string {
	retorno, erro := RetornaTabelaComponentes()
	if erro != nil {
		fmt.Println(erro)
	}

	for key, value := range retorno {
		fmt.Println("Key:", key, "Value:", value)
	}

	return "teste"
}
