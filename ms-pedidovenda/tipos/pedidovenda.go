package tipos

import (
	"mega/customizacaoservico"
	"strconv"
	"time"
)

const Versao = "v0"

type CustomTime *time.Time

type CustomFloat float64

func (f CustomFloat) MarshalJSON() ([]byte, error) {
	v := strconv.FormatFloat(float64(f), 'f', 1, 64)
	return []byte(v), nil
}

type Pedido struct {
	PedidoVenda PedidoVenda `json:"pedidovenda"`
}

type PedidoUk struct {
	Organizacao *int   `json:"organizacao,required" db:"ORG_IN_CODIGO" description:"Código da Organização"`
	Codigo      *int   `json:"numero" db:"PED_IN_CODIGO" description:"Código do Pedido. Informar 0 (zero) para inclusão."`
	Serie       string `json:"serie,required" db:"SER_ST_CODIGO" description:"Série da Numeração"`
}

type PedidoVenda struct {
	Operacao string `json:"-" db:"OPERACAO" xml:"OPERACAO" comp:"N"`
	IdPai    string `json:"-" db:"reg_St_idpai" xml:"REG_ST_IDPAI" comp:"N"`
	Id       string `json:"-" db:"reg_St_id" xml:"REG_ST_ID" comp:"N"`
	PedidoUk
	Filial                  *int                             `json:"filial,required" db:"FIL_IN_CODIGO" description:"Código da filial"`
	Cliente                 *int                             `json:"cliente" db:"CLI_IN_CODIGO" description:"Código do cliente (Agente tipo Cliente)"`
	ClienteCodAux           string                           `json:"clientecodaux" db:"CLI_ST_CODIGO" description:"Código do cliente auxiliar. Aceita Codigo Mega, Alternativo, CPF, CNPJ e DEPARA do integrador."`
	DataEmissao             *time.Time                       `json:"emissao,required" db:"PED_DT_EMISSAO" description:"Data de emissão (dd/mm/yyyy)"`
	CondPagto               string                           `json:"condpagto,required" db:"COND_ST_CODIGO" description:"Código da condição de pagamento"`
	Indice                  *int                             `json:"indice" db:"IND_IN_CODIGO" description:"Código do índice financeiro"`
	Representante           *int                             `json:"representante" db:"REP_IN_CODIGO" description:"código do representante (Agente tipo Representante)"`
	Equipe                  *int                             `json:"equipe" db:"EQU_IN_CODIGO" description:"Código da equipe do representate"`
	Transportadora          *int                             `json:"transportadora" db:"TRA_IN_CODIGO" description:"Código da transportadora (Agente tipo Transportadora)"`
	Redespacho              *int                             `json:"redespacho" db:"RED_IN_CODIGO" description:"Código da transportadora de redespacho (Agente tipo Transportadora)"`
	TotalFrete              float64                          `json:"totalfrete" db:"PED_RE_TOTALFRETE" description:"Valor total de frete"`
	TotalSeguro             float64                          `json:"totalseguro" db:"PED_RE_TOTALSEGURO" description:"Valor total de seguro"`
	DespAcessoria           float64                          `json:"totaldespacessoria" db:"PED_RE_TOTALDESPACESS" description:"Valor total de despesas acessórias"`
	TotalAcrescimo          float64                          `json:"totalacrescimo" db:"PED_RE_TOTALACRESCIMO" description:"Valor total de acréscimo"`
	TotalDesconto           float64                          `json:"totaldesconto" db:"PED_RE_TOTALDESCONTO" description:"Valor total de desconto"`
	Contato                 string                           `json:"contato" db:"PED_ST_CONTATO" description:"Nome do contato"`
	FoneContato             string                           `json:"contatofone" db:"PED_ST_CONTATOFONE" description:"Fone do contato"`
	CargoContato            *int                             `json:"contatocargo" db:"CAR_IN_CODIGO" description:"Cargo do contato"`
	CCustoTipo              string                           `json:"ccustotipo" db:"CCF_IDE_ST_CODIGO" description:"Identificador centro de custo"`
	CCusto                  *int                             `json:"ccusto" db:"CCF_IN_REDUZIDO" description:"código do centro de custo"`
	ProjetoTipo             string                           `json:"projetotipo" db:"PROJ_IDE_ST_CODIGO" description:"identificador do projeto"`
	Projeto                 *int                             `json:"projeto" db:"PROJ_IN_REDUZIDO" description:"Código do projeto"`
	ValorMoeda              float64                          `json:"valormoeda" db:"PED_RE_COTACAOMOE" description:"Valor de caução"`
	ValorCaucao             float64                          `json:"valorcaucao" db:"PED_RE_VALORCAUCAO" description:"Valor de cotação do índice/moeda"`
	PercentualCalcao        float64                          `json:"percentualcaucao" db:"PED_RE_PERCCAUCAO" description:"Percentual de caução"`
	PercentualAcrescimo     float64                          `json:"percentualacrescimo" db:"PED_RE_PERCACRESCIMO" description:"Percentual de acréscimo"`
	PercentualDesconto      float64                          `json:"percentualdesconto" db:"PED_RE_PERCDESCONTO" description:"Percentual de desconto"`
	Especial                string                           `json:"especial" db:"PED_BO_ESPECIAL" description:"Pedido especial"`
	PercentualDespAcessoria float64                          `json:"percentualdespacessoria" db:"PED_RE_PERCDESPESAS" description:"Percentual de despesas acessórias"`
	TipoDocumento           *int                             `json:"tipodocumento,required" db:"TPD_IN_CODIGO" description:"Código de tipo de documento"`
	CodigoAcao              *int                             `json:"codigoacao" db:"ACAO_IN_CODIGO" description:"Código de ação"`
	FretePorConta           *int                             `json:"freteporconta" db:"PED_IN_FRETEPCONTA" description:"Frete por conta: 1-Destinatario e 2-Remetente"`
	FreteEmbutidoDestacado  string                           `json:"freteembutidodestacado" db:"PED_ST_FRETEEMBDEST" description:"Valor de frete: E-Embutido e D-Destacado"`
	SequencialImportacao    *int                             `json:"sequencialimportacao" db:"PED_IN_SEQIMPORTACAO" description:"Código sequencial de importacao do pedido"`
	Usuario                 *int                             `json:"usuario" db:"USU_IN_CODIGO" description:"Usuário de importação"`
	DataEntrega             *time.Time                       `json:"entrega" db:"PED_DT_DATAENTREGA" description:"Data de entrega do pedido" comp:"N"`
	Observacao              []Observacao                     `json:"observacoes,omitempty" db:"-" description:"Observação do pedido"`
	Ocorrencia              []Ocorrencia                     `json:"ocorrencias,omitempty" db:"-" description:"Ocorrência do pedido"`
	Arquivo                 []Arquivo                        `json:"arquivos,omitempty" db:"-" description:"Anexos do pedido"`
	Parcela                 []ParcFinPedido                  `json:"parcelas,omitempty" db:"-" description:"Parcelas financeiras"`
	Item                    []Item                           `json:"itens,required" db:"-" description:"Itens do Pedido"`
	Customizacoes           customizacaoservico.Customizacao `json:"customizacao" db:"-" comp:"N"`
}

