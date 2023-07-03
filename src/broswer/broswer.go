package broswer

import (
	"fmt"
	"mygo/src/gameData"
	"mygo/src/remoteConfig"
)

func Start() {
	fmt.Println("broswer start")
	//gameData.GetEquipmentsBySkillOnType("攻击", 1, 1)
	gameData.BuildSkillMap()
	//gameData.ShowSkillMap()
	remoteConfig.AnalysisGlobalConfig()
}
