package tools

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"net/url"
	"runtime"
)

func WhichSystem() (ans string) {
	if "darwin" == runtime.GOOS {
		ans = "macos"
	} else {
		ans = "win"
	}
	return
}

var msgBox fyne.Window

func init() {
	msgBox = nil
}

//func MsgBox(msg string, title string, app fyne.App) {
//	if msgBox != nil {
//		msgBox.Close()
//	}
//	msgBox = app.NewWindow(title)
//	text := canvas.NewText(msg, color.White)
//	msgBox.SetContent(text)
//	msgBox.Show()
//
//}

func ParseURL(urlStr string) *url.URL {
	link, err := url.Parse(urlStr)
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}

	return link
}

func MsgBox(w fyne.Window, title string, txt string) {
	content := container.NewVBox(
		widget.NewLabel(txt),
	)
	customDialog1 := dialog.NewCustom(title, "dismiss", content, w)
	customDialog1.SetDismissText("确认") // 设置自定义按钮
	customDialog1.Show()

}
