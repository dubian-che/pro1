package broswer

import (
	"MHDict/src/util/tools"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"log"
)

func mainContainerFunc(w fyne.Window) fyne.CanvasObject {
	//g_mainApp = app.New()
	logoImg := canvas.NewImageFromFile("resource/mainWindow.jpg")
	logoImg.Translucency = 0.5
	logoImg.FillMode = canvas.ImageFillStretch
	logoImg.SetMinSize(fyne.NewSize(800, 600))
	text := canvas.NewText("Fyne Logo", color.White)

	menuItem1 := fyne.NewMenuItem("关于MHDict~", func() {
		tools.MsgBox(w, "1111111111", "title")
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
		layout.NewCenterLayout(),
		logoImg, text,
	)

	return container.New(
		layout.NewBorderLayout(toolbar, nil, nil, nil),
		toolbar, /*widget.NewLabel(`Lorem ipsum dolor,
				sit amet consectetur adipisicing elit.
		  	Quidem consectetur ipsam nesciunt,
				quasi sint expedita minus aut,
		  	porro iusto magnam ducimus voluptates cum vitae.
		  	Vero adipisci earum iure consequatur quidem.`),*/
		//welcomeContainer(),
		m,
		content,
	)

}
