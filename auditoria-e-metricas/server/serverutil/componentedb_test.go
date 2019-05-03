package serverutil

import (
	"../../auditutil"
	//	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"testing"
)

func criaTabelaComponente() error {
	db := auditutil.GetSession()

	c := mgo.Collection{db, "componentes", "test.componentes"}
	i := mgo.CollectionInfo{}
	err := c.Create(&i)
	index := mgo.Index{
		Key:      []string{"componente"},
		Unique:   true,
		DropDups: true,
	}
	err = c.EnsureIndex(index)
	if err != nil {
		return err
	}
	return nil
}

func InsereComponentes(arquivo string) error {
	err := criaTabelaComponente()
	componentejson, err := auditutil.LerArquivoTeste(arquivo)
	auditutil.CheckError(err)

	dbComponente := auditutil.GetSession().C("componentes")
	for i, v := range componentejson {
		alterados, erro := dbComponente.Upserted(v.componente, v.prefixo)
		if len(alterados) > 0 {
			fmt.Println(alterados[0], alterados[1])
		}
		auditutil.CheckError(err)
	}
	return nil

}

func TestRetornaTabelaComponentes(t *testing.T) {
	erro := InsereComponentes("./teste/Componentes.json")
	if erro != nil {
		t.Log(erro)
		t.Fail()
	}

	componenteDB, erro := RetornaTabelaComponentes()
	if erro != nil {
		t.Log(erro)
		t.Fail()
	}
	if len(componenteDB) <= 0 {
		t.Log(fmt.Printf("Registros não recuperados..."))
		t.Fail()
	}
	/*
		componenteDBjson, erroDB := json.Marshal(componenteDB)
		if erroDB != nil {
			t.Log(erroDB)
			t.Fail()
		}

		componenteArqjson, erroArq := auditutil.LerArquivoTeste("./teste/Componentes.json")
		if erroArq != nil {
			t.Log(erroArq)
			t.Fail()
		}

		/*	for i := range componenteArqjson {
				pos := Select()
				if DeepEqual(componenteDBjson, componenteArqjson) {
					t.Log(fmt.Printf("Resultados diferentes DB:[%s] \n Arq:[%s]\n"))
					t.Fail()
				}
			}
	*/
}