type Observacao struct {
	Operacao string `json:"-" db:"OPERACAO" comp:"N"`
	IdPai    string `json:"-" db:"reg_St_idPai" xml:"REG_ST_IDPAI" comp:"N"`
	Id       string `json:"-" db:"reg_St_id" xml:"REG_ST_ID" comp:"N"`
	PedidoUk
	TipoObservacao  string                           `json:"tipoobservacao" db:"pob_ch_tipoobservacao" xml:"POB_CH_TIPOOBSERVACAO" description:"Tipo de observação: P-Pedido, N-Nota Fiscal, F-Financeira, G-Geral"`
	TextoObservacao string                           `json:"textoobservacao" db:"pob_st_observacao" xml:"POB_ST_OBSERVACAO" description:"Texto da observação"`
	Customizacoes   customizacaoservico.Customizacao `json:"customizacao" db:"-" comp:"N"`
}

type Arquivo struct {
	Operacao string `json:"-" db:"OPERACAO" comp:"N"`
	IdPai    string `json:"-" db:"reg_St_idPai" xml:"REG_ST_IDPAI" comp:"N"`
	Id       string `json:"-" db:"reg_St_id" xml:"REG_ST_ID" comp:"N"`
	PedidoUk
	Sequencia     int                              `json:"sequencia" db:"PAR_IN_CODIGO" description:"Número sequencial"`
	Nome          string                           `json:"nome" db:"PAR_ST_NOME" description:"Nome do arquivo"`
	Conteudo      string                           `json:"conteudo" db:"PAR_ST_CONTENT" description:"Texto sobre o conteúdo do arquivo"`
	Dados         []byte                           `json:"dados" db:"PAR_BL_ARQUIVO" description:"Conteúdo do Arquivo"`
	Customizacoes customizacaoservico.Customizacao `json:"customizacao" db:"-" comp:"N"`
}

