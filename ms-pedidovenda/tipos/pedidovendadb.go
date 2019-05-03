package tipos

import (
	"mega/customizacaoservico"
	"time"
)

type PedidoPkDB struct {
	ORG_TAB_IN_CODIGO *int    `json:"-" db:"ORG_TAB_IN_CODIGO" xml:"ORG_TAB_IN_CODIGO" description:" " comp:"N" pk:"S"`
	ORG_PAD_IN_CODIGO *int    `json:"-" db:"ORG_PAD_IN_CODIGO" xml:"ORG_PAD_IN_CODIGO" description:" " comp:"N" pk:"S"`
	Organizacao       *int    `json:"organizacao" db:"ORG_IN_CODIGO" xml:"ORG_IN_CODIGO" description:" " comp:"N" pk:"S"`
	ORG_TAU_ST_CODIGO *string `json:"-" db:"ORG_TAU_ST_CODIGO" xml:"ORG_TAU_ST_CODIGO" description:" " comp:"N" pk:"S"`
	Serie             *string `json:"serie" db:"SER_ST_CODIGO" xml:"SER_ST_CODIGO" description:" " comp:"N" pk:"S"`
	Numero            *int    `json:"numero" db:"PED_IN_CODIGO" xml:"PED_IN_CODIGO" description:" " comp:"N" pk:"S"`
}

type PedidoVendaMemDB struct {
	Pedido PedidoVendaDB `json:"pedidovenda"`
}

