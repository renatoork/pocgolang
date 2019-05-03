package mgweb

import "time"

type Tipodoc struct {
	Serie   string `json:"serie" db:"TPD_SER_ST_CODIGO" xml:"TPD_SER_ST_CODIGO" description:"Serie do Tipo de Documento."`
	AltParc string `json:"altparc" db:"TPD_BO_ALTERAPARCELAS" xml:"TPD_BO_ALTERAPARCELAS" description:"Permite Alterar as parcelas do pedido."`
}

type Pedido struct {
	PedidoVenda PedidoVenda `json:"pedidovenda"`
}

type PedidoVenda struct {
	PED_IN_SEQUENCIA      *float64        `json:"sequencia" db:"PED_IN_SEQUENCIA" xml:"PED_IN_SEQUENCIA" description:"Sequencial da tabela do pedido a importar."`
	IND_IN_CODIGO         *int            `json:"indice" db:"IND_IN_CODIGO" xml:"IND_IN_CODIGO" description:"Código do Índice utilizado no pedido de venda."`
	CLI_ST_TIPOCODIGO     *string         `json:"tipocliente" db:"CLI_ST_TIPOCODIGO" xml:"CLI_ST_TIPOCODIGO" description:"Tipo do Código do cliente do Pedido."`
	FIL_IN_CODIGO         *int            `json:"filial" db:"FIL_IN_CODIGO" xml:"FIL_IN_CODIGO" description:"Código da Filial do Pedido."`
	ORG_IN_CODIGO         *int            `json:"organizacao" db:"-" xml:"ORG_IN_CODIGO" description:"Código da Organizacao do Pedido."`
	PED_IN_CODIGO         *int            `json:"numero" db:"PED_IN_CODIGO" xml:"PED_IN_CODIGO" description:"Código do Pedido de Venda"`
	ENA_IN_CODIGO         *int            `json:"enderecoentrega" db:"ENA_IN_CODIGO" xml:"ENA_IN_CODIGO" description:"Código do endereço de Entrega do pedido."`
	TPD_IN_CODIGO         *int            `json:"tipodocumento" db:"TPD_IN_CODIGO" xml:"TPD_IN_CODIGO" description:"Código do tipo de documento do pedido"`
	CLI_ST_CODIGO         *string         `json:"clientecodaux" db:"CLI_ST_CODIGO" xml:"CLI_ST_CODIGO" description:"Código do cliente Auxiliar. Aceita Codigo Mega, Alternativo, CPF, CNPJ e DEPARA do Integrador."`
	CLI_IN_CODIGO         int             `json:"cliente" db:"-" xml:"CLI_IN_CODIGO" description:"Código do Agente(cliente) do Pedido."`
	CLI_TAU_ST_CODIGO     *string         `json:"tipoagente" db:"CLI_TAU_ST_CODIGO" xml:"CLI_TAU_ST_CODIGO" description:"Tipo do Agente a importar."`
	COND_ST_CODIGO        *string         `json:"condpagto" db:"COND_ST_CODIGO" xml:"COND_ST_CODIGO" description:"Código da Condição de pagamento."`
	REP_IN_CODIGO         *int            `json:"representante" db:"REP_IN_CODIGO" xml:"REP_IN_CODIGO" description:"Código do Agente Representante."`
	PED_RE_VALORTOTAL     *float64        `json:"valortotal" db:"PED_RE_VALORTOTAL" xml:"PED_RE_VALORTOTAL" description:"Valor total do pedido"`
	UIN_IN_CODIGO         *int            `json:"usuarioinc" db:"UIN_IN_CODIGO" xml:"UIN_IN_CODIGO" description:"Código do usuário do pedido"`
	EQU_IN_CODIGO         *int            `json:"equipe" db:"EQU_IN_CODIGO" xml:"EQU_IN_CODIGO" description:"Código da Equipe do Representante."`
	UIN_DT_INCLUSAO       *time.Time      `json:"datainclusao" db:"UIN_DT_INCLUSAO" xml:"UIN_DT_INCLUSAO" description:"Data da inclusão do pedido."`
	PED_RE_VLMERCADORIA   *float64        `json:"totalmercadoria" db:"PED_RE_VLMERCADORIA" xml:"PED_RE_VLMERCADORIA" description:"Valor da Mercadoria."`
	PED_RE_MERCEMPREGADA  *float64        `json:"totalmercadoriaempregada" db:"PED_RE_MERCEMPREGADA" xml:"PED_RE_MERCEMPREGADA" description:"Valor da Mercadoria Empregada."`
	PED_RE_TOTALMAOOBRA   *float64        `json:"totalmaodeobra" db:"PED_RE_TOTALMAOOBRA" xml:"PED_RE_TOTALMAOOBRA" description:"Valor Total da Mão de Obra"`
	TRA_IN_CODIGO         *int            `json:"transportadora" db:"TRA_IN_CODIGO" xml:"TRA_IN_CODIGO" description:"Codigo do Agente Transportadora"`
	PED_RE_TOTALFRETE     *float64        `json:"totalfrete" db:"PED_RE_TOTALFRETE" xml:"PED_RE_TOTALFRETE" description:"Valor Total do Frete"`
	PED_RE_PERCACRESCIMO  *float64        `json:"percentualacrescimo" db:"PED_RE_PERCACRESCIMO" xml:"PED_RE_PERCACRESCIMO" description:"% do Acréscimo do Pedido"`
	PED_RE_TOTALACRESCIMO *float64        `json:"totalacrescimo" db:"PED_RE_TOTALACRESCIMO" xml:"PED_RE_TOTALACRESCIMO" description:"Valor total do acréscimo."`
	PED_RE_PERCDESCONTO   *float64        `json:"percentualdesconto" db:"PED_RE_PERCDESCONTO" xml:"PED_RE_PERCDESCONTO" description:"% de desconto do Pedido"`
	PED_RE_TOTALDESCONTO  *float64        `json:"totaldesconto" db:"PED_RE_TOTALDESCONTO" xml:"PED_RE_TOTALDESCONTO" description:"Valor total do desconto do pedido"`
	PED_ST_NOMECLIENTE    *string         `json:"nomecliente" db:"PED_ST_NOMECLIENTE" xml:"PED_ST_NOMECLIENTE" description:"Nome do cliente."`
	PED_BO_ERRO           *string         `json:"erro" db:"PED_BO_ERRO" xml:"PED_BO_ERRO" description:"Informa se houve algum erro na importação"`
	PED_BO_TEXTOERRO      *string         `json:"textoerro" db:"PED_BO_TEXTOERRO" xml:"PED_BO_TEXTOERRO" description:"Mensagem do erro ocorrido na importação"`
	ACAO_IN_CODIGO        *int            `json:"acao" db:"ACAO_IN_CODIGO" xml:"ACAO_IN_CODIGO" description:"Código de Ação do Mega2000"`
	PED_DT_EMISSAO        *time.Time      `json:"emissao" db:"PED_DT_EMISSAO" xml:"PED_DT_EMISSAO" description:"Data de Emissão do Pedido."`
	PED_CH_SITUACAO       *string         `json:"situacao" db:"PED_CH_SITUACAO" xml:"PED_CH_SITUACAO" description:"Situação"`
	PE_ORG_TAB_IN_CODIGO  *int            `json:"-" db:"PE_ORG_TAB_IN_CODIGO" xml:"PE_ORG_TAB_IN_CODIGO" description:"Esta coluna não precisa ser informado."`
	PE_ORG_PAD_IN_CODIGO  *int            `json:"-" db:"PE_ORG_PAD_IN_CODIGO" xml:"PE_ORG_PAD_IN_CODIGO" description:"Esta coluna não precisa ser informado."`
	PE_ORG_IN_CODIGO      *int            `json:"pkorganizacao" db:"PE_ORG_IN_CODIGO" xml:"PE_ORG_IN_CODIGO" description:"Esta coluna não precisa ser informado."`
	PE_ORG_TAU_ST_CODIGO  *string         `json:"-" db:"PE_ORG_TAU_ST_CODIGO" xml:"PE_ORG_TAU_ST_CODIGO" description:"Esta coluna não precisa ser informado."`
	PE_SER_ST_CODIGO      *string         `json:"pkserie" db:"PE_SER_ST_CODIGO" xml:"PE_SER_ST_CODIGO" description:"Esta coluna não precisa ser informado."`
	PE_PED_IN_CODIGO      *int            `json:"pkcodigo" db:"PE_PED_IN_CODIGO" xml:"PE_PED_IN_CODIGO" description:"Esta coluna não precisa ser informado."`
	GRU_IN_CODIGO         *int            `json:"usuario" db:"GRU_IN_CODIGO" xml:"GRU_IN_CODIGO" description:" "`
	COMP_ST_NOME          *string         `json:"computador" db:"COMP_ST_NOME" xml:"COMP_ST_NOME" description:" "`
	RED_IN_CODIGO         *int            `json:"redespacho" db:"RED_IN_CODIGO" xml:"RED_IN_CODIGO" description:"Código da transportadora do Redespacho."`
	PED_CH_STATUSIMP      *string         `json:"statusimportacao" db:"PED_CH_STATUSIMP" xml:"PED_CH_STATUSIMP" description:"Status do pedido no momento da importação."`
	SER_ST_CODIGO         *string         `json:"serie" db:"SER_ST_CODIGO" xml:"SER_ST_CODIGO" description:"Série do documento."`
	PED_IN_FRETEPCONTA    *int            `json:"freteporconta" db:"PED_IN_FRETEPCONTA" xml:"PED_IN_FRETEPCONTA" description:"Frete por Conta"`
	PED_RE_VALORCAUCAO    *float64        `json:"valorcaucao" db:"PED_RE_VALORCAUCAO" xml:"PED_RE_VALORCAUCAO" description:" "`
	PED_RE_PERCCAUCAO     *float64        `json:"percentualcaucao" db:"PED_RE_PERCCAUCAO" xml:"PED_RE_PERCCAUCAO" description:" "`
	PED_ST_FRETEEMBDEST   *string         `json:"freteembutidodestacado" db:"PED_ST_FRETEEMBDEST" xml:"PED_ST_FRETEEMBDEST" description:" "`
	PED_RE_TOTALSEGURO    *float64        `json:"totalseguro" db:"PED_RE_TOTALSEGURO" xml:"PED_RE_TOTALSEGURO" description:" "`
	PED_RE_TOTALDESPACESS *float64        `json:"totaldespacessoria" db:"PED_RE_TOTALDESPACESS" xml:"PED_RE_TOTALDESPACESS" description:" "`
	PED_CH_ROTINAIMP      *string         `json:"rotinaimportacao" db:"PED_CH_ROTINAIMP" xml:"PED_CH_ROTINAIMP" description:"Rotina de Importação de Pedido"`
	PED_DT_DATAENTREGA    *time.Time      `json:"entrega" db:"PED_DT_DATAENTREGA" xml:"PED_DT_DATAENTREGA" description:" "`
	CAR_IN_CODIGO         *int            `json:"cargo" db:"CAR_IN_CODIGO" xml:"CAR_IN_CODIGO" description:" "`
	PED_ST_CONTATO        *string         `json:"contato" db:"PED_ST_CONTATO" xml:"PED_ST_CONTATO" description:" "`
	PED_ST_CONTATOFONE    *string         `json:"fonecontato" db:"PED_ST_CONTATOFONE" xml:"PED_ST_CONTATOFONE" description:" "`
	PED_BO_ESPECIAL       *string         `json:"especial" db:"PED_BO_ESPECIAL" xml:"PED_BO_ESPECIAL" description:"Fab. Sob. Encomenda"`
	CCF_IDE_ST_CODIGO     *string         `json:"ccustotipo" db:"CCF_IDE_ST_CODIGO" xml:"CCF_IDE_ST_CODIGO" description:"Cód. Identificador"`
	CCF_IN_REDUZIDO       *int            `json:"ccusto" db:"CCF_IN_REDUZIDO" xml:"CCF_IN_REDUZIDO" description:"Centro Custo"`
	PROJ_IDE_ST_CODIGO    *string         `json:"projetotipo" db:"PROJ_IDE_ST_CODIGO" xml:"PROJ_IDE_ST_CODIGO" description:"Cód. Identificador"`
	PROJ_IN_REDUZIDO      *int            `json:"projeto" db:"PROJ_IN_REDUZIDO" xml:"PROJ_IN_REDUZIDO" description:"Projeto"`
	PED_IN_INDPRES        *int            `json:"indicadorpresenca" db:"PED_IN_INDPRES" xml:"PED_IN_INDPRES" description:"Indicador de presença"`
	PED_BO_INDFINAL       *string         `json:"indicadorconsumidorfinal" db:"PED_BO_INDFINAL" xml:"PED_BO_INDFINAL" description:"Indicador de Consumidor Final"`
	PED_IN_PRIORIDADE     *int            `json:"prioridade" db:"PED_IN_PRIORIDADE" xml:"PED_IN_PRIORIDADE" description:"Prioridade de importação dos pedidos"`
	Observacao            []ObsPedido     `json:"observacoes,omitempty" db:"-" description:"Observação do pedido"`
	Parcela               []ParcFinPedido `json:"parcelas,omitempty" db:"-" description:"Parcelas financeiras"`
	Item                  []Item          `json:"itens,required" db:"-" description:"Itens do Pedido"`
}