type Item struct {
	Operacao string `json:"-" db:"OPERACAO" comp:"N"`
	IdPai    string `json:"-" db:"reg_St_idPai" xml:"REG_ST_IDPAI" comp:"N"`
	Id       string `json:"-" db:"reg_St_id" xml:"REG_ST_ID" comp:"N"`
	PedidoUk
	Sequencia                  int                              `json:"sequencia" db:"itp_in_sequencia" xml:"ITP_IN_SEQUENCIA" description:"Sequencia do item"`
	Produto                    *int                             `json:"produto,required" db:"PRO_IN_CODIGO" xml:"PRO_IN_CODIGO" description:"Código do produto (informar somente se o Código Alternativo não for informado)"`
	ProdutoAlternativo         string                           `json:"produtoalternativo" db:"PRO_ST_ALTERNATIVO" xml:"PRO_ST_ALTERNATIVO" description:"Código alternativo do produto (informar somente se o Código do Produto não foi informado)" comp:"N"`
	Unidade                    string                           `json:"unidade,required" db:"UNI_ST_UNIDADE" xml:"UNI_ST_UNIDADE" description:"Código da unidade de venda do produto"`
	Descricao                  string                           `json:"descricao,required" db:"ITP_ST_DESCRICAO" xml:"ITP_ST_DESCRICAO" description:"Descrição do produto (assume a descrição do cadastro do produto se não informada)"`
	Complemento                string                           `json:"complemento" db:"ITP_ST_COMPLEMENTO" xml:"ITP_ST_COMPLEMENTO" description:"Texto Complementar do Item"`
	NCM                        *int                             `json:"ncm" db:"NCM_IN_CODIGO" xml:"NCM_IN_CODIGO" description:"Código do NCM"`
	Servico                    *int                             `json:"servico" db:"COS_IN_CODIGO" xml:"COS_IN_CODIGO" description:"Código de Serviço"`
	Aplicacao                  *int                             `json:"aplicacao" db:"APL_IN_CODIGO" xml:"APL_IN_CODIGO" description:"Código da Aplicação"`
	TabelaPreco                *int                             `json:"tabelapreco" db:"TPR_IN_CODIGO" xml:"TPR_IN_CODIGO" description:"Tabela de Preço"`
	TipoPreco                  *int                             `json:"tipopreco" db:"TPP_IN_CODIGO" xml:"TPP_IN_CODIGO" description:"Tipo de Preço (da tabela de preço)"`
	IdentificadorProjeto       string                           `json:"identificadorprojeto" db:"PROJ_IDE_ST_CODIGO" xml:"PROJ_IDE_ST_CODIGO" description:"Identificador do projeto"`
	Projeto                    *int                             `json:"projeto" db:"PROJ_IN_REDUZIDO" xml:"PROJ_IN_REDUZIDO" description:"Código do Projeto"`
	IdentificadorCentroCusto   string                           `json:"identificadorcentrocusto" db:"CUS_IDE_ST_CODIGO" xml:"CUS_IDE_ST_CODIGO" description:"Identificador do centro de custo"`
	CentroCusto                *int                             `json:"centrocusto" db:"CUS_IN_REDUZIDO" xml:"CUS_IN_REDUZIDO" description:"Código do centro de custo"`
	Quantidade                 float64                          `json:"quantidade" db:"ITP_RE_QUANTIDADE" description:"Quantidade"`
	ValorUnitario              float64                          `json:"valorunitario" db:"ITP_RE_VALORUNITARIO" xml:"ITP_RE_VALORUNITARIO" description:"Valor Unitário"`
	ValorMercadoria            float64                          `json:"valormercadoria" db:"ITP_RE_VALORMERCADORIA" xml:"ITP_RE_VALORMERCADORIA" description:"Valor total da mercadoria"`
	ValorMercadoriaEmpregada   float64                          `json:"valormercadoriaEmpregada" db:"ITP_RE_VALORMERCEMPREG" xml:"ITP_RE_VALORMERCEMPREG" description:"Valor da mercadoria empregada"`
	ValorMaoObra               float64                          `json:"valormaoobra" db:"ITP_RE_VALORMAOOBRA" xml:"ITP_RE_VALORMAOOBRA" description:"Valor de mão de obra"`
	Frete                      float64                          `json:"frete" db:"ITP_RE_FRETE" xml:"ITP_RE_FRETE" description:"Valor do frete"`
	Seguro                     float64                          `json:"seguro" db:"ITP_RE_SEGURO" xml:"ITP_RE_SEGURO" description:"Valor do seguro do frete"`
	DespesaAcessoria           float64                          `json:"despesaacessoria" db:"ITP_RE_DESPACESSORIA" xml:"ITP_RE_DESPACESSORIA" description:"Valor de despesas acessórias"`
	CodigoPedidoCliente        string                           `json:"codigopedidocliente" db:"ITP_ST_PEDIDOCLIENTE" xml:"ITP_ST_PEDIDOCLIENTE" description:"Código do pedido do cliente"`
	CodigoProdutoCliente       string                           `json:"codigoprodutocliente" db:"ITP_ST_CODPROCLI" xml:"ITP_ST_CODPROCLI" description:"Código do produto do cliente"`
	PercentualDesconto         float64                          `json:"percentualdesconto" db:"ITP_RE_PERCDESCONTO" xml:"ITP_RE_PERCDESCONTO" description:"% desconto"`
	ValorDesconto              float64                          `json:"valordesconto" db:"ITP_RE_VALORDESCONTO" xml:"ITP_RE_VALORDESCONTO" description:"Valor do desconto"`
	PercentualAcrescimo        float64                          `json:"percentualacrescimo" db:"ITP_RE_PERCACRESCIMO" xml:"ITP_RE_PERCACRESCIMO" description:"% acréscimo"`
	ValorAcrescimo             float64                          `json:"valoracrescimo" db:"ITP_RE_VALORACRESCIMO" xml:"ITP_RE_VALORACRESCIMO" description:"Valor de acréscimo"`
	Totalizadocumento          string                           `json:"totalizadocumento" db:"ITP_BO_TOTALIZA" xml:"ITP_BO_TOTALIZA" description:"Soma o valor do item no total do documento (S/N)"`
	ValorImportacao            float64                          `json:"valorimportacao" db:"ITP_RE_VALORIMPORTACAO" xml:"ITP_RE_VALORIMPORTACAO" description:"Valor de importação"`
	ValorCaucao                float64                          `json:"valorcaucao" db:"ITP_RE_VALORCAUCAO" xml:"ITP_RE_VALORCAUCAO" description:"Valor de caução"`
	Composicao                 *int                             `json:"composicao" db:"CPS_IN_CODIGO" xml:"CPS_IN_CODIGO" description:"Composição do item"`
	TipoClasse                 string                           `json:"tipoclasse" db:"TPC_ST_CLASSE" xml:"TPC_ST_CLASSE" description:"Tipo de classe financeira"`
	ItemEspecial               string                           `json:"itemespecial" db:"ITP_BO_ESPECIAL" xml:"ITP_BO_ESPECIAL" description:"Item especial"`
	FormatoConversao           string                           `json:"formatoconversao" db:"FMT_ST_CODIGO" xml:"FMT_ST_CODIGO" description:"Código do formato do item para calculo de conversão de unidade"`
	Embalagem                  *int                             `json:"embalagem" db:"EMB_IN_CODIGO" xml:"EMB_IN_CODIGO" description:"Código de embalagem"`
	UnidadeEmbalagem           string                           `json:"unidadeembalagem" db:"EMB_UNI_ST_UNIDADE" xml:"EMB_UNI_ST_UNIDADE" description:"Unidade da embalagem"`
	FormatoEmbalagem           string                           `json:"formatoembalagem" db:"EMB_FMT_ST_CODIGO" xml:"EMB_FMT_ST_CODIGO" description:"Código do formato da unidade da embalagem"`
	SequenciaPai               *int                             `json:"sequenciapai" db:"PAI_ITP_IN_SEQUENCIA" description:"Sequencia do item pai"`
	SevicoUF                   string                           `json:"servicouf" db:"UF_LOC_ST_SIGLA" description:"UF de aplicação do seriço"`
	MunicipioUF                *int                             `json:"servicomunicipio" db:"MUN_LOC_IN_CODIGO" description:"Município de aplicação do serviço"`
	GeraISSNF                  string                           `json:"geraissnf" db:"ITP_BO_GERISSNF" description:"Gera ISS na nota fiscal"`
	ValorTotal                 float64                          `json:"valortotal" db:"ITP_RE_VALORTOTAL" description:"Valor Total do Item"`
	ValorDespesasNaoTributada  float64                          `json:"valordespesasnaotributada" db:"ITP_RE_VALORDESPNTRIB" description:"Valor de despesas não tributária"`
	ValorDescontoRateado       float64                          `json:"valordescontoRateado" db:"ITP_RE_VALORDESCRATEIO" description:"Valor de desconto rateado por item"`
	ValorAcrescimoRateado      float64                          `json:"valoracrescimoRateado" db:"ITP_RE_VALORACRESCRATEIO" description:"Valor de acréscimo rateado por item"`
	QuantidadeConvertida       float64                          `json:"quantidadeconvertida" db:"ITP_RE_QTDECONVERTIDA" description:"Quantidade convertida conforme unidade de venda"`
	ValorUnitarioConvertido    float64                          `json:"valorunitarioconvertido" db:"ITP_RE_VALORUNITARIOCONV" description:"Valor Unitário convertido conforme unidade de venda"`
	PercentualDescontoUnitario float64                          `json:"percentualdescontoUnitario" db:"ITP_RE_PERDESCUNITARIO" description:"% de desconto unitário"`
	ValorDescontoUnitario      float64                          `json:"valordescontounitario" db:"ITP_RE_VALDESCUNITARIO" description:"Valor de desconto unitário"`
	ValorDescontoMicroEmpresa  float64                          `json:"valordescontomicroEmpresa" db:"ITP_RE_VALORDESCME" description:"Valor de desconto para Micro Empresa"`
	PercentualCaucao           float64                          `json:"percentualcaucao" db:"ITP_RE_PERCCAUCAO" description:"% de caução"`
	ValorCaucaoRateado         float64                          `json:"valorCaucaorateado" db:"ITP_RE_VALORCAUCAORATEIO" description:"Valor de caução rateado por item"`
	SituacaoTributariaA        string                           `json:"situacaotributariaa" db:"ITP_CH_STICMS_A" description:"Código da Situação tributária A" comp:"N"`
	SituacaoTributariaB        string                           `json:"situacaotributariab" db:"ITP_CH_STICMS_B" description:"Código da situação tributária B" comp:"N"`
	SituacaoTributariaIPI      string                           `json:"situacaotributariaipi" db:"ITP_ST_STIPI" description:"Código da Situação tributaria do IPI"`
	ObsItem                    []ObsItem                        `json:"obsitem,omitempty" db:"-" xml:"obsitem>obs" description:"Observação do Item"`
	PedProgEntrega             []PedProgEntrega                 `json:"pedprogentregas,omitempty" db:"-" xml:"progentrega>entrega" description:"Programação de Entrega do Item"`
	Customizacoes              customizacaoservico.Customizacao `json:"customizacao" db:"-" comp:"N"`
}

