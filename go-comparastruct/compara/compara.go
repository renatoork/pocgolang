package compara

import (
	"fmt"
	"mega/go-util/dbg"
	"reflect"
	"strings"
	"time"
)

var nivel int

func CompararStruct(str1 interface{}, str2 interface{}) (bool, string) {
	if reflect.DeepEqual(str1, str2) {
		return true, ""
	}

	retorno = ""
	diff = false

	nivel = 0
	ret := deepCompare(str1, str2)

	return !diff, ret
}

func CompararStruct2(str1 interface{}, str2 interface{}) (bool, string) {
	if reflect.DeepEqual(str1, str2) {
		return true, ""
	}

	retorno = ""
	diff = false

	nivel = 0
	ret := deepCompare2(str1, str2)

	return !diff, ret
}

// Retorno:
// []operação,
// []estrutura para aplicar a operção,
// []posição no new,
// []posição no old
func CompararSliceStruct(slicenew interface{}, sliceold interface{}, fieldschave []string, filtrooperacao []string) ([]string, []interface{}, []int, []int) {
	var operacoes []string
	var opatual string
	var estruturas []interface{}
	var posnew []int
	var posold []int

	ori := reflect.ValueOf(sliceold).Elem()
	alt := reflect.ValueOf(slicenew).Elem()

	encontrados := []interface{}{}

	for i := 0; i < alt.Len(); i++ {
		opatual = "U"
		achou := false
		pnew := i
		pold := -1
		for j := 0; j < ori.Len(); j++ {
			chaveok := true
			for _, v := range fieldschave {
				if !comparaReflectField(getFieldByName(alt.Index(i), v), getFieldByName(ori.Index(j), v)) {
					chaveok = false
					break
				}
			}
			if chaveok {
				encontrados = append(encontrados, j)
				pold = j
				achou = true
				ok, _ := CompararStruct(alt.Index(i).Addr().Interface(), ori.Index(j).Addr().Interface())
				if ok {
					opatual = "N"
				}
				break
			}
		}
		if !achou {
			opatual = "I"
		}

		if len(filtrooperacao) == 0 || ContainsStr(filtrooperacao, opatual) {
			operacoes = append(operacoes, opatual)
			estruturas = append(estruturas, alt.Index(i).Addr().Interface())
			posnew = append(posnew, pnew)
			posold = append(posold, pold)
		}
	}
	opatual = "D"
	if len(filtrooperacao) == 0 || ContainsStr(filtrooperacao, opatual) {
		for i := 0; i < ori.Len(); i++ {
			if !Contains(encontrados, i) {
				operacoes = append(operacoes, opatual)
				estruturas = append(estruturas, ori.Index(i).Addr().Interface())
				posnew = append(posnew, -1)
				posold = append(posold, i)
			}
		}
	}
	return operacoes, estruturas, posnew, posold
}

func getFieldByName(estru reflect.Value, campo string) reflect.Value {
	var camporetorno reflect.Value

	if strings.Contains(campo, ".") {
		pos := strings.Index(campo, ".")
		if dbg.GetDebug() {
			fmt.Println("campo[0:pos] => ", campo[0:pos])
			fmt.Println("campo[pos+1:len(campo)] => ", campo[pos+1:len(campo)])
		}
		camporetorno = getFieldByName(estru.FieldByName(campo[0:pos]), campo[pos+1:len(campo)])
	} else {
		camporetorno = estru.FieldByName(campo)

	}

	return camporetorno

}

func comparaReflectField(fil1 reflect.Value, fil2 reflect.Value) bool {
	var (
		valorOri interface{}
		valorAlt interface{}
	)
	if dbg.GetDebug() {
		if !fil1.IsValid() {
			fmt.Println(" fil1 inválido: ")

		}

		if !fil2.IsValid() {
			fmt.Println(" fil2 inválido: ")

		}
	}
	if fil1.Type() != fil2.Type() {
		return false
	} else {
		if fil1.Type().Kind() == reflect.Ptr {
			if !fil1.IsNil() {
				valorOri = reflect.Indirect(fil1).Interface()
			} else {
				valorOri = nil
			}

		} else {
			valorOri = fil1.Interface()
		}

		if fil2.Type().Kind() == reflect.Ptr {
			if !fil2.IsNil() {
				valorAlt = reflect.Indirect(fil2).Interface()
			} else {
				valorAlt = nil
			}
		} else {
			valorAlt = fil2.Interface()
		}

	}

	return reflect.DeepEqual(valorOri, valorAlt)

}