type Item struct {
	PED_IN_SEQUENCIA       *int             `json:"-" db:"PED_IN_SEQUENCIA" xml:"PED_IN_SEQUENCIA" description:"Sequencial do pedido"`
	ITP_ST_COMPLEMENTO     string           `json:"complemento" db:"ITP_ST_COMPLEMENTO" xml:"ITP_ST_COMPLEMENTO" description:"Texto do complemento do item de estoque."`
	ITP_IN_SEQUENCIA       *int             `json:"sequencia" db:"ITP_IN_SEQUENCIA" xml:"ITP_IN_SEQUENCIA" description:"Sequencial do item do pedido"`
	APL_IN_CODIGO          *int             `json:"aplicacao" db:"APL_IN_CODIGO" xml:"APL_IN_CODIGO" description:"Código da aplicação da venda"`
	PRO_ST_ALTERNATIVO     string           `json:"produtoalternativo" db:"PRO_ST_ALTERNATIVO" xml:"PRO_ST_ALTERNATIVO" description:"Código alternativo do item de estoque."`
	COS_IN_CODIGO          *int             `json:"servico" db:"COS_IN_CODIGO" xml:"COS_IN_CODIGO" description:"Código do tipo de Serviço."`
	PROJ_IDE_ST_CODIGO     string           `json:"projetotipo" db:"PROJ_IDE_ST_CODIGO" xml:"PROJ_IDE_ST_CODIGO" description:"Código identificador do projeto do item da venda."`
	PROJ_IN_REDUZIDO       *int             `json:"projeto" db:"PROJ_IN_REDUZIDO" xml:"PROJ_IN_REDUZIDO" description:"Código do projeto do item da venda."`
	TPR_IN_CODIGO          *int             `json:"tabelapreco" db:"TPR_IN_CODIGO" xml:"TPR_IN_CODIGO" description:"Código da tabela de preço."`
	TPP_IN_CODIGO          *int             `json:"tipopreco" db:"TPP_IN_CODIGO" xml:"TPP_IN_CODIGO" description:"Código do Tipo de preço (tabela de preço)."`
	CUS_IDE_ST_CODIGO      string           `json:"ccustotipo" db:"CUS_IDE_ST_CODIGO" xml:"CUS_IDE_ST_CODIGO" description:"Código identificador do centro de custo do item da venda."`
	CUS_IN_REDUZIDO        *int             `json:"ccusto" db:"CUS_IN_REDUZIDO" xml:"CUS_IN_REDUZIDO" description:"Código do projeto do item da venda."`
	ITP_CH_FRETEPCONTA     string           `json:"freteporconta" db:"ITP_CH_FRETEPCONTA" xml:"ITP_CH_FRETEPCONTA" description:"Tipo do Frete"`
	PRO_IN_CODIGO          *int             `json:"produto" db:"PRO_IN_CODIGO" xml:"PRO_IN_CODIGO" description:"Código Mega do item de estoque."`
	UIN_IN_CODIGO          *int             `json:"usuarioinclusao" db:"UIN_IN_CODIGO" xml:"UIN_IN_CODIGO" description:"Código do usuário"`
	UIN_DT_INCLUSAO        *time.Time       `json:"datainclusao" db:"UIN_DT_INCLUSAO" xml:"UIN_DT_INCLUSAO" description:"Data da inclusão."`
	FMT_ST_CODIGO          string           `json:"formato" db:"FMT_ST_CODIGO" xml:"FMT_ST_CODIGO" description:"Código do formato da conversão de unidade"`
	EMB_IN_CODIGO          *int             `json:"embalagem" db:"EMB_IN_CODIGO" xml:"EMB_IN_CODIGO" description:"Código da embalagem do item."`
	ALM_IN_CODIGO          *int             `json:"almoxarifado" db:"ALM_IN_CODIGO" xml:"ALM_IN_CODIGO" description:"Código do almoxarifado"`
	UNI_ST_UNIDADE         string           `json:"unidade" db:"UNI_ST_UNIDADE" xml:"UNI_ST_UNIDADE" description:"Código da Unidade de venda."`
	LOC_IN_CODIGO          *int             `json:"localizacao" db:"LOC_IN_CODIGO" xml:"LOC_IN_CODIGO" description:"Código do município de tributação do serviço."`
	ITP_ST_DESCRICAO       string           `json:"descricao" db:"ITP_ST_DESCRICAO" xml:"ITP_ST_DESCRICAO" description:"Descrição do item de estoque"`
	EMB_UNI_ST_UNIDADE     string           `json:"unidadeembalagem" db:"EMB_UNI_ST_UNIDADE" xml:"EMB_UNI_ST_UNIDADE" description:"Código da unidade da embalagem do item."`
	EMB_FMT_ST_CODIGO      string           `json:"formatoembalagem" db:"EMB_FMT_ST_CODIGO" xml:"EMB_FMT_ST_CODIGO" description:"Código do formato da embalagem."`
	ITP_RE_QUANTIDADE      *float64         `json:"quantidade" db:"ITP_RE_QUANTIDADE" xml:"ITP_RE_QUANTIDADE" description:"Quantidade da venda"`
	ITP_RE_QTDEAUX         *float64         `json:"quantidadeauxiliar" db:"ITP_RE_QTDEAUX" xml:"ITP_RE_QTDEAUX" description:"Campo não utilizado."`
	ITP_RE_VALORUNITARIO   *float64         `json:"valorunitario" db:"ITP_RE_VALORUNITARIO" xml:"ITP_RE_VALORUNITARIO" description:"Valor Unitário da Venda"`
	ITP_RE_LARGURA         *float64         `json:"largura" db:"ITP_RE_LARGURA" xml:"ITP_RE_LARGURA" description:"Campo não utilizado"`
	ITP_RE_COMPRIMENTO     *float64         `json:"comprimento" db:"ITP_RE_COMPRIMENTO" xml:"ITP_RE_COMPRIMENTO" description:"Campo não utilizado"`
	ITP_RE_VALORMERCADORIA *float64         `json:"valormercadoria" db:"ITP_RE_VALORMERCADORIA" xml:"ITP_RE_VALORMERCADORIA" description:"Valor das mercadorias"`
	ITP_RE_MEDIDAREAL      *float64         `json:"medidareal" db:"ITP_RE_MEDIDAREAL" xml:"ITP_RE_MEDIDAREAL" description:"Campo não utilizado"`
	ITP_RE_VALORMERCEMPREG *float64         `json:"valormercadoriaempregada" db:"ITP_RE_VALORMERCEMPREG" xml:"ITP_RE_VALORMERCEMPREG" description:"Valor das Mercadorias empregadas."`
	ITP_RE_VALORMAOOBRA    *float64         `json:"valormaodeobra" db:"ITP_RE_VALORMAOOBRA" xml:"ITP_RE_VALORMAOOBRA" description:"Valor da Mão de Obra"`
	ITP_RE_VALORTOTAL      *float64         `json:"valortotal" db:"ITP_RE_VALORTOTAL" xml:"ITP_RE_VALORTOTAL" description:"Valor Total"`
	ITP_RE_FRETE           *float64         `json:"frete" db:"ITP_RE_FRETE" xml:"ITP_RE_FRETE" description:"Valor Frete do item da venda"`
	ITP_ST_PEDIDOCLIENTE   string           `json:"codigopedidocliente" db:"ITP_ST_PEDIDOCLIENTE" xml:"ITP_ST_PEDIDOCLIENTE" description:"Número do pedido do cliente"`
	ITP_ST_CODPROCLI       string           `json:"codigoprodutocliente" db:"ITP_ST_CODPROCLI" xml:"ITP_ST_CODPROCLI" description:"Código do item da venda para o cliente."`
	ITP_RE_PERCDESCONTO    *float64         `json:"percentualdesconto" db:"ITP_RE_PERCDESCONTO" xml:"ITP_RE_PERCDESCONTO" description:"% de desconto do item da venda"`
	ITP_RE_VALORDESCONTO   *float64         `json:"valordesconto" db:"ITP_RE_VALORDESCONTO" xml:"ITP_RE_VALORDESCONTO" description:"Valor do desconto do item da venda"`
	ITP_RE_PERCACRESCIMO   *float64         `json:"percentualacrescimo" db:"ITP_RE_PERCACRESCIMO" xml:"ITP_RE_PERCACRESCIMO" description:"% de acréscimo do item da venda."`
	ITP_RE_VALORACRESCIMO  *float64         `json:"valoracrescimo" db:"ITP_RE_VALORACRESCIMO" xml:"ITP_RE_VALORACRESCIMO" description:"Valor do acréscimo do item da venda."`
	TPC_ST_CLASSE          string           `json:"classe" db:"TPC_ST_CLASSE" xml:"TPC_ST_CLASSE" description:"Tipo da classe na aplicação."`
	ITP_CH_STATUSIMP       string           `json:"statusimportacao" db:"ITP_CH_STATUSIMP" xml:"ITP_CH_STATUSIMP" description:"Status a importar(S/N)"`
	ITP_RE_VALORCAUCAO     *float64         `json:"valorcaucao-" db:"ITP_RE_VALORCAUCAO" xml:"ITP_RE_VALORCAUCAO" description:" "`
	ITP_RE_PERCCAUCAO      *float64         `json:"percentualcaucai" db:"ITP_RE_PERCCAUCAO" xml:"ITP_RE_PERCCAUCAO" description:" "`
	ITP_RE_ALIQIPI         *float64         `json:"aliquotaipi" db:"ITP_RE_ALIQIPI" xml:"ITP_RE_ALIQIPI" description:" "`
	ITP_RE_BASEIPI         *float64         `json:"baseipi" db:"ITP_RE_BASEIPI" xml:"ITP_RE_BASEIPI" description:" "`
	ITP_RE_ICMSRETIDO      *float64         `json:"icmsretido" db:"ITP_RE_ICMSRETIDO" xml:"ITP_RE_ICMSRETIDO" description:" "`
	ITP_RE_BASESUBTRIB     *float64         `json:"basesubstituicaotributaria" db:"ITP_RE_BASESUBTRIB" xml:"ITP_RE_BASESUBTRIB" description:" "`
	ITP_RE_ALIQICMS        *float64         `json:"aliquotaicms" db:"ITP_RE_ALIQICMS" xml:"ITP_RE_ALIQICMS" description:" "`
	ITP_RE_VALORICMS       *float64         `json:"valoricms" db:"ITP_RE_VALORICMS" xml:"ITP_RE_VALORICMS" description:" "`
	ITP_RE_BASEICMS        *float64         `json:"baseicms" db:"ITP_RE_BASEICMS" xml:"ITP_RE_BASEICMS" description:" "`
	ITP_RE_PRECOTABELA     *float64         `json:"precotabela" db:"ITP_RE_PRECOTABELA" xml:"ITP_RE_PRECOTABELA" description:" "`
	CFOP_ST_EXTENSO        string           `json:"cfop" db:"CFOP_ST_EXTENSO" xml:"CFOP_ST_EXTENSO" description:" "`
	ITP_RE_SEGURO          *float64         `json:"seguro" db:"ITP_RE_SEGURO" xml:"ITP_RE_SEGURO" description:" "`
	ITP_RE_DESPACESSORIA   *float64         `json:"despesaacessoria" db:"ITP_RE_DESPACESSORIA" xml:"ITP_RE_DESPACESSORIA" description:" "`
	ITP_RE_PERCREDIPI      *float64         `json:"percentualdescontoipi" db:"ITP_RE_PERCREDIPI" xml:"ITP_RE_PERCREDIPI" description:" "`
	ITP_RE_VALORIPI        *float64         `json:"valoripi" db:"ITP_RE_VALORIPI" xml:"ITP_RE_VALORIPI" description:" "`
	ITP_RE_PERCDESCDIG     *float64         `json:"percentualdescontodigitado" db:"ITP_RE_PERCDESCDIG" xml:"ITP_RE_PERCDESCDIG" description:" "`
	ITP_RE_VALDESCUNITARIO *float64         `json:"valordescontounitario" db:"ITP_RE_VALDESCUNITARIO" xml:"ITP_RE_VALDESCUNITARIO" description:" "`
	ITP_RE_PERDESCUNITARIO *float64         `json:"percentualdescontounitario" db:"ITP_RE_PERDESCUNITARIO" xml:"ITP_RE_PERDESCUNITARIO" description:" "`
	ITP_BO_ESPECIAL        string           `json:"especial" db:"ITP_BO_ESPECIAL" xml:"ITP_BO_ESPECIAL" description:"Fab. Sob. Encomenda"`
	CPS_IN_CODIGO          *int             `json:"composicao" db:"CPS_IN_CODIGO" xml:"CPS_IN_CODIGO" description:"Composição Produto Composição do Produto"`
	ObsItem                []ObsItem        `json:"obsitem,omitempty" db:"-" xml:"obsitem>obs" description:"Observação do Item"`
	PedProgEntrega         []PedProgEntrega `json:"pedprogentregas,omitempty" db:"-" xml:"progentrega>entrega" description:"Programação de Entrega do Item"`
}

