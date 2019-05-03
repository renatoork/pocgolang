package customizacaoservico

type Coluna struct {
	Nome  string      `json:"nome"`
	Valor interface{} `json:"valor"`
}

type Registro struct {
	Colunas []Coluna `json:"colunas"`
	Tabela  []Tabela `json:"tabela"`
}

type Tabela struct {
	Nometabela string     `json:"nometabela"`
	Registros  []Registro `json:"registros"`
}

type Customizacao struct {
	Tabela []Tabela `json:"tabela"`
}
