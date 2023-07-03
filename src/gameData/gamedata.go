package gameData

import (
	"fmt"
	"mygo/src/util/equipStruct"
	"mygo/src/util/jsonReader"
	"strconv"
	"strings"
)

type equipment struct {
	Data   equipStruct.EquipmentsData
	skills map[string]uint
}

var AllEquipments []equipment                               //装备表
var equipmentsinTypes map[uint][]uint                       //装备分类表
var EquipmentsMapOnSkillsOnLevel map[string]map[uint][]uint //装备技能表

func init() {
	fmt.Println("GD init")
	equipmentsinTypes = make(map[uint][]uint)
	AllEquipments = make([]equipment, 0, 0)
	var e []equipStruct.EquipmentsData = jsonReader.JsonReaderOnFile("test.json")
	for key, value := range e {
		equip := equipment{Data: value}
		equip.skills = make(map[string]uint)
		analysisEquipment(&equip)
		AllEquipments = append(AllEquipments, equip)
		equipmentsinTypes[equip.Data.Type] = append(equipmentsinTypes[equip.Data.Type], uint(key))
	}
	fmt.Println("GD init done")
}

func analysisEquipment(equip *equipment) {
	//println(equip.data.Detail)
	var skills = strings.Fields(equip.Data.Detail)
	for _, skill := range skills {
		if strings.Contains(skill, ":") {
			var detail = strings.Split(skill, ":")
			if len(detail) == 2 {
				level, err := strconv.Atoi(detail[1])
				if nil == err {
					equip.skills[detail[0]] = uint(level)
				}
			}
		}
	}

}

func GetEquipmentsBySkillOnType(skill string, level uint, typeIndex uint) (ans map[uint][]uint) {
	fmt.Println("GetEquipmentsBySkillOnType")
	for _, index := range equipmentsinTypes[typeIndex] {
		if AllEquipments[index].skills[skill] >= level {
			ans[AllEquipments[index].skills[skill]] = append(ans[AllEquipments[index].skills[skill]], index)
		}
	}
	return
}

func BuildSkillMap() {
	EquipmentsMapOnSkillsOnLevel = make(map[string]map[uint][]uint)
	for index, equip := range AllEquipments {
		for skill, level := range equip.skills {
			if nil == EquipmentsMapOnSkillsOnLevel[skill] {
				EquipmentsMapOnSkillsOnLevel[skill] = make(map[uint][]uint)
			}
			EquipmentsMapOnSkillsOnLevel[skill][level] = append(EquipmentsMapOnSkillsOnLevel[skill][level], uint(index))
		}
	}
}

func ShowSkillMap() {
	for key, value := range EquipmentsMapOnSkillsOnLevel {
		println("----------------")
		println(key)
		for k, v := range value {
			println("level:", k)
			for _, equipIndex := range v {
				print(AllEquipments[equipIndex].Data.Name, "\t")
			}
			println()
		}
	}
}
