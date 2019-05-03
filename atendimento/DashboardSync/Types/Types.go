package Types

import (
	"database/sql"
	"strings"
	"time"
)

type CustomTime time.Time

func (t *CustomTime) UnmarshalJSON(b []byte) error {
	timeStr := string(b)
	if timeStr != "" {
		x := strings.Split(strings.Trim(timeStr, `"`), "Z")
		ts, err := time.Parse("2006-01-02T15:04:05", x[0])
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
		return []byte(`"` + tt.Format("02/01/2006 15:04") + `"`), nil
	} else {
		return []byte(`"01/01/2000"`), nil
	}
}

type DbManager struct {
	UserDb string
	PassDb string
	UrlDb  string

	Conn   *sql.DB
	Stmt   *sql.Stmt
	TxConn *sql.Tx
	Teste  string
}

type TicketComments struct {
	Comments []struct {
		Id            int        `json:"id"`
		Type          string     `json:"type"`
		AuthorId      int        `json:"author_id"`
		Body          string     `json:"body"`
		HtmlBody      string     `json:"html_body"`
		Public        bool       `json:"public"`
		DataCreatedAt CustomTime `json:"created_at"`
		CreatedAt     time.Time  `json:"-"`
	}
	NextPage string `json:"next_page"`
	Count    int    `josn:"count"`
	Error    string `json:"error"`
}

type IssueTicket struct {
	Links []struct {
		IssueId string `json:"issue_id"`
	} `json:"links"`
}

type CustomValue string

type CustomFields struct {
	Id    int         `json:"id"`
	Value CustomValue `json:"value"`
}

type SatisfactionRating struct {
	Id      int    `json:"id"`
	Score   string `json:"score"`
	Comment string `json:"comment"`
}

type Group struct {
	Id            int        `json:"id"`
	Url           string     `json:"url"`
	Name          string     `json:"name"`
	Deleted       bool       `json:"deleted"`
	DataCreatedAt CustomTime `json:"created_at"`
	CreatedAt     time.Time  `json:"-"`
	DataUpdatedAt CustomTime `json:"updated_at"`
	UpdatedAt     time.Time  `json:"-"`
}

type User struct {
	Id                   int        `json:"id"`
	Url                  string     `json:"url"`
	Name                 string     `json:"name"`
	ExternalID           string     `json:"external_id"`
	Alias                string     `json:"alias"`
	DataCreatedAt        CustomTime `json:"created_at"`
	CreatedAt            time.Time  `json:"-"`
	DataUpdatedAt        CustomTime `json:"updated_at"`
	UpdatedAt            time.Time  `json:"-"`
	Active               bool       `json:"active"`
	Verified             bool       `json:"verified"`
	Shared               bool       `json:"shared"`
	SharedAgent          bool       `json:"shared_agent"`
	Locale               string     `json:"locale"`
	LocaleID             int        `json:"locale_id"`
	TimeZone             string     `json:"time_zone"`
	DataLastLoginAt      CustomTime `json:"last_login_at"`
	LastLoginAt          time.Time  `json:"-"`
	TwoFactorAuthEnabled bool       `json:"two_factor_auth_enabled"`
	Email                string     `json:"email"`
	Phone                string     `json:"phone"`
	Signature            string     `json:"signature"`
	Details              string     `json:"details"`
	Notes                string     `json:"notes"`
	OrganizationID       int        `json:"organization_id"`
	Role                 string     `json:"role"`
	CustomRoleID         int        `json:"custom_role_id"`
	Moderator            bool       `json:"moderator"`
	TicketRestriction    string     `json:"ticket_restriction"`
	OnlyPrivateComments  bool       `json:"only_private_comments"`
	Tags                 []string   `json:"tags"`
	Suspended            bool       `json:"suspended"`
	RestrictedAgent      bool       `json:"restricted_agent"`
	UserFields           struct {
		Cargo        string `json:"cargo"`
		Celular      int    `json:"celular"`
		Departamento string `json:"departamento"`
		Kcs          string `json:"kcs"`
	} `json:"user_fields"`
}

