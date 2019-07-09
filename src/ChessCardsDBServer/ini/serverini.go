package mx_ini

/*
* Desc:服务器配置相关文件读取
 */

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/glog"
)

type ChessCardsDBServerIni struct {
	USERNAME string
	PASSWORD string
	NETWORK  string
	SERVER   string
	PORT     int
	DATABASE string
}

var ChessCardsDBServerIni *ChessCardsDBServerIni

func GetFileServerIni() *ChessCardsDBServerIni {
	if nil == ChessCardsDBServerIni {
		ChessCardsDBServerIni = &ChessCardsDBServerIni{}
	}
	return ChessCardsDBServerIni
}

//LoadConfig  读取配置信息
func (this *ChessCardsDBServerIni) LoadIni() bool {
	jsonFile, err := os.Open("ini/fileserver.json")
	defer jsonFile.Close()
	if err != nil {
		glog.Error("[Config] ", err)
		return false
	}

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		glog.Error("[Config] ", err)
		return false
	}

	err = json.Unmarshal(jsonData, this)
	if err != nil {
		glog.Error("[Config] ", err)
		return false
	}
	return true
}
