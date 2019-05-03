package serverutil

import (
	"../../auditutil"
	"gopkg.in/mgo.v2/bson"
)

type Componente struct {
	componente string
	prefixo    string
}

func RetornaTabelaComponentes() (bson.M, error) {
	dbComponente := auditutil.GetSession().C("componentes")

	var componentes bson.M
	err := dbComponente.Find(nil).All(&componentes)
	auditutil.CheckError(err)
	return componentes, nil
}
