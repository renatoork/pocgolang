package Util

import (
	"mega/atendimento/DashBoardSync/LogError"

	"bytes"
	"net/http"
	"strconv"
)

func GetJson(url string, metodo string) ([]byte, string) {

	client := &http.Client{}

	reqJson, err := http.NewRequest(metodo, url, nil)
	reqJson.SetBasicAuth("qualiteam@mega.com.br/token", "pzwsg7ItHBIQqX1wzV66sYjl36awkhCXLgyMcjMx")
	reqJson.Header.Set("Content-Type", "application/json")

	reqJson.Close = true

	jsonResp, err := client.Do(reqJson)
	LogError.Log("GetJson", err)

	defer jsonResp.Body.Close()

	rAfter := jsonResp.Header.Get("Retry-After")

	b := new(bytes.Buffer)
	b.ReadFrom(jsonResp.Body)

	return b.Bytes(), rAfter

}

func ArIntToString(array []int) string {

	var ret string

	for i, v := range array {

		if i == 0 {
			ret = strconv.Itoa(v)
		} else {
			ret += ", " + strconv.Itoa(v)
		}

	}

	if len(array) <= 0 {
		ret = " "
	}

	return ret

}

func ArStrToString(array []string) string {

	var ret string

	for i, v := range array {

		if i == 0 {
			ret = v
		} else {
			ret += ", " + v
		}

	}

	if len(array) <= 0 {
		ret = " "
	}

	return ret

}

func BoolToStr(b bool) string {

	if b {
		return "true"
	} else {
		return "false"
	}

}
