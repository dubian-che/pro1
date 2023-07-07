package broswer

import (
	"MHDict/resource"
	"MHDict/src/util/tools"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
)

func init() {
	g_skillSelected = make(map[string]uint)
	g_skillSelected["攻击"] = 6
	g_skillSelected["会心"] = 7
}

var (
	g_skillSelected   map[string]uint
	g_selectedContail *fyne.Container
)

func ChangeSkillSelected(skill string, add bool) {
	if add {
		g_skillSelected[skill]++
		fmt.Println("+++")
	} else {
		v, ok := g_skillSelected[skill]
		if !ok || v == 0 {
			return
		}
		g_skillSelected[skill]--
		fmt.Println("---")
	}
	RefreshSelectedContain()
	//g_skillSelected.refresh()
	fmt.Println(skill)
	g_selectedContail.Refresh()
}

func mainContainerFunc(_ fyne.Window) fyne.CanvasObject {
	item := container.NewVBox(widget.NewSeparator(), widget.NewSeparator())

	dataList := binding.BindIntList(&[]int{0, 1, 2})

	button := widget.NewButton("Append", func() {
		dataList.Append(int(dataList.Length() + 1))
	})

	list := widget.NewListWithData(dataList,
		func() fyne.CanvasObject {
			return container.NewBorder(nil, nil, widget.NewButton("-", nil), widget.NewButton("+", nil),
				widget.NewLabel("item 1"))
		},
		func(item binding.DataItem, obj fyne.CanvasObject) {
			f := item.(binding.Int)
			text := obj.(*fyne.Container).Objects[0].(*widget.Label)
			text.Bind(binding.IntToStringWithFormat(f, "item %d"))

			btn1 := obj.(*fyne.Container).Objects[1].(*widget.Button)
			btn2 := obj.(*fyne.Container).Objects[2].(*widget.Button)
			btn1.OnTapped = func() {
				val, _ := f.Get()
				fmt.Println(val)
				_ = f.Set(val - 1)
			}
			btn2.OnTapped = func() {
				val, _ := f.Get()
				fmt.Println(val)
				_ = f.Set(val + 1)
			}
		})

	//list.Resize(fyne.NewSize(400, 400))
	listPanel := container.NewBorder(button, nil, nil, nil, list)
	content := fyne.NewContainerWithLayout(
		layout.NewCenterLayout(),
		listPanel,
	)
	//listPanel.Resize(fyne.NewSize(400, 400))
	return container.NewBorder(item, nil, nil, nil, container.NewGridWithColumns(5, content))
}

func RefreshSelectedContain() {
	g_selectedContail = container.NewVBox()
	//for skill, level := range g_skillSelected {
	//	sk := skill
	//	//f := container.NewHBox(widget.NewButton("+", func(){})), canvas.NewText(skill + level, color.Black), widget.NewButton("-", func(){}))
	//	f := container.NewHBox(
	//		widget.NewButton("-", func() {
	//			ChangeSkillSelected(sk, false)
	//		}), canvas.NewText(sk+strconv.Itoa(int(level)), color.Black),
	//		widget.NewButton("+", func() {
	//			ChangeSkillSelected(sk, true)
	//		}))
	//	g_selectedContail.Add(f)
	//}
}

func skillSelectedContain() fyne.CanvasObject {
	//f := make([]fyne.CanvasObject, 0)

	RefreshSelectedContain()
	return g_selectedContail
}

func mainContain() fyne.CanvasObject {
	provinceSelect := widget.NewSelect([]string{"anhui", "zhejiang", "shanghai"}, func(value string) {
		fmt.Println("province:", value)
	})
	RefreshSelectedContain()
	return container.NewVBox(provinceSelect, g_selectedContail)
}

func mainContainerFunc1(w fyne.Window) fyne.CanvasObject {
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
