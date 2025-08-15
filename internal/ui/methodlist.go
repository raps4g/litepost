package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/raps4g/litepost/internal/core"
	"github.com/rivo/tview"
)

func (ui *Ui) NewMethodList(req *core.Request) *tview.List {
    MethodList := tview.NewList()
    MethodList.SetMainTextStyle(defaultStyle).
        ShowSecondaryText(false).
        SetSelectedFocusOnly(true).
        SetHighlightFullLine(true).
        SetSelectedFunc(func(i int, s1, s2 string, r rune) {
            req.SelectedMethod = i
        }).
        SetShortcutStyle(placeholderStyle).
        SetTitle(" Method ").
        SetTitleColor(focusedTitleColor).
        SetBorder(true).
        SetBackgroundColor(backgroundColor).
        SetBorderStyle(borderStyle)

    for i, m := range req.Methods {
        MethodList.AddItem(m, "", rune('1'+i), nil)
    }

    return MethodList
}

func (ui *Ui) methodListInputCapture(event *tcell.EventKey) *tcell.EventKey {
    if event.Key() == tcell.KeyEnter {
    }    
    return event
}
