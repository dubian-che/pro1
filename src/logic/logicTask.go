package logic

import "MHDict/src/gameData"

type taskAlgorithm struct {
	target map[string]uint //当前目标技能
	taskId uint
}

type dictAlgorithm struct {
	target      taskAlgorithm //当前目标技能
	target_next taskAlgorithm //下一个目标技能
	//state            	uint          //工作状态 0 空任务，随时响应， 1 工作中,无新任务, 2 工作中
	ans              [][]uint
	bead             []uint
	ansnTargetNumber uint
	stoping_flag     bool
}

var (
	g_dictAlgorithm dictAlgorithm
)

func init() {
	g_dictAlgorithm.ansnTargetNumber = 10
	//g_dictAlgorithm.state = 0
	g_dictAlgorithm.stoping_flag = false
}

func fitSkill(tmp []uint, index uint) bool {
	if len(gameData.AllEquipments[index].Skills) == 1 {
		_, ok := gameData.AllEquipments[index].Skills["hole"]
		if ok {
			return true
		}
	}
	hasSkill := make(map[string]uint)
	for _, index := range tmp {
		for skill, level := range gameData.AllEquipments[index].Skills {
			hasSkill[skill] += level
		}
	}
	for skill, level := range hasSkill {
		_, ok := gameData.AllEquipments[index].Skills[skill]
		if ok && level < g_dictAlgorithm.target.target[skill] {
			return true
		}
	}
	return false
}

func fitAllSkill(tmp []uint) bool {
	hasSkill := make(map[string]uint)
	for _, index := range tmp {
		for skill, level := range gameData.AllEquipments[index].Skills {
			hasSkill[skill] += level
		}
	}
	for skill, level := range g_dictAlgorithm.target.target {
		l, ok := hasSkill[skill]
		if !ok || l < level {
			return false
		}
	}
	return true
}

func algorithmFunc(tmp []uint) {
	if g_dictAlgorithm.target.taskId == 0 {
		if g_dictAlgorithm.ans != nil {
			g_dictAlgorithm.ans = nil
		}
		return
	}

	if len(g_dictAlgorithm.ans) == int(g_dictAlgorithm.ansnTargetNumber) {
		return
	}

	if len(tmp) == 5 {
		if fitAllSkill(tmp) {
			g_dictAlgorithm.ans = append(g_dictAlgorithm.ans, tmp)
			g_dictAlgorithm.bead = nil
			return
		}
		//todo 加珠子

		g_dictAlgorithm.ans = nil //无结果
		return
	} else {
		for index, _ := range gameData.AllEquipments {
			if fitSkill(tmp, uint(index)) {
				tmp = append(tmp, uint(index))
				algorithmFunc(tmp)
				tmp = append(tmp[:len(tmp)-1])
			}
		}
	}

}

func addAlgorithmFuncTask(target map[string]uint) {
	g_dictAlgorithm.target_next = taskAlgorithm{target, 0}
}

func taskRun() {
	for !logicStop {
		if g_dictAlgorithm.target_next.taskId > g_dictAlgorithm.target.taskId {
			g_dictAlgorithm.target = g_dictAlgorithm.target_next
			g_dictAlgorithm.target_next.taskId = 0
		}
		if g_dictAlgorithm.target.taskId != 0 {
			algorithmFunc(make([]uint, 0))
		}
	}
}
