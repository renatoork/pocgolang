package compara

import (
	"testing"
)

type Comp struct {
	Coluna1 int `comp:"S"`
	Coluna2 string
	Coluna3 int `comp:"N"`
	Estrut  Estrut
}

type Estrut struct {
	Coluna1 int
	Coluna2 string
}

func TestComparaIguais(t *testing.T) {
	str1 := Comp{1, "teste1", 0, Estrut{1, "compara 1"}}
	str2 := Comp{1, "teste1", 0, Estrut{1, "compara 1"}}
	ok, resultado := CompararStruct(&str1, &str2)
	if !ok {
		t.Error("Estruturas deveriam ser Iguais: \n", resultado)
	}

}

func TestComparaIguaisTag(t *testing.T) {
	str1 := Comp{1, "teste1", 0, Estrut{1, "compara 1"}}
	str2 := Comp{1, "teste1", 1, Estrut{1, "compara 1"}}
	ok, resultado := CompararStruct(&str1, &str2)
	if !ok {
		t.Error("Estruturas deveriam ser Iguais: \n", resultado)
	}
}

func TestComparaDiferenteInt(t *testing.T) {
	str1 := Comp{1, "teste1", 0, Estrut{1, "compara 1"}}
	str2 := Comp{2, "teste1", 0, Estrut{1, "compara 1"}}
	ok, resultado := CompararStruct(&str1, &str2)
	if ok {
		t.Error("Estruturas deveriam ser Diferentes: \n", resultado)
	}
}

func TestComparaDiferenteString(t *testing.T) {
	str1 := Comp{1, "teste1", 0, Estrut{1, "compara 1"}}
	str2 := Comp{1, "teste2", 0, Estrut{1, "compara 1"}}
	ok, resultado := CompararStruct(&str1, &str2)
	if ok {
		t.Error("Estruturas deveriam ser Diferentes: \n", resultado)
	}

}

func TestComparaDiferenteStruct(t *testing.T) {
	str1 := Comp{1, "teste1", 0, Estrut{1, "compara 1"}}
	str2 := Comp{1, "teste1", 0, Estrut{1, "compara 2"}}
	ok, resultado := CompararStruct(&str1, &str2)
	if ok {
		t.Error("Estruturas deveriam ser Diferentes: \n", resultado)
	}
}
