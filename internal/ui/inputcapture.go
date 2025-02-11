package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/raps4g/litepost/internal/core"
)


func (s *State) appHandler(event *tcell.EventKey) *tcell.EventKey {

    switch {
    case event.Modifiers() == tcell.ModAlt &&
        event.Rune() == 'm':
        s.addFooter("method_footer", "Enter: Select \t\t Esc: Close", placeholderStyle)
        s.addModal("method_modal", s.MethodList, 0, 10, 40, 9)
    
    case event.Modifiers() == tcell.ModAlt &&
        event.Rune() == 'h' :
        s.addFooter("header_footer", "Enter: Edit field \t\t A: Add header \t\t D: Delete header \t\t Esc: Close", placeholderStyle)
        s.addModal("header_table", s.HeaderTable, 0, 10, 60, 8)
    
    case event.Key() == tcell.KeyCtrlSpace:
        s.handleSendRequest()
        s.saveCurrentRequest()
        s.addHistoryList(0)
    } 
    return event
}

func (s *State) handleSendRequest() {
    resp, err := core.SendHttpRequest(s.Request)
    if err != nil {
        panic(err)
    }

    s.mapResponseToUi(resp)
    s.ResponseView.SetText(s.Response)
    s.prettyRequestInput()

}

