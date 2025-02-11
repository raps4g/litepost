package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (s *State) SetHistoryList() *State {
    s.HistoryList = tview.NewList()
    s.HistoryList.
        SetMainTextStyle(defaultStyle).
        SetSecondaryTextStyle(defaultStyle).
        SetSelectedFocusOnly(true).
        SetTitle(" History ").
        SetTitleColor(unfocusedTitleColor).
        SetTitleAlign(0).
        SetBorder(true).
        SetBackgroundColor(backgroundColor).
        SetBorderStyle(borderStyle).
        SetBlurFunc(func () {
            s.HistoryList.SetTitleColor(unfocusedTitleColor)
        }).
        SetFocusFunc(func () {
            s.HistoryList.SetTitleColor(focusedTitleColor)
        }).
        SetInputCapture(s.historyInputCapture)

    return s
}

func (s *State) historyInputCapture(event *tcell.EventKey) *tcell.EventKey {

    if event.Modifiers() == tcell.ModCtrl && 
        event.Key() == tcell.KeyRight {
        s.App.SetFocus(s.RequestGrid)   
    } 
    return event
}

func (s *State) addHistoryList(i int) {
    name := s.Methods[s.History[i].SelectedMethod]
    name += " - "
    name += s.History[i].Url
    s.HistoryList.AddItem(name, "/path/foo/barr", 0, nil) 
}
