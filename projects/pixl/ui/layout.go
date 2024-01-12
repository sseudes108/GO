package ui

import "fyne.io/fyne/v2/container"

func Setup(app *AppInit) {
	SetUpMenus(app)

	swatchesContainer := BuildSwatches(app)
	colorPicker := SetupColorPicker(app)
	rightSideContainer := container.NewVSplit(colorPicker, swatchesContainer)

	appLayout := container.NewBorder(nil, nil, nil, rightSideContainer, app.PixlCanvas)

	app.PixlWindow.SetContent(appLayout)
}
