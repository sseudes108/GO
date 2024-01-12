package pxcanvas

import (
	"fyne.io/fyne/v2"
)

func (pxCanvas *PxCanvas) Scale(direction int) {
	switch {
	case direction > 0:
		pxCanvas.PxSize++
	case direction < 0:
		if pxCanvas.PxSize > 2 {
			pxCanvas.PxSize--
		}
	default:
		pxCanvas.PxSize = 10
	}
}

func (pxCanvas *PxCanvas) Pan(previousCoord, currentCoord fyne.PointEvent) {
	xDifference := currentCoord.Position.X - previousCoord.Position.X
	yDifference := currentCoord.Position.Y - previousCoord.Position.Y

	pxCanvas.CanvasOffset.X += xDifference
	pxCanvas.CanvasOffset.Y += yDifference
	pxCanvas.Refresh()
}
