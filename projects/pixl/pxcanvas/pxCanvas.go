package pxcanvas

import (
	"image"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
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

func NewBlankImage(cols, rows int, color color.Color) image.Image {
	newImage := image.NewNRGBA(image.Rect(0, 0, cols, rows))
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			newImage.Set(x, y, color)
		}
	}
	return newImage
}

func newPxCanvas(state *apptype.State, config apptype.PxCanvasConfig) *PxCanvas {
	pxCanvas := &PxCanvas{
		PxCanvasConfig: config,
		appState:       state,
	}
	pxCanvas.PixelData = NewBlankImage(pxCanvas.PxCols, pxCanvas.PxRows, color.NRGBA64{128, 128, 128, 128})
	pxCanvas.ExtendBaseWidget(pxCanvas)
	return pxCanvas
}

func (pxCanvas *PxCanvas) CreateRenderer() fyne.WidgetRenderer {
	canvasImage := canvas.NewImageFromImage(pxCanvas.PixelData)
	canvasImage.ScaleMode = canvas.ImageScalePixels
	canvasImage.FillMode = canvas.ImageFillContain

	canvasBorder := make([]canvas.Line, 4)
	for i := 0; i < len(canvasBorder); i++ {
		canvasBorder[i].StrokeColor = color.NRGBA64{100, 100, 100, 255}
		canvasBorder[i].StrokeWidth = 2
	}

	renderer := &PxCanvasRenderer{
		pxCanvas:     pxCanvas,
		canvasImage:  canvasImage,
		canvasBorder: canvasBorder,
	}

	pxCanvas.renderer = renderer
	return renderer
}
