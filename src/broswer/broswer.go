package broswer

import (
	"MHDict/src/gameData"
	"MHDict/src/remoteConfig"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
)

var g_mainApp fyne.App
var g_mainWindow fyne.Window

func Start() {
	fmt.Println("broswer start")
	//gameData.GetEquipmentsBySkillOnType("攻击", 1, 1)
	gameData.BuildSkillMap()
	//gameData.ShowSkillMap()
	remoteConfig.AnalysisGlobalConfig()

	APPMain()
}

func APPMain() {
	g_mainApp = app.New()
	g_mainWindow = g_mainApp.NewWindow("Hello, i am MHDict!")
	// 创建按钮
	//button1 := widget.NewButton("按钮1", func() {
	//	// 处理按钮1的点击事件
	//})
	//button2 := widget.NewButton("按钮2", func() {
	//	// 处理按钮2的点击事件
	//})
	//
	//widget.NewMenu(
	//	widget.Menu{})
	// 创建下拉菜单
	//menu := widget.NewMenu(
	//	widget.NewMenuItem("菜单项1", func() {
	//		// 处理菜单项1的点击事件
	//	}),
	//	widget.NewMenuItem("菜单项2", func() {
	//		// 处理菜单项2的点击事件
	//	}),
	//)

	menuItem1 := fyne.NewMenuItem("c1", func() {
		// 处理菜单项1的点击事件
	})

	// 创建菜单项2的点击事件处理程序
	menuItem2 := fyne.NewMenuItem("c2", func() {
		// 处理菜单项2的点击事件
	})

	m := widget.NewMenu(fyne.NewMenu(
		"",
		menuItem1,
		menuItem2,
	))
	m.Hide()

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.QuestionIcon(), func() {
			if m.Hidden {
				m.Show()
			} else {
				m.Hide()
			}
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {

			fmt.Println("Cut")
		}),
		//widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {
			fmt.Println("Copy")
		}),
		//widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {
			fmt.Println("Paste")
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			log.Println("Display help")
		}),
	)

	content := fyne.NewContainerWithLayout(
		layout.NewBorderLayout(toolbar, nil, nil, nil),
		toolbar, widget.NewLabel(`Lorem ipsum dolor, 
		sit amet consectetur adipisicing elit.
    	Quidem consectetur ipsam nesciunt,
		quasi sint expedita minus aut,
    	porro iusto magnam ducimus voluptates cum vitae.
    	Vero adipisci earum iure consequatur quidem.`),
		m,
	)
	g_mainWindow.SetContent(content)
	g_mainWindow.Resize(fyne.NewSize(500, 500))
	g_mainWindow.ShowAndRun()
}