type PedidoVendaDB struct {
	PedidoPkDB
	CLI_TAB_IN_CODIGO                  *int                             `json:"-" db:"CLI_TAB_IN_CODIGO" xml:"CLI_TAB_IN_CODIGO" description:" "`
	Agentepad                          *int                             `json:"clientepadrao" db:"CLI_PAD_IN_CODIGO" xml:"CLI_PAD_IN_CODIGO" description:" "`
	Cliente                            *int                             `json:"cliente" db:"CLI_IN_CODIGO" xml:"CLI_IN_CODIGO" description:" "`
	CLI_TAU_ST_CODIGO                  *string                          `json:"tipoagente" db:"CLI_TAU_ST_CODIGO" xml:"CLI_TAU_ST_CODIGO" description:" "`
	EnderecoEntrega                    *int                             `json:"enderecoentrega" db:"ENA_IN_CODIGO" xml:"ENA_IN_CODIGO" description:" "`
	Filial                             *int                             `json:"filial" db:"FIL_IN_CODIGO" xml:"FIL_IN_CODIGO" description:" "`
	TPD_TAB_IN_CODIGO                  *int                             `json:"-" db:"TPD_TAB_IN_CODIGO" xml:"TPD_TAB_IN_CODIGO" description:" "`
	TipoDocumentoPad                   *int                             `json:"tipodocumentopadrao" db:"TPD_PAD_IN_CODIGO" xml:"TPD_PAD_IN_CODIGO" description:" "`
	TipoDocumento                      *int                             `json:"tipodocumento" db:"TPD_IN_CODIGO" xml:"TPD_IN_CODIGO" description:" "`
	ACAO_TAB_IN_CODIGO                 *int                             `json:"-" db:"ACAO_TAB_IN_CODIGO" xml:"ACAO_TAB_IN_CODIGO" description:" "`
	AcaoPad                            *int                             `json:"acaopadrao" db:"ACAO_PAD_IN_CODIGO" xml:"ACAO_PAD_IN_CODIGO" description:" "`
	Acao                               *int                             `json:"codigoacao" db:"ACAO_IN_CODIGO" xml:"ACAO_IN_CODIGO" description:" "`
	Emissao                            *time.Time                       `json:"emissao" db:"PED_DT_EMISSAO" xml:"PED_DT_EMISSAO" description:" "`
	Situacao                           *string                          `json:"situacao" db:"PED_CH_SITUACAO" xml:"PED_CH_SITUACAO" description:"Situacão" comp:"N"`
	COND_TAB_IN_CODIGO                 *int                             `json:"-" db:"COND_TAB_IN_CODIGO" xml:"COND_TAB_IN_CODIGO" description:" "`
	CondicaoPagamentoPad               *int                             `json:"condpagtopad" db:"COND_PAD_IN_CODIGO" xml:"COND_PAD_IN_CODIGO" description:" "`
	Condicaopagamento                  *string                          `json:"condpagto" db:"COND_ST_CODIGO" xml:"COND_ST_CODIGO" description:" "`
	Indicemoeda                        *int                             `json:"indice" db:"IND_IN_CODIGO" xml:"IND_IN_CODIGO" description:" "`
	REP_TAB_IN_CODIGO                  *int                             `json:"-" db:"REP_TAB_IN_CODIGO" xml:"REP_TAB_IN_CODIGO" description:" "`
	REP_PAD_IN_CODIGO                  *int                             `json:"-" db:"REP_PAD_IN_CODIGO" xml:"REP_PAD_IN_CODIGO" description:" "`
	Representante                      *int                             `json:"representante" db:"REP_IN_CODIGO" xml:"REP_IN_CODIGO" description:" "`
	REP_TAU_ST_CODIGO                  *string                          `json:"-" db:"REP_TAU_ST_CODIGO" xml:"REP_TAU_ST_CODIGO" description:" "`
	EQU_TAB_IN_CODIGO                  *int                             `json:"-" db:"EQU_TAB_IN_CODIGO" xml:"EQU_TAB_IN_CODIGO" description:" "`
	EQU_PAD_IN_CODIGO                  *int                             `json:"-" db:"EQU_PAD_IN_CODIGO" xml:"EQU_PAD_IN_CODIGO" description:" "`
	Equipe                             *int                             `json:"equipe" db:"EQU_IN_CODIGO" xml:"EQU_IN_CODIGO" description:" "`
	TRA_TAB_IN_CODIGO                  *int                             `json:"-" db:"TRA_TAB_IN_CODIGO" xml:"TRA_TAB_IN_CODIGO" description:" "`
	TRA_PAD_IN_CODIGO                  *int                             `json:"-" db:"TRA_PAD_IN_CODIGO" xml:"TRA_PAD_IN_CODIGO" description:" "`
	Transportadora                     *int                             `json:"transportadora" db:"TRA_IN_CODIGO" xml:"TRA_IN_CODIGO" description:" "`
	TRA_TAU_ST_CODIGO                  *string                          `json:"-" db:"TRA_TAU_ST_CODIGO" xml:"TRA_TAU_ST_CODIGO" description:" "`
	RED_TAB_IN_CODIGO                  *int                             `json:"-" db:"RED_TAB_IN_CODIGO" xml:"RED_TAB_IN_CODIGO" description:" "`
	RED_PAD_IN_CODIGO                  *int                             `json:"-" db:"RED_PAD_IN_CODIGO" xml:"RED_PAD_IN_CODIGO" description:" "`
	Redespacho                         *int                             `json:"redespacho" db:"RED_IN_CODIGO" xml:"RED_IN_CODIGO" description:" "`
	RED_TAU_ST_CODIGO                  *string                          `json:"-" db:"RED_TAU_ST_CODIGO" xml:"RED_TAU_ST_CODIGO" description:" "`
	Acaomovimento                      *int                             `json:"acaomovimento" db:"ACAOM_IN_SEQUENCIA" xml:"ACAOM_IN_SEQUENCIA" description:" " comp:"N"`
	Tipocalculofrete                   *string                          `json:"tipocalculofrete" db:"PED_CH_TIPOCALCULOFRETE" xml:"PED_CH_TIPOCALCULOFRETE" description:" "`
	Tipopedido                         *string                          `json:"tipopedido" db:"PED_CH_TIPOPEDIDO" xml:"PED_CH_TIPOPEDIDO" description:" "`
	Valortotal                         *float64                         `json:"valortotal" db:"PED_RE_VALORTOTAL" xml:"PED_RE_VALORTOTAL" description:" "`
	Usuario                            *int                             `json:"usuario" db:"UIN_IN_CODIGO" xml:"UIN_IN_CODIGO" description:" "`
	Datainclusao                       *time.Time                       `json:"datainclusao" db:"UIN_DT_INCLUSAO" xml:"UIN_DT_INCLUSAO" description:" " comp:"N"`
	Usuarioalteracao                   *int                             `json:"usuarioalteracao" db:"UAL_IN_CODIGO" xml:"UAL_IN_CODIGO" description:" " comp:"N"`
	Dataalteracao                      *time.Time                       `json:"dataalteracao" db:"UAL_DT_ALTERACAO" xml:"UAL_DT_ALTERACAO" description:" " comp:"N"`
	Mercadoria                         *float64                         `json:"mercadoria" db:"PED_RE_VLMERCADORIA" xml:"PED_RE_VLMERCADORIA" description:" " comp:"N"`
	Mercadoriaempregada                *float64                         `json:"mercadoriaempregada" db:"PED_RE_MERCEMPREGADA" xml:"PED_RE_MERCEMPREGADA" description:" "`
	Maoobraaplicada                    *float64                         `json:"maoobraaplicada" db:"PED_RE_TOTALMAOOBRA" xml:"PED_RE_TOTALMAOOBRA" description:" "`
	Baseicms                           *float64                         `json:"baseicms" db:"PED_RE_BASEICMS" xml:"PED_RE_BASEICMS" description:" "`
	Icms                               *float64                         `json:"icms" db:"PED_RE_VLICMS" xml:"PED_RE_VLICMS" description:" "`
	Baseipi                            *float64                         `json:"baseipi" db:"PED_RE_BASEIPI" xml:"PED_RE_BASEIPI" description:" "`
	Ipi                                *float64                         `json:"ipi" db:"PED_RE_VLIPI" xml:"PED_RE_VLIPI" description:" "`
	Substituicaotributaria             *float64                         `json:"substituicaotributaria" db:"PED_RE_BASESUBTRIBUT" xml:"PED_RE_BASESUBTRIBUT" description:" "`
	Icmsretido                         *float64                         `json:"icmsretido" db:"PED_RE_VLICMSRETIDO" xml:"PED_RE_VLICMSRETIDO" description:" "`
	Iss                                *float64                         `json:"iss" db:"PED_RE_TOTALISS" xml:"PED_RE_TOTALISS" description:" "`
	Irrf                               *float64                         `json:"irrf" db:"PED_RE_TOTALIRRF" xml:"PED_RE_TOTALIRRF" description:" "`
	Inss                               *float64                         `json:"inss" db:"PED_RE_TOTALINSS" xml:"PED_RE_TOTALINSS" description:" "`
	Frete                              *float64                         `json:"totalfrete" db:"PED_RE_TOTALFRETE" xml:"PED_RE_TOTALFRETE" description:" "`
	Seguro                             *float64                         `json:"totalseguro" db:"PED_RE_TOTALSEGURO" xml:"PED_RE_TOTALSEGURO" description:" "`
	Percentualdespesas                 *float64                         `json:"percentualdespacessoria" db:"PED_RE_PERCDESPESAS" xml:"PED_RE_PERCDESPESAS" description:" "`
	Despesasacessorias                 *float64                         `json:"totaldespacessoria" db:"PED_RE_TOTALDESPACESS" xml:"PED_RE_TOTALDESPACESS" description:" "`
	Percentualacrescimo                *float64                         `json:"percentualacrescimo" db:"PED_RE_PERCACRESCIMO" xml:"PED_RE_PERCACRESCIMO" description:" "`
	Acrescimo                          *float64                         `json:"totalacrescimo" db:"PED_RE_TOTALACRESCIMO" xml:"PED_RE_TOTALACRESCIMO" description:" "`
	Percentualdesconto                 *float64                         `json:"percentualdesconto" db:"PED_RE_PERCDESCONTO" xml:"PED_RE_PERCDESCONTO" description:" "`
	Desconto                           *float64                         `json:"totaldesconto" db:"PED_RE_TOTALDESCONTO" xml:"PED_RE_TOTALDESCONTO" description:" "`
	Contato                            *string                          `json:"contato" db:"PED_ST_CONTATO" xml:"PED_ST_CONTATO" description:" "`
	Cargocontato                       *int                             `json:"contatocargo" db:"CAR_IN_CODIGO" xml:"CAR_IN_CODIGO" description:" "`
	Fonecontato                        *string                          `json:"contatofone" db:"PED_ST_CONTATOFONE" xml:"PED_ST_CONTATOFONE" description:" "`
	Valorimportacao                    *float64                         `json:"valorimportacao" db:"PED_RE_VALORIMPORTACAO" xml:"PED_RE_VALORIMPORTACAO" description:" "`
	Despesasnaotributaria              *float64                         `json:"despesasnaotributaria" db:"PED_RE_VALORDESPNTRIB" xml:"PED_RE_VALORDESPNTRIB" description:" "`
	Faturamentoparcial                 *string                          `json:"faturamentoparcial" db:"PED_BO_FATPARCIAL" xml:"PED_BO_FATPARCIAL" description:" "`
	Integramrp                         *string                          `json:"integramrp" db:"PED_BO_INTEGRAMRP" xml:"PED_BO_INTEGRAMRP" description:" "`
	Quantidadeimpressao                *int                             `json:"quantidadeimpressao" db:"PED_IN_QTDEIMPRESSAO" xml:"PED_IN_QTDEIMPRESSAO" description:" "`
	Expedicao                          *string                          `json:"expedicao" db:"PED_BO_EXPEDICAO" xml:"PED_BO_EXPEDICAO" description:" "`
	Caixa                              *string                          `json:"caixa" db:"PED_BO_CAIXA" xml:"PED_BO_CAIXA" description:" "`
	Nomecliente                        *string                          `json:"nomecliente" db:"PED_ST_NOMECLIENTE" xml:"PED_ST_NOMECLIENTE" description:" "`
	Freteporconta                      *int                             `json:"freteporconta" db:"PED_IN_FRETEPCONTA" xml:"PED_IN_FRETEPCONTA" description:" "`
	Freteembutidodestacado             *string                          `json:"freteembutidodestacado" db:"PED_ST_FRETEEMBDEST" xml:"PED_ST_FRETEEMBDEST" description:" "`
	Tiporeserva                        *string                          `json:"tiporeserva" db:"PED_CH_TIPORESERVA" xml:"PED_CH_TIPORESERVA" description:"tipo da reserva"`
	Csll                               *float64                         `json:"csll" db:"PED_RE_VALORCSLL" xml:"PED_RE_VALORCSLL" description:" "`
	Basecofins                         *float64                         `json:"basecofins" db:"PED_RE_VLBASECOFINS" xml:"PED_RE_VLBASECOFINS" description:" "`
	Basepis                            *float64                         `json:"basepis" db:"PED_RE_VLBASEPIS" xml:"PED_RE_VLBASEPIS" description:" "`
	Cofins                             *float64                         `json:"cofins" db:"PED_RE_VLCOFINS" xml:"PED_RE_VLCOFINS" description:" "`
	Cofinsretido                       *float64                         `json:"cofinsretido" db:"PED_RE_VLCOFINSRETIDO" xml:"PED_RE_VLCOFINSRETIDO" description:" "`
	Pis                                *float64                         `json:"pis" db:"PED_RE_VLPIS" xml:"PED_RE_VLPIS" description:" "`
	Pisretido                          *float64                         `json:"pisretido" db:"PED_RE_VLPISRETIDO" xml:"PED_RE_VLPISRETIDO" description:" "`
	Suframa                            *float64                         `json:"suframa" db:"PED_RE_VALORSUFRAMA" xml:"PED_RE_VALORSUFRAMA" description:"Desconto de ICMS ZFM - Suframa"`
	Suframapis                         *float64                         `json:"suframapis" db:"PED_RE_VLSUFRAMAPIS" xml:"PED_RE_VLSUFRAMAPIS" description:"Desconto de PIS ZFM - Suframa"`
	Suframacofins                      *float64                         `json:"suframacofins" db:"PED_RE_VLSUFRAMACOFINS" xml:"PED_RE_VLSUFRAMACOFINS" description:"Desconto de COFINS ZFM - Suframa"`
	CCF_TAB_IN_CODIGO                  *int                             `json:"-" db:"CCF_TAB_IN_CODIGO" xml:"CCF_TAB_IN_CODIGO" description:"Tabela"`
	CCF_PAD_IN_CODIGO                  *int                             `json:"-" db:"CCF_PAD_IN_CODIGO" xml:"CCF_PAD_IN_CODIGO" description:"Padrão"`
	Identificadorcentrocusto           *string                          `json:"ccustotipo" db:"CCF_IDE_ST_CODIGO" xml:"CCF_IDE_ST_CODIGO" description:"Identificador"`
	Centrocusto                        *int                             `json:"ccusto" db:"CCF_IN_REDUZIDO" xml:"CCF_IN_REDUZIDO" description:"Centro Custo"`
	PROJ_TAB_IN_CODIGO                 *int                             `json:"-" db:"PROJ_TAB_IN_CODIGO" xml:"PROJ_TAB_IN_CODIGO" description:"Cód. Tabela"`
	PROJ_PAD_IN_CODIGO                 *int                             `json:"-" db:"PROJ_PAD_IN_CODIGO" xml:"PROJ_PAD_IN_CODIGO" description:"Cód. Padrão"`
	Identificadorprojeto               *string                          `json:"projetotipo" db:"PROJ_IDE_ST_CODIGO" xml:"PROJ_IDE_ST_CODIGO" description:"Cód. Identificador"`
	Projeto                            *int                             `json:"projeto" db:"PROJ_IN_REDUZIDO" xml:"PROJ_IN_REDUZIDO" description:"Projeto"`
	Basesubstituicaotributariaanterior *float64                         `json:"basesubstituicaotributariaanterior" db:"PED_RE_BASESUBTRIBANT" xml:"PED_RE_BASESUBTRIBANT" description:"Base da Substituição Trib. Op. Anterior no pedido"`
	Icmsretidoanterior                 *float64                         `json:"icmsretidoanterior" db:"PED_RE_ICMSRETIDOANT" xml:"PED_RE_ICMSRETIDOANT" description:"Valor da Substituição Trib. Op. Anterior no pedido"`
	Especial                           *string                          `json:"especial" db:"PED_BO_ESPECIAL" xml:"PED_BO_ESPECIAL" description:"Pedido com todos os itens com fabricação sob encomenda"`
	Cotacaomoeda                       *float64                         `json:"valormoeda" db:"PED_RE_COTACAOMOE" xml:"PED_RE_COTACAOMOE" description:" "`
	Valorcaucao                        *float64                         `json:"valorcaucao" db:"PED_RE_VALORCAUCAO" xml:"PED_RE_VALORCAUCAO" description:" "`
	Percentualcaucao                   *float64                         `json:"percentualcaucao" db:"PED_RE_PERCCAUCAO" xml:"PED_RE_PERCCAUCAO" description:" "`
	Repasseicms                        *float64                         `json:"repasseicms" db:"PED_RE_VLREPASSEICMS" xml:"PED_RE_VLREPASSEICMS" description:"Repasse de ICMS"`
	Inssrural                          *float64                         `json:"inssrural" db:"PED_RE_VALORINSSRURAL" xml:"PED_RE_VALORINSSRURAL" description:"Valor do INSS RURAL"`
	Baseinssrural                      *float64                         `json:"baseinssrural" db:"PED_RE_BASEINSSRURAL" xml:"PED_RE_BASEINSSRURAL" description:"Base de cálculo do INSS RURAL"`
	Sequenciaimportacao                *int                             `json:"sequenciaimportacao" db:"PED_IN_SEQIMPORTACAO" xml:"PED_IN_SEQIMPORTACAO" description:"Sequencia da tabela do pedido importado."`
	Pisst                              *float64                         `json:"pisst" db:"PED_RE_PISVALORST" xml:"PED_RE_PISVALORST" description:"Valor Total do PIS ST."`
	Cofinsst                           *float64                         `json:"cofinsst" db:"PED_RE_COFINSVALORST" xml:"PED_RE_COFINSVALORST" description:"Valor Total do COFINS ST."`
	Pagina                             *int                             `json:"pagina" db:"PAG_IN_CODIGO" xml:"PAG_IN_CODIGO" description:"Cód. Contato do Agente"`
	Clientepresente                    *int                             `json:"clientepresente" db:"PED_IN_INDPRES" xml:"PED_IN_INDPRES" description:"Indicador de presença"`
	Clienteconsumidorfinal             *string                          `json:"clienteconsumidorfinal" db:"PED_BO_INDFINAL" xml:"PED_BO_INDFINAL" description:"Indicador de Consumidor Final"`
	Observacao                         []ObsPedidoDB                    `json:"observacoes" db:"-" xml:"-" description:"Observações do Pedido" comp:"N"`
	Ocorrencia                         []OcorrenciaDB                   `json:"ocorrencias" db:"-" xml:"-" description:"Ocorrencias do Pedido" comp:"N"`
	Arquivo                            []ArquivoDB                      `json:"arquivos" db:"-" xml:"-" description:"Arquivo do Pedido" comp:"N"`
	Parcelas                           []ParcFinPedidoDB                `json:"parcelas" db:"-" xml:"-" description:"Parcelas Financeiras do Pedido" comp:"N"`
	Item                               []ItemDB                         `json:"itens" db:"-" xml:"-" description:"Items do Pedido"`
	Customizacoes                      customizacaoservico.Customizacao `json:"customizacao" db:"-" comp:"N"`
}

