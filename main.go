package main

import (
	"MHDict/src/broswer"
	"MHDict/src/logic"
	"fmt"
	"github.com/flopp/go-findfont"
	"os"
	"strings"
)

func initMain() {
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		//fmt.Println(path)
		//楷体:simkai.ttf
		//黑体:simhei.ttf
		if strings.Contains(path, "msyh.ttf") ||
			strings.Contains(path, "simhei.ttf") ||
			strings.Contains(path, "simsun.ttc") ||
			strings.Contains(path, "simkai.ttf") ||
			strings.Contains(path, "STHeiti Light") ||
			strings.Contains(path, "Medium.ttc") ||
			strings.Contains(path, "STSong") {
			fmt.Println("set language", path)
			os.Setenv("FYNE_FONT", path)
			//fmt.Println("set ok")
			break
		}
	}
}

func main() {
	//jsonReader.JsonReaderOnFile("test.json")
	//initMain()
	go logic.LogicStart()
	broswer.Start()
	//httpUtil.Test()
	logic.LogicEnd()

}
