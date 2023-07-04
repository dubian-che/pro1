package main

import (
	"MHDict/src/broswer"
	"fmt"
	"github.com/flopp/go-findfont"
	"os"
	"strings"
)

func initMain() {
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		fmt.Println(path)
		//楷体:simkai.ttf
		//黑体:simhei.ttf
		if strings.Contains(path, "simkai.ttf") {
			os.Setenv("FYNE_FONT", path)
			break
		}
	}
}

func main() {
	//jsonReader.JsonReaderOnFile("test.json")
	initMain()
	broswer.Start()
	//httpUtil.Test()
	fmt.Println("hello")
}
