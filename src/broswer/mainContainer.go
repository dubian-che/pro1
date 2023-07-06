package broswer

import (
	"MHDict/resource"
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
	"strconv"
)

func init() {
	g_skillSelected = make(map[string]uint)
	g_skillSelected["攻击"] = 6
	g_skillSelected["会心"] = 7
}

var (
	g_skillSelected map[string]uint
)

func ChangeSkillSelected(skill string, add bool) {
	if add {
		g_skillSelected[skill]++
	} else {
		v, ok := g_skillSelected[skill]
		if !ok || v == 0 {
			return
		}
		g_skillSelected[skill]--
	}

	g_mainWindow.SetContent(split)
}

func SkillSelectedContain() fyne.CanvasObject {
	//f := make([]fyne.CanvasObject, 0)
	c := container.NewVBox()
	for skill, level := range g_skillSelected {
		//f := container.NewHBox(widget.NewButton("+", func(){})), canvas.NewText(skill + level, color.Black), widget.NewButton("-", func(){}))
		f := container.NewHBox(widget.NewButton("-", func() {
			ChangeSkillSelected(skill, true)
		}), canvas.NewText(skill+strconv.Itoa(int(level)), color.Black), widget.NewButton("+", func() {
			ChangeSkillSelected(skill, false)
		}))
		c.Add(f)
	}

	return c
}

func mainContain() fyne.CanvasObject {
	provinceSelect := widget.NewSelect([]string{"anhui", "zhejiang", "shanghai"}, func(value string) {
		fmt.Println("province:", value)
	})
	return container.NewVBox(provinceSelect, SkillSelectedContain())
}

func mainContainerFunc(w fyne.Window) fyne.CanvasObject {
	//g_mainApp = app.New()
	logoImg := resource.Img_mainWindow1
	logoImg.Translucency = 0.5
	logoImg.FillMode = canvas.ImageFillStretch
	logoImg.SetMinSize(fyne.NewSize(800, 600))
	mainContain := mainContain()

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
		logoImg, mainContain,
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
