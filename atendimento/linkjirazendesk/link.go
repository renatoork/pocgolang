package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://jiraplugin.zendesk.com/integrations/jira/account/megasistemas/links", nil)
	req.SetBasicAuth("qualiteam@mega.com.br/token", "pzwsg7ItHBIQqX1wzV66sYjl36awkhCXLgyMcjMx")
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println("erro: ", err.Error())
	}

	resp, err1 := client.Do(req)
	if err1 != nil {
		fmt.Println("erro: ", err1.Error())
	}

	b := new(bytes.Buffer)
	b.ReadFrom(resp.Body)

	fmt.Println("Resp: ", b.String())

}
