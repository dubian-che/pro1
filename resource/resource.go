package resource

import (
	"fyne.io/fyne/v2/canvas"
)

var (
	Img_welcome1    *canvas.Image
	Img_mainWindow1 *canvas.Image
)

func init() {
	Img_welcome1 = canvas.NewImageFromFile("rsc/img/mhdict1.jpg")
	Img_mainWindow1 = canvas.NewImageFromFile("rsc/img/mainWindow.jpg")
}
