package pxcanvas

import (
	"image"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"zerotomastery.io/pixel/apptype"
)

type PxCanvasMouseState struct {
	previousMouseCoord *fyne.PointEvent
}

type PxCanvas struct {
	widget.BaseWidget
	apptype.PxCanvasConfig
	renderer    *PxCanvasRenderer
	PixelData   image.Image
	mouseState  PxCanvasMouseState
	appState    *apptype.State
	reloadImage bool
}

func (PxCanvas *PxCanvas) Bounds() image.Rectangle {
	x0 := int(PxCanvas.CanvasOffset.X)
	y0 := int(PxCanvas.CanvasOffset.Y)
	x1 := int(PxCanvas.PxCols*PxCanvas.PxSize + int(PxCanvas.CanvasOffset.X))
	y1 := int(PxCanvas.PxRows*PxCanvas.PxSize + int(PxCanvas.CanvasOffset.Y))
	return image.Rect(x0, y0, x1, y1)
}

func CheckMousePositionInBounds(mousePosition fyne.Position, bounds image.Rectangle) bool {
	if mousePosition.X >= float32(bounds.Min.X) && mousePosition.X < float32(bounds.Max.X) &&
		mousePosition.Y >= float32(bounds.Min.Y) && mousePosition.Y < float32(bounds.Max.Y) {
		return true
	}
	return false
}

func NewPxCanvas(state *apptype.State, config apptype.PxCanvasConfig) *PxCanvas {
	pxCanvas := &PxCanvas{
		PxCanvasConfig: config,
		appState:       state,
	}
	pxCanvas.PixelData = NewBlankImage(pxCanvas.PxCols, pxCanvas.PxRows, color.NRGBA{128, 128, 128, 128})
	pxCanvas.ExtendBaseWidget(pxCanvas)
	return pxCanvas
}

func NewBlankImage(cols, rows int, color color.Color) image.Image {
	newImage := image.NewNRGBA(image.Rect(0, 0, cols, rows))
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			newImage.Set(x, y, color)
		}
	}
	return newImage
}

func (pxCanvas *PxCanvas) CreateRenderer() fyne.WidgetRenderer {
	canvasImage := canvas.NewImageFromImage(pxCanvas.PixelData)
	canvasImage.ScaleMode = canvas.ImageScalePixels
	canvasImage.FillMode = canvas.ImageFillContain

	canvasBorder := make([]canvas.Line, 4)
	for i := 0; i < len(canvasBorder); i++ {
		canvasBorder[i].StrokeWidth = 2
		canvasBorder[i].StrokeColor = color.NRGBA{255, 255, 255, 255}
	}

	renderer := &PxCanvasRenderer{
		pxCanvas:     pxCanvas,
		canvasImage:  canvasImage,
		canvasBorder: canvasBorder,
	}

	pxCanvas.renderer = renderer
	return renderer
}

func (pxCanvas *PxCanvas) TryPan(previousCoord *fyne.PointEvent, event *desktop.MouseEvent) {
	if previousCoord != nil && event.Button == desktop.MouseButtonTertiary {
		pxCanvas.Pan(*previousCoord, event.PointEvent)
	}
}

// Brushable interface
func (pxCanvas *PxCanvas) SetColor(newColor color.Color, x, y int) {

	if nrgba, ok := pxCanvas.PixelData.(*image.NRGBA); ok {
		nrgba.Set(x, y, newColor)
	}

	if rgba, ok := pxCanvas.PixelData.(*image.RGBA); ok {
		rgba.Set(x, y, newColor)
	}

	pxCanvas.Refresh()
}

func (pxCanvas *PxCanvas) MouseToCanvasXY(event *desktop.MouseEvent) (*int, *int) {
	bounds := pxCanvas.Bounds()
	if !CheckMousePositionInBounds(event.Position, bounds) {
		return nil, nil
	}
	pxSize := float32(pxCanvas.PxSize)
	xOffeset := pxCanvas.CanvasOffset.X
	yOffeset := pxCanvas.CanvasOffset.Y

	x := int((event.Position.X - xOffeset) / pxSize)
	y := int((event.Position.Y - yOffeset) / pxSize)

	return &x, &y
}

func (pxCanvas *PxCanvas) LoadImage(img image.Image) {
	dimensions := img.Bounds()

	pxCanvas.PxCanvasConfig.PxCols = dimensions.Dx()
	pxCanvas.PxCanvasConfig.PxRows = dimensions.Dy()

	pxCanvas.PixelData = img
	pxCanvas.reloadImage = true
	pxCanvas.Refresh()
}

func (pxCanvas *PxCanvas) NewDrawing(cols, rows int) {
	pxCanvas.appState.SetFilePath("")
	pxCanvas.PxCols = cols
	pxCanvas.PxRows = rows
	pixelData := NewBlankImage(cols, rows, color.NRGBA{128, 128, 128, 255})
	pxCanvas.LoadImage(pixelData)
}
