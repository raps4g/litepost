package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (s *State) SetMethodList() *State {
    s.MethodList = tview.NewList()
    s.MethodList.SetMainTextStyle(defaultStyle).
        ShowSecondaryText(false).
        SetSelectedFocusOnly(true).
        SetHighlightFullLine(true).
        SetSelectedFunc(func(i int, s1, s2 string, r rune) {
            s.SelectedMethod = i
        }).
        SetShortcutStyle(placeholderStyle).
        SetTitle(" Method ").
        SetTitleColor(focusedTitleColor).
        SetBorder(true).
        SetBackgroundColor(backgroundColor).
        SetBorderStyle(borderStyle).
        SetInputCapture(s.methodListInputCapture)

    for i, m := range s.Methods {
        s.MethodList.AddItem(m, "", rune('1'+i), nil)
    }

    return s
}

func (s *State) methodListInputCapture(event *tcell.EventKey) *tcell.EventKey {
    if event.Key() == tcell.KeyEnter {
        s.removeModal("method_modal")
        s.removeFooter("method_footer")
    }    
    return event
}