type Organization struct {
	Id                 int        `json:"id"`
	Url                string     `json:"url"`
	Name               string     `json:"name"`
	SharedTickets      bool       `json:"shared_tickets"`
	SharedComments     bool       `json:"shared_comments"`
	ExternalID         string     `json:"external_id"`
	DataCreatedAt      CustomTime `json:"created_at"`
	CreatedAt          time.Time  `json:"-"`
	DataUpdatedAt      CustomTime `json:"updated_at"`
	UpdatedAt          time.Time  `json:"-"`
	DomainNames        []string   `json:"domain_names"`
	Details            string     `json:"details"`
	Notes              string     `json:"notes"`
	GroupID            int        `json:"group_id"`
	Tags               []string   `json:"tags"`
	OrganizationFields struct {
		AtualizacaoDeVersao       string `json:"atualizacao_de_versao"`
		CanalDeAtendimento        string `json:"canal_de_atendimento"`
		CdigoCliente              int    `json:"cdigo_cliente"`
		Cnpj                      string `json:"cnpj"`
		CoordenadorDeProjeto      string `json:"coordenador_de_projeto"`
		Cpf                       string `json:"cpf"`
		EmailGerenteConta         string `json:"email_gerente_conta"`
		EquipeTecnica             string `json:"equipe_tecnica"`
		EquipeTecnicaEMail        string `json:"equipe_tecnica_e_mail"`
		EMailCoordenadorDeProjeto string `json:"e_mail_coordenador_de_projeto"`
		FocalPointEmail           string `json:"focal_point_email"`
		FocalPointNome            string `json:"focal_point_nome"`
		GerenteContas             string `json:"gerente_contas"`
		ModulosAtendidos          string `json:"modulos_atendidos"`
		ModulosLicenciados        string `json:"modulos_licenciados"`
		NomeFantasia              string `json:"nome_fantasia"`
		RazaoSocial               string `json:"razao_social"`
		Representante             string `json:"representante"`
		Segmento                  string `json:"segmento"`
		Servicos                  string `json:"servicos"`
		ServicoCloud              string `json:"servico_cloud"`
		Status                    string `json:"status"`
		TelefoneOrg               string `json:"telefone_org"`
		TipoDaConta               string `json:"tipo_da_conta"`
	} `json:"organization_fields"`
}

type Ticket struct {
	Id               int        `json:"id"`
	Url              string     `json:"url"`
	External_id      string     `json:"external_id"`
	Type             string     `json:"type"`
	Subject          string     `json:"subject"`
	Raw_subject      string     `json:"raw_subject"`
	Description      string     `json:"description"`
	Priority         string     `json:"priority"`
	Status           string     `json:"status"`
	Recipient        string     `json:"recipient"`
	Requester_id     int        `json:"requester_id"`
	Submitter_id     int        `json:"submitter_id"`
	Assignee_id      int        `json:"assignee_id"`
	Organization_id  int        `json:"organization_id"`
	Group_id         int        `json:"group_id"`
	Collaborator_ids []int      `json:"collaborator_ids"`
	Forum_topic_id   int        `json:"forum_topic_id"`
	Problem_id       int        `json:"problem_id"`
	Has_incidents    bool       `json:"has_incidents"`
	DataDue_at       CustomTime `json:"due_at"`
	Due_at           time.Time  `json:"-"`
	Tags             []string   `json:"tags"`
	//Via                   Via      `json:"via"`
	Custom_fields         []CustomFields     `json:"custom_fields"`
	Satisfaction_rating   SatisfactionRating `json:"satisfaction_rating"`
	Sharing_agreement_ids []int              `json:"sharing_agreement_ids"`
	//Followup_ids          array              `json:"followup_ids"`
	//Ticket_form_id        int                `json:"ticket_form_id"`
	//Brand_id              int                `json:"brand_id"`
	DataCreated_at CustomTime `json:"created_at"`
	Created_at     time.Time  `json:"-"`
	DataUpdated_at CustomTime `json:"updated_at"`
	Updated_at     time.Time  `json:"-"`
}

type IncrementalGroup struct {
	Groups      []Group `json:"groups"`
	Count       int     `json:"count"`
	NextPage    string  `json:"next_page"`
	EndTime     int     `json:"end_time"`
	Error       string  `json:"error"`
	Description string  `json:"description"`
}

type IncrementalTicket struct {
	Tickets     []Ticket `json:"tickets"`
	Count       int      `json:"count"`
	NextPage    string   `json:"next_page"`
	EndTime     int      `json:"end_time"`
	Error       string   `json:"error"`
	Description string   `json:"description"`
}

type IncrementalOrganization struct {
	Organization []Organization `json:"organizations"`
	Count        int            `json:"count"`
	NextPage     string         `json:"next_page"`
	EndTime      int            `json:"end_time"`
	Error        string         `json:"error"`
	Description  string         `json:"description"`
}

type IncrementalUser struct {
	User        []User `json:"users"`
	Count       int    `json:"count"`
	NextPage    string `json:"next_page"`
	EndTime     int    `json:"end_time"`
	Error       string `json:"error"`
	Description string `json:"description"`
}

func (c *CustomValue) UnmarshalJSON(b []byte) error {
	v := string(b)

	*c = CustomValue(v)

	return nil
}
