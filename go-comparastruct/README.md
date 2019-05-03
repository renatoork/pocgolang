# Go-ComparaStruct #

Essa package trata a comparação de duas estruturas Go (struct)

# Instalação e uso #

Fazer o clone do repositório para a pasta **%GOPATH%/src/mega**.

Importar as packages no seu projeto:

``` go
import (
	"mega/go-comparastruct/struct"
)
```

# Funções #

## mega/go-comparastruct/compara ##

### ComparaStruct ###

Essa função recebe duas struct e mostra os valores diferentes entre elas. As tags que não houverem diferenças é mostrado um '-'. Quando houver um slice onde a quantidade de itens estiver diferente, a rotina mostra a informação que há diferenças, mas não mostra o conteúdo do slice para estas linhas. Se na definição da struct de origem houver uma TAG [comp:"N"] o comparador ignora a linha. Na ausencia da TAG ou se [comp:"S"], a rotina compara a linha.

Exemplo:

``` go
resultado, diferecas := compara.compararStruct(&structOrigem, &structAlterada)
```

O retorno é um bool sendo true para arquivos iguais e false para diferente. E o segundo retorno é uma string contendo as diferenças se houver.

