package definicao

import (
	"mega/servico/banco"
	"time"
)

func AchaPadrao(tabela int, data time.Time) (int, error) {
	padrao, err := banco.Dbmap.SelectInt("select mgglo.pck_mega.AchaPadraoDaTabela(:filial, :tabela, :tipoagente, :data, 'N') from dual", 3, 53, "", data)
	return int(padrao), err
}