type ObsPedidoDB struct {
	PedidoPkDB
	Tipoobservacao *string                          `json:"tipoobservacao" db:"POB_CH_TIPOOBSERVACAO" xml:"POB_CH_TIPOOBSERVACAO" description:" "`
	Observacao     string                           `json:"textoobservacao" db:"POB_ST_OBSERVACAO" xml:"POB_ST_OBSERVACAO" description:" "`
	Customizacoes  customizacaoservico.Customizacao `json:"customizacao" db:"-" comp:"N"`
}

type OcorrenciaDB struct {
	PedidoPkDB
	DataOcorrencia *time.Time                       `json:"dataocorrencia" db:"OCP_DT_DATAOCORRENCIA"`
	Ocorrencia     *int                             `json:"ocorrencia" db:"OCO_IN_CODIGO" descricao:"Código da ocorrência"`
	Observacao     string                           `json:"observacao" db:"OCP_ST_OBSERVACAO" descricao:"Texto da ocorrência"`
	Customizacoes  customizacaoservico.Customizacao `json:"customizacao" db:"-" comp:"N"`
}

type ArquivoDB struct {
	PedidoPkDB
	Usuario        *int                             `json:"usuario" db:"USU_IN_INCLUSAO" xml:"USU_IN_INCLUSAO" description:"Cód. Grupo/Usuário"`
	Sequencia      *int                             `json:"sequencia" db:"PAR_IN_CODIGO" xml:"PAR_IN_CODIGO" description:"Sequência do arquivo"`
	Nomearquivo    *string                          `json:"nome" db:"PAR_ST_NOME" xml:"PAR_ST_NOME" description:"Nome do Arquivo"`
	Resumoconteudo *string                          `json:"conteudo" db:"PAR_ST_CONTENT" xml:"PAR_ST_CONTENT" description:"Tipo do Arquivo"`
	Conteudo       []byte                           `json:"dados" db:"PAR_BL_ARQUIVO" xml:"PAR_BL_ARQUIVO" description:"Conteúdo do Arquivo"`
	DatainclusaoO  *time.Time                       `json:"datainclusao" db:"PAR_DT_INCLUSAO" xml:"PAR_DT_INCLUSAO" description:"Data"`
	Customizacoes  customizacaoservico.Customizacao `json:"customizacao" db:"-" comp:"N"`
}

type ParcFinPedidoDB struct {
	PedidoPkDB
	Sequencia            *int                             `json:"sequencia" db:"PFP_IN_SEQUENCIA" xml:"PFP_IN_SEQUENCIA" description:" "`
	Documento            *string                          `json:"documento" db:"PFP_ST_DOCUMENTO" xml:"PFP_ST_DOCUMENTO" description:" "`
	Parcela              *string                          `json:"parcela" db:"PFP_ST_PARCELA" xml:"PFP_ST_PARCELA" description:" "`
	Vencimento           *time.Time                       `json:"vencimento" db:"PFP_DT_VENCTO" xml:"PFP_DT_VENCTO" description:" "`
	Diadasemana          *string                          `json:"diadasemana" db:"PFP_ST_DIASEMANA" xml:"PFP_ST_DIASEMANA" description:" "`
	Valormoeda           *float64                         `json:"valormoeda" db:"PFP_RE_VALORMOE" xml:"PFP_RE_VALORMOE" description:" "`
	Percentualparcela    *float64                         `json:"percentualparcela" db:"PFP_RE_PERC" xml:"PFP_RE_PERC" description:" "`
	Tipocobranca         *int                             `json:"tipocobranca" db:"PFP_HCOB_IN_SEQUENCIA" xml:"PFP_HCOB_IN_SEQUENCIA" description:" "`
	Database             *time.Time                       `json:"database" db:"PFP_DT_BASE" xml:"PFP_DT_BASE" description:" "`
	Numerodiasvencimento *int                             `json:"numerodiasvencimento" db:"PFP_IN_DIASVENC" xml:"PFP_IN_DIASVENC" description:"Numero Dias"`
	Customizacoes        customizacaoservico.Customizacao `json:"customizacao" db:"-" comp:"N"`
}

