package pxcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"zerotomastery.io/pixel/pxcanvas/brush"
)

func (pxCanvas *PxCanvas) Scrolled(event *fyne.ScrollEvent) {
	pxCanvas.Scale(int(event.Scrolled.DY))
	pxCanvas.Refresh()
}

func (pxCanvas *PxCanvas) MouseMoved(event *desktop.MouseEvent) {

	if x, y := pxCanvas.MouseToCanvasXY(event); x != nil && y != nil {
		brush.TryBrush(pxCanvas.appState, pxCanvas, event)
		cursor := brush.Cursor(pxCanvas.PxCanvasConfig, pxCanvas.appState.BrushType, event, *x, *y)
		pxCanvas.renderer.SetCursor(cursor)
	} else {
		pxCanvas.renderer.SetCursor(make([]fyne.CanvasObject, 0))
	}

	pxCanvas.TryPan(pxCanvas.mouseState.previousMouseCoord, event)
	pxCanvas.Refresh()
	pxCanvas.mouseState.previousMouseCoord = &event.PointEvent
}

func (pxCanvas *PxCanvas) MouseIn(event *desktop.MouseEvent) {}
func (pxCanvas *PxCanvas) MouseOut()                         {}

func (pxCanvas *PxCanvas) MouseUp(event *desktop.MouseEvent) {}

func (pxCanvas *PxCanvas) MouseDown(event *desktop.MouseEvent) {
	brush.TryBrush(pxCanvas.appState, pxCanvas, event)
	pxCanvas.Refresh()
}
