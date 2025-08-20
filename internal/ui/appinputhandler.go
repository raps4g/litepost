package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/raps4g/litepost/internal/core"
)

func (ui *Ui) SetAppInputCapture(req *core.Request, variables *map[string]string, history *[]core.Request) {
	ui.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// ui.ResponseBodyView.SetText(fmt.Sprintf("Key: %v, Rune: %q, Modifiers: %v\n", event.Key(), event.Rune(), event.Modifiers()))
		if event.Modifiers() == tcell.ModCtrl {
			switch event.Key() {
			case tcell.KeyUp:
				ui.FocusLayout.MoveUp()
				ui.App.SetFocus(ui.FocusLayout.OnFocus)
			case tcell.KeyDown:
				ui.FocusLayout.MoveDown()
				ui.App.SetFocus(ui.FocusLayout.OnFocus)
			case tcell.KeyLeft:
				ui.FocusLayout.MoveLeft()
				ui.App.SetFocus(ui.FocusLayout.OnFocus)
			case tcell.KeyRight:
				ui.FocusLayout.MoveRight()
				ui.App.SetFocus(ui.FocusLayout.OnFocus)
			}
		}

		if event.Key() == tcell.KeyCtrlT {
			ui.addModal("method_list", ui.MethodList, 0, 5, 40, 10, true)
		}

		if event.Key() == tcell.KeyCtrlV {
			ui.addModal("variable_list", ui.ParsedVariablesTable, 0, 0, 60, 0, false)
		}

		_, _, h, w := ui.MainPages.GetRect()
		if event.Key() == tcell.KeyF1 {
			ui.addModal("help_page", ui.HelpPage, 0, 0, h-2, w-4, false)
		}

		if event.Key() == tcell.KeyCtrlSpace {
			ui.SendRequestHandler(req, variables)
		}
		return event
	})

}

func (ui *Ui) SendRequestHandler(req *core.Request, variables *map[string]string) {
	err := core.SendHttpRequest(req, variables)

	if err != nil {
		req.RespBody = "[red]" + core.GetErrorDescription(err)
		req.Status = ""
		ui.loadResponse(req)
		return
	}
	ui.AddHistoryEntry(req)
	ui.loadResponse(req)
}
