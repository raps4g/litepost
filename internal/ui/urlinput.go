package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (s *State) SetUrlInput() *State {
    
    urlPlaceholder := "Enter URL"
    s.UrlInput = tview.NewInputField()
    s.UrlInput.
        SetFieldStyle(defaultStyle).
        SetChangedFunc(func(text string) {
            s.Request.Url = text
        }).
        SetText("https://google.com").
        SetLabelStyle(defaultStyle).
        SetPlaceholderStyle(placeholderStyle).
        SetPlaceholder(urlPlaceholder).
        SetDrawFunc(func(screen tcell.Screen, x int, y int, width int, height int) (int, int, int, int) {
            tview.Print(screen, s.Request.Methods[s.SelectedMethod], x+2, y+1, 8, 0, foregroundColor)
            screen.SetContent(x+9, y, tview.BoxDrawingsLightDownAndHorizontal, nil, borderStyle)
            screen.SetContent(x+9, y+1, tview.BoxDrawingsLightVertical, nil, borderStyle)
            screen.SetContent(x+9, y+2, tview.BoxDrawingsLightUpAndHorizontal, nil, borderStyle)
            return x + 11, y + 1, width - 12, height - 2
        }).
        SetBackgroundColor(backgroundColor).
        SetBorderStyle(borderStyle).
        SetBorder(true).
        SetTitle(" URL ").
        SetTitleColor(unfocusedTitleColor).
        SetTitleAlign(0).
        SetBlurFunc(func () {
            s.UrlInput.SetPlaceholder(urlPlaceholder).
                SetTitleColor(unfocusedTitleColor)
        }).
        SetTitleColor(tcell.ColorBlue).
        SetFocusFunc(func () {
            s.UrlInput.SetPlaceholder("").
                SetTitleColor(focusedTitleColor)
        }).
        SetInputCapture(s.urlInputInputpCapture)

    return s
}

func (s *State) urlInputInputpCapture(event *tcell.EventKey) *tcell.EventKey {

    if event.Modifiers() == tcell.ModCtrl {
       switch event.Key() {
        case tcell.KeyLeft:
            s.App.SetFocus(s.HistoryList)   
        case tcell.KeyDown:
            s.App.SetFocus(s.RequestInput)   
       } 
    }
    return event
}
