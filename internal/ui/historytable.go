package ui

import (
	"net/url"

	"github.com/gdamore/tcell/v2"
	"github.com/raps4g/litepost/internal/core"
	"github.com/rivo/tview"
)

func (ui *Ui) NewHistoryTable(req *core.Request) *tview.Table {
    HistoryTable := tview.NewTable()
    HistoryTable.
		SetSelectable(false, false).
		SetTitle(" History ").
        SetTitleAlign(0).
		SetTitleColor(bluredTitleColor).
        SetBorder(true).
		SetBorderStyle(borderStyle).
		SetBackgroundColor(backgroundColor).
        SetFocusFunc(func() {
            HistoryTable.SetSelectable(true,false).
            SetTitleColor(focusedTitleColor)
        }).
        SetBlurFunc(func() {
            HistoryTable.SetSelectable(false,false).
            SetTitleColor(bluredTitleColor)
        }).
        SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
            if event.Key() == tcell.KeyEnter {
                ui.loadSelection(req)
            }
            return event
        })

    return HistoryTable
}

func (ui *Ui) AddHistoryEntry(req *core.Request) {
    u, _ := url.Parse(req.Url)
    host := u.Host
    path := u.Path
    method := req.Methods[req.SelectedMethod]

    newReq := *req

    newReq.ReqHeaders = make(map[string]string)
    for key, value := range req.ReqHeaders {
        newReq.ReqHeaders[key] = value
    }
    
    newReq.RespHeaders = make(map[string]string)
    for key, value := range req.RespHeaders {
        newReq.RespHeaders[key] = value
    }
    newReq.ParsedVariables = make(map[string]string)
    for key, value := range req.ParsedVariables {
        newReq.ParsedVariables[key] = value
    }

    ui.HistoryTable.InsertRow(0)
    ui.HistoryTable.SetCell(0, 0, tview.NewTableCell( " " + method + " " + host ).SetReference(newReq))
    if path != "" {
        ui.HistoryTable.InsertRow(1)
        ui.HistoryTable.SetCell(1, 0, tview.NewTableCell("  ► " + path).SetSelectable(false))
    }
}
func (ui *Ui) loadSelection(req *core.Request) {
    row, col := ui.HistoryTable.GetSelection()
    
    cell := ui.HistoryTable.GetCell(row, col)
    if cell != nil {
        if request, ok := cell.GetReference().(core.Request); ok {
            req.Url = request.Url
            req.SelectedMethod = request.SelectedMethod
            req.ReqBody = request.ReqBody
            req.RespBody = request.RespBody
            req.Status = request.Status
            req.Methods = request.Methods

            req.ReqHeaders = make(map[string]string)
            for k, v := range request.ReqHeaders {
                req.ReqHeaders[k] = v
            }

            req.RespHeaders = make(map[string]string)
            for k, v := range request.RespHeaders {
                req.RespHeaders[k] = v
            }
            req.ParsedVariables = make(map[string]string)
            for k, v := range request.ParsedVariables {
                req.ParsedVariables[k] = v
            }

            ui.loadResponse(req)
            ui.loadRequest(req)
        } else {
            panic("Error loading history entry")
        }
    }
}

