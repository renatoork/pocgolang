package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"
	_ "time"
)

type Scfe struct {
	XMLName xml.Name `xml:"CFe"`
	InfCFe  SinfCFe  `xml:"infCFe"`
}

type SinfCFe struct {
	XMLName  xml.Name `xml:"infCFe"`
	Versao   string   `xml:"versao,attr"`
	VersaoSB string   `xml:"versaoSB,attr"`
	Ide      Side     `xml:"ide"`
	Emit     Semit    `xml:"emit"`
	Det      []Sdet   `xml:"det"`
	Total    Stotal   `xml:"total"`
	Pgto     Spgto    `xml:"pgto"`
	InfAdic  SinfAdic `xml:"infAdic"`
}

type Side struct {
	XMLName     xml.Name `xml:"ide"`
	CUF         int      `xml:"cUF"`
	CNF         int      `xml:"cNF"`
	Mod         string   `xml:"mod"`
	NserieSAT   int      `xml:"nserieSAT"`
	NCFe        string   `xml:"nCFE"`
	DEmi        string   `xml:"dEmi"`
	HEmi        string   `xml:"hEmi"`
	CDV         int      `xml:"cDV"`
	TpAmb       int      `xml:"tpAmb"`
	CNPJ        int      `xml:"CNPJ"`
	SignAC      string   `xml:"signAC"`
	NumeroCaixa int      `xml:"numeroCaixa"`
}

type Semit struct {
	XMLName       xml.Name `xml:"emit"`
	Cnpj          string   `xml:"CNPJ"`
	XNome         string   `xml:"xNome"`
	XFant         string   `xml:"xFant"`
	XLgr          string   `xml:"enderEmit>xLgr"`
	Nro           int      `xml:"enderEmit>nro"`
	XCpl          int      `xml:"enderEmit>xCpl"`
	XBairro       string   `xml:"enderEmit>xBairro"`
	XMun          string   `xml:"enderEmit>xMun"`
	Cep           string   `xml:"enderEmit>CEP"`
	Ie            int      `xml:"IE"`
	Im            int      `xml:"IM"`
	CRegTrib      int      `xml:"cRegTrib"`
	CRegTribISSQN int      `xml:"cRegTribISSQN"`
	IndRatISSQN   string   `xml:"indRatISSQN"`
}

type Sdet struct {
	XMLName       xml.Name `xml:"det"`
	NItem         string   `xml:"nItem,attr"`
	Cprod         int      `xml:"cProd"`
	XProd         string   `xml:"xProd"`
	CFOP          int      `xml:"CFPO"`
	UCom          string   `xml:"uCom"`
	QCom          float64  `xml:"qCom"`
	VUnCom        float64  `xml:"vUnCom"`
	VProd         float64  `xml:"vProd"`
	IndRegra      string   `xml:"indRegra"`
	VItem         float64  `xml:"vItem"`
	ICMSOrig      int      `xml:"imposto>ICMS>ICMS00>Orig"`
	ICMSCST       string   `xml:"imposto>ICMS>ICMS00>CST"`
	ICMSpICMS     float64  `xml:"imposto>ICMS>ICMS00>pICMS"`
	ICMSvICMS     float64  `xml:"imposto>ICMS>ICMS00>vICMS"`
	PISCST        string   `xml:"imposto>PIS>PISAliq>CST"`
	PISvBC        float64  `xml:"imposto>PIS>PISAliq>vBC"`
	PISpPIS       float64  `xml:"imposto>PIS>PISAliq>pPIS"`
	PISvPIS       float64  `xml:"imposto>PIS>PISAliq>vPIS"`
	COFINSCST     string   `xml:"imposto>COFINS>COFINSAliq>CST"`
	COFINSvBC     float64  `xml:"imposto>COFINS>COFINSAliq>vBC"`
	COFINSpCOFINS float64  `xml:"imposto>COFINS>COFINSAliq>pCOFINS"`
	COFINSvCOFINS float64  `xml:"imposto>COFINS>COFINSAliq>vCOFINS"`
}