var (
	retorno      string
	diff         bool
	nomeestrupai string
)

func deepCompare(ori interface{}, alt interface{}) string {

	vOri := reflect.ValueOf(ori).Elem()
	vAlt := reflect.ValueOf(alt).Elem()

	if vOri.Type() != vAlt.Type() {
		retorno = retorno + fmt.Sprintf("\nEstruturas com tipos diferentes: %s e %s", vOri.Type(), vAlt.Type())
		diff = true
	} else {

		for i := 0; i < vOri.NumField(); i++ {
			valueFieldOri := vOri.Field(i)
			typeFieldOri := vOri.Type().Field(i)
			valueFieldAlt := vAlt.Field(i)

			tag := typeFieldOri.Tag
			if tag.Get("comp") != "N" {

				var (
					valorOri interface{}
					valorAlt interface{}
				)

				if vOri.Type() != vAlt.Type() {
					retorno = retorno + fmt.Sprintf("\nTipos diferentes: (%v)  /  (%v)", vOri.Type(), vAlt.Type())
					diff = true
				} else {
					if valueFieldOri.Type().Kind() == reflect.Ptr {
						if !valueFieldOri.IsNil() {
							valorOri = reflect.Indirect(valueFieldOri).Interface()
						} else {
							valorOri = nil
						}

					} else {
						valorOri = valueFieldOri.Interface()
					}

					if valueFieldAlt.Type().Kind() == reflect.Ptr {
						if !valueFieldAlt.IsNil() {
							valorAlt = reflect.Indirect(valueFieldAlt).Interface()
						} else {
							valorAlt = nil
						}
					} else {
						valorAlt = valueFieldAlt.Interface()
					}

				}

				if valueFieldOri.Type().Kind() == reflect.Struct && valueFieldOri.Type() != reflect.TypeOf(time.Now()) {
					nivel++
					deepCompare(valueFieldOri.Addr().Interface(), valueFieldAlt.Addr().Interface())
				} else if valueFieldOri.Type().Kind() == reflect.Slice {
					nivel++
					for k := 0; k < valueFieldOri.Len(); k++ {
						if valueFieldOri.Index(k).Type().Kind() == reflect.Struct && valueFieldAlt.Len() > k {

							deepCompare(valueFieldOri.Index(k).Addr().Interface(), valueFieldAlt.Index(k).Addr().Interface())
						}
					}

					if valueFieldAlt.Len() > valueFieldOri.Len() || valueFieldAlt.Len() < valueFieldOri.Len() {
						diff = true
						retorno = retorno + fmt.Sprintf("\nQuantidades diferentes: %v.%s: (%v) / (%v)", vOri.Type(), typeFieldOri.Name, valueFieldOri.Len(), valueFieldAlt.Len())
					}
				} else {

					if !reflect.DeepEqual(valorOri, valorAlt) {
						diff = true
						retorno = retorno + fmt.Sprintf("\n %v.%s: (%v) / (%v)", vOri.Type(), typeFieldOri.Name, valorOri, valorAlt)
					}
				}
			}
		}
	}
	return retorno
}

