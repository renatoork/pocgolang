package agente

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/jinzhu/copier"
	"mega/go-util/erro"
	"mega/servico"
	"mega/servico/banco"
	"mega/servico/custom"
	"mega/servico/definicao"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Chave struct {
	Codigo int `db:"AGN_IN_CODIGO"`
	Padrao int `db:"AGN_PAD_IN_CODIGO"`
	Tabela int `db:"AGN_TAB_IN_CODIGO"`
}

type Agente struct {
	Nome       string `db:"AGN_ST_NOME"`
	Fantasia   string `db:"AGN_ST_FANTASIA"`
	TipoPessoa string `db:"AGN_CH_TIPOPESSOAFJ"`
	Chave
	Consolidador       string                 `db:"AGN_BO_CONSOLIDADOR"`
	NaturezaJuridica   int                    `db:"AGN_IN_NATJURID"`
	Pais               string                 `db:"PA_ST_SIGLA"`
	Cnpj               string                 `db:"AGN_ST_CGC"`
	CodigoPai          int                    `db:"PAI_AGN_IN_CODIGO"`
	PadraoPai          int                    `db:"PAI_AGN_PAD_IN_CODIGO"`
	TabelaPai          int                    `db:"PAI_AGN_TAB_IN_CODIGO"`
	Cnae               string                 `db:"CNAE_ST_CODIGO"`
	UsuarioInclusao    int                    `db:"USU_IN_INCLUSAO"`
	Data               time.Time              `db:"-" json:"Data"`
	Tipos              []TipoAgente           `db:"-" json:"Tipos"`
	CamposCustomizados map[string]interface{} `db:"-"`
}

type TipoAgente struct {
	Chave
	Tipo              string `db:"AGN_TAU_ST_CODIGO"`
	CodigoAlternativo string `db:"AGN_ST_CODIGOALT"`
	AtivadoSaldo      string `db:"AGN_BO_ATIVADOSALDO"`
	ExcecaoFiscal     string `db:"AGN_BO_EXCECAOFISCAL"`
}

var padrao int

func Inicializa(router *mux.Router) {

	if banco.Dbmap == nil {
		panic("Não conectado")
	}
	banco.Dbmap.AddTableWithNameAndSchema(Agente{}, "mgglo", "glo_agentes").SetKeys(false, "AGN_IN_CODIGO", "AGN_PAD_IN_CODIGO")
	banco.Dbmap.AddTableWithNameAndSchema(TipoAgente{}, "mgglo", "glo_agentes_id").SetKeys(false, "AGN_IN_CODIGO", "AGN_PAD_IN_CODIGO", "AGN_TAU_ST_CODIGO")
	var err error
	padrao, err = definicao.AchaPadrao(53, time.Now())
	if err != nil {
		panic(err)
	}
	router.HandleFunc("/api/agente/{id}", AgenteHandler).Methods("PUT", "GET")
	router.HandleFunc("/api/agente", AgenteHandler).Methods("POST", "GET")
	router.HandleFunc("/api/agente/{id}/tipo/{tipo}", TipoHandler).Methods("GET", "PUT")
	router.HandleFunc("/api/agente/{id}/tipo", TipoHandler).Methods("POST", "GET")
}

func CriaComDefaults() (Agente, error) {
	ultimoCodigo, err := banco.Dbmap.SelectInt("select max(agn_in_codigo) from mgglo.glo_agentes")

	if erro.Trata(err) {
		return Agente{}, err
	}

	proximoCodigo := int(ultimoCodigo) + 1

	chave := Chave{Codigo: proximoCodigo,
		Tabela: 53,
		Padrao: padrao}

	agente := Agente{Chave: chave,
		Consolidador:     "E",
		TipoPessoa:       "F",
		NaturezaJuridica: 2992,
		Pais:             "BRA",
		Cnpj:             "10013",
		Cnae:             "9900-8/00",
		UsuarioInclusao:  1,
		TabelaPai:        53,
		PadraoPai:        padrao,
		CodigoPai:        proximoCodigo}

	return agente, nil
}

func NewTipoAgente(chave Chave, tipo TipoAgente) TipoAgente {
	tipoAgente := TipoAgente{Chave: chave,
		CodigoAlternativo: strconv.Itoa(chave.Codigo),
		AtivadoSaldo:      "S",
		ExcecaoFiscal:     "N"}

	copier.Copy(&tipoAgente, &tipo)

	return tipoAgente
}

func AlimentaFantasia(agente *Agente) {
	if agente.Fantasia == "" {
		agente.Fantasia = strings.ToUpper(agente.Nome)
	} else {
		agente.Fantasia = strings.ToUpper(agente.Fantasia)
	}
	println(agente.Fantasia)
}

func ValidaTiposAgente(agente *Agente) error {
	if len(agente.Tipos) == 0 {
		return errors.New("Deve ser informado pelo menos um tipo de agente")
	}
	return nil
}

func AplicaRegraNegocio(agente *Agente) error {

	var err error

	AlimentaFantasia(agente)

	err = ValidaTiposAgente(agente)

	return err

}

