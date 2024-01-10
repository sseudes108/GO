package swatch

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"zerotomastery.io/pixel/apptype"
)

type Swatch struct {
	widget.BaseWidget
	Selected     bool
	Color        color.Color
	SwatchIndex  int
	ClickHandler func(swatch *Swatch)
}

func (swatch *Swatch) SetColor(newColor color.Color) {
	swatch.Color = newColor
}

func NewSwatch(state *apptype.State, color color.Color, swatchIndex int, clickHandler func(swatch *Swatch)) *Swatch {
	swatch := &Swatch{
		Selected:     false,
		Color:        color,
		ClickHandler: clickHandler,
		SwatchIndex:  swatchIndex,
	}
	swatch.ExtendBaseWidget(swatch)
	return swatch
}

func (swatch *Swatch) CreateRenderer() fyne.WidgetRenderer {
	square := canvas.NewRectangle(swatch.Color)
	objects := []fyne.CanvasObject{square}
	return &SwatchRenderer{
		square:  square,
		objects: objects,
		parent:  swatch,
	}
}
