package customizacaoservico

import (
	"errors"
	"fmt"
	"mega/go-util/dbg"
	"reflect"
	"strings"

	_ "gopkg.in/goracle.v1"
	_ "gopkg.in/goracle.v1/oracle"
	"gopkg.in/gorp.v1"
)

func InsereCustomizacao(ownerpai string, nometabelapai string, estruturacomcustomizacoes interface{}, trans *gorp.Transaction) error {
	valcustom := reflect.ValueOf(estruturacomcustomizacoes).Elem()

	if valcustom.Type().Kind() == reflect.Slice {
		for i := 0; i < valcustom.Len(); i++ {
			if valcustom.Index(i).FieldByName("Customizacoes").IsValid() {
				err := InsereCustomizacao(ownerpai, nometabelapai, valcustom.Index(i).Addr().Interface(), trans)
				if err != nil {
					return err
				}
			}
		}
	} else if valcustom.FieldByName("Customizacoes").IsValid() {
		colunas, valores := estruturaParaColunasValores(estruturacomcustomizacoes)
		return insereCustomizacaoInterno(ownerpai, nometabelapai, colunas, valores, valcustom.FieldByName("Customizacoes").Addr().Interface().(*Customizacao).Tabela, trans)
	}

	return nil
}

func insereCustomizacaoInterno(ownerpai string, nometabelapai string, colunaspai []string, valorespai []interface{}, customizacoes []Tabela, trans *gorp.Transaction) error {

	for i := 0; i < len(customizacoes); i++ {
		ownertabela := strings.Split(customizacoes[i].Nometabela, ".")

		fkpai, fkfilho, err := carregaCamposFkCustom(ownerpai, nometabelapai, ownertabela[0], ownertabela[1], trans)
		if err != nil {
			return err
		}

		fkvalores := filtraValoresColunas(colunaspai, valorespai, fkpai)

		err = insereTabelaCustomizada(ownertabela[0], ownertabela[1], customizacoes[i].Registros, fkfilho, fkvalores, trans)
		if err != nil {
			return err
		}
	}

	return nil
}

func insereTabelaCustomizada(owner string, nometabela string, registros []Registro, colunasfk []string, valoresfk []interface{}, trans *gorp.Transaction) error {

	for i := 0; i < len(registros); i++ {
		colunas := estrutcolunasParaColunas(registros[i].Colunas)
		valores := estrutvaloresParaValores(registros[i].Colunas)

		for k := 0; k < len(colunasfk); k++ {
			achou := false
			for j := 0; j < len(colunas); j++ {
				if colunasfk[k] == colunas[j] {
					valores[j] = valoresfk[k]
					achou = true
				}
			}
			if !achou && valoresfk[k] != nil {
				colunas = append(colunas, colunasfk[k])
				valores = append(valores, valoresfk[k])
			}
		}

		err := insereRegistroCustomizado(owner, nometabela, colunas, valores, trans)
		if err != nil {
			return err
		}

		err = insereCustomizacaoInterno(owner, nometabela, colunas, valores, registros[i].Tabela, trans)
		if err != nil {
			return err
		}
	}

	return nil
}

func insereRegistroCustomizado(owner string, nometabela string, colunas []string, valores []interface{}, trans *gorp.Transaction) error {

	sqlcmd := geraInsertCustom(owner, nometabela, colunas)
	if dbg.GetDebug() {
		fmt.Println("COMANDO: ", sqlcmd)
		fmt.Println("VALORES: ", valores)
	}
	stmt, err := trans.Prepare(sqlcmd)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(valores...)
	if err != nil {
		return err
	}

	return nil
}

func CarregaCustomizacao(ownerpai string, nometabelapai string, estruturapai interface{}, mapeamento map[string][]string, estruturacomcustomizacoes interface{}, trans *gorp.Transaction) error {

	if len(mapeamento[ownerpai+"."+nometabelapai]) > 0 {

		val := reflect.ValueOf(estruturapai).Elem()
		valcustom := reflect.ValueOf(estruturacomcustomizacoes).Elem()

		if val.Type().Kind() == reflect.Slice {
			for i := 0; i < val.Len(); i++ {
				if valcustom.Index(i).FieldByName("Customizacoes").IsValid() {
					err := CarregaCustomizacao(ownerpai, nometabelapai, val.Index(i).Addr().Interface(), mapeamento, &valcustom.Index(i).FieldByName("Customizacoes").Addr().Interface().(*Customizacao).Tabela, trans)
					if err != nil {
						return err
					}
				}
			}
		} else {
			colunas, valores := estruturaParaColunasValores(estruturapai)
			return carregaCustomizacaoInterno(ownerpai, nometabelapai, colunas, valores, mapeamento, valcustom.Addr().Interface().(*[]Tabela), trans)
		}
	}

	return nil
}

