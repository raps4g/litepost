package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (ui *Ui) addModal(name string, p tview.Primitive, x, y, width, height int, enterCloses bool) {
    modal := tview.NewGrid().
		SetColumns(x, width, 0).
		SetRows(y, height, 0).
		AddItem(p, 1, 1, 1, 1, 0, 0, true)

    modal.
        SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
            if (event.Key() == tcell.KeyEnter && enterCloses) ||
                event.Key() == tcell.KeyESC {
                primitiveHandler := p.InputHandler()
                primitiveHandler(event, nil)
                ui.FocusLayout.OpenDialog = nil
                ui.MainPages.RemovePage(name)
                ui.App.SetFocus(ui.FocusLayout.OnFocus)
            }
            return event
        })
    ui.FocusLayout.OpenDialog = modal
    ui.MainPages.AddPage(name, modal, true, true)
}
