package auditutil

import (
	"gopkg.in/mgo.v2/bson"
)

type LinhaResultado struct {
	Linha    int
	Mensagem string
}

type Resultado struct {
	Id         bson.ObjectId `bson:"_id"`
	Resultados []LinhaResultado
}

func AuditaArquivo() bson.ObjectId {
	dbAudit := GetSession().C("audit")
	id := bson.NewObjectId()
	var resultados = []LinhaResultado{LinhaResultado{Linha: 10, Mensagem: "Variável fora do padrão"}, LinhaResultado{Linha: 255, Mensagem: "Parâmetro pTeste não usado"}}
	dbAudit.Insert(bson.M{"_id": id, "resultados": &resultados})
	return id
}

func RetornaResultadoAuditoria(id string) (Resultado, error) {
	dbAudit := GetSession().C("audit")
	if bson.IsObjectIdHex(id) {
		objectId := bson.ObjectIdHex(id)
		var resultado Resultado
		err := dbAudit.Find(bson.M{"_id": objectId}).One(&resultado)
		if err != nil {
			return Resultado{}, &IdNotFoundError{id}
		}
		return resultado, err
	} else {
		return Resultado{}, &InvalidIdError{id}
	}
}
