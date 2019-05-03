package main

type CasoTeste struct {
	Descricao  string
	Computador string
	Usuario    int
	Setup      Setup
	Entrada    Entrada
	Saida      Saida
}

type Setup struct {
	Rotina     string
	Parametros string
}

type Saida struct {
	Retorno Retorno
}
type Retorno struct {
	Pdv  string
	Item []Item
}

type Item struct {
	Valor string
	Teste string
}

type Entrada struct {
	Pedidovenda Pedidovenda
}
type Pedidovenda struct {
	Operacao            string
	Ped_in_sequencia    string
	Ped_in_codigo       string
	Ped_dt_emissao      string
	Fil_in_codigo       string
	Cli_st_tipocodigo   string
	Cli_tau_st_codigo   string
	Cli_st_codigo       string
	Tpd_in_codigo       string
	Cond_st_codigo      string
	Rep_in_codigo       string
	Ena_in_codigo       string
	Equ_in_codigo       string
	Tra_in_codigo       string
	Ped_re_percdesconto string
	Ped_in_fretepconta  string
	Ped_dt_dataentrega  string
	Ped_st_contato      string
	Ped_re_valortotal   float32
	Observacao          ObsPedido
	Parcela             []ParcFinPedido
	Itempedido          []ItemPedido
}

type ObsPedido struct {
	Operacao              string
	Ped_in_sequencia      string
	Pob_ch_tipoobservacao string
	Pob_st_observacao     string
}

type ParcFinPedido struct {
	Operacao          string
	Ped_in_sequencia  string
	Pfp_in_sequencia  string
	Mov_st_documento  string
	Mov_st_parcela    string
	Mov_dt_vencto     string
	Mov_re_valormoe   string
	Mov_re_perc       string
	Hcob_in_sequencia string
}

type ItemPedido struct {
	Operacao              string
	Ped_in_sequencia      string
	Itp_in_sequencia      string
	Pro_in_codigo         string
	Pro_st_alternativo    string
	Uni_st_unidade        string
	Itp_st_descricao      string
	Itp_re_quantidade     string
	Itp_re_valorunitario  string
	Itp_re_valormaoobra   string
	Itp_re_percdesconto   string
	Itp_re_valordesconto  string
	Itp_st_complemento    string
	Cos_in_codigo         string
	Tpr_in_codigo         string
	Tpp_in_codigo         string
	Itp_re_frete          string
	Itp_st_pedidocliente  string
	Itp_st_codprocli      string
	Itp_re_percacrescimo  string
	Itp_re_valoracrescimo string
	Itp_re_seguro         string
	Itp_re_despacessoria  string
	Itp_re_precotabELA    string
	Entrega               []ItemProgEntrega
	Observacao            []ItemObservacao
	Representante         []ItemRepresentante
}

type ItemProgEntrega struct {
	Operacao              string
	Ped_in_sequencia      string
	Itp_in_sequencia      string
	Ipe_in_sequencia      string
	Ipe_dt_dataentrega    string
	Ipe_ch_tipoentrega    string
	Ipe_ch_entregaparcial string
	Ipe_re_quantidade     string
	Ipe_ch_tipodata       string
}

type ItemObservacao struct {
	Operacao              string
	Ped_in_sequencia      string
	Itp_in_sequencia      string
	Pob_ch_tipoobservacao string
	Pob_st_observacao     string
}

type ItemRepresentante struct {
	Operacao                 string
	Ped_in_sequencia         string
	Itp_in_sequencia         string
	Rip_in_sequencia         string
	Rep_in_codigo            string
	Rip_re_basecomissao      string
	Rip_re_perccomissaocalc  string
	Rip_re_valorcomissaocalc string
}

const insert_PedidoVenda string = "insert into mgven.ven_pedidovenda_int (Operacao,Ped_in_sequencia,Ped_in_codigo,Ped_dt_emissao,Fil_in_codigo,Cli_st_tipocodigo,Cli_tau_st_codigo,Cli_st_codigo) values (?,?,?,?,?,?,?,?)"
