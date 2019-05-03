# Go-Util #

Essa package tem algumas funções para debug e tratamento de erro

# Instalação e uso #

Fazer o clone do repositório para a pasta **%GOPATH%/src/mega**.

Depois é só importar as packages no seu projeto:

``` go
import (
	"mega/go-util/dbg"
	"mega/go-util/erro"
)
```

# Funções #

## mega/go-util/erro ##

### Trata ###

Essa função recebe um erro como parâmetro. Se ele for diferente de nil, será impresso e a função retorna true, caso contrário retorna false. Se o debug estiver ligado (ver abaixo), essa função também imprime o call stack.

Exemplo:

``` go
arquivo, err := os.Create(magnetico.Destino)
erro.Trata(err)
```

O retorno é útil para fazer algum tratamento se der erro, como por exemplo sair da aplicação após imprimir o erro:

```go
db, err := sql.Open("oci8", connectString)
if erro.Trata(err) {
	os.Exit(1)
}
```

## mega/go-debug/dbg ##

### SetDebug ###

Liga ou desliga o debug, conforme o parâmetro bool passado. Isso é gravado na variável de ambiente GO-DEBUG.

### GetDebug ###

Retorna true se o debug estiver ligado. Lê da variável de ambiente GO-DEBUG.

### Print ###

Imprime um título e a representação JSON de um objeto passado como parâmetro, por exemplo:

``` go
package main

import (
	"mega/go-util/dbg"
)

type Pessoa struct {
	Nome  string
	Idade int
}

func main() {
	dbg.SetDebug(true)
	dbg.Print("pessoa", Pessoa{Nome: "Fulano", Idade: 30})
}
```

O resultado gerado é:
``` json
pessoa
{
  "Nome": "Fulano",
  "Idade": 30
}
```

### Trace ###
Imprime o tempo entre a hora passada e a hora atual. É útil para medir o tempo que uma função  demora para ser executada. É possível fazer tudo em uma linha, adicionado uma chamada ao início da função, conforme exemplo:

``` go
package main

import (
	"mega/go-util/dbg"
	"time"
)

func main() {
	defer dbg.Trace(time.Now())
	a := 0
	for x := 0; x < 1000000000; x++ {
		a++
	}
}
```