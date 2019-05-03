package main

import (
	"encoding/xml"
)

type TestSuite struct {
	XMLName   xml.Name   `xml:"testsuite"`
	TestCases []TestCase `xml:"testcase"`
}

type TestCase struct {
	XMLName     xml.Name `xml:"testcase"`
	ClassName   string   `xml:"classname,attr"`
	Name        string   `xml:"name,attr"`
	TestFailure string   `xml:"failure,omitempty"`
}

type ResultCT struct {
	NomeCT    string
	Resultado string
	Falha     string
}

type Casosteste struct {
	CasoTeste Casoteste `json:"casoteste"`
}

type Casoteste struct {
	Nome          string      `json:"nome"`
	Grupo         string      `json:"grupo"`
	Descricao     string      `json:"descricao"`
	Rotina        string      `json:"rotina"`
	Retorno       string      `json:"retorno"`
	TempoExecucao float64     `json:"tempoexecucao"`
	Entrada       interface{} `json:"entrada"`
}
