package pedidovenda

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"io"
	"io/ioutil"
	"mega/go-util/dbg"
	_ "mega/go-util/dbg"
	"mega/ms-consul/consul"
	"mega/ms-pedidovenda/tipos"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func mockServico(metodo string, url string, contentType string, body io.Reader) (*httptest.ResponseRecorder, *http.Request) {
	os.Setenv("NLS_LANG", ".AL32UTF8")
	r, _ := http.NewRequest(metodo, url, body)
	w := httptest.NewRecorder()
	r.Header.Set("Content-Type", contentType)
	w.Code = 400
	return w, r
}

func getUrl(nomeSrv string) string {
	return fmt.Sprintf("/api/%s/%s", tipos.Versao, nomeSrv)
}

func TestCheckServico(t *testing.T) {
	Convey("Check do serviço do PDV", t, func() {
		w, r := mockServico("GET", getUrl("check"), "", nil)
		consul.CheckServico(w, r)
		So(w.Code, ShouldEqual, http.StatusOK)
	})
}

func lerArquivo(nomeArq string) string {
	arq, _ := ioutil.ReadFile(nomeArq)
	return string(arq)
}

func executaTesteDir(dir string, exsrv func(http.ResponseWriter, *http.Request)) {
	readerDir, _ := ioutil.ReadDir(dir)
	if len(readerDir) > 0 {
		for _, fileInfo := range readerDir {
			if !fileInfo.IsDir() {
				Convey(fileInfo.Name(), func() {
					arq := lerArquivo(fmt.Sprintf("%s\\%s", dir, fileInfo.Name()))
					body := strings.NewReader(arq)
					w, r := mockServico("POST", getUrl("pedido"), "application/json; param=value", body)

					exsrv(w, r)

					So(w.Code, ShouldEqual, http.StatusOK)

				})
			}
		}
	} else {
		So(len(readerDir), ShouldBeGreaterThan, 0)
	}
}

func TestSetPedido(t *testing.T) {
	dbg.SetDebug(false)
	os.Setenv("NLS_LANG", ".AL32UTF8")
	os.Setenv("NLS_NUMERIC_CHARACTERS", ",.")
	DbMap = InitDb()
	Convey("Inclusão", t, func() {
		executaTesteDir("../testes/inclusao", SetPedido)
	})
}

/*
func getPedido(res http.ResponseWriter, req *http.Request) {
	emConstrucao(res, req, varPed(mux.Vars(req)))
}

func putPedido(res http.ResponseWriter, req *http.Request) {
	emConstrucao(res, req, varPed(mux.Vars(req)))
}
func delPedido(res http.ResponseWriter, req *http.Request) {
	emConstrucao(res, req, varPed(mux.Vars(req)))
}

func setObs(res http.ResponseWriter, req *http.Request) {
	emConstrucao(res, req, varPed(mux.Vars(req)))
}
func getObs(res http.ResponseWriter, req *http.Request) {
	emConstrucao(res, req, varPed(mux.Vars(req)))
}
func putObs(res http.ResponseWriter, req *http.Request) {
	emConstrucao(res, req, varPed(mux.Vars(req)))
}
func delObs(res http.ResponseWriter, req *http.Request) {
	emConstrucao(res, req, varPed(mux.Vars(req)))
}

func setItems(res http.ResponseWriter, req *http.Request) {
	emConstrucao(res, req, varPed(mux.Vars(req)))
}
func getItems(res http.ResponseWriter, req *http.Request) {
	emConstrucao(res, req, varPed(mux.Vars(req)))
}
func putItem(res http.ResponseWriter, req *http.Request) {
	emConstrucao(res, req, varItem(mux.Vars(req)))
}
func delItem(res http.ResponseWriter, req *http.Request) {
	emConstrucao(res, req, varItem(mux.Vars(req)))
}

func setItemObs(res http.ResponseWriter, req *http.Request) {
	emConstrucao(res, req, varItem(mux.Vars(req)))
}
func getItemObs(res http.ResponseWriter, req *http.Request) {
	emConstrucao(res, req, varItem(mux.Vars(req)))
}
func putItemObs(res http.ResponseWriter, req *http.Request) {
	emConstrucao(res, req, varItemObs(mux.Vars(req)))
}
func delItemObs(res http.ResponseWriter, req *http.Request) {
	emConstrucao(res, req, varItemObs(mux.Vars(req)))
}

func setItemProg(res http.ResponseWriter, req *http.Request) {
	emConstrucao(res, req, varItemProg(mux.Vars(req)))
}
func getItemProg(res http.ResponseWriter, req *http.Request) {
	emConstrucao(res, req, varItemProg(mux.Vars(req)))
}
func putItemProg(res http.ResponseWriter, req *http.Request) {
	emConstrucao(res, req, varItemProg(mux.Vars(req)))
}
func delItemProg(res http.ResponseWriter, req *http.Request) {
	emConstrucao(res, req, varItemProg(mux.Vars(req)))
}

func setItemProgEstoque(res http.ResponseWriter, req *http.Request) {
	emConstrucao(res, req, varItemProgEstoque(mux.Vars(req)))
}
func getItemProgEstoque(res http.ResponseWriter, req *http.Request) {
	emConstrucao(res, req, varItemProgEstoque(mux.Vars(req)))
}
func putItemProgEstoque(res http.ResponseWriter, req *http.Request) {
	emConstrucao(res, req, varItemProgEstoque(mux.Vars(req)))
}
func delItemProgEstoque(res http.ResponseWriter, req *http.Request) {
	emConstrucao(res, req, varItemProgEstoque(mux.Vars(req)))
}

*/
