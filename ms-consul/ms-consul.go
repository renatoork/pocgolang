package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mega/go-util/erro"
	"os"
	"os/exec"
	"strings"
)

type server struct {
	pid        int
	dataDir    string
	configFile string
}

type portConfig struct {
	DNS     int `json:"dns,omitempty"`
	HTTP    int `json:"http,omitempty"`
	RPC     int `json:"rpc,omitempty"`
	SerfLan int `json:"serf_lan,omitempty"`
	SerfWan int `json:"serf_wan,omitempty"`
	Server  int `json:"server,omitempty"`
}

type addressConfig struct {
	HTTP string `json:"http,omitempty"`
}

type serverConfig struct {
	Bootstrap bool           `json:"bootstrap,omitempty"`
	Server    bool           `json:"server,omitempty"`
	DataDir   string         `json:"data_dir,omitempty"`
	LogLevel  string         `json:"log_level,omitempty"`
	Addresses *addressConfig `json:"addresses,omitempty"`
	Ports     portConfig     `json:"ports,omitempty"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Execute: ms-consul <parametro> \n - carregar: para carregar o sevico Consul.")
		os.Exit(1)
	}
	switch strings.ToUpper(os.Args[1]) {
	case "CARREGAR":
		fmt.Println(carregarServico())
	}
}

func DefaultConfig() *serverConfig {
	return &serverConfig{
		Bootstrap: true,
		Server:    true,
		LogLevel:  "debug",
		Ports: portConfig{
			DNS:     53,
			HTTP:    18800,
			RPC:     18600,
			SerfLan: 18200,
			SerfWan: 18400,
			Server:  18000,
		},
	}
}

func carregarServico() *server {
	if path, err := exec.LookPath("consul"); err != nil || path == "" {
		fmt.Println("O serviço Consul não esta configurado corretamente. \n Verifique se o caminho do executável do Consul esta no PATH.")
		os.Exit(1)
	}

	pidFile, err := ioutil.TempFile("", "consul")
	if erro.Trata(err) {
		os.Exit(1)
	}

	pidFile.Close()
	os.Remove(pidFile.Name())

	dataDir, err := ioutil.TempDir("", "consul")
	if erro.Trata(err) {
		os.Exit(1)
	}

	configFile, err := ioutil.TempFile("", "consul")
	if erro.Trata(err) {
		os.Exit(1)
	}

	consulConfig := DefaultConfig()
	consulConfig.DataDir = dataDir

	configContent, err := json.Marshal(consulConfig)
	if erro.Trata(err) {
		os.Exit(1)
	}

	if _, err := configFile.Write(configContent); erro.Trata(err) {
		os.Exit(1)
	}
	configFile.Close()

	// Start the server
	cmd := exec.Command("consul", "agent", "-config-file", configFile.Name())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); erro.Trata(err) {
		os.Exit(1)
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Println(err.Error())
	}

	return &server{
		pid:        cmd.Process.Pid,
		dataDir:    dataDir,
		configFile: configFile.Name(),
	}
}
