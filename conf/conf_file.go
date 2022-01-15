package conf

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

type CessClient struct {
	BoardInfo BoardInfo `yaml:"boardInfo"`
	ChainData ChainData `yaml:"chainData"`
}

type BoardInfo struct {
	BoardPath string `yaml:"boardPath"`
}
type ChainData struct {
	CessRpcAddr           string `yaml:"cessRpcAddr"`
	IdAccountPhraseOrSeed string `yaml:"idAccountPhraseOrSeed"`
	FaucetAddress         string `yaml:"faucetAddress"`
}

var ClientConf = new(CessClient)
var ConfFilePath string

func InitConf() {
	if ConfFilePath == "" {
		ConfFilePath = Conf_File_Path_D
	}
	_, err := os.Stat(ConfFilePath)
	if err != nil {
		fmt.Printf("\x1b[%dm[err]\x1b[0m The '%v' config file does not exist\n", 41, ConfFilePath)
		os.Exit(Exit_CmdLineParaErr)
	}
	yamlFile, err := ioutil.ReadFile(ConfFilePath)
	if err != nil {
		fmt.Printf("\x1b[%dm[err]\x1b[0m The '%v' file read error\n", 41, ConfFilePath)
		os.Exit(Exit_ConfErr)
	}
	err = yaml.Unmarshal(yamlFile, ClientConf)
	if err != nil {
		fmt.Printf("\x1b[%dm[err]\x1b[0m The '%v' file format error\n", 41, ConfFilePath)
		os.Exit(Exit_ConfErr)
	}
}
