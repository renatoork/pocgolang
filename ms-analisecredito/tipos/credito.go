package tipos

import (
	"strings"
	"time"
)

const Versao = "v0"

type CustomFloat string

func (c *CustomFloat) UnmarshalJSON(b []byte) error {
	v := string(b)
	if v != "" {
		parse := strings.Replace(v, ",", ".", -1)
		*c = CustomFloat(parse)
	}
	return nil
}

type CustomTime time.Time

func (t *CustomTime) UnmarshalJSON(b []byte) error {
	timeStr := string(b)
	if timeStr != "" {
		ts, err := time.Parse("02/01/2006", (strings.Trim(timeStr, `"`)))
		if err == nil {
			*t = CustomTime(ts)
		}
	}

	return nil
}

func (t CustomTime) MarshalJSON() ([]byte, error) {
	tt := time.Time(t)
	parse, err := time.Parse("02/01/2006", "01/01/0001")
	if err == nil && !tt.Equal(parse) {
		return []byte(`"` + tt.Format("02/01/2006") + `"`), nil
	} else {
		return []byte(`"01/01/2015"`), nil
	}
}

type Documento struct {
	Documento Doc `json:"pedidovenda"`
}

type Doc struct {
	Organizacao            string    `json:"organizacao" db:"org_in_codigo" xml:"ORG_IN_CODIGO"`
	Filial                 string    `json:"filial" db:"fil_in_codigo" xml:"FIL_IN_CODIGO"`
	Acao                   string    `json:"acao" db:"acao_in_codigo" xml:"ACAO_IN_CODIGO"`
	Serie                  string    `json:"serie" db:"ser_st_codigo" xml:"SER_ST_CODIGO"`
	Numero                 string    `json:"numero" db:"ped_in_codigo" xml:"PED_IN_CODIGO"`
	Cliente                string    `json:"cliente" db:"cli_in_codigo" xml:"CLI_IN_CODIGO"`
	DataEmissao            time.Time `json:"emissao" db:"ped_dt_emissao"`
	CondPagto              string    `json:"condpagto" db:"cond_st_codigo" xml:"COND_ST_CODIGO"`
	Representante          string    `json:"representante" db:"rep_in_codigo" xml:"REP_IN_CODIGO"`
	Equipe                 string    `json:"equipe" db:"equ_in_codigo" xml:"EQU_IN_CODIGO"`
	Transportadora         string    `json:"transportadora" db:"tra_in_codigo" xml:"TRA_IN_CODIGO"`
	Redespacho             string    `json:"redespacho" db:"red_in_codigo" xml:"RED_IN_CODIGO"`
	TipoPedido             string    `json:"tipopedido" db:"ped_ch_tipopedido" xml:"PED_CH_TIPOPEDIDO"`
	Contato                string    `json:"contato" db:"ped_st_contato" xml:"PED_ST_CONTATO"`
	ContatoFone            string    `json:"contatofone" db:"ped_st_contatofone" xml:"PED_ST_CONTATOFONE"`
	ContatoCargo           string    `json:"contatocargo" db:"car_in_codigo" xml:"CAR_IN_CODIGO"`
	IndiceMoeda            string    `json:"indicemoeda" db:"ind_in_codigo" xml:"IND_IN_CODIGO"`
	ValorMoeda             string    `json:"valormoeda" db:"ped_re_cotacaomoe" xml:"PED_RE_COTACAOMOE"`
	PedidoEspecial         string    `json:"pedidoespecial" db:"ped_bo_especial" xml:"PED_BO_ESPECIAL"`
	TipoDocumento          string    `json:"tipodocumento" db:"tpd_in_codigo" xml:"TPD_IN_CODIGO"`
	FretePorConta          string    `json:"freteporconta" db:"ped_in_fretepconta" xml:"PED_IN_FRETEPCONTA"`
	FreteEmbutidoDestacado string    `json:"freteembutidodestacado" db:"ped_st_freteembdest" xml:"PED_ST_FRETEEMBDEST"`
	Usuario                string    `json:"usuario" db:"usu_in_codigo" xml:"USU_IN_CODIGO"`
	IndicaPresenca         string    `json:"indicapresenca" db:"ped_in_indpres" xml:"PED_IN_INDPRES"`
	IndicaConsFinal        string    `json:"indicaconsfinal" db:"ped_bo_indfinal" xml:"PED_BO_INDFINAL"`
	ValorMercadorias       float64   `json:"valormercadoria" db:"ped_re_valormercadoria"`
	ValorPedido            float64   `json:"valorpedido" db:"ped_re_valortotal"`
	ValorIPI               float64   `json:"valoripi" db:"ped_re_valoripi"`
	ValorICMS              float64   `json:"valoricms" db:"ped_re_valoricms"`
	//ItemDocumento []Item      `json:"itens" db:"-" xml:"itemdocumento>item"`
}

