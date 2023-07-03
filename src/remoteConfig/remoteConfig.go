package remoteConfig

import (
	"fmt"
	"mygo/src/util/configStruct"
	"mygo/src/util/httpUtil"
	"mygo/src/util/jsonReader"
)

var globalConfig string
var globalConfigInfo configStruct.ConfigFileStruct

func AnalysisGlobalConfig() {
	//println("global config:\n", globalConfig)
	jsonReader.JsonReaderToStruct(globalConfig, &globalConfigInfo)
	//println(globalConfigInfo.Config.Version)
	//println(globalConfigInfo.Config.Downloadurl)
}

func init() {
	c, code := httpUtil.HttpUtil_Req("http://my-json-server.typicode.com/dubian-che/pro1/posts/1")
	if code == 200 {
		fmt.Println("global config init success")
		globalConfig = c
	} else {
		fmt.Println("global config init error")
		return
	}
	//AnalysisGlobalConfig()
}
