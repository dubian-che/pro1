package broswer

import (
	"MHDict/src/util/tools"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func welcomeContainer(w fyne.Window) fyne.CanvasObject {
	//g_mainApp = app.New()
	logo := canvas.NewImageFromFile("resource/mhdict1.jpg")
	logo.Translucency = 0.8
	//logo := canvas.NewImageFromResource(data.FyneScene)
	logo.FillMode = canvas.ImageFillContain
	logo.SetMinSize(fyne.NewSize(400, 300))

	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("MHDict配装器，年轻人的第一款配装器", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		logo,
		container.NewHBox(
			widget.NewHyperlink("项目地址", tools.ParseURL("https://github.com/dubian-che/pro1/")),
			widget.NewLabel("        "),
			widget.NewButton("关于MHDict", func() {
				tools.MsgBox(w, "ttt", "132123")
			}),
			//widget.NewHyperlink("fyne.io", tools.ParseURL("https://fyne.io/")),
			//widget.NewHyperlink("documentation", tools.ParseURL("https://developer.fyne.io/")),
			//widget.NewLabel("-"),
			//widget.NewHyperlink("sponsor", tools.ParseURL("https://fyne.io/sponsor/")),
		),
		widget.NewLabel(""), // balance the header on the tutorial screen we leave blank on this content
	))
}
