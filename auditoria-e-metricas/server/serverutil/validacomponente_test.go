package serverutil

import (
	"../../auditutil"
	"errors"
	"fmt"
	"strings"
	"testing"
)

type str [][]string

func TestSepararComponentesDoFonteSlice(t *testing.T) {
	arquivo, erro := auditutil.LerArquivoTeste("./teste/uFormSelOrgProjeto.pas")
	if erro != nil {
		t.Log(erro)
		t.Fail()
	}

	res := separarDeclaracaoComponenteSlice(arquivo)

	auditutil.Assert(t, len(res) == 0, "Arquivo vazio..")
	/*
		if len(res) <= 0 {
			t.Log(fmt.Printf("Arquivo vazio.."))
			t.Fail()
		}
	*/
}

func TestSepararComponentesDoFonte(t *testing.T) {

	arquivo, erro := auditutil.LerArquivoTeste("./teste/uFormSelOrgProjeto.pas")
	auditutil.Ok(t, erro)

	esperado := "class(TFormMegaManutencao)\r\n" +
		"    Gd_Usu: TMgDBGrid;\r\n" +
		"    Gb_Orgs: TmgGroupBox;\r\n" +
		"    Gd_OrgSel: TMgDBGrid;\r\n" +
		"    mgSplitter1: TmgSplitter;\r\n" +
		"    Rg_Org_Default: TmgRadioGroup;\r\n" +
		"    Ed_Org_Default: TmgMaskEdit;\r\n" +
		"    Tm_Bi: TmgTimer;\r\n" +
		"    Ed_Fil_In_Codigo1: TMgDBEdit;\r\n" +
		"    Ed_GRU_IN_CODIGO1: TMgDBEdit;\r\n" +
		"    procedure FormCreate(Sender: TObject);\r\n" +
		"    procedure FormClose(Sender: TObject; var Action: TCloseAction);\r\n" +
		"    procedure Rg_Org_DefaultClick(Sender: TObject);\r\n" +
		"    procedure FormShow(Sender: TObject);\r\n" +
		"    procedure Tm_BiTimer(Sender: TObject);\r\n" +
		"    procedure Ed_Fil_In_Codigo1Change(Sender: TObject);\r\n" +
		"    procedure Ed_GRU_IN_CODIGO1Change(Sender: TObject);\r\n" +
		"  private"

	res := separarDeclaracaoComponente(arquivo)

	auditutil.Equals(t, esperado, res)

	//auditutil.Assert(t, esperado == res, fmt.Printf("Arquivo: [%s]\n[Esperado:%s]\n", res, esperado))
	/*
		if esperado != res {
			t.Log(fmt.Printf("Arquivo: [%s]\n[Esperado:%s]\n", res, esperado))
			t.Fail()
		}
	*/
}

func SepararComponentePadrao(arq string) error {

	arquivo, err := auditutil.LerArquivoTeste(arq)
	auditutil.CheckError(err)

	esperado := str{
		[]string{"Gd_Usu: TMgDBGrid;", "Gd_Usu", "TMgDBGrid"},
		[]string{"Gb_Orgs: TmgGroupBox;", "Gb_Orgs", "TmgGroupBox"},
		[]string{"Gd_OrgSel: TMgDBGrid;", "Gd_OrgSel", "TMgDBGrid"},
		[]string{"mgSplitter1: TmgSplitter;", "mgSplitter1", "TmgSplitter"},
		[]string{"Rg_Org_Default: TmgRadioGroup;", "Rg_Org_Default", "TmgRadioGroup"},
		[]string{"Ed_Org_Default: TmgMaskEdit;", "Ed_Org_Default", "TmgMaskEdit;"},
		[]string{"Tm_Bi: TmgTimer;", "Tm_Bi", "TmgTimer;"},
		[]string{"Ed_Fil_In_Codigo1: TMgDBEdit;", "Ed_Fil_In_Codigo1", "TMgDBEdit;"},
		[]string{"Ed_GRU_IN_CODIGO1: TMgDBEdit;", "Ed_GRU_IN_CODIGO1", "TMgDBEdit;"},
	}

	res := separarComponente(separarDeclaracaoComponente(arquivo))

	if len(res) == 0 || res == nil {
		return errors.New("Tabela de componentes vazio!")
	} else {
		if len(res) != len(esperado) {
			return errors.New(fmt.Sprintf("Resultados Diferentes res[%d] <> esperado [%d].\n", len(res), len(esperado)))
		} else {
			for i, v := range res {
				if strings.Trim(v[0], " ") != esperado[i][0] {
					return errors.New(fmt.Sprintf("Arquivo: %d\nEsperado: %d\n", v[0], esperado[i][0]))
				}
			}
		}
	}
	return nil
}

func TestSepararComponente(t *testing.T) {
	erro := SepararComponentePadrao("./teste/uFormSelOrgProjeto.pas")
	if erro != nil {
		t.Log(erro)
		t.Fail()
	}
}

func TestSepararComponenteComentario(t *testing.T) {
	erro := SepararComponentePadrao("./teste/uFormSelOrgProjetoCom.pas")
	if erro != nil {
		t.Log(erro)
		t.Fail()
	}
}

func TestValidaNomenclatura(t *testing.T) {
	retorno := validaNomenclatura()
	if retorno != "Teste" {
		t.Log("Falou")
		t.Fail()
	}
}