type ObsItem struct {
	Operacao string `json:"-" db:"OPERACAO" comp:"N"`
	IdPai    string `json:"-" db:"reg_St_idPai" xml:"REG_ST_IDPAI" comp:"N"`
	Id       string `json:"-" db:"reg_St_id" xml:"REG_ST_ID" comp:"N"`
	PedidoUk
	SequenciaItem   int                              `json:"sequenciaitem" db:"itp_in_sequencia" xml:"ITP_IN_SEQUENCIA"`
	TipoObservacao  string                           `json:"tipoobservacao" db:"oip_ch_tipoobservacao" xml:"OIP_CH_TIPOOBSERVACAO" description:"Tipo de observação: P-Pedido, N-Nota Fiscal, F-Financeira, G-Geral"`
	TextoObservacao string                           `json:"textoobservacao" db:"ito_st_observacao" xml:"ITO_ST_OBSERVACAO" descricao:"Texto da Observação"`
	Customizacoes   customizacaoservico.Customizacao `json:"customizacao" db:"-"  comp:"N"`
}

type PedProgEntrega struct {
	Operacao string `json:"-" db:"OPERACAO" comp:"N"`
	IdPai    string `json:"-" db:"reg_St_idPai" xml:"REG_ST_IDPAI" comp:"N"`
	Id       string `json:"-" db:"reg_St_id" xml:"REG_ST_ID" comp:"N"`
	PedidoUk
	Sequencia               int                              `json:"sequencia" db:"ipe_in_sequencia" xml:"IPE_IN_SEQUENCIA"`
	SequenciaItem           int                              `json:"sequenciaitem" db:"itp_in_sequencia" xml:"ITP_IN_SEQUENCIA"`
	Cliente                 *int                             `json:"cliente" db:"CLI_IN_CODIGO" xml:"CLI_IN_CODIGO" descricao:"Código do Agente (tipo Cliente)"`
	EnderecoEntrega         *int                             `json:"enderecoentrega" db:"ENA_IN_CODIGO" xml:"ENA_IN_CODIGO" descricao:"Código do endereço de entrega"`
	Quantidade              float64                          `json:"quantidade" db:"IPE_RE_QUANTIDADE" descricao:"Quantidade"`
	Unidade                 string                           `json:"unidade" db:"UNI_ST_UNIDADE" xml:"UNI_ST_UNIDADE" descricao:"Unidade"`
	FormatoConversao        string                           `json:"formatoconversao" db:"FMT_ST_CODIGO" xml:"FMT_ST_CODIGO" descricao:"Formato da unidade para conversão de unidade"`
	Embalagem               *int                             `json:"embalagem" db:"EMB_IN_CODIGO" xml:"EMB_IN_CODIGO" descricao:"Código da embalagem"`
	DataEntrega             *time.Time                       `json:"dataentrega" db:"IPE_DT_DATAENTREGA" descricao:"Data de entrega"`
	TipoEntrega             string                           `json:"tipoentrega" db:"IPE_CH_TIPOENTREGA" xml:"IPE_CH_TIPOENTREGA" descricao:"Tipo de Entrega"`
	TipoData                string                           `json:"tipodata" db:"IPE_CH_TIPODATA" xml:"IPE_CH_TIPODATA" descricao:"Tipo da data de entrega"`
	PedProgEstoque          []PedProgEstoque                 `json:"pedprogestoque,omitempty" db:"-" xml:"progestoque>estoque" descricao:"Dados de estoque da programação"`
	NumeroOE                string                           `json:"numeroOE" db:"IPE_ST_NUMEROORDEM" descricao:"Número da OE vinculada a programação"`
	DataEmissao             *time.Time                       `json:"emissao" db:"IPE_DT_DATAEMISSAO" descricao:"Data de emissão da OE"`
	QuantidadeConvertida    float64                          `json:"quantidadeConvertida" db:"IPE_RE_QTDECONVERTIDA" descricao:"Quantidade convertida"`
	ValorUnitario           float64                          `json:"valorUnitario" db:"IPE_RE_VALORUNITARIO" descricao:"Valor unitário"`
	ValorUnitarioConvertido float64                          `json:"valorUnitarioConvertido" db:"IPE_RE_VALORUNITARIOCONV" descricao:"Valor unitário convertido"`
	DataExpedicao           *time.Time                       `json:"expedicao" db:"IPE_DT_DATAEXPEDICAO" descricao:"Data da Expedição"`
	Observacao              string                           `json:"observacao" db:"IPH_ST_OBSERVACAO" descricao:"Texto de observação" comp:"N"`
	Customizacoes           customizacaoservico.Customizacao `json:"customizacao" db:"-" comp:"N"`
}