type Item struct {
	Sequencia                string  `json:"-" db:"itp_in_sequencia" xml:"ITP_IN_SEQUENCIA"`
	Produto                  string  `json:"produto" db:"PRO_IN_CODIGO" xml:"PRO_IN_CODIGO"`
	ProdutoAlternativo       string  `json:"produtoalternativo" db:"PRO_ST_ALTERNATIVO" xml:"PRO_ST_ALTERNATIVO"`
	Unidade                  string  `json:"unidade" db:"UNI_ST_UNIDADE" xml:"UNI_ST_UNIDADE"`
	Descricao                string  `json:"descricao" db:"ITP_ST_DESCRICAO" xml:"ITP_ST_DESCRICAO"`
	Complemento              string  `json:"complemento" db:"ITP_ST_COMPLEMENTO" xml:"ITP_ST_COMPLEMENTO"`
	NCM                      string  `json:"ncm" db:"NCM_IN_CODIGO" xml:"NCM_IN_CODIGO"`
	Servico                  string  `json:"servico" db:"COS_IN_CODIGO" xml:"COS_IN_CODIGO"`
	Aplicacao                string  `json:"aplicacao" db:"APL_IN_CODIGO" xml:"APL_IN_CODIGO"`
	TabelaPreco              string  `json:"tabelapreco" db:"TPR_IN_CODIGO" xml:"TPR_IN_CODIGO"`
	TipoPreco                string  `json:"tipopreco" db:"TPP_IN_CODIGO" xml:"TPP_IN_CODIGO"`
	IdentificadorProjeto     string  `json:"identificadorprojeto" db:"PROJ_IDE_ST_CODIGO" xml:"PROJ_IDE_ST_CODIGO"`
	Projeto                  string  `json:"projeto" db:"PROJ_IN_REDUZIDO" xml:"PROJ_IN_REDUZIDO"`
	IdentificadorCentroCusto string  `json:"identificadorcentrocusto" db:"CUS_IDE_ST_CODIGO" xml:"CUS_IDE_ST_CODIGO"`
	CentroCusto              string  `json:"centrocusto" db:"CUS_IN_REDUZIDO" xml:"CUS_IN_REDUZIDO"`
	Qtde                     float64 `json:"quantidade" db:"ITP_RE_QUANTIDADE" xml:"ITP_RE_QUANTIDADE"`
	Quantidade               string  `json:"-" db:"ITP_RE_QUANTIDADE"`
	ValorUnitario            string  `json:"valorunitario" db:"ITP_RE_VALORUNITARIO" xml:"ITP_RE_VALORUNITARIO"`
	ValorMercadoria          string  `json:"valormercadoria" db:"ITP_RE_VALORMERCADORIA" xml:"ITP_RE_VALORMERCADORIA"`
	ValorMercadoriaEmpregada string  `json:"valormercadoriaEmpregada" db:"ITP_RE_VALORMERCEMPREG" xml:"ITP_RE_VALORMERCEMPREG"`
	ValorMaoObra             string  `json:"valormaoobra" db:"ITP_RE_VALORMAOOBRA" xml:"ITP_RE_VALORMAOOBRA"`
	Frete                    string  `json:"frete" db:"ITP_RE_FRETE" xml:"ITP_RE_FRETE"`
	Seguro                   string  `json:"seguro" db:"ITP_RE_SEGURO" xml:"ITP_RE_SEGURO"`
	DespesaAcessoria         string  `json:"despesaacessoria" db:"ITP_RE_DESPACESSORIA" xml:"ITP_RE_DESPACESSORIA"`
	CodigoPedidoCliente      string  `json:"codigopedidocliente" db:"ITP_ST_PEDIDOCLIENTE" xml:"ITP_ST_PEDIDOCLIENTE"`
	CodigoProdutoCliente     string  `json:"codigoprodutocliente" db:"ITP_ST_CODPROCLI" xml:"ITP_ST_CODPROCLI"`
	PercentualDesconto       string  `json:"percentualdesconto" db:"ITP_RE_PERCDESCONTO" xml:"ITP_RE_PERCDESCONTO"`
	ValorDesconto            string  `json:"valordesconto" db:"ITP_RE_VALORDESCONTO" xml:"ITP_RE_VALORDESCONTO"`
	PercentualAcrescimo      string  `json:"percentualacrescimo" db:"ITP_RE_PERCACRESCIMO" xml:"ITP_RE_PERCACRESCIMO"`
	ValorAcrescimo           string  `json:"valoracrescimo" db:"ITP_RE_VALORACRESCIMO" xml:"ITP_RE_VALORACRESCIMO"`
	TotalizaItem             string  `json:"totalizaitem" db:"ITP_BO_TOTALIZA" xml:"ITP_BO_TOTALIZA"`
	ValorImportacao          string  `json:"valorimportacao" db:"ITP_RE_VALORIMPORTACAO" xml:"ITP_RE_VALORIMPORTACAO"`
	ValorCaucao              string  `json:"valorcaucao" db:"ITP_RE_VALORCAUCAO" xml:"ITP_RE_VALORCAUCAO"`
	Composicao               string  `json:"composicao" db:"CPS_IN_CODIGO" xml:"CPS_IN_CODIGO"`
	TipoClasse               string  `json:"tipoclasse" db:"TPC_ST_CLASSE" xml:"TPC_ST_CLASSE"`
	ItemEspecial             string  `json:"itemespecial" db:"ITP_BO_ESPECIAL" xml:"ITP_BO_ESPECIAL"`
	FormatoConversao         string  `json:"formatoconversao" db:"FMT_ST_CODIGO" xml:"FMT_ST_CODIGO"`
	Embalagem                string  `json:"embalagem" db:"EMB_IN_CODIGO" xml:"EMB_IN_CODIGO"`
	UnidadeEmbalagem         string  `json:"unidadeembalagem" db:"EMB_UNI_ST_UNIDADE" xml:"EMB_UNI_ST_UNIDADE"`
	FormatoEmbalagem         string  `json:"formatoembalagem" db:"EMB_FMT_ST_CODIGO" xml:"EMB_FMT_ST_CODIGO"`
	ValorMercadorias         float64 `json:"valormercadoria" db:"itp_re_valormercadoria"`
	ValorPedido              float64 `json:"valorpedido" db:"itp_re_valortotal"`
	ValorIPI                 float64 `json:"valoripi" db:"itp_re_valoripi"`
	ValorICMS                float64 `json:"valoricms" db:"itp_re_valoricms"`
}