type ObsPedido struct {
	POB_CH_TIPOOBSERVACAO string `json:"tipoobservacao" db:"POB_CH_TIPOOBSERVACAO" xml:"POB_CH_TIPOOBSERVACAO" description:"Tipo Observação"`
	PED_IN_SEQUENCIA      *int   `json:"-" db:"PED_IN_SEQUENCIA" xml:"PED_IN_SEQUENCIA" description:"Sequencia para importação do Pedido."`
	POB_ST_OBSERVACAO     string `json:"observacao" db:"POB_ST_OBSERVACAO" xml:"POB_ST_OBSERVACAO" description:"Texto da Observação do Item do Pedido."`
}

type ObsItem struct {
	PED_IN_SEQUENCIA      *int   `json:"-" db:"PED_IN_SEQUENCIA" xml:"PED_IN_SEQUENCIA" description:"Sequencia para importação do Pedido."`
	ITP_IN_SEQUENCIA      *int   `json:"-" db:"ITP_IN_SEQUENCIA" xml:"ITP_IN_SEQUENCIA" description:"Número da Sequência do Produto no Pedido"`
	OIP_CH_TIPOOBSERVACAO string `json:"tipoobservacao" db:"OIP_CH_TIPOOBSERVACAO" xml:"OIP_CH_TIPOOBSERVACAO" description:"Tipo da Observação do Item do Pedido."`
	ITO_ST_OBSERVACAO     string `json:"observacao" db:"ITO_ST_OBSERVACAO" xml:"ITO_ST_OBSERVACAO" description:"Texto da Observação do Item do Pedido."`
}

