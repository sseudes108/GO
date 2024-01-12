package brush

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"zerotomastery.io/pixel/apptype"
)

const (
	Pixel = iota
)

func Cursor(config apptype.PxCanvasConfig, brush apptype.BrushType, event *desktop.MouseEvent, x, y int) []fyne.CanvasObject {
	var objects []fyne.CanvasObject
	switch {
	case brush == Pixel:
		pxSize := float32(config.PxSize)
		xOrigin := (float32(x) * pxSize) + config.CanvasOffset.X
		yOrigin := (float32(y) * pxSize) + config.CanvasOffset.Y

		cursorColor := color.RGBA{100, 100, 100, 255}

		left := canvas.NewLine(cursorColor)
		left.StrokeWidth = 3
		left.Position1 = fyne.NewPos(xOrigin, yOrigin)
		left.Position2 = fyne.NewPos(xOrigin, yOrigin+pxSize)

		top := canvas.NewLine(cursorColor)
		top.StrokeWidth = 3
		top.Position1 = fyne.NewPos(xOrigin, yOrigin)
		top.Position2 = fyne.NewPos(xOrigin+pxSize, yOrigin)

		right := canvas.NewLine(cursorColor)
		right.StrokeWidth = 3
		right.Position1 = fyne.NewPos(xOrigin+pxSize, yOrigin)
		right.Position2 = fyne.NewPos(xOrigin+pxSize, yOrigin+pxSize)

		bottom := canvas.NewLine(cursorColor)
		bottom.StrokeWidth = 3
		bottom.Position1 = fyne.NewPos(xOrigin, yOrigin+pxSize)
		bottom.Position2 = fyne.NewPos(xOrigin+pxSize, yOrigin+pxSize)

		objects = append(objects, left, top, right, bottom)
	}
	return objects
}

func TryBrush(appState *apptype.State, canvas apptype.Brushable, event *desktop.MouseEvent) bool {
	switch {
	case appState.BrushType == Pixel:
		return tryPaintPixel(appState, canvas, event)
	default:
		return false
	}
}

func tryPaintPixel(appState *apptype.State, canvas apptype.Brushable, event *desktop.MouseEvent) bool {
	x, y := canvas.MouseToCanvasXY(event)
	if x != nil && y != nil && event.Button == desktop.MouseButtonPrimary {
		canvas.SetColor(appState.BrushColor, *x, *y)
		return true
	}
	return false
}
