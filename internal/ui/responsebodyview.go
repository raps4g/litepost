package ui

import (
	"unicode/utf8"

	"github.com/gdamore/tcell/v2"
	"github.com/raps4g/litepost/internal/core"
	"github.com/rivo/tview"
)

func (ui *Ui) NewResponseBodyView(req *core.Request) *tview.TextView{
    ResponseBodyView := tview.NewTextView()
    ResponseBodyView.
        SetToggleHighlights(true).
        SetDynamicColors(true).
        SetTextStyle(defaultStyle).
        SetBackgroundColor(backgroundColor).
        SetDrawFunc(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
            maxWidth := utf8.RuneCountInString(req.Status) 
            if maxWidth > 0 {
                tview.Print(screen, " " + req.Status + " " , x+width-maxWidth-4, y-1, maxWidth+2,0, tcell.ColorPink)
            } 
            return x, y, width-1, height-1
        })

    return ResponseBodyView
}