type ParcFinPedido struct {
	PED_IN_SEQUENCIA      *int       `json:"-" db:"PED_IN_SEQUENCIA" xml:"PED_IN_SEQUENCIA" description:" "`
	PFP_IN_SEQUENCIA      *int       `json:"sequencia" db:"PFP_IN_SEQUENCIA" xml:"PFP_IN_SEQUENCIA" description:" "`
	PFP_ST_DOCUMENTO      string     `json:"documento" db:"PFP_ST_DOCUMENTO" xml:"PFP_ST_DOCUMENTO" description:" "`
	PFP_ST_PARCELA        string     `json:"parcela" db:"PFP_ST_PARCELA" xml:"PFP_ST_PARCELA" description:" "`
	PFP_DT_VENCTO         *time.Time `json:"vencimento" db:"PFP_DT_VENCTO" xml:"PFP_DT_VENCTO" description:" "`
	PFP_RE_VALORMOE       *float64   `json:"valormoeda" db:"PFP_RE_VALORMOE" xml:"PFP_RE_VALORMOE" description:" "`
	PFP_RE_PERC           *float64   `json:"percentualparcela" db:"PFP_RE_PERC" xml:"PFP_RE_PERC" description:" "`
	PFP_HCOB_IN_SEQUENCIA *int       `json:"tipocobranca" db:"PFP_HCOB_IN_SEQUENCIA" xml:"PFP_HCOB_IN_SEQUENCIA" description:" "`
}

