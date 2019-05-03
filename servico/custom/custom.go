package custom

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"mega/servico/banco"
	"net/http"
	"reflect"
)

type ValorCustom struct {
	Campo string `db:"CAMPO"`
	Valor string `db:"VALOR"`
	Tipo  string `db:"TIPO"`
}

type CampoCustom struct {
	Campo  string
	Campos []CampoCustom `json:"Campos,omitempty"`
}

type CampoCustomBanco struct {
	Id        int64
	Tabela_Id int64
	Campo     string
	Tipo      string
	Pai       sql.NullInt64
}

func Inicializa(router *mux.Router) {
	banco.Dbmap.AddTableWithNameAndSchema(ValorCustom{}, "mgglo", "glo_valorcustom")
	banco.Dbmap.AddTableWithNameAndSchema(CampoCustomBanco{}, "mgglo", "glo_campocustom")
	router.HandleFunc("/api/custom/{id}", customHandler)
}

func GetValores(tabela int, registro int) map[string]interface{} {
	var valores []ValorCustom
	query := `select campo,
	       valor,
	       tipo
	  from mgglo.glo_valorcustom valor,
	       mgglo.glo_campocustom campo
	 where campo.id = valor.campo_id
	       and tabela_id = :tabela
	       and registro_id = :registro`
	_, err := banco.Dbmap.Select(&valores, query, tabela, registro)

	if err != nil {
		panic(err)
	}

	retorno := make(map[string]interface{})

	for _, valor := range valores {
		if valor.Tipo != "json" {
			retorno[valor.Campo] = valor.Valor
		} else {
			var valorJson interface{}
			err := json.Unmarshal([]byte(valor.Valor), &valorJson)
			if err != nil {
				panic(err)
			}
			retorno[valor.Campo] = valorJson
		}
	}

	return retorno
}

func GravaValores(tabela int, registro int, valores map[string]interface{}) error {
	comando := `merge into mgglo.glo_valorcustom valor
	using (
	  select id
	  from mgglo.glo_campocustom campo
	 where tabela_id = 1
	       and campo = :nomecampo
	) e
	on (e.id = valor.campo_id and registro_id = :registro)
	when matched then
	  update set valor.valor = :valor
	when not matched then
	  insert (valor.registro_id, valor.campo_id, valor.valor)
	  values (:registro, e.id, :valor)`
	for chave, valor := range valores {
		if reflect.TypeOf(valor).Kind() == reflect.Slice {
			valorJson, err := json.MarshalIndent(&valor, "", "  ")
			if err != nil {
				return err
			}
			valor = string(valorJson)

		}
		if _, err := banco.Dbmap.Exec(comando, chave, registro, valor, registro, valor); err != nil {
			return err
		}
	}
	return nil
}

func GetDefinicao(tabela string) ([]CampoCustom, error) {
	var camposBanco []CampoCustomBanco
	_, err := banco.Dbmap.Select(&camposBanco, "select * from mgglo.glo_campocustom where tabela_id = :tabela order by id", 1)
	if err != nil {
		return nil, err
	}

	camposMap := make(map[int64]*CampoCustom)
	for _, campo := range camposBanco {
		novoCampo := CampoCustom{Campo: campo.Campo}
		if !campo.Pai.Valid {
			camposMap[campo.Id] = &novoCampo
		} else {
			pai := camposMap[campo.Pai.Int64]
			if pai.Campos == nil {
				pai.Campos = make([]CampoCustom, 0, 1)
			}
			pai.Campos = append(pai.Campos, novoCampo)
			fmt.Println(pai)
			fmt.Println(novoCampo)
		}
	}
	campos := make([]CampoCustom, 0, len(camposMap))
	for _, campo := range camposMap {
		campos = append(campos, *campo)
	}
	fmt.Println(camposMap)
	return campos, nil
}

func customHandler(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	definicao, err := GetDefinicao(vars["id"])
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	}

	err = json.NewEncoder(res).Encode(definicao)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	}
}
