package credito

import (
	"encoding/json"
	"net/http"
)

// APIs handlers
func GetAnaliseCredito(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if origin := req.Header.Get("Origin"); origin != "" {
		res.Header().Set("Access-Control-Allow-Origin", origin)
	}
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type")
	res.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	if req.Method == "OPTIONS" {
		return
	}
	ret := execAnalise(res, req)
	mens, _ := json.Marshal(ret)
	res.Write(mens)
}