type Stotal struct {
	XMLName      xml.Name `xml:"total"`
	VICMS        float64  `xml:"ICMSTot>vICMS"`
	VProd        float64  `xml:"ICMSTot>vProd"`
	VDesc        float64  `xml:"ICMSTot>vDesc"`
	VPIS         float64  `xml:"ICMSTot>vPIS"`
	VCOFINS      float64  `xml:"ICMSTot>vCOFINS"`
	VPISST       float64  `xml:"ICMSTot>vPISST"`
	VCOFINSST    float64  `xml:"ICMSTot>vCOFINSST"`
	VOutro       float64  `xml:"ICMSTot>vOutro"`
	VCFe         float64  `xml:"vCFe"`
	VCFeLei12741 float64  `xml:"vCFeLei12741"`
}

type Spgto struct {
	XMLName xml.Name `xml:"pgto"`
	CMP     string   `xml:"MP>cMP"`
	VMP     float64  `xml:"MP>vMP"`
	VTroco  float64  `xml:"vTroco"`
}

type SinfAdic struct {
	XMLName  xml.Name `xml:"infAdic"`
	InfCpl   string   `xml:"infCpl"`
	obsFisco string   `xml:"obsFisco>xTexto"`
}

type Cupom struct {
	SATC_IN_SEQUENCIA      int     `json:"-" db:"SATC_IN_SEQUENCIA" xml:"SATC_IN_SEQUENCIA" description:"Sequencial"`
	ORG_TAB_IN_CODIGO      int     `json:"-" db:"ORG_TAB_IN_CODIGO" xml:"ORG_TAB_IN_CODIGO" description:"Tabela Organização"`
	ORG_PAD_IN_CODIGO      int     `json:"-" db:"ORG_PAD_IN_CODIGO" xml:"ORG_PAD_IN_CODIGO" description:"Padrão Organização"`
	ORG_IN_CODIGO          int     `json:"-" db:"ORG_IN_CODIGO" xml:"ORG_IN_CODIGO" description:"Código Organização"`
	ORG_TAU_ST_CODIGO      string  `json:"-" db:"ORG_TAU_ST_CODIGO" xml:"ORG_TAU_ST_CODIGO" description:"Código Aux. Organização"`
	FIL_IN_CODIGO          int     `json:"-" db:"FIL_IN_CODIGO" xml:"FIL_IN_CODIGO" description:"Código Filial"`
	ACAO_TAB_IN_CODIGO     int     `json:"-" db:"ACAO_TAB_IN_CODIGO" xml:"ACAO_TAB_IN_CODIGO" description:"Tabela Ação"`
	ACAO_PAD_IN_CODIGO     int     `json:"-" db:"ACAO_PAD_IN_CODIGO" xml:"ACAO_PAD_IN_CODIGO" description:"Padrão Ação"`
	ACAO_IN_CODIGO         int     `json:"-" db:"ACAO_IN_CODIGO" xml:"ACAO_IN_CODIGO" description:"Código Ação"`
	AGN_TAB_IN_CODIGO      int     `json:"-" db:"AGN_TAB_IN_CODIGO" xml:"AGN_TAB_IN_CODIGO" description:"Tabela Agente"`
	AGN_PAD_IN_CODIGO      int     `json:"-" db:"AGN_PAD_IN_CODIGO" xml:"AGN_PAD_IN_CODIGO" description:"Padrão Agente"`
	AGN_IN_CODIGO          int     `json:"-" db:"AGN_IN_CODIGO" xml:"AGN_IN_CODIGO" description:"Código Agente"`
	AGN_TAU_ST_CODIGO      string  `json:"-" db:"AGN_TAU_ST_CODIGO" xml:"AGN_TAU_ST_CODIGO" description:"Código Aux. Agente"`
	AGN_ST_CPFCNPJ         string  `json:"-" db:"AGN_ST_CPFCNPJ" xml:"AGN_ST_CPFCNPJ" description:"CNPJ CPF"`
	AGN_ST_NOMECONSUMIDOR  string  `json:"-" db:"AGN_ST_NOMECONSUMIDOR" xml:"AGN_ST_NOMECONSUMIDOR" description:"Nome"`
	TDF_IN_CODIGO          int     `json:"-" db:"TDF_IN_CODIGO" xml:"TDF_IN_CODIGO" description:"Tipo de Documento Fiscal"`
	SATC_IN_VERSAOSB       string  `json:"-" db:"SATC_IN_VERSAOSB" xml:"SATC_IN_VERSAOSB" description:"Versão SAT"`
	SATC_IN_VERSAO         string  `json:"-" db:"SATC_IN_VERSAO" xml:"SATC_IN_VERSAO" description:"Versão CF-e"`
	SATC_ST_CFE            string  `json:"-" db:"SATC_ST_CFE" xml:"SATC_ST_CFE" description:"Número CF-e"`
	SATC_IN_NUMEROCUPOM    int     `json:"-" db:"SATC_IN_NUMEROCUPOM" xml:"SATC_IN_NUMEROCUPOM" description:"Número Cupom"`
	SATC_IN_NUMEROCAIXA    int     `json:"-" db:"SATC_IN_NUMEROCAIXA" xml:"SATC_IN_NUMEROCAIXA" description:"Número Caixa"`
	SATC_ST_MODELOFISCAL   string  `json:"-" db:"SATC_ST_MODELOFISCAL" xml:"SATC_ST_MODELOFISCAL" description:"Modelo Cupom Fiscal"`
	SATC_IN_NSERIESAT      int     `json:"-" db:"SATC_IN_NSERIESAT" xml:"SATC_IN_NSERIESAT" description:"Número SAT"`
	SATC_DT_EMISSAO        string  `json:"-" db:"SATC_DT_EMISSAO" xml:"SATC_DT_EMISSAO" description:"Data emissão"`
	SATC_DT_HORAEMISSAO    string  `json:"-" db:"SATC_DT_HORAEMISSAO" xml:"SATC_DT_HORAEMISSAO" description:"Hora emissão"`
	SATC_IN_CODSITUACAO    int     `json:"-" db:"SATC_IN_CODSITUACAO" xml:"SATC_IN_CODSITUACAO" description:"Código da situação do Cupom Fiscal"`
	SATC_RE_VALORTOTAL     float64 `json:"-" db:"SATC_RE_VALORTOTAL" xml:"SATC_RE_VALORTOTAL" description:"Valor Total"`
	SATC_RE_VALORDESCONTO  float64 `json:"-" db:"SATC_RE_VALORDESCONTO" xml:"SATC_RE_VALORDESCONTO" description:"Desconto"`
	SATC_RE_VALORACRESCIMO float64 `json:"-" db:"SATC_RE_VALORACRESCIMO" xml:"SATC_RE_VALORACRESCIMO" description:"Valor Acréscimo"`
	SATC_RE_VALORCFE       float64 `json:"-" db:"SATC_RE_VALORCFE" xml:"SATC_RE_VALORCFE" description:"Valor CFe"`
	SATC_RE_VALORICMS      float64 `json:"-" db:"SATC_RE_VALORICMS" xml:"SATC_RE_VALORICMS" description:"Valor ICMS"`
	SATC_RE_VALORPIS       float64 `json:"-" db:"SATC_RE_VALORPIS" xml:"SATC_RE_VALORPIS" description:"Valor PIS"`
	SATC_RE_VALORCOFINS    float64 `json:"-" db:"SATC_RE_VALORCOFINS" xml:"SATC_RE_VALORCOFINS" description:"Valor COFINS"`
	SATC_RE_VALORPISST     float64 `json:"-" db:"SATC_RE_VALORPISST" xml:"SATC_RE_VALORPISST" description:"Valor PISST"`
	SATC_RE_VALORCOFINSST  float64 `json:"-" db:"SATC_RE_VALORCOFINSST" xml:"SATC_RE_VALORCOFINSST" description:"Valor COFINSST"`
	SATC_RE_BASECALCISS    float64 `json:"-" db:"SATC_RE_BASECALCISS" xml:"SATC_RE_BASECALCISS" description:"Valor base calculo ISS"`
	SATC_RE_VALORISS       float64 `json:"-" db:"SATC_RE_VALORISS" xml:"SATC_RE_VALORISS" description:"Valor ISS"`
	SATC_RE_ICMSOUTRAS     float64 `json:"-" db:"SATC_RE_ICMSOUTRAS" xml:"SATC_RE_ICMSOUTRAS" description:"ICMS Outras"`
	SATC_RE_ICMSISENTAS    float64 `json:"-" db:"SATC_RE_ICMSISENTAS" xml:"SATC_RE_ICMSISENTAS" description:"ICMS Isentas"`
	Item                   []Item
}

