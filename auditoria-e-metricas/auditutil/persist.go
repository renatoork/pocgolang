package auditutil

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

type InvalidIdError struct {
	Id string
}

type IdNotFoundError struct {
	Id string
}

func (p InvalidIdError) Error() string {
	return fmt.Sprintf("chave inválida (%s). Favor informar uma chave válida no parâmetro id.", p.Id)
}

func (p IdNotFoundError) Error() string {
	return fmt.Sprintf("resultado com id () não encontrado", p.Id)
}

var (
	session *mgo.Session
)

func GetSession() *mgo.Database {
	if session == nil {
		var err error
		session, err = mgo.Dial("localhost")
		if err != nil {
			panic(err) // no, not really
		}
	}
	return session.Clone().DB("test")
}
