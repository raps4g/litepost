package ui

import "github.com/rivo/tview"

func (s *State) addModal(name string, p tview.Primitive, x, y, width, height int) {
    modal := tview.NewGrid().
		SetColumns(x, width, 0).
		SetRows(y, height, 0).
		AddItem(p, 1, 1, 1, 1, 0, 0, true)
    s.Pages.AddPage(name, modal, true, true)
}

func (s *State) removeModal(name string) {
    s.Pages.RemovePage(name)
}