type Item struct {
	SATC_IN_SEQUENCIA      int     `json:"-" db:"SATC_IN_SEQUENCIA" xml:"SATC_IN_SEQUENCIA" description:"Lançamento"`
	SATI_IN_NUMEROITENS    int     `json:"-" db:"SATI_IN_NUMEROITENS" xml:"SATI_IN_NUMEROITENS" description:"Nº Item"`
	APL_TAB_IN_CODIGO      int     `json:"-" db:"APL_TAB_IN_CODIGO" xml:"APL_TAB_IN_CODIGO" description:"Tabela Aplicação"`
	APL_PAD_IN_CODIGO      int     `json:"-" db:"APL_PAD_IN_CODIGO" xml:"APL_PAD_IN_CODIGO" description:"Padrão Aplicação"`
	APL_IN_CODIGO          int     `json:"-" db:"APL_IN_CODIGO" xml:"APL_IN_CODIGO" description:"Aplicação"`
	PRO_TAB_IN_CODIGO      int     `json:"-" db:"PRO_TAB_IN_CODIGO" xml:"PRO_TAB_IN_CODIGO" description:"Tabela Itens"`
	PRO_PAD_IN_CODIGO      int     `json:"-" db:"PRO_PAD_IN_CODIGO" xml:"PRO_PAD_IN_CODIGO" description:"Padrão Itens"`
	PRO_IN_CODIGO          int     `json:"-" db:"PRO_IN_CODIGO" xml:"PRO_IN_CODIGO" description:"Cód. Item"`
	UNI_TAB_IN_CODIGO      int     `json:"-" db:"UNI_TAB_IN_CODIGO" xml:"UNI_TAB_IN_CODIGO" description:"Tabela Unidade"`
	UNI_PAD_IN_CODIGO      int     `json:"-" db:"UNI_PAD_IN_CODIGO" xml:"UNI_PAD_IN_CODIGO" description:"Padrão Unidade"`
	UNI_ST_UNIDADE         string  `json:"-" db:"UNI_ST_UNIDADE" xml:"UNI_ST_UNIDADE" description:"Unidade"`
	CFOP_TAB_IN_CODIGO     int     `json:"-" db:"CFOP_TAB_IN_CODIGO" xml:"CFOP_TAB_IN_CODIGO" description:"Tabela CFOP"`
	CFOP_PAD_IN_CODIGO     int     `json:"-" db:"CFOP_PAD_IN_CODIGO" xml:"CFOP_PAD_IN_CODIGO" description:"Padrão CFOP"`
	CFOP_IDE_ST_CODIGO     string  `json:"-" db:"CFOP_IDE_ST_CODIGO" xml:"CFOP_IDE_ST_CODIGO" description:"Identificador CFOP"`
	CFOP_IN_CODIGO         int     `json:"-" db:"CFOP_IN_CODIGO" xml:"CFOP_IN_CODIGO" description:"CFOP"`
	NCM_TAB_IN_CODIGO      int     `json:"-" db:"NCM_TAB_IN_CODIGO" xml:"NCM_TAB_IN_CODIGO" description:"Tabela NCM"`
	NCM_PAD_IN_CODIGO      int     `json:"-" db:"NCM_PAD_IN_CODIGO" xml:"NCM_PAD_IN_CODIGO" description:"Padrão NCM"`
	NCM_IN_CODIGO          int     `json:"-" db:"NCM_IN_CODIGO" xml:"NCM_IN_CODIGO" description:"Cód. NCM"`
	COS_IN_CODIGO          int     `json:"-" db:"COS_IN_CODIGO" xml:"COS_IN_CODIGO" description:"Cód. Serviço Prestado"`
	COSM_IN_CODIGO         int     `json:"-" db:"COSM_IN_CODIGO" xml:"COSM_IN_CODIGO" description:"Cód. Serviço Prestado Município"`
	SATI_RE_QUANTIDADE     float64 `json:"-" db:"SATI_RE_QUANTIDADE" xml:"SATI_RE_QUANTIDADE" description:"Quantidade"`
	SATI_RE_VALORUNITARIO  float64 `json:"-" db:"SATI_RE_VALORUNITARIO" xml:"SATI_RE_VALORUNITARIO" description:"Preço Unitário"`
	SATI_RE_VALORTOTAL     float64 `json:"-" db:"SATI_RE_VALORTOTAL" xml:"SATI_RE_VALORTOTAL" description:"Valor Total"`
	SATI_RE_VALORDESCONTO  float64 `json:"-" db:"SATI_RE_VALORDESCONTO" xml:"SATI_RE_VALORDESCONTO" description:"Valor Desconto"`
	SATI_RE_VALORACRESCIMO float64 `json:"-" db:"SATI_RE_VALORACRESCIMO" xml:"SATI_RE_VALORACRESCIMO" description:"Valor Acréscimo"`
	SATI_RE_VALORLIQUIDO   float64 `json:"-" db:"SATI_RE_VALORLIQUIDO" xml:"SATI_RE_VALORLIQUIDO" description:"Valor Líquido Item"`
	SATI_ST_CSTICMS        string  `json:"-" db:"SATI_ST_CSTICMS" xml:"SATI_ST_CSTICMS" description:"CST ICMS"`
	SATI_ST_CSOSN          string  `json:"-" db:"SATI_ST_CSOSN" xml:"SATI_ST_CSOSN" description:"CST CSOSN"`
	SATI_RE_BASEICMS       float64 `json:"-" db:"SATI_RE_BASEICMS" xml:"SATI_RE_BASEICMS" description:"Base ICMS"`
	SATI_RE_ALIQICMS       float64 `json:"-" db:"SATI_RE_ALIQICMS" xml:"SATI_RE_ALIQICMS" description:"Alíq. ICMS"`
	SATI_RE_VALORICMS      float64 `json:"-" db:"SATI_RE_VALORICMS" xml:"SATI_RE_VALORICMS" description:"Valor ICMS"`
	SATI_RE_BASEISS        float64 `json:"-" db:"SATI_RE_BASEISS" xml:"SATI_RE_BASEISS" description:"Base ISS"`
	SATI_RE_ALIQISS        float64 `json:"-" db:"SATI_RE_ALIQISS" xml:"SATI_RE_ALIQISS" description:"Alíq. ISS"`
	SATI_RE_VALORISS       float64 `json:"-" db:"SATI_RE_VALORISS" xml:"SATI_RE_VALORISS" description:"Valor ISS"`
	STP_ST_CSTPIS          string  `json:"-" db:"STP_ST_CSTPIS" xml:"STP_ST_CSTPIS" description:"CST PIS"`
	SATI_RE_BASEPIS        float64 `json:"-" db:"SATI_RE_BASEPIS" xml:"SATI_RE_BASEPIS" description:"Base PIS"`
	SATI_RE_ALIQPIS        float64 `json:"-" db:"SATI_RE_ALIQPIS" xml:"SATI_RE_ALIQPIS" description:"Alíq. PIS"`
	SATI_RE_PAUTAPIS       float64 `json:"-" db:"SATI_RE_PAUTAPIS" xml:"SATI_RE_PAUTAPIS" description:"Pauta PIS"`
	SATI_RE_VALORPIS       float64 `json:"-" db:"SATI_RE_VALORPIS" xml:"SATI_RE_VALORPIS" description:"Valor PIS"`
	SATI_RE_VALORPISST     float64 `json:"-" db:"SATI_RE_VALORPISST" xml:"SATI_RE_VALORPISST" description:"Valor PISST"`
	STC_ST_CSTCOFINS       string  `json:"-" db:"STC_ST_CSTCOFINS" xml:"STC_ST_CSTCOFINS" description:"CST COFINS"`
	SATI_RE_BASECOFINS     float64 `json:"-" db:"SATI_RE_BASECOFINS" xml:"SATI_RE_BASECOFINS" description:"Base COFINS"`
	SATI_RE_ALIQCOFINS     float64 `json:"-" db:"SATI_RE_ALIQCOFINS" xml:"SATI_RE_ALIQCOFINS" description:"Alíq. COFINS"`
	SATI_RE_PAUTACOFINS    float64 `json:"-" db:"SATI_RE_PAUTACOFINS" xml:"SATI_RE_PAUTACOFINS" description:"Pauta COFINS"`
	SATI_RE_VALORCOFINS    float64 `json:"-" db:"SATI_RE_VALORCOFINS" xml:"SATI_RE_VALORCOFINS" description:"Valor COFINS"`
	SATI_RE_VALORCOFINSST  float64 `json:"-" db:"SATI_RE_VALORCOFINSST" xml:"SATI_RE_VALORCOFINSST" description:"Valor COFINSST"`
	SATI_IN_MUNCODIGOFG    int     `json:"-" db:"SATI_IN_MUNCODIGOFG" xml:"SATI_IN_MUNCODIGOFG" description:"Cód. Município Ocorrência Fato Gerador ISS"`
	SATI_ST_ITEMLISTSERV   string  `json:"-" db:"SATI_ST_ITEMLISTSERV" xml:"SATI_ST_ITEMLISTSERV" description:"Item Lista Serviços"`
	SATI_IN_CODTRIBMUN     int     `json:"-" db:"SATI_IN_CODTRIBMUN" xml:"SATI_IN_CODTRIBMUN" description:"Cód. Serviço Prestado Município ISS"`
	SATI_ST_NATOPERACAO    string  `json:"-" db:"SATI_ST_NATOPERACAO" xml:"SATI_ST_NATOPERACAO" description:"Natureza Operação ISS"`
	SATI_IN_SEQUENCIA      int     `json:"-" db:"SATI_IN_SEQUENCIA" xml:"SATI_IN_SEQUENCIA" description:"Sequencial Único do Item"`
	SATI_RE_ICMSOUTRAS     float64 `json:"-" db:"SATI_RE_ICMSOUTRAS" xml:"SATI_RE_ICMSOUTRAS" description:"ICMS Outras"`
	SATI_RE_ICMSISENTAS    float64 `json:"-" db:"SATI_RE_ICMSISENTAS" xml:"SATI_RE_ICMSISENTAS" description:"ICMS Isentas"`
}

