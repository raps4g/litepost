package ui

import (
    "io"
    "net/http"
    "strconv"
    "unicode/utf8"

    "github.com/gdamore/tcell/v2"
    "github.com/rivo/tview"
)

func (s *State) SetResponseView() *State {
    s.ResponseView = tview.NewTextView()
    s.ResponseView.
        SetToggleHighlights(true).
        SetDynamicColors(true).
        SetTextStyle(defaultStyle).
        SetTitle(" Response ").
        SetTitleColor(unfocusedTitleColor).
        SetTitleAlign(0).
        SetBorder(true).
        SetBackgroundColor(backgroundColor).
        SetBorderStyle(borderStyle).
        SetBlurFunc(func () {
            s.ResponseView.SetTitleColor(unfocusedTitleColor)
        }).
        SetFocusFunc(func () {
            s.ResponseView.SetTitleColor(focusedTitleColor)
        }).SetDrawFunc(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {

            maxWidth := utf8.RuneCountInString(s.Status) 
            tview.Print(screen, s.Status, x+width-maxWidth-2, y+1, maxWidth,0, tcell.ColorPink)
            return x+1, y+1, width-2, height-2
        }).
        SetInputCapture(s.responseViewInputpCapture)

    return s
}

func (s *State) mapResponseToUi(resp *http.Response) {
    bodyBytes, err := io.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    responseBody, _ := prettyPrintJSON(string(bodyBytes))
    s.Status = strconv.Itoa(resp.StatusCode) + " " + http.StatusText(resp.StatusCode)
    s.Response = responseBody
}

func (s *State) responseViewInputpCapture(event *tcell.EventKey) *tcell.EventKey {
    if event.Modifiers() == tcell.ModCtrl {
        switch event.Key() {
        case tcell.KeyLeft:
            s.App.SetFocus(s.HistoryList)   
        case tcell.KeyUp: 
            s.App.SetFocus(s.RequestInput)   
        }
    }
    return event
}