type PedProgEntrega struct {
	PED_IN_SEQUENCIA      *int             `json:"-" db:"PED_IN_SEQUENCIA" xml:"PED_IN_SEQUENCIA" description:"Sequencia para importação do Pedido."`
	ITP_IN_SEQUENCIA      *int             `json:"-" db:"ITP_IN_SEQUENCIA" xml:"ITP_IN_SEQUENCIA" description:"Número da Sequência do Item no Pedido"`
	IPE_IN_SEQUENCIA      *int             `json:"sequencia" db:"IPE_IN_SEQUENCIA" xml:"IPE_IN_SEQUENCIA" description:"Sequencia da Programação de Entrega."`
	UNI_ST_UNIDADE        string           `json:"unidade" db:"UNI_ST_UNIDADE" xml:"UNI_ST_UNIDADE" description:"Código Mega da Unidade do Produto"`
	CLI_ST_CODIGO         *int             `json:"cliente" db:"CLI_ST_CODIGO" xml:"CLI_ST_CODIGO" description:"Código do Agente da programação de entrega, pode ser diferente do agente definido no pedido. Pode ser:  Codigo Mega, Código Alternativo Mega, CGC ou CPF."`
	ENA_IN_CODIGO         *int             `json:"enderecoentrega-" db:"ENA_IN_CODIGO" xml:"ENA_IN_CODIGO" description:"Código do endereço do Agente da programação de entrega."`
	CLI_ST_TIPOCODIGO     string           `json:"tipocodigo" db:"CLI_ST_TIPOCODIGO" xml:"CLI_ST_TIPOCODIGO" description:"Determina o tipo do código do cliente da programação de entrega. Pode ser: C - código Mega, A - código Alternativo Mega, G - CGC, F - CPF"`
	CLI_TAU_ST_CODIGO     string           `json:"tipocliente" db:"CLI_TAU_ST_CODIGO" xml:"CLI_TAU_ST_CODIGO" description:"Código da tabela auxiliar do Agente da programação de entrega: C - Cliente, R - Representante, T - Transportadora, C - Colaborador, F - Fornecedor, O - Outros."`
	IPE_DT_DATAENTREGA    *time.Time       `json:"dataentrega" db:"IPE_DT_DATAENTREGA" xml:"IPE_DT_DATAENTREGA" description:"Data da Entrega."`
	IPE_CH_TIPOENTREGA    string           `json:"tipoentrega" db:"IPE_CH_TIPOENTREGA" xml:"IPE_CH_TIPOENTREGA" description:"Tipo da Entrega."`
	IPE_CH_ENTREGAPARCIAL string           `json:"entregaparcial" db:"IPE_CH_ENTREGAPARCIAL" xml:"IPE_CH_ENTREGAPARCIAL" description:"Programação de entrega permite entrega parcial para este item."`
	IPE_ST_NUMEROORDEM    string           `json:"numerooe" db:"IPE_ST_NUMEROORDEM" xml:"IPE_ST_NUMEROORDEM" description:"Número da Ordem de Entrega."`
	IPE_RE_QUANTIDADE     *float64         `json:"quantidade" db:"IPE_RE_QUANTIDADE" xml:"IPE_RE_QUANTIDADE" description:"Quantidade Rateada para a Entrega do Item do Pedido."`
	IPE_DT_DATAEMISSAO    *time.Time       `json:"dataemissao" db:"IPE_DT_DATAEMISSAO" xml:"IPE_DT_DATAEMISSAO" description:"Data da Emissão da Ordem de Entrega."`
	UIN_IN_CODIGO         *int             `json:"usuarioinclusao" db:"UIN_IN_CODIGO" xml:"UIN_IN_CODIGO" description:"Código Mega do usuários da inclusão da programação."`
	UIN_DT_INCLUSAO       *time.Time       `json:"datainclusao" db:"UIN_DT_INCLUSAO" xml:"UIN_DT_INCLUSAO" description:"Data da  inclusão do Item da Prog Entrega."`
	IPE_CH_TIPODATA       string           `json:"tipodata" db:"IPE_CH_TIPODATA" xml:"IPE_CH_TIPODATA" description:"Tipo Data Entrega"`
	IPE_DT_DATAORIGINAL   *time.Time       `json:"dataoriginal" db:"IPE_DT_DATAORIGINAL" xml:"IPE_DT_DATAORIGINAL" description:"Data original da programação de entrega do item de pedido."`
	IPE_DT_DATAPLANEJADA  *time.Time       `json:"dataplanejada" db:"IPE_DT_DATAPLANEJADA" xml:"IPE_DT_DATAPLANEJADA" description:"Data planejada da programação de entrega do item de pedido."`
	FMT_ST_CODIGO         string           `json:"formatoconversao" db:"FMT_ST_CODIGO" xml:"FMT_ST_CODIGO" description:"Cód. Conversor"`
	EMB_IN_CODIGO         *int             `json:"embalagem" db:"EMB_IN_CODIGO" xml:"EMB_IN_CODIGO" description:"Embalagem"`
	PedProgEstoque        []PedProgEstoque `json:"pedprogestoque,omitempty" db:"-" xml:"progestoque>estoque" descricao:"Dados de estoque da programação"`
}

