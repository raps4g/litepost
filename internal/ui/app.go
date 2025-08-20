package ui

import (
	"github.com/raps4g/litepost/internal/core"
	"github.com/rivo/tview"
)

func SetApp(req *core.Request, vars *map[string]string, hist *[]core.Request) *tview.Application {
   
    box := tview.NewBox().SetBackgroundColor(backgroundColor)
    ui := Ui{}
    ui.App = tview.NewApplication()
    ui.UrlInput = ui.NewUrlInput(req)
    ui.RequestBodyInput = ui. NewRequestBodyInput(req)
    ui.RequestHeadersTable = ui.NewKeyValueTable(&req.ReqHeaders, "header", true)
    ui.ResponseBodyView = ui.NewResponseBodyView(req)
    ui.ResponseHeadersTable = ui.NewKeyValueTable(&req.RespHeaders, "header", false)
    ui.VariablesTable = ui.NewKeyValueTable(vars, "variable", true)
    ui.HistoryTable = ui.NewHistoryTable(req)
    ui.RequestPages = ui.NewTabPages("Request")
    ui.RequestPages.AddTab("Variables", ui.VariablesTable.Table)
    ui.RequestPages.AddTab("Headers", ui.RequestHeadersTable.Table)
    ui.RequestPages.AddTab("Body", ui.RequestBodyInput)
    ui.ResponsePages = ui.NewTabPages("Response")
    ui.ResponsePages.AddTab("Cookies", box)
    ui.ResponsePages.AddTab("Headers", ui.ResponseHeadersTable.Table)
    ui.ResponsePages.AddTab("Body", ui.ResponseBodyView)
    ui.MethodList = ui.NewMethodList(req)
    ui.ParsedVariablesTable = ui.NewKeyValueTable(&req.ParsedVariables, "variable", false)
    ui.ParsedVariablesTable.SetSelectable(true,false).SetBorder(true).SetTitle(" Parsed Variables ").SetTitleColor(focusedTitleColor).SetBorderStyle(borderStyle)
    ui.HelpPage = ui.NewHelpPage()

    ui.FooterPages = tview.NewPages()
    setCustomBorders()

    ui.newFooter(ui.FooterPages, "help_footer", "  F1: Help", placeholderStyle)

    mainGrid := tview.NewGrid().
        SetColumns(25,0).
        SetRows(3,0,0,1).
        AddItem(ui.HistoryTable, 0,0,3,1,0,0,false).
        AddItem(ui.UrlInput, 0,1,1,1,0,0,true).
        AddItem(ui.RequestPages, 1,1,1,1,0,0,false).
        AddItem(ui.ResponsePages, 2,1,1,1,0,0,false).
        AddItem(ui.FooterPages, 3,0,1,2,0,0,false)
        

    ui.MainPages = tview.NewPages()
    ui.MainPages.AddPage("main_grid", mainGrid, true, true)
    ui.FocusLayout = FocusLayout{}
    ui.FocusLayout.Matrix = [][]tview.Primitive{
        {ui.HistoryTable, ui.UrlInput},
        {ui.HistoryTable, ui.RequestPages},
        {ui.HistoryTable, ui.ResponsePages},
    }
    ui.FocusLayout.X = 0
    ui.FocusLayout.Y = 1
    ui.FocusLayout.OnFocus = ui.UrlInput 
    ui.SetFocusHandlers()
    ui.SetAppInputCapture(req, vars, hist)

    ui.loadResponse(req)

    ui.App.SetRoot(ui.MainPages, true)
    return ui.App
}
