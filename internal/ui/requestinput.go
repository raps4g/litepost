package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (s *State) SetRequestInput() *State {

    bodyPlaceholder := "Enter request body"
    s.RequestInput = tview.NewTextArea()
    s.RequestInput.
        SetPlaceholder(bodyPlaceholder).
        SetTextStyle(defaultStyle).
        SetChangedFunc(func() {
            s.Body = s.RequestInput.GetText()  
        }).
        SetLabelStyle(defaultStyle).
        SetPlaceholderStyle(placeholderStyle).
        SetBorderStyle(borderStyle).
        SetTitle(" Request body ").
        SetTitleColor(unfocusedTitleColor).
        SetTitleAlign(0).
        SetBorder(true).
        SetBlurFunc(func () {
            s.RequestInput.SetPlaceholder(bodyPlaceholder).
                SetTitleColor(unfocusedTitleColor)
        }).
        SetFocusFunc(func () {
            s.RequestInput.SetPlaceholder("").
                SetTitleColor(focusedTitleColor)
        }).
        SetInputCapture(s.requestInputInputpCapture)

    return s
}

func (s *State) requestInputInputpCapture(event *tcell.EventKey) *tcell.EventKey {

    if event.Modifiers() == tcell.ModCtrl {
        switch event.Key() {
        case tcell.KeyLeft:
            s.App.SetFocus(s.HistoryList)   
        case tcell.KeyUp:
            s.App.SetFocus(s.UrlInput)   
        case tcell.KeyDown:
            s.App.SetFocus(s.ResponseView)   
        }
    }
    return event
}

func (s *State) prettyRequestInput() {
    prettyJson, _ := prettyPrintJSON(s.Body)
    s.RequestInput.SetText(prettyJson, true)
}
