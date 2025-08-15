package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/raps4g/litepost/internal/core"
	"github.com/rivo/tview"
)

func (ui *Ui) NewUrlInput(req *core.Request) *tview.InputField {
    
    UrlInput := tview.NewInputField()
    UrlInput.
        SetFieldStyle(defaultStyle).
        SetChangedFunc(func(text string) {
            req.Url = text
        }).
        //SetText("http://localhost:8080/login").
        SetLabelStyle(defaultStyle).
        SetPlaceholderStyle(placeholderStyle).
        SetDrawFunc(func(screen tcell.Screen, x int, y int, width int, height int) (int, int, int, int) {
            tview.Print(screen, req.Methods[req.SelectedMethod], x+2, y+1, 8, 0, foregroundColor)
            screen.SetContent(x+9, y, tview.BoxDrawingsLightDownAndHorizontal, nil, borderStyle)
            screen.SetContent(x+9, y+1, tview.BoxDrawingsLightVertical, nil, borderStyle)
            screen.SetContent(x+9, y+2, tview.BoxDrawingsLightUpAndHorizontal, nil, borderStyle)
            return x + 11, y + 1, width - 12, height - 2
        }).
        SetBackgroundColor(backgroundColor).
        SetBorderStyle(borderStyle).
        SetBorder(true).
        SetTitle(" URL ").
        SetTitleColor(bluredTitleColor).
        SetTitleAlign(0)

    return UrlInput
}
