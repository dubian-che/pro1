package gameData

import (
	"MHDict/src/util/equipStruct"
	"MHDict/src/util/jsonReader"
	"fmt"
	"strconv"
	"strings"
)

type Equipment struct {
	Data   equipStruct.EquipmentsData
	Skills map[string]uint
}

var AllEquipments []Equipment         //装备表
var equipmentsinTypes map[uint][]uint //装备分类表
var equipmentsinTypes_inSkills map[uint]map[string][]uint
var EquipmentsMapOnSkillsOnLevel map[string]map[uint][]uint //装备技能表

func init() {
	fmt.Println("GD init")
	equipmentsinTypes = make(map[uint][]uint)
	AllEquipments = make([]Equipment, 0, 0)
	var e []equipStruct.EquipmentsData = jsonReader.JsonReaderOnFile("../test.json")
	for key, value := range e {
		equip := Equipment{Data: value}
		equip.Skills = make(map[string]uint)
		analysisEquipment(&equip)
		AllEquipments = append(AllEquipments, equip)
		equipmentsinTypes[equip.Data.Type] = append(equipmentsinTypes[equip.Data.Type], uint(key))

	}
	fmt.Println("GD init done")
}

func analysisEquipment(equip *Equipment) {
	//println(equip.data.Detail)
	var skills = strings.Fields(equip.Data.Detail)
	for _, skill := range skills {
		if strings.Contains(skill, ":") {
			var detail = strings.Split(skill, ":")
			if len(detail) == 2 {
				level, err := strconv.Atoi(detail[1])
				if nil == err {
					equip.Skills[detail[0]] = uint(level)
				}
			}
		}
	}

}

func GetEquipmentsBySkillOnType(skill string, level uint, typeIndex uint) (ans map[uint][]uint) {
	fmt.Println("GetEquipmentsBySkillOnType")
	for _, index := range equipmentsinTypes[typeIndex] {
		if AllEquipments[index].Skills[skill] >= level {
			ans[AllEquipments[index].Skills[skill]] = append(ans[AllEquipments[index].Skills[skill]], index)
		}
	}
	return
}

func BuildSkillMap() {
	EquipmentsMapOnSkillsOnLevel = make(map[string]map[uint][]uint)
	equipmentsinTypes_inSkills = make(map[uint]map[string][]uint)
	for index, equip := range AllEquipments {
		if equipmentsinTypes_inSkills[equip.Data.Type] == nil {
			equipmentsinTypes_inSkills[equip.Data.Type] = make(map[string][]uint)
		}
		for skill, level := range equip.Skills {
			if nil == EquipmentsMapOnSkillsOnLevel[skill] {
				EquipmentsMapOnSkillsOnLevel[skill] = make(map[uint][]uint)
			}
			EquipmentsMapOnSkillsOnLevel[skill][level] = append(EquipmentsMapOnSkillsOnLevel[skill][level], uint(index))

			equipmentsinTypes_inSkills[equip.Data.Type][skill] = append(equipmentsinTypes_inSkills[equip.Data.Type][skill], uint(index))
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