type PedProgEstoque struct {
	PED_IN_SEQUENCIA  *int       `json:"-" db:"PED_IN_SEQUENCIA" xml:"PED_IN_SEQUENCIA" description:"Sequencia para importação do Pedido."`
	ITP_IN_SEQUENCIA  *int       `json:"-" db:"ITP_IN_SEQUENCIA" xml:"ITP_IN_SEQUENCIA" description:"Número da Sequência do Item no Pedido"`
	IPE_IN_SEQUENCIA  *int       `json:"-" db:"IPE_IN_SEQUENCIA" xml:"IPE_IN_SEQUENCIA" description:"Sequencia da Programação de Entrega."`
	IPPE_IN_SEQUENCIA *int       `json:"sequencia" db:"IPPE_IN_SEQUENCIA" xml:"IPPE_IN_SEQUENCIA" description:"Sequencial"`
	ALM_IN_CODIGO     *int       `json:"almoxarifado" db:"ALM_IN_CODIGO" xml:"ALM_IN_CODIGO" description:" "`
	LOC_IN_CODIGO     *int       `json:"localizacao" db:"LOC_IN_CODIGO" xml:"LOC_IN_CODIGO" description:" "`
	NAT_ST_CODIGO     string     `json:"natureza" db:"NAT_ST_CODIGO" xml:"NAT_ST_CODIGO" description:"Cód.Natureza Estoque"`
	MVS_ST_REFERENCIA string     `json:"referencia" db:"MVS_ST_REFERENCIA" xml:"MVS_ST_REFERENCIA" description:"Cód.Referência"`
	MVS_ST_LOTEFORNE  string     `json:"lotefornecedor" db:"MVS_ST_LOTEFORNE" xml:"MVS_ST_LOTEFORNE" description:"Nº Lote"`
	MVS_DT_ENTRADA    *time.Time `json:"entrada" db:"MVS_DT_ENTRADA" xml:"MVS_DT_ENTRADA" description:"Data Entrada"`
	MVS_DT_VALIDADE   *time.Time `json:"validade" db:"MVS_DT_VALIDADE" xml:"MVS_DT_VALIDADE" description:"Data Validade"`
	LMS_RE_QUANTIDADE *int       `json:"quantidade" db:"LMS_RE_QUANTIDADE" xml:"LMS_RE_QUANTIDADE" description:"Quantidade"`
}