type ItemDB struct {
	PedidoPkDB
	Sequencia                          *int                             `json:"sequencia" db:"ITP_IN_SEQUENCIA" xml:"ITP_IN_SEQUENCIA" description:" "`
	NCM_TAB_IN_CODIGO                  *int                             `json:"-" db:"NCM_TAB_IN_CODIGO" xml:"NCM_TAB_IN_CODIGO" description:" "`
	NCM_PAD_IN_CODIGO                  *int                             `json:"-" db:"NCM_PAD_IN_CODIGO" xml:"NCM_PAD_IN_CODIGO" description:" "`
	Ncm                                *int                             `json:"ncm" db:"NCM_IN_CODIGO" xml:"NCM_IN_CODIGO" description:" "`
	Tiposervico                        *int                             `json:"tiposervico" db:"TSE_IN_CODIGO" xml:"TSE_IN_CODIGO" description:" "`
	Servico                            *int                             `json:"servico" db:"COS_IN_CODIGO" xml:"COS_IN_CODIGO" description:" "`
	Ufservico                          *string                          `json:"ufservico" db:"UF_LOC_ST_SIGLA" xml:"UF_LOC_ST_SIGLA" description:" "`
	Municipioservico                   *int                             `json:"servicomunicipio" db:"MUN_LOC_IN_CODIGO" xml:"MUN_LOC_IN_CODIGO" description:" "`
	APL_TAB_IN_CODIGO                  *int                             `json:"-" db:"APL_TAB_IN_CODIGO" xml:"APL_TAB_IN_CODIGO" description:" "`
	APL_PAD_IN_CODIGO                  *int                             `json:"-" db:"APL_PAD_IN_CODIGO" xml:"APL_PAD_IN_CODIGO" description:" "`
	Aplicacao                          *int                             `json:"aplicacao" db:"APL_IN_CODIGO" xml:"APL_IN_CODIGO" description:" "`
	TPR_TAB_IN_CODIGO                  *int                             `json:"-" db:"TPR_TAB_IN_CODIGO" xml:"TPR_TAB_IN_CODIGO" description:" "`
	TPR_PAD_IN_CODIGO                  *int                             `json:"-" db:"TPR_PAD_IN_CODIGO" xml:"TPR_PAD_IN_CODIGO" description:" "`
	Tabelapreco                        *int                             `json:"tabelapreco" db:"TPR_IN_CODIGO" xml:"TPR_IN_CODIGO" description:" "`
	Tipopreco                          *int                             `json:"tipopreco" db:"TPP_IN_CODIGO" xml:"TPP_IN_CODIGO" description:" "`
	PROJ_TAB_IN_CODIGO                 *int                             `json:"-" db:"PROJ_TAB_IN_CODIGO" xml:"PROJ_TAB_IN_CODIGO" description:" "`
	PROJ_PAD_IN_CODIGO                 *int                             `json:"-" db:"PROJ_PAD_IN_CODIGO" xml:"PROJ_PAD_IN_CODIGO" description:" "`
	Identificadorprojeto               *string                          `json:"identificadorprojeto" db:"PROJ_IDE_ST_CODIGO" xml:"PROJ_IDE_ST_CODIGO" description:" "`
	Projeto                            *int                             `json:"projeto" db:"PROJ_IN_REDUZIDO" xml:"PROJ_IN_REDUZIDO" description:" "`
	CUS_TAB_IN_CODIGO                  *int                             `json:"-" db:"CUS_TAB_IN_CODIGO" xml:"CUS_TAB_IN_CODIGO" description:" "`
	CUS_PAD_IN_CODIGO                  *int                             `json:"-" db:"CUS_PAD_IN_CODIGO" xml:"CUS_PAD_IN_CODIGO" description:" "`
	Identificadorcentrocusto           *string                          `json:"identificadorcentrocusto" db:"CUS_IDE_ST_CODIGO" xml:"CUS_IDE_ST_CODIGO" description:" "`
	Centrocusto                        *int                             `json:"centrocusto" db:"CUS_IN_REDUZIDO" xml:"CUS_IN_REDUZIDO" description:" "`
	COND_TAB_IN_CODIGO                 *int                             `json:"-" db:"COND_TAB_IN_CODIGO" xml:"COND_TAB_IN_CODIGO" description:" "`
	COND_PAD_IN_CODIGO                 *int                             `json:"-" db:"COND_PAD_IN_CODIGO" xml:"COND_PAD_IN_CODIGO" description:" "`
	Condicaopagamento                  *string                          `json:"condicaopagamento" db:"COND_ST_CODIGO" xml:"COND_ST_CODIGO" description:" "`
	PRO_TAB_IN_CODIGO                  *int                             `json:"-" db:"PRO_TAB_IN_CODIGO" xml:"PRO_TAB_IN_CODIGO" description:" "`
	PRO_PAD_IN_CODIGO                  *int                             `json:"-" db:"PRO_PAD_IN_CODIGO" xml:"PRO_PAD_IN_CODIGO" description:" "`
	Produto                            *int                             `json:"produto" db:"PRO_IN_CODIGO" xml:"PRO_IN_CODIGO" description:" "`
	UNI_TAB_IN_CODIGO                  *int                             `json:"-" db:"UNI_TAB_IN_CODIGO" xml:"UNI_TAB_IN_CODIGO" description:" "`
	UNI_PAD_IN_CODIGO                  *int                             `json:"-" db:"UNI_PAD_IN_CODIGO" xml:"UNI_PAD_IN_CODIGO" description:" "`
	Unidade                            *string                          `json:"unidade" db:"UNI_ST_UNIDADE" xml:"UNI_ST_UNIDADE" description:" "`
	Descricao                          *string                          `json:"descricao" db:"ITP_ST_DESCRICAO" xml:"ITP_ST_DESCRICAO" description:" "`
	Complemento                        *string                          `json:"complemento" db:"ITP_ST_COMPLEMENTO" xml:"ITP_ST_COMPLEMENTO" description:" "`
	Quantidade                         *float64                         `json:"quantidade" db:"ITP_RE_QUANTIDADE" xml:"ITP_RE_QUANTIDADE" description:" "`
	Quantidadeentregue                 *float64                         `json:"quantidadeentregue" db:"ITP_RE_QTDEENTREGUE" xml:"ITP_RE_QTDEENTREGUE" description:" "`
	Quantidadefaturada                 *float64                         `json:"quantidadefaturada" db:"ITP_RE_QTDEFATURADA" xml:"ITP_RE_QTDEFATURADA" description:" "`
	Valorunitario                      *float64                         `json:"valorunitario" db:"ITP_RE_VALORUNITARIO" xml:"ITP_RE_VALORUNITARIO" description:" "`
	Valormercadoria                    *float64                         `json:"valormercadoria" db:"ITP_RE_VALORMERCADORIA" xml:"ITP_RE_VALORMERCADORIA" description:" "`
	Valormercadoriaempregada           *float64                         `json:"valormercadoriaempregada" db:"ITP_RE_VALORMERCEMPREG" xml:"ITP_RE_VALORMERCEMPREG" description:" "`
	Valormaoobra                       *float64                         `json:"valormaoobra" db:"ITP_RE_VALORMAOOBRA" xml:"ITP_RE_VALORMAOOBRA" description:" "`
	Valortotal                         *float64                         `json:"valortotal" db:"ITP_RE_VALORTOTAL" xml:"ITP_RE_VALORTOTAL" description:" "`
	Baseicms                           *float64                         `json:"baseicms" db:"ITP_RE_BASEICMS" xml:"ITP_RE_BASEICMS" description:" "`
	Valoricms                          *float64                         `json:"valoricms" db:"ITP_RE_VALORICMS" xml:"ITP_RE_VALORICMS" description:" "`
	Aliquotaicms                       *float64                         `json:"aliquotaicms" db:"ITP_RE_ALIQICMS" xml:"ITP_RE_ALIQICMS" description:" "`
	Basesubstituicaotributaria         *float64                         `json:"basesubstituicaotributaria" db:"ITP_RE_BASESUBTRIB" xml:"ITP_RE_BASESUBTRIB" description:" "`
	Icmsretido                         *float64                         `json:"icmsretido" db:"ITP_RE_ICMSRETIDO" xml:"ITP_RE_ICMSRETIDO" description:" "`
	Baseipi                            *float64                         `json:"baseipi" db:"ITP_RE_BASEIPI" xml:"ITP_RE_BASEIPI" description:" "`
	Valoripi                           *float64                         `json:"valoripi" db:"ITP_RE_VALORIPI" xml:"ITP_RE_VALORIPI" description:" "`
	Aliquotaipi                        *float64                         `json:"aliquotaipi" db:"ITP_RE_ALIQIPI" xml:"ITP_RE_ALIQIPI" description:" "`
	Baseiss                            *float64                         `json:"baseiss" db:"ITP_RE_BASEISS" xml:"ITP_RE_BASEISS" description:" "`
	Aliquotaiss                        *float64                         `json:"aliquotaiss" db:"ITP_RE_ALIQISS" xml:"ITP_RE_ALIQISS" description:" "`
	Valoriss                           *float64                         `json:"valoriss" db:"ITP_RE_VALORISS" xml:"ITP_RE_VALORISS" description:" "`
	Baseirrf                           *float64                         `json:"baseirrf" db:"ITP_RE_BASEIRRF" xml:"ITP_RE_BASEIRRF" description:" "`
	Aliquotairrf                       *float64                         `json:"aliquotairrf" db:"ITP_RE_ALIQIRRF" xml:"ITP_RE_ALIQIRRF" description:" "`
	Valorirrf                          *float64                         `json:"valorirrf" db:"ITP_RE_VALORIRRF" xml:"ITP_RE_VALORIRRF" description:" "`
	Baseinss                           *float64                         `json:"baseinss" db:"ITP_RE_BASEINSS" xml:"ITP_RE_BASEINSS" description:" "`
	Aliquoitainss                      *float64                         `json:"aliquoitainss" db:"ITP_RE_ALIQINSS" xml:"ITP_RE_ALIQINSS" description:" "`
	Valorinss                          *float64                         `json:"valorinss" db:"ITP_RE_VALORINSS" xml:"ITP_RE_VALORINSS" description:" "`
	Freteporconta                      *string                          `json:"freteporconta" db:"ITP_CH_FRETEPCONTA" xml:"ITP_CH_FRETEPCONTA" description:" "`
	Embalagemfretedestino              *string                          `json:"embalagemfretedestino" db:"ITP_CH_FRETEEMBDEST" xml:"ITP_CH_FRETEEMBDEST" description:" "`
	Valorunitariofrete                 *float64                         `json:"valorunitariofrete" db:"ITP_RE_VALORUNIFRE" xml:"ITP_RE_VALORUNIFRE" description:" "`
	Frete                              *float64                         `json:"frete" db:"ITP_RE_FRETE" xml:"ITP_RE_FRETE" description:" "`
	Seguro                             *float64                         `json:"seguro" db:"ITP_RE_SEGURO" xml:"ITP_RE_SEGURO" description:" "`
	Despesasacessorias                 *float64                         `json:"despesasacessorias" db:"ITP_RE_DESPACESSORIA" xml:"ITP_RE_DESPACESSORIA" description:" "`
	Situacao                           *string                          `json:"situacao" db:"ITP_ST_SITUACAO" xml:"ITP_ST_SITUACAO" description:"Situacão" comp:"N"`
	Quantidadeconvertidaestoque        *float64                         `json:"quantidadeconvertidaestoque" db:"ITP_RE_QTDECONVEST" xml:"ITP_RE_QTDECONVEST" description:" "`
	Valorconvertido                    *float64                         `json:"valorconvertido" db:"ITP_RE_VALORCONVEST" xml:"ITP_RE_VALORCONVEST" description:" "`
	Numeropedidocliente                *string                          `json:"numeropedidocliente" db:"ITP_ST_PEDIDOCLIENTE" xml:"ITP_ST_PEDIDOCLIENTE" description:" "`
	Codigoprodutocliente               *string                          `json:"codigoprodutocliente" db:"ITP_ST_CODPROCLI" xml:"ITP_ST_CODPROCLI" description:" "`
	Percentualdesconto                 *float64                         `json:"percentualdesconto" db:"ITP_RE_PERCDESCONTO" xml:"ITP_RE_PERCDESCONTO" description:" "`
	Valordesconto                      *float64                         `json:"valordesconto" db:"ITP_RE_VALORDESCONTO" xml:"ITP_RE_VALORDESCONTO" description:" "`
	Percentualacrescmio                *float64                         `json:"percentualacrescmio" db:"ITP_RE_PERCACRESCIMO" xml:"ITP_RE_PERCACRESCIMO" description:" "`
	Valoracrescimo                     *float64                         `json:"valoracrescimo" db:"ITP_RE_VALORACRESCIMO" xml:"ITP_RE_VALORACRESCIMO" description:" "`
	Precopadrao                        *float64                         `json:"precopadrao" db:"ITP_RE_PRECOPADRAO" xml:"ITP_RE_PRECOPADRAO" description:" "`
	Comissaopadrao                     *float64                         `json:"comissaopadrao" db:"ITP_RE_COMISSAOPADRAO" xml:"ITP_RE_COMISSAOPADRAO" description:" "`
	Comissaofinal                      *float64                         `json:"comissaofinal" db:"ITP_RE_COMISSAOFINAL" xml:"ITP_RE_COMISSAOFINAL" description:" "`
	Basecomissao                       *float64                         `json:"basecomissao" db:"ITP_RE_BASECOMISSAO" xml:"ITP_RE_BASECOMISSAO" description:" "`
	Valorcomissao                      *float64                         `json:"valorcomissao" db:"ITP_RE_VALORCOMISSAO" xml:"ITP_RE_VALORCOMISSAO" description:" "`
	Usuarioinclusao                    *int                             `json:"usuarioinclusao" db:"UIN_IN_CODIGO" xml:"UIN_IN_CODIGO" description:" " comp:"N"`
	Datainclusao                       *time.Time                       `json:"datainclusao" db:"UIN_DT_INCLUSAO" xml:"UIN_DT_INCLUSAO" description:" " comp:"N"`
	Usuarioalteracao                   *int                             `json:"usuarioalteracao" db:"UAL_IN_CODIGO" xml:"UAL_IN_CODIGO" description:" " comp:"N"`
	Dataalteracao                      *time.Time                       `json:"dataalteracao" db:"UAL_DT_ALTERACAO" xml:"UAL_DT_ALTERACAO" description:" " comp:"N"`
	Totalizadocumento                  *string                          `json:"totalizadocumento" db:"ITP_BO_TOTALIZA" xml:"ITP_BO_TOTALIZA" description:" "`
	Valorimportacao                    *float64                         `json:"valorimportacao" db:"ITP_RE_VALORIMPORTACAO" xml:"ITP_RE_VALORIMPORTACAO" description:" "`
	Despesasnaotributaria              *float64                         `json:"despesasnaotributaria" db:"ITP_RE_VALORDESPNTRIB" xml:"ITP_RE_VALORDESPNTRIB" description:" "`
	Descontorateado                    *float64                         `json:"descontorateado" db:"ITP_RE_VALORDESCRATEIO" xml:"ITP_RE_VALORDESCRATEIO" description:" "`
	AcrescimorateadoO                  *float64                         `json:"acrescimorateado" db:"ITP_RE_VALORACRESCRATEIO" xml:"ITP_RE_VALORACRESCRATEIO" description:" "`
	Composicao                         *int                             `json:"composicao" db:"CPS_IN_CODIGO" xml:"CPS_IN_CODIGO" description:" "`
	TPC_TAB_IN_CODIGO                  *int                             `json:"-" db:"TPC_TAB_IN_CODIGO" xml:"TPC_TAB_IN_CODIGO" description:" "`
	TPC_PAD_IN_CODIGO                  *int                             `json:"-" db:"TPC_PAD_IN_CODIGO" xml:"TPC_PAD_IN_CODIGO" description:" "`
	Tipoclasse                         *string                          `json:"tipoclasse" db:"TPC_ST_CLASSE" xml:"TPC_ST_CLASSE" description:" "`
	Itemespecial                       *string                          `json:"itemespecial" db:"ITP_BO_ESPECIAL" xml:"ITP_BO_ESPECIAL" description:" "`
	Reservaestoque                     *string                          `json:"reservaestoque" db:"ITP_BO_RESERVAESTOQUE" xml:"ITP_BO_RESERVAESTOQUE" description:" "`
	CFOP_TAB_IN_CODIGO                 *int                             `json:"-" db:"CFOP_TAB_IN_CODIGO" xml:"CFOP_TAB_IN_CODIGO" description:" "`
	CFOP_PAD_IN_CODIGO                 *int                             `json:"-" db:"CFOP_PAD_IN_CODIGO" xml:"CFOP_PAD_IN_CODIGO" description:" "`
	Identificadorcfop                  *string                          `json:"identificadorcfop" db:"CFOP_IDE_ST_CODIGO" xml:"CFOP_IDE_ST_CODIGO" description:" "`
	Cfop                               *int                             `json:"cfop" db:"CFOP_IN_CODIGO" xml:"CFOP_IN_CODIGO" description:" "`
	FMT_TAB_IN_CODIGO                  *int                             `json:"-" db:"FMT_TAB_IN_CODIGO" xml:"FMT_TAB_IN_CODIGO" description:" "`
	FMT_PAD_IN_CODIGO                  *int                             `json:"-" db:"FMT_PAD_IN_CODIGO" xml:"FMT_PAD_IN_CODIGO" description:" "`
	Formato                            *string                          `json:"formato" db:"FMT_ST_CODIGO" xml:"FMT_ST_CODIGO" description:" "`
	Quantidadeconvertida               *float64                         `json:"quantidadeconvertida" db:"ITP_RE_QTDECONVERTIDA" xml:"ITP_RE_QTDECONVERTIDA" description:" "`
	ValorunitarioconvertidoV           *float64                         `json:"valorunitarioconvertido" db:"ITP_RE_VALORUNITARIOCONV" xml:"ITP_RE_VALORUNITARIOCONV" description:" "`
	EMB_TAB_IN_CODIGO                  *int                             `json:"-" db:"EMB_TAB_IN_CODIGO" xml:"EMB_TAB_IN_CODIGO" description:" "`
	EMB_PAD_IN_CODIGO                  *int                             `json:"-" db:"EMB_PAD_IN_CODIGO" xml:"EMB_PAD_IN_CODIGO" description:" "`
	Embalagem                          *int                             `json:"embalagem" db:"EMB_IN_CODIGO" xml:"EMB_IN_CODIGO" description:" "`
	ALM_TAB_IN_CODIGO                  *int                             `json:"-" db:"ALM_TAB_IN_CODIGO" xml:"ALM_TAB_IN_CODIGO" description:" "`
	ALM_PAD_IN_CODIGO                  *int                             `json:"-" db:"ALM_PAD_IN_CODIGO" xml:"ALM_PAD_IN_CODIGO" description:" "`
	Almoxarifado                       *int                             `json:"almoxarifado" db:"ALM_IN_CODIGO" xml:"ALM_IN_CODIGO" description:" "`
	Localizacao                        *int                             `json:"localizacao" db:"LOC_IN_CODIGO" xml:"LOC_IN_CODIGO" description:" "`
	EMB_UNI_TAB_IN_CODIGO              *int                             `json:"-" db:"EMB_UNI_TAB_IN_CODIGO" xml:"EMB_UNI_TAB_IN_CODIGO" description:" "`
	EMB_UNI_PAD_IN_CODIGO              *int                             `json:"-" db:"EMB_UNI_PAD_IN_CODIGO" xml:"EMB_UNI_PAD_IN_CODIGO" description:" "`
	Unidadeembalagem                   *string                          `json:"unidadeembalagem" db:"EMB_UNI_ST_UNIDADE" xml:"EMB_UNI_ST_UNIDADE" description:" "`
	EMB_FMT_TAB_IN_CODIGO              *int                             `json:"-" db:"EMB_FMT_TAB_IN_CODIGO" xml:"EMB_FMT_TAB_IN_CODIGO" description:" "`
	EMB_FMT_PAD_IN_CODIGO              *int                             `json:"-" db:"EMB_FMT_PAD_IN_CODIGO" xml:"EMB_FMT_PAD_IN_CODIGO" description:" "`
	Formatoembalagem                   *string                          `json:"formatoembalagem" db:"EMB_FMT_ST_CODIGO" xml:"EMB_FMT_ST_CODIGO" description:" "`
	Quantidadeauxiliar                 *float64                         `json:"quantidadeauxiliar" db:"ITP_RE_QTDEAUX" xml:"ITP_RE_QTDEAUX" description:" "`
	Largura                            *float64                         `json:"largura" db:"ITP_RE_LARGURA" xml:"ITP_RE_LARGURA" description:" "`
	Comprimento                        *float64                         `json:"comprimento" db:"ITP_RE_COMPRIMENTO" xml:"ITP_RE_COMPRIMENTO" description:" "`
	Medidareal                         *float64                         `json:"medidareal" db:"ITP_RE_MEDIDAREAL" xml:"ITP_RE_MEDIDAREAL" description:" "`
	Percentualdescontounitario         *float64                         `json:"percentualdescontounitario" db:"ITP_RE_PERDESCUNITARIO" xml:"ITP_RE_PERDESCUNITARIO" description:" "`
	Descontounitario                   *float64                         `json:"descontounitario" db:"ITP_RE_VALDESCUNITARIO" xml:"ITP_RE_VALDESCUNITARIO" description:" "`
	Percentualdescontoinformado        *float64                         `json:"percentualdescontoinformado" db:"ITP_RE_PERCDESCDIG" xml:"ITP_RE_PERCDESCDIG" description:"% Desconto Percentual de desconto digitado pelo usuario."`
	Precotabelapreco                   *float64                         `json:"precotabelapreco" db:"ITP_RE_PRECOTABELA" xml:"ITP_RE_PRECOTABELA" description:"Preco tabela"`
	Percentualcofins                   *float64                         `json:"percentualcofins" db:"ITP_RE_PERCCOFINS" xml:"ITP_RE_PERCCOFINS" description:" "`
	Percentualpis                      *float64                         `json:"percentualpis" db:"ITP_RE_PERCPIS" xml:"ITP_RE_PERCPIS" description:" "`
	Percentualcsll                     *float64                         `json:"percentualcsll" db:"ITP_RE_PERCSLL" xml:"ITP_RE_PERCSLL" description:" "`
	Basecofins                         *float64                         `json:"basecofins" db:"ITP_RE_VLBASECOFINS" xml:"ITP_RE_VLBASECOFINS" description:" "`
	Basecsll                           *float64                         `json:"basecsll" db:"ITP_RE_VLBASECSLL" xml:"ITP_RE_VLBASECSLL" description:" "`
	Basepis                            *float64                         `json:"basepis" db:"ITP_RE_VLBASEPIS" xml:"ITP_RE_VLBASEPIS" description:" "`
	Cofins                             *float64                         `json:"cofins" db:"ITP_RE_VLCOFINS" xml:"ITP_RE_VLCOFINS" description:" "`
	Cofinsretido                       *float64                         `json:"cofinsretido" db:"ITP_RE_VLCOFINSRETIDO" xml:"ITP_RE_VLCOFINSRETIDO" description:" "`
	Csll                               *float64                         `json:"csll" db:"ITP_RE_VLCSLL" xml:"ITP_RE_VLCSLL" description:" "`
	Pis                                *float64                         `json:"pis" db:"ITP_RE_VLPIS" xml:"ITP_RE_VLPIS" description:" "`
	Pisretido                          *float64                         `json:"pisretido" db:"ITP_RE_VLPISRETIDO" xml:"ITP_RE_VLPISRETIDO" description:" "`
	VSO_TAB_IN_CODIGO                  *int                             `json:"-" db:"VSO_TAB_IN_CODIGO" xml:"VSO_TAB_IN_CODIGO" description:"Cód. Tabela"`
	VSO_PAD_IN_CODIGO                  *int                             `json:"-" db:"VSO_PAD_IN_CODIGO" xml:"VSO_PAD_IN_CODIGO" description:"Cód. Padrão"`
	Serieos                            *string                          `json:"serieos" db:"VSO_ST_CODIGO" xml:"VSO_ST_CODIGO" description:"Série OS"`
	SequenciaPai                       *int                             `json:"sequenciapai" db:"PAI_ITP_IN_SEQUENCIA" xml:"PAI_ITP_IN_SEQUENCIA" description:" "`
	Numeroos                           *string                          `json:"numeroos" db:"ITP_ST_CODIGOOS" xml:"ITP_ST_CODIGOOS" description:"Número OS"`
	Codigobuscaproduto                 *string                          `json:"codigobuscaproduto" db:"ITP_ST_CODBUSCAPROD" xml:"ITP_ST_CODBUSCAPROD" description:"Código da busca"`
	Suframa                            *float64                         `json:"suframa" db:"ITP_RE_VLSUFRAMA" xml:"ITP_RE_VLSUFRAMA" description:"Desconto de ICMS ZFM - Suframa"`
	Suframapis                         *float64                         `json:"suframapis" db:"ITP_RE_VLSUFRAMAPIS" xml:"ITP_RE_VLSUFRAMAPIS" description:"Desconto de PIS ZFM - Suframa"`
	Suframacofins                      *float64                         `json:"suframacofins" db:"ITP_RE_VLSUFRAMACOFINS" xml:"ITP_RE_VLSUFRAMACOFINS" description:"Desconto de COFINS ZFM - Suframa"`
	Quantidadefaturadaconvertida       *float64                         `json:"quantidadefaturadaconvertida" db:"ITP_RE_QTDEFATURADACONV" xml:"ITP_RE_QTDEFATURADACONV" description:" "`
	Quantidadeentregueconvertida       *float64                         `json:"quantidadeentregueconvertida" db:"ITP_RE_QTDEENTREGUECONV" xml:"ITP_RE_QTDEENTREGUECONV" description:" "`
	Precominimo                        *float64                         `json:"precominimo" db:"ITP_RE_CF_PRECOMIN" xml:"ITP_RE_CF_PRECOMIN" description:"Preco minimo da tabela de preco padrao da Centroflora"`
	Precoobjetivo                      *float64                         `json:"precoobjetivo" db:"ITP_RE_CF_PRECOOBJ" xml:"ITP_RE_CF_PRECOOBJ" description:"Preco objetivo da tabela de preco padrao da Centroflora"`
	Tabelaprecopadrao                  *string                          `json:"tabelaprecopadrao" db:"ITP_ST_CF_TABPRECO" xml:"ITP_ST_CF_TABPRECO" description:"Tabela padrao correspondente ao produto no momento da inclusao do pedido"`
	Descontomicroempresa               *float64                         `json:"descontomicroempresa" db:"ITP_RE_VALORDESCME" xml:"ITP_RE_VALORDESCME" description:"Valor do desconto que será utilizado para abater no preço do item."`
	Basesubstituicaotributariaanterior *float64                         `json:"basesubstituicaotributariaanterior" db:"ITP_RE_BASESUBTRIBANT" xml:"ITP_RE_BASESUBTRIBANT" description:"Base da Substituição Trib. Op. Anterior no item do pedido"`
	Icmsretidoanterior                 *float64                         `json:"icmsretidoanterior" db:"ITP_RE_ICMSRETIDOANT" xml:"ITP_RE_ICMSRETIDOANT" description:"Valor da Substituição Trib. Op. Anterior no item do pedido"`
	Servicomunicipio                   *int                             `json:"servicomunicipio" db:"COSM_IN_CODIGO" xml:"COSM_IN_CODIGO" description:"Código de serviço por município."`
	Caucao                             *float64                         `json:"caucao" db:"ITP_RE_VALORCAUCAO" xml:"ITP_RE_VALORCAUCAO" description:" "`
	Percentualcaucao                   *float64                         `json:"percentualcaucao" db:"ITP_RE_PERCCAUCAO" xml:"ITP_RE_PERCCAUCAO" description:" "`
	CaucaorateadoO                     *float64                         `json:"caucaorateado" db:"ITP_RE_VALORCAUCAORATEIO" xml:"ITP_RE_VALORCAUCAORATEIO" description:" "`
	Repasseicms                        *float64                         `json:"repasseicms" db:"ITP_RE_VLREPASSEICMS" xml:"ITP_RE_VLREPASSEICMS" description:"Repasse de ICMS"`
	Aliquotainssrural                  *float64                         `json:"aliquotainssrural" db:"ITP_RE_ALIQINSSRURAL" xml:"ITP_RE_ALIQINSSRURAL" description:"Alíquota de INSS RURAL"`
	Inssrural                          *float64                         `json:"inssrural" db:"ITP_RE_VALORINSSRURAL" xml:"ITP_RE_VALORINSSRURAL" description:"Valor do INSS RURAL"`
	Baseinssrural                      *float64                         `json:"baseinssrural" db:"ITP_RE_BASEINSSRURAL" xml:"ITP_RE_BASEINSSRURAL" description:"Base de cálculo do INSS RURAL"`
	Isentoicms                         *float64                         `json:"isentoicms" db:"ITP_RE_ISENTOICMS" xml:"ITP_RE_ISENTOICMS" description:"Isento ICMS"`
	Outrosicms                         *float64                         `json:"outrosicms" db:"ITP_RE_OUTROSICMS" xml:"ITP_RE_OUTROSICMS" description:"Outros ICMS"`
	Icmsrecuperado                     *float64                         `json:"icmsrecuperado" db:"ITP_RE_RECUPERADOICMS" xml:"ITP_RE_RECUPERADOICMS" description:"ICMS Recuperado"`
	Ipiisento                          *float64                         `json:"ipiisento" db:"ITP_RE_ISENTOIPI" xml:"ITP_RE_ISENTOIPI" description:"Isento IPI"`
	Ipioutros                          *float64                         `json:"ipioutros" db:"ITP_RE_OUTROSIPI" xml:"ITP_RE_OUTROSIPI" description:"Outros IPI"`
	Ipirecuperado                      *float64                         `json:"ipirecuperado" db:"ITP_RE_RECUPERADOIPI" xml:"ITP_RE_RECUPERADOIPI" description:"IPI Recuperado"`
	Pisrecuperado                      *float64                         `json:"pisrecuperado" db:"ITP_RE_RECUPERADOPIS" xml:"ITP_RE_RECUPERADOPIS" description:"Valor PIS Recuperado"`
	Cofinsrecuperado                   *float64                         `json:"cofinsrecuperado" db:"ITP_RE_RECUPERADOCOFINS" xml:"ITP_RE_RECUPERADOCOFINS" description:"Valor COFINS Recuperado"`
	Percentuallucro                    *float64                         `json:"percentuallucro" db:"ITP_RE_PERCLUCRO" xml:"ITP_RE_PERCLUCRO" description:"Valor percentual lucro"`
	Cstpis                             *string                          `json:"cstpis" db:"STP_ST_CSTPIS" xml:"STP_ST_CSTPIS" description:"CST PIS Código da Situação Tributária do PIS"`
	Cstcofins                          *string                          `json:"cstcofins" db:"STC_ST_CSTCOFINS" xml:"STC_ST_CSTCOFINS" description:"CST COFINS Código da Situação Tributária do COFINS"`
	Sticmsa                            *string                          `json:"sticmsa" db:"ITP_CH_STICMS_A" xml:"ITP_CH_STICMS_A" description:" "`
	Sticmsb                            *string                          `json:"sticmsb" db:"ITP_CH_STICMS_B" xml:"ITP_CH_STICMS_B" description:" "`
	Situacaotributariaipi              *string                          `json:"situacaotributariaipi" db:"ITP_ST_STIPI" xml:"ITP_ST_STIPI" description:"Situação Tributária do IPI"`
	Quantidadepadraotributada          *float64                         `json:"quantidadepadraotributada" db:"ITP_RE_QTDEPADRAOTRIB" xml:"ITP_RE_QTDEPADRAOTRIB" description:"Quantidade Padrão Tributável."`
	Pisbasest                          *float64                         `json:"pisbasest" db:"ITP_RE_PISBASEST" xml:"ITP_RE_PISBASEST" description:"Valor da Base do PIS ST."`
	Pisaliquotast                      *float64                         `json:"pisaliquotast" db:"ITP_RE_PISALIQST" xml:"ITP_RE_PISALIQST" description:"Alíquota do PIS ST."`
	Pisaliquotavalorst                 *float64                         `json:"pisaliquotavalorst" db:"ITP_RE_PISALIQVALORST" xml:"ITP_RE_PISALIQVALORST" description:"Alíquota em Valor do PIS ST."`
	Pisst                              *float64                         `json:"pisst" db:"ITP_RE_PISVALORST" xml:"ITP_RE_PISVALORST" description:"Valor do PIS ST."`
	Pispercentuallucro                 *float64                         `json:"pispercentuallucro" db:"ITP_RE_PISPERCLUCRO" xml:"ITP_RE_PISPERCLUCRO" description:"Valor da Margem de Lucro (IVA) do PIS ST."`
	Cofinsbasest                       *float64                         `json:"cofinsbasest" db:"ITP_RE_COFINSBASEST" xml:"ITP_RE_COFINSBASEST" description:"Valor da Base do COFINS ST."`
	Cofinsaliqst                       *float64                         `json:"cofinsaliqst" db:"ITP_RE_COFINSALIQST" xml:"ITP_RE_COFINSALIQST" description:"Alíquota do COFINS ST."`
	CofinsaliquotavalorstT             *float64                         `json:"cofinsaliquotavalorst" db:"ITP_RE_COFINSALIQVALORST" xml:"ITP_RE_COFINSALIQVALORST" description:"Alíquota em Valor do COFINS ST."`
	Cofinsst                           *float64                         `json:"cofinsst" db:"ITP_RE_COFINSVALORST" xml:"ITP_RE_COFINSVALORST" description:"Valor do COFINS ST."`
	Cofinspercentuallucro              *float64                         `json:"cofinspercentuallucro" db:"ITP_RE_COFINSPERCLUCRO" xml:"ITP_RE_COFINSPERCLUCRO" description:"Valor da Margem de Lucro (IVA) do COFINS ST."`
	Definicaoipi                       *string                          `json:"definicaoipi" db:"ITP_CH_DEFIPI" xml:"ITP_CH_DEFIPI" description:"Definição IPI"`
	Pautaipi                           *float64                         `json:"pautaipi" db:"ITP_RE_PAUTAIPI" xml:"ITP_RE_PAUTAIPI" description:"Pauta do IPI"`
	Observacao                         []ObsItemDB                      `json:"obsitem" db:"-" xml:"-" description:"Observações do Item do Pedido" comp:"N"`
	ProgEntrega                        []PedProgEntregaDB               `json:"pedprogentregas" db:"-" xml:"-" description:"Programação de Entrega"`
	Customizacoes                      customizacaoservico.Customizacao `json:"customizacao" db:"-" comp:"N"`
}