func consomeSat(xml *Scfe) {
	var cupom Cupom
	var item Item

	cupom.SATC_IN_SEQUENCIA = 1
	cupom.ORG_IN_CODIGO = 2
	cupom.FIL_IN_CODIGO = 3
	cupom.ACAO_IN_CODIGO = 115
	cupom.AGN_IN_CODIGO = 254
	cupom.AGN_TAU_ST_CODIGO = "C"
	cupom.AGN_ST_CPFCNPJ = xml.InfCFe.Emit.Cnpj
	cupom.AGN_ST_NOMECONSUMIDOR = xml.InfCFe.Emit.XFant
	cupom.TDF_IN_CODIGO = 1
	cupom.SATC_IN_VERSAOSB = xml.InfCFe.VersaoSB
	cupom.SATC_IN_VERSAO = xml.InfCFe.Versao
	cupom.SATC_ST_CFE = xml.InfCFe.Ide.NCFe
	cupom.SATC_IN_NUMEROCUPOM = xml.InfCFe.Ide.CNF
	cupom.SATC_IN_NUMEROCAIXA = xml.InfCFe.Ide.NumeroCaixa
	cupom.SATC_ST_MODELOFISCAL = xml.InfCFe.Ide.Mod
	cupom.SATC_IN_NSERIESAT = xml.InfCFe.Ide.NserieSAT
	cupom.SATC_DT_EMISSAO = xml.InfCFe.Ide.DEmi
	cupom.SATC_DT_HORAEMISSAO = xml.InfCFe.Ide.HEmi
	cupom.SATC_IN_CODSITUACAO = xml.InfCFe.Ide.TpAmb
	cupom.SATC_RE_VALORTOTAL = xml.InfCFe.Total.VProd
	cupom.SATC_RE_VALORDESCONTO = xml.InfCFe.Total.VDesc
	cupom.SATC_RE_VALORACRESCIMO = 0
	cupom.SATC_RE_VALORCFE = xml.InfCFe.Total.VCFe
	cupom.SATC_RE_VALORICMS = xml.InfCFe.Total.VICMS
	cupom.SATC_RE_VALORPIS = xml.InfCFe.Total.VPIS
	cupom.SATC_RE_VALORCOFINS = xml.InfCFe.Total.VCOFINS
	cupom.SATC_RE_VALORPISST = xml.InfCFe.Total.VPISST
	cupom.SATC_RE_VALORCOFINSST = xml.InfCFe.Total.VCOFINSST
	cupom.SATC_RE_BASECALCISS = 0
	cupom.SATC_RE_VALORISS = 0
	cupom.SATC_RE_ICMSOUTRAS = xml.InfCFe.Total.VOutro
	cupom.SATC_RE_ICMSISENTAS = 0

	for i, v := range xml.InfCFe.Det {
		item.SATC_IN_SEQUENCIA = 1
		item.SATI_IN_NUMEROITENS = len(xml.InfCFe.Det)
		//item.APL_IN_CODIGO =
		item.PRO_IN_CODIGO = v.Cprod
		item.UNI_ST_UNIDADE = v.UCom
		//item.CFOP_IDE_ST_CODIGO = v
		item.CFOP_IN_CODIGO = v.CFOP
		//item.NCM_IN_CODIGO
		//item.COS_IN_CODIGO
		//item.COSM_IN_CODIGO
		item.SATI_RE_QUANTIDADE = v.QCom
		item.SATI_RE_VALORUNITARIO = v.VUnCom
		item.SATI_RE_VALORTOTAL = v.VItem
		//item.SATI_RE_VALORACRESCIMO
		//item.SATI_RE_VALORLIQUIDO
		item.SATI_ST_CSTICMS = v.ICMSCST
		//item.SATI_ST_CSOSN
		//item.SATI_RE_BASEICMS
		item.SATI_RE_ALIQICMS = v.ICMSpICMS
		item.SATI_RE_VALORICMS = v.ICMSvICMS
		//item.SATI_RE_BASEISS
		//item.SATI_RE_ALIQISS
		//item.SATI_RE_VALORISS
		item.STP_ST_CSTPIS = v.PISCST
		item.SATI_RE_BASEPIS = v.PISvBC
		item.SATI_RE_ALIQPIS = v.PISpPIS
		//item.SATI_RE_PAUTAPIS
		item.SATI_RE_VALORPIS = v.PISvPIS
		//item.SATI_RE_VALORPISST
		item.STC_ST_CSTCOFINS = v.COFINSCST
		item.SATI_RE_BASECOFINS = v.COFINSvBC
		item.SATI_RE_ALIQCOFINS = v.COFINSpCOFINS
		//item.SATI_RE_PAUTACOFINS
		item.SATI_RE_VALORCOFINS = v.COFINSvCOFINS
		//item.SATI_RE_VALORCOFINSST
		//item.SATI_IN_MUNCODIGOFG
		//item.SATI_ST_ITEMLISTSERV
		//item.SATI_IN_CODTRIBMUN
		//item.SATI_ST_NATOPERACAO
		item.SATI_IN_SEQUENCIA = i + 1
		//item.SATI_RE_ICMSOUTRAS
		//item.SATI_RE_ICMSISENTAS
		cupom.Item = append(cupom.Item, item)
	}
	fmt.Println(cupom)
}

func main() {

	aux := new(Scfe)

	readerDir, _ := ioutil.ReadDir(".\\")
	if len(readerDir) > 0 {
		for _, fileInfo := range readerDir {
			if !fileInfo.IsDir() && strings.Contains(fileInfo.Name(), ".xml") {
				arq, _ := ioutil.ReadFile(fileInfo.Name())
				errj := xml.Unmarshal(arq, aux)
				if errj != nil {
					fmt.Println("ERRO: ", errj.Error())
				} else {
					consomeSat(aux)
				}
			}
		}
	}
}