type PedProgEstoque struct {
	Operacao string `json:"-" db:"OPERACAO" comp:"N"`
	IdPai    string `json:"-" db:"reg_St_idPai" xml:"REG_ST_IDPAI" comp:"N"`
	Id       string `json:"-" db:"reg_St_id" xml:"REG_ST_ID" comp:"N"`
	PedidoUk
	SequenciaItem         int                              `json:"sequenciaitem" db:"itp_in_sequencia" xml:"ITP_IN_SEQUENCIA"`
	SequenciaProg         int                              `json:"sequenciaprog" db:"ipe_in_sequencia" xml:"IPE_IN_SEQUENCIA"`
	Sequencia             int                              `json:"sequencia" db:"ppe_in_sequencia" xml:"PPE_IN_SEQUENCIA"`
	Almoxarifado          *int                             `json:"almoxarifado" db:"ALM_IN_CODIGO" xml:"ALM_IN_CODIGO" descricao:"Código do almoxarifado"`
	Localizacao           *int                             `json:"localizacao" db:"LOC_IN_CODIGO" xml:"LOC_IN_CODIGO" descricao:"C´pdigo da localização"`
	Natureza              string                           `json:"natureza" db:"NAT_ST_CODIGO" xml:"NAT_ST_CODIGO" descricao:"Código da natureza"`
	Referencia            string                           `json:"referencia" db:"MVS_ST_REFERENCIA" xml:"MVS_ST_REFERENCIA" descricao:"Referencia de estoque do item"`
	Lote                  string                           `json:"lote" db:"MVS_ST_LOTEFORNE" xml:"MVS_ST_LOTEFORNE" descricao:"Código do Lote interno"`
	Quantidade            float64                          `json:"quantidade" db:"LMS_RE_QUANTIDADE" descricao:"Quantidade"`
	QuantidadeSolicitacao float64                          `json:"quantidadesolicitacao" db:"SOI_RE_QUANTIDADESOL" descricao:"Quantidade Solicitada ao materiais"`
	SequenciaDocumento    int                              `json:"sequenciadocumento" db:"MIE_IN_SEQMATTITDOC" descricao:"Sequencia do documento para controle de terceiros"`
	Customizacoes         customizacaoservico.Customizacao `json:"customizacao" db:"-" comp:"N"`
}