type ObsItemDB struct {
	PedidoPkDB
	Sequenciaitem   *int                             `json:"sequenciaitem" db:"ITP_IN_SEQUENCIA" xml:"ITP_IN_SEQUENCIA" description:" "`
	TipoobservacaoO *string                          `json:"tipoobservacao" db:"OIP_CH_TIPOOBSERVACAO" xml:"OIP_CH_TIPOOBSERVACAO" description:" "`
	Observacao      string                           `json:"textoobservacao" db:"ITO_ST_OBSERVACAO" xml:"ITO_ST_OBSERVACAO" description:" "`
	Customizacoes   customizacaoservico.Customizacao `json:"customizacao" db:"-" comp:"N"`
}

type PedProgEntregaDB struct {
	PedidoPkDB
	Sequenciaitem                *int                             `json:"sequenciaitem" db:"ITP_IN_SEQUENCIA" xml:"ITP_IN_SEQUENCIA" description:" "`
	Sequencia                    *int                             `json:"sequencia" db:"IPE_IN_SEQUENCIA" xml:"IPE_IN_SEQUENCIA" description:" "`
	UNI_TAB_IN_CODIGO            *int                             `json:"-" db:"UNI_TAB_IN_CODIGO" xml:"UNI_TAB_IN_CODIGO" description:" "`
	UNI_PAD_IN_CODIGO            *int                             `json:"-" db:"UNI_PAD_IN_CODIGO" xml:"UNI_PAD_IN_CODIGO" description:" "`
	Unidade                      *string                          `json:"unidade" db:"UNI_ST_UNIDADE" xml:"UNI_ST_UNIDADE" description:" "`
	CLI_TAB_IN_CODIGO            *int                             `json:"-" db:"CLI_TAB_IN_CODIGO" xml:"CLI_TAB_IN_CODIGO" description:" "`
	CLI_PAD_IN_CODIGO            *int                             `json:"-" db:"CLI_PAD_IN_CODIGO" xml:"CLI_PAD_IN_CODIGO" description:" "`
	Agente                       *int                             `json:"cliente" db:"CLI_IN_CODIGO" xml:"CLI_IN_CODIGO" description:" "`
	Tipoagente                   *string                          `json:"tipoagente" db:"CLI_TAU_ST_CODIGO" xml:"CLI_TAU_ST_CODIGO" description:" "`
	Enderecoentrega              *int                             `json:"enderecoentrega" db:"ENA_IN_CODIGO" xml:"ENA_IN_CODIGO" description:" "`
	Quantidade                   *float64                         `json:"quantidade" db:"IPE_RE_QUANTIDADE" xml:"IPE_RE_QUANTIDADE" description:" "`
	Quantidadefaturada           *float64                         `json:"quantidadefaturada" db:"IPE_RE_QTDEFATURADA" xml:"IPE_RE_QTDEFATURADA" description:" "`
	Quantidadeentregue           *float64                         `json:"quantidadeentregue" db:"IPE_RE_QTDEENTREGUE" xml:"IPE_RE_QTDEENTREGUE" description:" "`
	Situacao                     *string                          `json:"situacao" db:"IPE_CH_SITUACAO" xml:"IPE_CH_SITUACAO" description:" " comp:"N"`
	Dataentrega                  *time.Time                       `json:"dataentrega" db:"IPE_DT_DATAENTREGA" xml:"IPE_DT_DATAENTREGA" description:" "`
	Tipoentrega                  *string                          `json:"tipoentrega" db:"IPE_CH_TIPOENTREGA" xml:"IPE_CH_TIPOENTREGA" description:" "`
	Permiteentregaparcial        *string                          `json:"permiteentregaparcial" db:"IPE_CH_ENTREGAPARCIAL" xml:"IPE_CH_ENTREGAPARCIAL" description:" "`
	Numeroordemexpedicao         *string                          `json:"numeroordemexpedicao" db:"IPE_ST_NUMEROORDEM" xml:"IPE_ST_NUMEROORDEM" description:" "`
	Dataemissaoordemexpedicao    *time.Time                       `json:"dataemissaoordemexpedicao" db:"IPE_DT_DATAEMISSAO" xml:"IPE_DT_DATAEMISSAO" description:" " comp:"N"`
	Usuarioinclusao              *int                             `json:"usuarioinclusao" db:"UIN_IN_CODIGO" xml:"UIN_IN_CODIGO" description:" " comp:"N"`
	Datainclusao                 *time.Time                       `json:"datainclusao" db:"UIN_DT_INCLUSAO" xml:"UIN_DT_INCLUSAO" description:" " comp:"N"`
	Usuarioalteracao             *int                             `json:"usuarioalteracao" db:"UAL_IN_CODIGO" xml:"UAL_IN_CODIGO" description:" " comp:"N"`
	Dataalteracao                *time.Time                       `json:"dataalteracao" db:"UAL_DT_ALTERACAO" xml:"UAL_DT_ALTERACAO" description:" " comp:"N"`
	Tipodataentrega              *string                          `json:"tipodata" db:"IPE_CH_TIPODATA" xml:"IPE_CH_TIPODATA" description:" "`
	Dataoriginal                 *time.Time                       `json:"dataoriginal" db:"IPE_DT_DATAORIGINAL" xml:"IPE_DT_DATAORIGINAL" description:" "`
	Dataplanejada                *time.Time                       `json:"dataplanejada" db:"IPE_DT_DATAPLANEJADA" xml:"IPE_DT_DATAPLANEJADA" description:" "`
	FMT_TAB_IN_CODIGO            *int                             `json:"-" db:"FMT_TAB_IN_CODIGO" xml:"FMT_TAB_IN_CODIGO" description:" "`
	FMT_PAD_IN_CODIGO            *int                             `json:"-" db:"FMT_PAD_IN_CODIGO" xml:"FMT_PAD_IN_CODIGO" description:" "`
	Formato                      *string                          `json:"formato" db:"FMT_ST_CODIGO" xml:"FMT_ST_CODIGO" description:" "`
	Quantidadeconvertida         *float64                         `json:"quantidadeconvertida" db:"IPE_RE_QTDECONVERTIDA" xml:"IPE_RE_QTDECONVERTIDA" description:" "`
	EMB_TAB_IN_CODIGO            *int                             `json:"-" db:"EMB_TAB_IN_CODIGO" xml:"EMB_TAB_IN_CODIGO" description:" "`
	EMB_PAD_IN_CODIGO            *int                             `json:"-" db:"EMB_PAD_IN_CODIGO" xml:"EMB_PAD_IN_CODIGO" description:" "`
	Embalagem                    *int                             `json:"embalagem" db:"EMB_IN_CODIGO" xml:"EMB_IN_CODIGO" description:" "`
	Valorunitario                *float64                         `json:"valorunitario" db:"IPE_RE_VALORUNITARIO" xml:"IPE_RE_VALORUNITARIO" description:" "`
	ValorunitarioconvertidoV     *float64                         `json:"valorunitarioconvertido" db:"IPE_RE_VALORUNITARIOCONV" xml:"IPE_RE_VALORUNITARIOCONV" description:" "`
	Saldomrp                     *float64                         `json:"saldomrp" db:"IPE_RE_SALDOMRP" xml:"IPE_RE_SALDOMRP" description:" "`
	Dataexpedicao                *time.Time                       `json:"expedicao" db:"IPE_DT_DATAEXPEDICAO" xml:"IPE_DT_DATAEXPEDICAO" description:"Data de Expedic?o" comp:"N"`
	Quantidadefaturadaconvertida *float64                         `json:"quantidadefaturadaconvertida" db:"IPE_RE_QTDEFATURADACONV" xml:"IPE_RE_QTDEFATURADACONV" description:"Qtde Fat.Conv." comp:"N"`
	Quantidadeentregueconvertida *float64                         `json:"quantidadeentregueconvertida" db:"IPE_RE_QTDEENTREGUECONV" xml:"IPE_RE_QTDEENTREGUECONV" description:"Qtde Ent.Conv." comp:"N"`
	ProgEstoque                  []PedProgEstoqueDB               `json:"pedprogestoque" db:"-" xml:"-" description:"Dados de estoque da programação de entrega" comp:"N"`
	Customizacoes                customizacaoservico.Customizacao `json:"customizacao" db:"-" comp:"N"`
}

