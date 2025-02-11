package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (s *State) addFooter(name, text string, style tcell.Style) {
    if s.Footers == nil {
        s.Footers = tview.NewPages()
    }
    footer := tview.NewTextView()
    footer.SetBackgroundColor(backgroundColor)
    footer.SetTextStyle(style)
    footer.SetText(text)
    s.Footers.AddPage(name, footer, true, true)
    s.Footers.SendToFront(name)
}

func (s *State) removeFooter(name string) {
    s.Footers.RemovePage(name)
}
