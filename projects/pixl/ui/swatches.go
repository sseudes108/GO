package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"zerotomastery.io/pixel/swatch"
)

func BuildSwatches(app *AppInit) *fyne.Container {
	canvasSwatches := make([]fyne.CanvasObject, 0, 64)
	for i := 0; i < cap(app.Swatches); i++ {
		initialColor := color.NRGBA{255, 255, 255, 255}
		newSwatch := swatch.NewSwatch(app.State, initialColor, i, func(swatch *swatch.Swatch) {
			for j := 0; j < len(app.Swatches); j++ {
				app.Swatches[j].Selected = false
				canvasSwatches[j].Refresh()
			}
			app.State.SwatchSelected = swatch.SwatchIndex
			app.State.BrushColor = swatch.Color
		})
		if i == 0 {
			newSwatch.Selected = true
			app.State.SwatchSelected = 0
		}
		app.Swatches = append(app.Swatches, newSwatch)
		canvasSwatches = append(canvasSwatches, newSwatch)
		newSwatch.Refresh()
	}
	return container.NewGridWrap(fyne.NewSize(20, 20), canvasSwatches...)
}