type PedProgEstoqueDB struct {
	PPE_IN_SEQUENCIA *int `json:"sequencia" db:"PPE_IN_SEQUENCIA" xml:"PPE_IN_SEQUENCIA" description:"Sequencial"`
	PedidoPkDB
	Sequenciaitem         *int                             `json:"sequenciaitem" db:"ITP_IN_SEQUENCIA" xml:"ITP_IN_SEQUENCIA" description:" "`
	Sequenciaprogramacao  *int                             `json:"sequenciaprogramacao" db:"IPE_IN_SEQUENCIA" xml:"IPE_IN_SEQUENCIA" description:" "`
	ALM_TAB_IN_CODIGO     *int                             `json:"-" db:"ALM_TAB_IN_CODIGO" xml:"ALM_TAB_IN_CODIGO" description:" "`
	ALM_PAD_IN_CODIGO     *int                             `json:"-" db:"ALM_PAD_IN_CODIGO" xml:"ALM_PAD_IN_CODIGO" description:" "`
	Almoxarifado          *int                             `json:"almoxarifado" db:"ALM_IN_CODIGO" xml:"ALM_IN_CODIGO" description:" "`
	Localizacao           *int                             `json:"localizacao" db:"LOC_IN_CODIGO" xml:"LOC_IN_CODIGO" description:" "`
	Reserva               *int                             `json:"reserva" db:"MVS_IN_RESERVA" xml:"MVS_IN_RESERVA" description:"Cód.Reserva"`
	NAT_TAB_IN_CODIGO     *int                             `json:"-" db:"NAT_TAB_IN_CODIGO" xml:"NAT_TAB_IN_CODIGO" description:"Cód.Tabela Natureza Estoque"`
	NAT_PAD_IN_CODIGO     *int                             `json:"-" db:"NAT_PAD_IN_CODIGO" xml:"NAT_PAD_IN_CODIGO" description:"Padrão Natureza de Estoque"`
	Natureza              *string                          `json:"natureza" db:"NAT_ST_CODIGO" xml:"NAT_ST_CODIGO" description:"Cód.Natureza Estoque"`
	Referencia            *string                          `json:"referencia" db:"MVS_ST_REFERENCIA" xml:"MVS_ST_REFERENCIA" description:"Cód.Referência"`
	Lote                  *string                          `json:"lote" db:"MVS_ST_LOTEFORNE" xml:"MVS_ST_LOTEFORNE" description:"Nº Lote"`
	Dataentradalote       *time.Time                       `json:"dataentradalote" db:"MVS_DT_ENTRADA" xml:"MVS_DT_ENTRADA" description:"Data Entrada"`
	Datavalidadelote      *time.Time                       `json:"datavalidadelote" db:"MVS_DT_VALIDADE" xml:"MVS_DT_VALIDADE" description:"Data Validade"`
	Tiporeserva           *string                          `json:"tiporeserva" db:"IRE_CH_TIPORESERVA" xml:"IRE_CH_TIPORESERVA" description:"Tipo Reserva"`
	Quantidade            *int                             `json:"quantidade" db:"LMS_RE_QUANTIDADE" xml:"LMS_RE_QUANTIDADE" description:"Quantidade"`
	Quantidadefaturada    *float64                         `json:"quantidadefaturada" db:"IPE_RE_QTDEFATURADA" xml:"IPE_RE_QTDEFATURADA" description:" "`
	Quantidadeentregue    *float64                         `json:"quantidadeentregue" db:"IPE_RE_QTDEENTREGUE" xml:"IPE_RE_QTDEENTREGUE" description:" "`
	QuantidadesolicitadaL *float64                         `json:"quantidadesolicitada" db:"SOI_RE_QUANTIDADESOL" xml:"SOI_RE_QUANTIDADESOL" description:"Qtde.Solicitada"`
	Sequenciamatterceito  *int                             `json:"sequenciamatterceito" db:"MIE_IN_SEQMATTITDOC" xml:"MIE_IN_SEQMATTITDOC" description:"Sequencia da tabela de item nota Material terceiro"`
	Customizacoes         customizacaoservico.Customizacao `json:"customizacao" db:"-" comp:"N"`
}

type TipoDocumentoDB struct {
	TPD_TAB_IN_CODIGO *int    `json:"-" db:"TPD_TAB_IN_CODIGO" xml:"-" description:" "`
	TPD_PAD_IN_CODIGO *int    `json:"-" db:"TPD_PAD_IN_CODIGO" xml:"-" description:" "`
	Codigo            *int    `json:"codigo" db:"TPD_IN_CODIGO" xml:"TPD_IN_CODIGO" description:"Código do tipo de documento."`
	SujeitoAprovacao  *string `json:"TPD_BO_APROVACAO" db:"TPD_BO_APROVACAO" xml:"TPD_BO_APROVACAO" description:"Pedido sujeito a aprovação? (S/N)."`
}