type RepresentantePedido struct {
	PED_IN_SEQUENCIA     *int     `json:"-" db:"PED_IN_SEQUENCIA" xml:"PED_IN_SEQUENCIA" description:" "`
	RPE_IN_SEQUENCIA     *int     `json:"-" db:"RPE_IN_SEQUENCIA" xml:"RPE_IN_SEQUENCIA" description:" "`
	REP_IN_CODIGO        *int     `json:"-" db:"REP_IN_CODIGO" xml:"REP_IN_CODIGO" description:" "`
	EQU_IN_CODIGO        *int     `json:"-" db:"EQU_IN_CODIGO" xml:"EQU_IN_CODIGO" description:" "`
	RPE_RE_VALORCOMISSAO *float64 `json:"-" db:"RPE_RE_VALORCOMISSAO" xml:"RPE_RE_VALORCOMISSAO" description:" "`
	RPE_RE_BASECOMISSAO  *float64 `json:"-" db:"RPE_RE_BASECOMISSAO" xml:"RPE_RE_BASECOMISSAO" description:" "`
	RPE_RE_PERCCOMISSAO  *float64 `json:"-" db:"RPE_RE_PERCCOMISSAO" xml:"RPE_RE_PERCCOMISSAO" description:" "`
}

type RepresentanteItem struct {
	PED_IN_SEQUENCIA         *int     `json:"-" db:"PED_IN_SEQUENCIA" xml:"PED_IN_SEQUENCIA" description:" "`
	ITP_IN_SEQUENCIA         *int     `json:"-" db:"ITP_IN_SEQUENCIA" xml:"ITP_IN_SEQUENCIA" description:" "`
	RIP_IN_SEQUENCIA         *int     `json:"-" db:"RIP_IN_SEQUENCIA" xml:"RIP_IN_SEQUENCIA" description:" "`
	REP_IN_CODIGO            *int     `json:"-" db:"REP_IN_CODIGO" xml:"REP_IN_CODIGO" description:" "`
	RIP_RE_BASECOMISSAO      *float64 `json:"-" db:"RIP_RE_BASECOMISSAO" xml:"RIP_RE_BASECOMISSAO" description:" "`
	RIP_RE_PERCCOMISSAOCALC  *float64 `json:"-" db:"RIP_RE_PERCCOMISSAOCALC" xml:"RIP_RE_PERCCOMISSAOCALC" description:" "`
	RIP_RE_VALORCOMISSAOCALC *float64 `json:"-" db:"RIP_RE_VALORCOMISSAOCALC" xml:"RIP_RE_VALORCOMISSAOCALC" description:" "`
}
