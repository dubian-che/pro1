package jsonReader

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"mygo/src/util/configStruct"
	"mygo/src/util/equipStruct"
	"os"
)

type Equip struct {
	EquipmentsDatas []equipStruct.EquipmentsData `json:"equipments"`
}

func JsonReaderOnFile(file string) (ans []equipStruct.EquipmentsData) {
	fileContent, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
		return
	}

	//fmt.Println("The File is opened successfully...")

	defer fileContent.Close()
	byteResult, _ := ioutil.ReadAll(fileContent)
	var equip Equip
	json.Unmarshal(byteResult, &equip)
	ans = equip.EquipmentsDatas

	for i := 0; i < len(equip.EquipmentsDatas); i++ {
		//fmt.Println("Equip Name: " + equip.EquipmentsDatas[i].Name)
		//fmt.Println("Equip Detail: " + equip.EquipmentsDatas[i].Detail)
	}
	return
}

func JsonReaderToStruct(jsonString string, item *configStruct.ConfigFileStruct) error {
	if nil == item {
		return errors.New("item nil")
	}
	return json.Unmarshal([]byte(jsonString), item)
}
