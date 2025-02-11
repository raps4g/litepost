package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/raps4g/litepost/internal/core"
	"github.com/rivo/tview"
)

var (
    defaultStyle = tcell.StyleDefault
    foregroundColor, backgroundColor, _ = defaultStyle.Decompose()
    placeholderColor = tcell.ColorGray
    placeholderStyle = tcell.StyleDefault.Foreground(placeholderColor)
    selectedStyle = tcell.StyleDefault
    borderColor = tcell.ColorGray
    borderStyle = tcell.StyleDefault.Foreground(borderColor)
    focusedTitleColor = tcell.ColorGreen
    unfocusedTitleColor = tcell.ColorBlue
)

type Ui struct {
    Pages *tview.Pages
    MainGrid *tview.Grid
    RequestGrid *tview.Grid
    UrlInput *tview.InputField
    RequestInput *tview.TextArea
    ResponseView *tview.TextView
    HistoryList *tview.List
    MethodList *tview.List
    HeaderTable *tview.Table
    Footers *tview.Pages
}

type State struct {
    App *tview.Application
    Ui
    core.Request
    History []core.Request
}

func InitUi() *State {

    setCustomBorders()

    s := &State{} 
    s.Methods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
    s.SelectedMethod = 0
    s.History = []core.Request{}
    
    s.Headers = map[string]string {
        "Content-Type": "application/json",
        "User-Agent": "Go-http-client/1.1",
    }

    s.App = tview.NewApplication()
    s.SetUrlInput().
        SetMethodList().
        SetRequestInput().
        SetResponseView().
        SetHistoryList().
        SetHeaderTable()

    s.addFooter("main_footer", 
        " Ctrl + ⬅/⬇/⬆/➡: Navigate \t\t Alt + H: Edit headers \t\t Alt + M: Select Method \t\t Ctrl + Enter: Send",
        placeholderStyle)
    
    s.RequestGrid = tview.NewGrid().
        SetRows(3,0,0).
        AddItem(s.UrlInput,0,0,1,1,1,0,true).
        AddItem(s.RequestInput,1,0,1,1,0,0,false).
        AddItem(s.ResponseView,2,0,1,1,1,2,false)

    s.MainGrid = tview.NewGrid().
        SetColumns(30,0).
        SetRows(0,1).
        AddItem(s.HistoryList,0,0,1,1,0,80,false).
        AddItem(s.RequestGrid,0,1,1,1,0,80,true).
        AddItem(s.RequestGrid,0,0,1,2,0,0,true).
        AddItem(s.Footers,1,0,1,2,0,0,false)

    s.Pages = tview.NewPages().
        AddPage("main_grid", s.MainGrid, true, true)
    s.App.SetInputCapture(s.appHandler)
    s.App.SetRoot(s.Pages, true)
    
    return s
}
