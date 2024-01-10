package pxcanvas

import (
	"fyne.io/fyne"
	"fyne.io/fyne/v2/canvas"
)

type PxCanvasRenderer struct {
	pxCanvas     *PxCanvas
	canvasImage  *canvas.Image
	canvasBorder []canvas.Line
}

// WidgetRenderer interface implementation.
func (pxCanvasRenderer *PxCanvasRenderer) MinSize() fyne.Size {
	return pxCanvasRenderer.pxCanvas.DrawingArea
}

// WidgetRenderer interface implementation.
func (pxCanvasRenderer *PxCanvasRenderer) Objects() []fyne.CanvasObject {
	objects := make([]fyne.CanvasObject, 0, 5)
	for i := 0; i < len(pxCanvasRenderer.canvasBorder); i++ {
		objects = append(objects, &pxCanvasRenderer.canvasBorder[i])
	}
	objects = append(objects, pxCanvasRenderer.canvasImage)
	return objects
}

// WidgetRenderer interface implementation.
func (pxCanvasRenderer *PxCanvasRenderer) Destroy() {}

func (pxCanvasRenderer *PxCanvasRenderer) Layout(size fyne.Size) {
	pxCanvasRenderer.LayoutCanvas(size)
	pxCanvasRenderer.LayoutBorders(size)
}
func (pxCanvasRenderer *PxCanvasRenderer) Refresh(size fyne.Size) {
	if pxCanvasRenderer.pxCanvas.reloadImage {
		pxCanvasRenderer.canvasImage = canvas.NewImageFromImage(pxCanvasRenderer.pxCanvas.PixelData)
		pxCanvasRenderer.canvasImage.ScaleMode = canvas.ImageScalePixels
		pxCanvasRenderer.canvasImage.FillMode = canvas.ImageFillContain
		pxCanvasRenderer.pxCanvas.reloadImage = false
	}
	pxCanvasRenderer.Layout(pxCanvasRenderer.pxCanvas.Size())
	canvas.Refresh(pxCanvasRenderer.canvasImage)
}

func (pxCanvasRenderer *PxCanvasRenderer) LayoutCanvas(size fyne.Size) {
	imgPxWidth := pxCanvasRenderer.pxCanvas.PxCols
	imgPxHeight := pxCanvasRenderer.pxCanvas.PxRows
	pxSize := pxCanvasRenderer.pxCanvas.PxSize
	pxCanvasRenderer.canvasImage.Move(fyne.NewPos(pxCanvasRenderer.pxCanvas.CanvasOffset.X, pxCanvasRenderer.pxCanvas.CanvasOffset.Y))
	pxCanvasRenderer.canvasImage.Resize(fyne.NewSize(float32(imgPxWidth*pxSize), fyne.NewSize(float32(imgPxHeight*pxSize))))
}

func (pxCanvasRenderer *PxCanvasRenderer) LayoutBorders(size fyne.Size) {
	offset := pxCanvasRenderer.pxCanvas.CanvasOffset
	imgWidth := pxCanvasRenderer.pxCanvas.Size().Width
	imgHeight := pxCanvasRenderer.pxCanvas.Size().Height

	left := &pxCanvasRenderer.canvasBorder[0]
	left.Position1 = fyne.NewPos(offset.X, offset.Y)
	left.Position2 = fyne.NewPos(offset.X, offset.Y+imgHeight)

	top := &pxCanvasRenderer.canvasBorder[1]
	top.Position1 = fyne.NewPos(offset.X+imgWidth, offset.Y)
	top.Position2 = fyne.NewPos(offset.X+imgWidth, offset.Y+imgHeight)

	right := &pxCanvasRenderer.canvasBorder[2]
	right.Position1 = fyne.NewPos(offset.X, offset.Y)
	right.Position2 = fyne.NewPos(offset.X, offset.Y-imgHeight)

	bottom := &pxCanvasRenderer.canvasBorder[3]
	bottom.Position1 = fyne.NewPos(offset.X, offset.Y+imgHeight)
	bottom.Position2 = fyne.NewPos(offset.X+imgWidth, offset.Y)
}