func carregaCustomizacaoInterno(ownerpai string, nometabelapai string, colunaspai []string, valorespai []interface{}, mapeamento map[string][]string, customizacoes *[]Tabela, trans *gorp.Transaction) error {
	nomecompletotabelapai := ownerpai + "." + nometabelapai

	if len(mapeamento[nomecompletotabelapai]) > 0 {
		for _, v := range mapeamento[nomecompletotabelapai] {
			ownertabela := strings.Split(v, ".")

			fkpai, fkfilho, err := carregaCamposFkCustom(ownerpai, nometabelapai, ownertabela[0], ownertabela[1], trans)
			if err != nil {
				return err
			}
			fkvalores := filtraValoresColunas(colunaspai, valorespai, fkpai)

			tabcustom, err := carregaTabelaCustomizada(ownertabela[0], ownertabela[1], fkfilho, fkvalores, trans)
			if err != nil {
				return err
			}
			*customizacoes = append(*customizacoes, tabcustom)

			for i := 0; i < len(tabcustom.Registros); i++ {
				colunas := estrutcolunasParaColunas(tabcustom.Registros[i].Colunas)
				valores := estrutvaloresParaValores(tabcustom.Registros[i].Colunas)
				err = carregaCustomizacaoInterno(ownertabela[0], ownertabela[1], colunas, valores, mapeamento, &tabcustom.Registros[i].Tabela, trans)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func carregaTabelaCustomizada(owner string, nometabela string, colunaswhere []string, valoreswhere []interface{}, trans *gorp.Transaction) (Tabela, error) {

	var tabelaestru Tabela = Tabela{}
	tabelaestru.Nometabela = owner + "." + nometabela

	sqlcmd := geraSelectCustom(owner, nometabela, colunaswhere)

	stmt, err := trans.Prepare(sqlcmd)
	if err != nil {
		return tabelaestru, err
	}

	rows, err := stmt.Query(valoreswhere...)
	if err != nil {
		return tabelaestru, err
	}
	defer rows.Close()

	colunas, err := rows.Columns()
	if err != nil {
		return tabelaestru, err
	}

	var valores = make([]interface{}, len(colunas))
	var valoresptrs = make([]interface{}, len(colunas))

	for rows.Next() {
		var reg Registro

		for i, _ := range valores {
			valoresptrs[i] = &valores[i]
		}

		rows.Scan(valoresptrs...)

		for i := 0; i < len(colunas); i++ {
			reg.Colunas = append(reg.Colunas, Coluna{Nome: colunas[i], Valor: valores[i]})
		}

		tabelaestru.Registros = append(tabelaestru.Registros, reg)
	}

	return tabelaestru, nil
}

func AtualizaCustomizacao(ownerpai string, nometabelapai string, estruturapaiori interface{}, estruturapaialt interface{}, trans *gorp.Transaction) error {

	var valori reflect.Value
	var valalt reflect.Value

	if reflect.ValueOf(estruturapaiori).IsValid() {
		valori = reflect.ValueOf(estruturapaiori).Elem()
	} else {
		valori = reflect.ValueOf(estruturapaiori)
	}

	if reflect.ValueOf(estruturapaialt).IsValid() {
		valalt = reflect.ValueOf(estruturapaialt).Elem()
	} else {
		valalt = reflect.ValueOf(estruturapaialt)
	}

	if (valori.IsValid() && valori.Type().Kind() == reflect.Slice) && (valalt.IsValid() && valalt.Type().Kind() == reflect.Slice) {
		if valalt.IsValid() {
			for i := 0; i < valalt.Len(); i++ {
				if valalt.Index(i).FieldByName("Customizacoes").IsValid() {
					err := AtualizaCustomizacao(ownerpai, nometabelapai, estruturapaiori, valalt.Index(i).Addr().Interface(), trans)
					if err != nil {
						return err
					}
				}
			}
		}
		if valori.IsValid() {
			for i := 0; i < valori.Len(); i++ {
				if valori.Index(i).FieldByName("Customizacoes").IsValid() {
					err := AtualizaCustomizacao(ownerpai, nometabelapai, valori.Index(i).Addr().Interface(), estruturapaialt, trans)
					if err != nil {
						return err
					}
				}
			}
		}

	} else if (valori.IsValid() && valori.Type().Kind() != reflect.Slice) && (!valalt.IsValid() || valalt.Type().Kind() == reflect.Slice) {
		achou := false
		colunasori, valoresori := estruturaParaColunasValores(estruturapaiori)

		if valalt.IsValid() {

			pk, err := carregaCamposPkCustom(ownerpai, nometabelapai, trans)
			if err != nil {
				return err
			}

			valorespkori := filtraValoresColunas(colunasori, valoresori, pk)

			if TemNilNoArray(valorespkori) {
				return errors.New("Não foram enviados os dados da chave primária da tabela " + ownerpai + "." + nometabelapai + " na estrutura original.")
			}

			for i := 0; i < valalt.Len(); i++ {
				colunasalt, valoresalt := estruturaParaColunasValores(valalt.Index(i).Addr().Interface())
				valorespkalt := filtraValoresColunas(colunasalt, valoresalt, pk)

				if TemNilNoArray(valorespkalt) {
					return errors.New("Não foram enviados os dados da chave primária da tabela " + ownerpai + "." + nometabelapai + " na estrutura alterada.")
				}

				if ComparaArrayInterface(valorespkori, valorespkalt) {
					achou = true
					break
				}
			}
		}

		if !achou {
			err := removeCustomizacaoInterno(ownerpai, nometabelapai, colunasori, valoresori, valori.FieldByName("Customizacoes").Addr().Interface().(*Customizacao).Tabela, trans)
			if err != nil {
				return err
			}
		}

	} else if (!valori.IsValid() || valori.Type().Kind() == reflect.Slice) && (valalt.IsValid() && valalt.Type().Kind() != reflect.Slice) {
		achou := false
		colunasalt, valoresalt := estruturaParaColunasValores(estruturapaialt)

		if valori.IsValid() {
			pk, err := carregaCamposPkCustom(ownerpai, nometabelapai, trans)
			if err != nil {
				return err
			}

			valorespkalt := filtraValoresColunas(colunasalt, valoresalt, pk)

			if TemNilNoArray(valorespkalt) {
				return errors.New("Não foram enviados os dados da chave primária da tabela " + ownerpai + "." + nometabelapai + " na estrutura alterada.")
			}

			for i := 0; i < valori.Len(); i++ {
				colunasori, valoresori := estruturaParaColunasValores(valori.Index(i).Addr().Interface())
				valorespkori := filtraValoresColunas(colunasori, valoresori, pk)

				if TemNilNoArray(valorespkori) {
					return errors.New("Não foram enviados os dados da chave primária da tabela " + ownerpai + "." + nometabelapai + " na estrutura original.")
				}

				if ComparaArrayInterface(valorespkori, valorespkalt) {
					achou = true
					if valori.Index(i).FieldByName("Customizacoes").IsValid() {
						err := atualizaCustomizacaoInterno(ownerpai, nometabelapai, colunasalt, valoresalt, valori.Index(i).FieldByName("Customizacoes").Addr().Interface().(*Customizacao).Tabela, valalt.FieldByName("Customizacoes").Addr().Interface().(*Customizacao).Tabela, trans)
						if err != nil {
							return err
						}
					}
					break
				}
			}

			if !achou {
				err := insereCustomizacaoInterno(ownerpai, nometabelapai, colunasalt, valoresalt, valalt.FieldByName("Customizacoes").Addr().Interface().(*Customizacao).Tabela, trans)
				if err != nil {
					return err
				}
			}
		}

	} else if valori.IsValid() || valalt.IsValid() {
		if valori.IsValid() && !valalt.IsValid() {
			colunasori, valoresori := estruturaParaColunasValores(estruturapaiori)
			err := removeCustomizacaoInterno(ownerpai, nometabelapai, colunasori, valoresori, valori.FieldByName("Customizacoes").Addr().Interface().(*Customizacao).Tabela, trans)
			if err != nil {
				return err
			}
		} else if !valori.IsValid() && valalt.IsValid() {
			colunasalt, valoresalt := estruturaParaColunasValores(estruturapaialt)
			err := insereCustomizacaoInterno(ownerpai, nometabelapai, colunasalt, valoresalt, valalt.FieldByName("Customizacoes").Addr().Interface().(*Customizacao).Tabela, trans)
			if err != nil {
				return err
			}
		} else {

			pk, err := carregaCamposPkCustom(ownerpai, nometabelapai, trans)
			if err != nil {
				return err
			}

			colunasalt, valoresalt := estruturaParaColunasValores(estruturapaialt)
			valorespkalt := filtraValoresColunas(colunasalt, valoresalt, pk)

			colunasori, valoresori := estruturaParaColunasValores(estruturapaiori)
			valorespkori := filtraValoresColunas(colunasori, valoresori, pk)

			if TemNilNoArray(valorespkori) {
				return errors.New("Não foram enviados os dados da chave primária da tabela " + ownerpai + "." + nometabelapai + " na estrutura original.")
			}

			if TemNilNoArray(valorespkalt) {
				return errors.New("Não foram enviados os dados da chave primária da tabela " + ownerpai + "." + nometabelapai + " na estrutura alterada.")
			}

			if ComparaArrayInterface(valorespkori, valorespkalt) {
				err := atualizaCustomizacaoInterno(ownerpai, nometabelapai, colunasalt, valoresalt, valori.FieldByName("Customizacoes").Addr().Interface().(*Customizacao).Tabela, valalt.FieldByName("Customizacoes").Addr().Interface().(*Customizacao).Tabela, trans)
				if err != nil {
					return err
				}
			} else {
				err := AtualizaCustomizacao(ownerpai, nometabelapai, estruturapaiori, nil, trans)
				if err != nil {
					return err
				}

				err = AtualizaCustomizacao(ownerpai, nometabelapai, nil, estruturapaialt, trans)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func atualizaCustomizacaoInterno(ownerpai string, nometabelapai string, colunaspai []string, valorespai []interface{}, customizacoesori []Tabela, customizacoesalt []Tabela, trans *gorp.Transaction) error {

	for i := 0; i < len(customizacoesori); i++ {
		indicetabela := -1

		ownertabela := strings.Split(customizacoesori[i].Nometabela, ".")
		colunaswhere, err := carregaCamposPkCustom(ownertabela[0], ownertabela[1], trans)
		if err != nil {
			return err
		}

		for k := 0; k < len(customizacoesalt); k++ {
			if strings.EqualFold(customizacoesori[i].Nometabela, customizacoesalt[k].Nometabela) {
				indicetabela = k
				break
			}
		}

		if indicetabela > -1 {
			err = atualizaTabelaCustomizada(ownertabela[0], ownertabela[1], customizacoesori[i].Registros, customizacoesalt[indicetabela].Registros, colunaswhere, trans)
			if err != nil {
				return err
			}
		} else {
			fkpai, fkfilho, err := carregaCamposFkCustom(ownerpai, nometabelapai, ownertabela[0], ownertabela[1], trans)
			if err != nil {
				return err
			}
			fkvalores := filtraValoresColunas(colunaspai, valorespai, fkpai)

			err = removeTabelaCustomizada(ownertabela[0], ownertabela[1], customizacoesori[i].Registros, fkfilho, fkvalores, trans)
			if err != nil {
				return err
			}
		}
	}

	for i := 0; i < len(customizacoesalt); i++ {
		indicetabela := -1

		ownertabela := strings.Split(customizacoesalt[i].Nometabela, ".")
		for k := 0; k < len(customizacoesori); k++ {
			if strings.EqualFold(customizacoesalt[i].Nometabela, customizacoesori[k].Nometabela) {
				indicetabela = k
				break
			}
		}

		if indicetabela == -1 {
			fkpai, fkfilho, err := carregaCamposFkCustom(ownerpai, nometabelapai, ownertabela[0], ownertabela[1], trans)
			if err != nil {
				return err
			}

			fkvalores := filtraValoresColunas(colunaspai, valorespai, fkpai)

			err = insereTabelaCustomizada(ownertabela[0], ownertabela[1], customizacoesalt[i].Registros, fkfilho, fkvalores, trans)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func atualizaTabelaCustomizada(owner string, nometabela string, registrosori []Registro, registrosalt []Registro, colunaswhere []string, trans *gorp.Transaction) error {

	for _, v := range registrosalt {
		colunas := estrutcolunasParaColunas(v.Colunas)
		valores := estrutvaloresParaValores(v.Colunas)

		valoreswhere := filtraValoresColunas(colunas, valores, colunaswhere)

		if len(colunaswhere) != len(valoreswhere) {
			return errors.New("Não foram enviados os dados da chave primária da tabela " + owner + "." + nometabela)
		}

		achou := false
		for _, d := range registrosori {

			colunasori := estrutcolunasParaColunas(d.Colunas)
			valoresori := estrutvaloresParaValores(d.Colunas)

			valoreswhereori := filtraValoresColunas(colunasori, valoresori, colunaswhere)

			if len(colunaswhere) != len(valoreswhereori) {
				return errors.New("Não foram enviados os dados da chave primária da tabela " + owner + "." + nometabela)
			}

			if ComparaArrayInterface(valoreswhereori, valoreswhere) {
				achou = true

				if !ComparaArrayInterface(valoresori, valores) {
					err := atualizaRegistroCustomizado(owner, nometabela, colunas, valores, colunaswhere, valoreswhere, trans)
					if err != nil {
						return err
					}
				}

				err := atualizaCustomizacaoInterno(owner, nometabela, colunas, valores, d.Tabela, v.Tabela, trans)
				if err != nil {
					return err
				}

				break
			}

		}

		if !achou {
			err := insereRegistroCustomizado(owner, nometabela, colunas, valores, trans)
			if err != nil {
				return err
			}

			err = insereCustomizacaoInterno(owner, nometabela, colunas, valores, v.Tabela, trans)
			if err != nil {
				return err
			}
		}

	}

	for _, v := range registrosori {
		colunas := estrutcolunasParaColunas(v.Colunas)
		valores := estrutvaloresParaValores(v.Colunas)

		valoreswhere := filtraValoresColunas(colunas, valores, colunaswhere)

		achou := false
		for _, d := range registrosalt {

			colunasalt := estrutcolunasParaColunas(d.Colunas)
			valoresalt := estrutvaloresParaValores(d.Colunas)

			valoreswherealt := filtraValoresColunas(colunasalt, valoresalt, colunaswhere)

			if ComparaArrayInterface(valoreswhere, valoreswherealt) {
				achou = true
				break
			}

		}

		if !achou {
			err := removeRegistroCustomizado(owner, nometabela, colunaswhere, valoreswhere, trans)
			if err != nil {
				return err
			}

		}

	}

	return nil
}

func atualizaRegistroCustomizado(owner string, nometabela string, colunas []string, valores []interface{}, colunaswhere []string, valoreswhere []interface{}, trans *gorp.Transaction) error {

	sqlcmd := geraUpdateCustom(owner, nometabela, colunas, colunaswhere)
	if dbg.GetDebug() {
		fmt.Println("COMANDO: ", sqlcmd)
		fmt.Println("VALORES: ", valores)
		fmt.Println("VALORES WHERE: ", valoreswhere)
	}
	stmt, err := trans.Prepare(sqlcmd)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(append(valores, valoreswhere...)...)
	if err != nil {
		return err
	}

	return nil
}

func removeCustomizacaoInterno(ownerpai string, nometabelapai string, colunaspai []string, valorespai []interface{}, customizacoes []Tabela, trans *gorp.Transaction) error {

	for i := 0; i < len(customizacoes); i++ {
		ownertabela := strings.Split(customizacoes[i].Nometabela, ".")

		fkpai, fkfilho, err := carregaCamposFkCustom(ownerpai, nometabelapai, ownertabela[0], ownertabela[1], trans)
		if err != nil {
			return err
		}
		fkvalores := filtraValoresColunas(colunaspai, valorespai, fkpai)

		err = removeTabelaCustomizada(ownertabela[0], ownertabela[1], customizacoes[i].Registros, fkfilho, fkvalores, trans)
		if err != nil {
			return err
		}
	}

	return nil

}

func removeTabelaCustomizada(owner string, nometabela string, registros []Registro, colunasfk []string, valoresfk []interface{}, trans *gorp.Transaction) error {

	for _, v := range registros {
		colunas := estrutcolunasParaColunas(v.Colunas)
		valores := estrutvaloresParaValores(v.Colunas)

		err := removeCustomizacaoInterno(owner, nometabela, colunas, valores, v.Tabela, trans)
		if err != nil {
			return err
		}
	}

	err := removeRegistroCustomizado(owner, nometabela, colunasfk, valoresfk, trans)
	if err != nil {
		return err
	}

	return nil
}

func removeRegistroCustomizado(owner string, nometabela string, colunaswhere []string, valoreswhere []interface{}, trans *gorp.Transaction) error {

	sqlcmd := geraDeleteCustom(owner, nometabela, colunaswhere)
	if dbg.GetDebug() {
		fmt.Println("COMANDO: ", sqlcmd)
		fmt.Println("VALORES WHERE: ", valoreswhere)
	}
	stmt, err := trans.Prepare(sqlcmd)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(valoreswhere...)
	if err != nil {
		return err
	}

	return nil
}
