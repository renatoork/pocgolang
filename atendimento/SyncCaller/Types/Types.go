package Types

import (
	"strconv"
	"time"
)

type Caller struct {
	Url     string
	Metodo  string
	Retorno []byte
	Erro    string
}

type RecLog struct {
	FileName string
	Log      string
}

type LastError struct {
	Log string
}

type IntervalSync struct {
	SyncInterval *time.Ticker
	Interval     int
}

type SyncConfig struct {
	SyncConfig struct {
		Host             string `json:"host"`
		Port             string `json:"port"`
		TicketPath       string `json:"ticket_path"`
		OrganizationPath string `json:"organization_path"`
		UserPath         string `json:"user_path"`
		GroupPath        string `json:"group_path"`
	} `json:"sync_config"`
}

func (c *IntervalSync) StopInterval() {
	c.SyncInterval.Stop()
	c.SyncInterval = nil
}

func (c *IntervalSync) DefineIntervalTicket() {
	c = new(IntervalSync)
}

func (c *IntervalSync) CreateInterval() {
	if c.SyncInterval == nil {
		c.SyncInterval = time.NewTicker(time.Minute * time.Duration(c.Interval))
	}
}

func (c *SyncConfig) GetUrl(tipo string) string {

	var url string

	url = c.SyncConfig.Host + ":" + c.SyncConfig.Port + "/"

	switch tipo {
	case "ticket":
		url = url + c.SyncConfig.TicketPath
	case "organization":
		url = url + c.SyncConfig.OrganizationPath
	case "group":
		url = url + c.SyncConfig.GroupPath
	case "user":
		url = url + c.SyncConfig.UserPath
	default:
		return url
	}

	return url

}

func (c *SyncConfig) GetCustomUrl() string {

	var url string

	url = c.SyncConfig.Host + ":" + c.SyncConfig.Port + "/"

	return url

}

func (c *IntervalSync) SetInterval(n string) {

	nm, err := strconv.Atoi(n)

	if err == nil {
		c.Interval = nm
	}

}