func deepCompare2(ori interface{}, alt interface{}) string {

	vOri := reflect.ValueOf(ori).Elem()

	var vAlt reflect.Value
	if reflect.ValueOf(alt).IsValid() {
		vAlt = reflect.ValueOf(alt).Elem()
	} else {
		vAlt = reflect.ValueOf(alt)
	}

	if vAlt.IsValid() && (vOri.Type() != vAlt.Type()) {
		retorno = retorno + fmt.Sprintf("\nEstruturas com tipos diferentes: %s e %s", vOri.Type(), vAlt.Type())
	} else {

		for i := 0; i < vOri.NumField(); i++ {
			valueFieldOri := vOri.Field(i)
			typeFieldOri := vOri.Type().Field(i)

			if nivel == 0 {
				nomeestrupai = vOri.Type().String()
			}

			var valueFieldAlt reflect.Value
			if vAlt.IsValid() {
				valueFieldAlt = vAlt.Field(i)
			}

			tag := typeFieldOri.Tag
			if tag.Get("comp") != "N" {

				var (
					valorOri interface{}
					valorAlt interface{}
				)

				if vAlt.IsValid() && (valueFieldOri.Type() != valueFieldAlt.Type()) {
					retorno = retorno + fmt.Sprintf("\nTipos diferentes: (%v)  /  (%v)", vOri.Type(), vAlt.Type())
				} else {
					if valueFieldOri.Type().Kind() == reflect.Ptr {
						if !valueFieldOri.IsNil() {
							valorOri = reflect.Indirect(valueFieldOri).Interface()
						} else {
							valorOri = nil
						}

					} else {
						valorOri = valueFieldOri.Interface()
					}

					if vAlt.IsValid() {
						if valueFieldAlt.Type().Kind() == reflect.Ptr {
							if !valueFieldAlt.IsNil() {
								valorAlt = reflect.Indirect(valueFieldAlt).Interface()
							} else {
								valorAlt = nil
							}
						} else {
							valorAlt = valueFieldAlt.Interface()
						}
					} else {
						valorAlt = nil
					}

				}

				if valueFieldOri.Type().Kind() == reflect.Struct && valueFieldOri.Type() != reflect.TypeOf(time.Now()) {
					nivel++
					//retorno = retorno + fmt.Sprintf("\n [%v.%s]", vOri.Type(), typeFieldOri.Name)
					//nomeestrupaiant := nomeestrupai
					//nomeestrupai = fmt.Sprintf("%s.%s", nomeestrupaiant, typeFieldOri.Name)

					if vAlt.IsValid() {
						deepCompare2(valueFieldOri.Addr().Interface(), valueFieldAlt.Addr().Interface())
					} else {
						deepCompare2(valueFieldOri.Addr().Interface(), nil)

					}
					//nomeestrupai = nomeestrupaiant
					nivel--

				} else if valueFieldOri.Type().Kind() == reflect.Slice {
					nivel++
					nomeestrupaiant := nomeestrupai
					if vAlt.IsValid() {
						for k := 0; k < valueFieldOri.Len(); k++ {

							//retorno = retorno + fmt.Sprintf("\n [%d] <<%v.%s>>", k, vOri.Type(), typeFieldOri.Name)
							nomeestrupai = fmt.Sprintf("%s.%s[%d]", nomeestrupaiant, typeFieldOri.Name, k)

							if valueFieldOri.Index(k).Type().Kind() == reflect.Struct {

								if valueFieldAlt.Len() > k {
									deepCompare2(valueFieldOri.Index(k).Addr().Interface(), valueFieldAlt.Index(k).Addr().Interface())
								} else {
									deepCompare2(valueFieldOri.Index(k).Addr().Interface(), nil)
								}
							}

							nomeestrupai = nomeestrupaiant
						}
					} else {
						for k := 0; k < valueFieldOri.Len(); k++ {

							//retorno = retorno + fmt.Sprintf("\n [%d] <<%v.%s>>", k, vOri.Type(), typeFieldOri.Name)
							nomeestrupai = fmt.Sprintf("%s.%s[%d]", nomeestrupaiant, typeFieldOri.Name, k)

							if valueFieldOri.Index(k).Type().Kind() == reflect.Struct {

								deepCompare2(valueFieldOri.Index(k).Addr().Interface(), nil)
							}

							nomeestrupai = nomeestrupaiant
						}
					}
					nivel--
				} else {

					if !reflect.DeepEqual(valorOri, valorAlt) {
						diff = true
						retorno = retorno + fmt.Sprintf("\n %v.%s: (%v) / (%v)", nomeestrupai /*vOri.Type()*/, typeFieldOri.Name, valorOri, valorAlt)
					}
				}
			}
		}
	}
	return retorno
}

func Contains(slice []interface{}, item interface{}) bool {
	ok := false

	for _, s := range slice {
		if s == item {
			ok = true
		}
	}

	return ok
}

func ContainsStr(slice []string, item string) bool {
	ok := false

	for _, s := range slice {
		if s == item {
			ok = true
		}
	}

	return ok
}
