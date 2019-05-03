package customizacaoservico

import (
	"fmt"
	"mega/go-util/dbg"
	"net/http"
	"reflect"
	"strings"
	"time"
)

func estrutcolunasParaColunas(estcolunas []Coluna) []string {
	var colunas []string

	for _, v := range estcolunas {
		colunas = append(colunas, strings.ToUpper(v.Nome))
	}

	return colunas
}

func estrutvaloresParaValores(estcvalores []Coluna) []interface{} {
	var valores []interface{}

	for _, v := range estcvalores {
		if !strings.Contains(v.Nome, "_DT_") {
			valores = append(valores, v.Valor)
		} else {
			var d interface{}
			switch t := v.Valor.(type) {
			case string:
				d, _ = time.Parse(time.RFC3339, t)
			default:
				d = t
			}
			valores = append(valores, d)
		}
	}

	return valores
}

func filtraValoresColunas(colunas []string, valores []interface{}, colunasfiltro []string) []interface{} {
	var valoresretorno []interface{}

	if len(colunasfiltro) > 0 {
		for k := 0; k < len(colunasfiltro); k++ {
			valoresretorno = append(valoresretorno, nil)
			for j := 0; j < len(colunas); j++ {
				if colunasfiltro[k] == colunas[j] {
					valoresretorno[k] = valores[j]
					break
				}
			}
		}
	}

	return valoresretorno
}

func Debug_estruturaParaColunasValores(estrut interface{}) ([]string, []interface{}) {
	return estruturaParaColunasValores(estrut)
}

func estruturaParaColunasValores(estrut interface{}) ([]string, []interface{}) {
	var colunas []string
	var valores []interface{}

	val := reflect.ValueOf(estrut).Elem()
	if val.IsValid() {
		for i := 0; i < val.NumField(); i++ {
			valueField := val.Field(i)
			typeField := val.Type().Field(i)
			tag := typeField.Tag

			if (tag.Get("db") != "-") && (tag.Get("db") != "") {

				var valor interface{}
				if valueField.Type().Kind() == reflect.Ptr {
					if !valueField.IsNil() {
						valor = reflect.Indirect(valueField).Interface()
					} else {
						valor = nil
					}

				} else {
					valor = valueField.Interface()
				}

				colunas = append(colunas, strings.ToUpper(tag.Get("db")))
				valores = append(valores, valor)

			} else if valueField.Type().Kind() == reflect.Struct {

				colunasestrut, valoresestrut := estruturaParaColunasValores(valueField.Addr().Interface())
				colunas = append(colunas, colunasestrut...)
				valores = append(valores, valoresestrut...)

			}

		}
	}
	return colunas, valores
}

func CarregaTabelasCustomDoHeader(req *http.Request) map[string][]string {
	tabelasmapeadas := make(map[string][]string)

	custom := strings.ToUpper(req.Header.Get("Customizacoes"))
	tabelas := strings.Split(custom, ";")
	for _, v := range tabelas {
		tabelasmapeadas[req.Header.Get(v)] = append(tabelasmapeadas[req.Header.Get(v)], v)
	}

	return tabelasmapeadas
}

func ComparaArrayInterface(src1 []interface{}, src2 []interface{}) bool {
	if len(src1) != len(src2) {
		if dbg.GetDebug() {
			fmt.Println("Quantidade diferente: ", len(src1), len(src2))
		}
		return false
	}

	for i := 0; i < len(src1); i++ {
		var v1 interface{}
		var v2 interface{}

		if src1[i] != nil {
			if reflect.TypeOf(src1[i]).Kind() == reflect.Int32 {
				v1 = float64(src1[i].(int32))
			} else if reflect.TypeOf(src1[i]).Kind() == reflect.Int64 {
				v1 = float64(src1[i].(int64))
			} else if reflect.TypeOf(src1[i]) == reflect.TypeOf(time.Now()) {
				v1 = src1[i].(time.Time).Format(time.RFC3339 /*"2006-01-01 15:04:05"*/)
			} else {
				v1 = src1[i]
			}
		}
		if src2[i] != nil {
			if reflect.TypeOf(src2[i]).Kind() == reflect.Int32 {
				v2 = float64(src2[i].(int32))
			} else if reflect.TypeOf(src2[i]).Kind() == reflect.Int64 {
				v2 = float64(src2[i].(int64))
			} else if reflect.TypeOf(src2[i]) == reflect.TypeOf(time.Now()) {
				v2 = src2[i].(time.Time).Format(time.RFC3339 /*"2006-01-01 15:04:05"*/)
			} else {
				v2 = src2[i]
			}
		}

		if v1 != v2 {
			if dbg.GetDebug() {
				fmt.Println("v1: ", v1, "tipo: ", reflect.TypeOf(src1[i]))
				fmt.Println("v2: ", v2, "tipo: ", reflect.TypeOf(src2[i]))
			}
			return false
		}
	}

	return true
}

func TemNilNoArray(a []interface{}) bool {
	for _, v := range a {
		if v == nil {
			return true
		}
	}
	return false
}
