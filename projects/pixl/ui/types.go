package ui

import (
	"fyne.io/fyne/v2"
	"zerotomastery.io/pixel/apptype"
	"zerotomastery.io/pixel/pxcanvas"
	"zerotomastery.io/pixel/swatch"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