func (agente Agente) Insert(vars servico.Params) (interface{}, error) {
	novoAgente, err := CriaComDefaults()

	if err != nil {
		return nil, err
	}

	copier.Copy(&novoAgente, &agente)

	if err := AplicaRegraNegocio(&novoAgente); err != nil {
		return nil, err
	}

	trans, err := banco.Dbmap.Begin()

	if err = trans.Insert(&novoAgente); err != nil {
		trans.Rollback()
		return nil, err
	}

	for _, tipo := range novoAgente.Tipos {
		tipoAgente := NewTipoAgente(novoAgente.Chave, tipo)
		if err = trans.Insert(&tipoAgente); err != nil {
			trans.Rollback()
			return nil, err
		}
	}

	if err = custom.GravaValores(1, novoAgente.Chave.Codigo, novoAgente.CamposCustomizados); err != nil {
		return nil, err
	}

	return &novoAgente, trans.Commit()
}

func InsereTipoAgente(chave Chave, tipo TipoAgente) error {
	return banco.Dbmap.Insert(&tipo)
}

func (agente Agente) Get(vars servico.Params) (interface{}, error) {
	id := vars["id"]

	if id != "" {
		retorno, err := banco.Dbmap.Get(Agente{}, id, padrao)
		if err != nil {
			return Agente{}, err
		}
		agenteRetorno := retorno.(*Agente)

		agenteRetorno.Tipos, err = GetTiposAgente(id, "")

		agenteRetorno.CamposCustomizados = custom.GetValores(1, agenteRetorno.Chave.Codigo)

		return agenteRetorno, nil
	} else {
		var retorno []Agente
		_, err := banco.Dbmap.Select(&retorno, `select "AGN_ST_NOME","AGN_ST_FANTASIA","AGN_CH_TIPOPESSOAFJ","AGN_IN_CODIGO","AGN_PAD_IN_CODIGO","AGN_TAB_IN_CODIGO","AGN_BO_CONSOLIDADOR","AGN_IN_NATJURID","PA_ST_SIGLA","AGN_ST_CGC","PAI_AGN_IN_CODIGO","PAI_AGN_PAD_IN_CODIGO","PAI_AGN_TAB_IN_CODIGO","CNAE_ST_CODIGO","USU_IN_INCLUSAO" from mgglo."GLO_AGENTES" where rownum <= 200`)
		if err != nil {
			return Agente{}, err
		} else {
			return retorno, nil
		}
	}
}

func (agente Agente) Update(vars servico.Params) (interface{}, error) {
	id := vars["id"]

	idInt, err := strconv.Atoi(id)

	agente.Chave.Codigo = idInt
	agente.Chave.Padrao = padrao

	AlimentaFantasia(&agente)

	count, err := banco.Dbmap.Update(&agente)
	if count == 0 {
		return nil, errors.New("Agente não encontrado")
	}
	if err = custom.GravaValores(1, agente.Chave.Codigo, agente.CamposCustomizados); err != nil {
		return nil, err
	}
	return agente, err

}

func (_ TipoAgente) Get(vars servico.Params) (interface{}, error) {
	var tipos []TipoAgente
	id := vars["id"]
	tipo := vars["tipo"]
	params := []interface{}{id, padrao}
	query := "select * from mgglo.glo_agentes_id where agn_in_codigo = :codigo and agn_pad_in_codigo = :padrao"
	if tipo != "" {
		query += " and agn_tau_st_codigo = :tipo"
		params = append(params, tipo)
	}
	_, err := banco.Dbmap.Select(&tipos, query, params...)
	println(strings.Index(err.Error(), "gorp: No fields("))
	if strings.Index(err.Error(), "gorp: No fields [") == 0 {
		err = nil
	}
	return tipos, err
}

func (tipo TipoAgente) AlimentaChave(vars servico.Params) error {
	idInt, err := strconv.Atoi(vars["id"])
	if err != nil {
		return err
	}
	tipo.Chave.Codigo = idInt
	tipo.Chave.Padrao = padrao
	tipo.Chave.Tabela = 53
	return nil
}

func (tipo TipoAgente) Insert(vars servico.Params) (interface{}, error) {
	err := tipo.AlimentaChave(vars)
	if err != nil {
		return nil, err
	}
	err = banco.Dbmap.Insert(&tipo)
	return tipo, err
}

func (tipo TipoAgente) Update(vars servico.Params) (interface{}, error) {
	tipo.AlimentaChave(vars)
	_, err := banco.Dbmap.Update(&tipo)
	return tipo, err
}

func GetTiposAgente(id string, tipo string) ([]TipoAgente, error) {
	var tipos []TipoAgente
	params := []interface{}{id, padrao}
	query := "select * from mgglo.glo_agentes_id where agn_in_codigo = :codigo and agn_pad_in_codigo = :padrao"
	if tipo != "" {
		query += " and agn_tau_st_codigo = :tipo"
		params = append(params, tipo)
	}
	_, err := banco.Dbmap.Select(&tipos, query, params...)
	println(strings.Index(err.Error(), "gorp: No fields("))
	if strings.Index(err.Error(), "gorp: No fields [") == 0 {
		err = nil
	}
	return tipos, err
}

func AgenteHandler(res http.ResponseWriter, req *http.Request) {
	var agente Agente
	servico.Handler(res, req, &agente)
}

func TipoHandler(res http.ResponseWriter, req *http.Request) {
	var tipo TipoAgente
	servico.Handler(res, req, &tipo)
}