type Ocorrencia struct {
	Operacao string `json:"-" db:"OPERACAO" comp:"N"`
	IdPai    string `json:"-" db:"reg_St_idPai" xml:"REG_ST_IDPAI" comp:"N"`
	Id       string `json:"-" db:"reg_St_id" xml:"REG_ST_ID" comp:"N"`
	PedidoUk
	DataOcorrencia *time.Time                       `json:"dataocorrencia" db:"OCP_DT_DATAOCORRENCIA" descricao:"Data da ocorrência"`
	Ocorrencia     *int                             `json:"ocorrencia" db:"OCO_IN_CODIGO" descricao:"Código da ocorrência"`
	Observacao     string                           `json:"observacao" db:"OCP_ST_OBSERVACAO" descricao:"Texto da ocorrência"`
	Customizacoes  customizacaoservico.Customizacao `json:"customizacao" db:"-" comp:"N"`
}

type ParcFinPedido struct {
	Operacao string `json:"-" db:"OPERACAO" comp:"N"`
	IdPai    string `json:"-" db:"reg_St_idPai" xml:"REG_ST_IDPAI" comp:"N"`
	Id       string `json:"-" db:"reg_St_id" xml:"REG_ST_ID" comp:"N"`
	PedidoUk
	Sequencia         int                              `json:"sequencia" db:"PFP_IN_SEQUENCIA" descricao:"Sequencia da parcela"`
	Documento         string                           `json:"documento" db:"PFP_ST_DOCUMENTO" descricao:"Número do documento"`
	Parcela           string                           `json:"parcela" db:"PFP_ST_PARCELA" descricao:"Número da parcela"`
	DataVencimento    *time.Time                       `json:"vencimento" db:"PFP_DT_VENCTO" descricao:"Data de Vencimento"`
	ValorMoeda        float64                          `json:"valormoeda" db:"PFP_RE_VALORMOE" descricao:"Valor"`
	PercentualParcela float64                          `json:"percentualparcela" db:"PFP_RE_PERC" descricao:"%"`
	TipoCobranca      *int                             `json:"tipocobranca" db:"PFP_HCOB_IN_SEQUENCIA" descricao:"Código da Cobrança"`
	Customizacoes     customizacaoservico.Customizacao `json:"customizacao" db:"-" comp:"N"`
}
