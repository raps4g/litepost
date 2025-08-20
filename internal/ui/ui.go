package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/raps4g/litepost/internal/core"
	"github.com/rivo/tview"
)

var (
	defaultStyle                        = tcell.StyleDefault
	foregroundColor, backgroundColor, _ = defaultStyle.Decompose()
	placeholderColor                    = tcell.ColorGray
	placeholderStyle                    = tcell.StyleDefault.Foreground(placeholderColor)
	selectedStyle                       = tcell.StyleDefault
	borderColor                         = tcell.ColorGray
	borderStyle                         = tcell.StyleDefault.Foreground(borderColor)
	focusedTitleColor                   = tcell.ColorGreen
	bluredTitleColor                      = tcell.ColorBlue
)

type Ui struct {
	App                  *tview.Application
	UrlInput             *tview.InputField
	RequestBodyInput     *tview.TextArea
	RequestHeadersTable  *KeyValueTable
	ResponseBodyView     *tview.TextView
	ResponseHeadersTable *KeyValueTable
	HistoryTable         *tview.Table
	MethodList           *tview.List
	VariablesTable       *KeyValueTable
	RequestPages         *TabPage
	ResponsePages        *TabPage
	FocusLayout          FocusLayout
	FooterPages          *tview.Pages
    MainPages            *tview.Pages
    ParsedVariablesTable *KeyValueTable
    HelpPage             *tview.TextView
}

func setCustomBorders() {
	tview.Borders.TopLeft = tview.BoxDrawingsLightDownAndRight
	tview.Borders.TopRight = tview.BoxDrawingsLightDownAndLeft
	tview.Borders.BottomLeft = tview.BoxDrawingsLightUpAndRight
	tview.Borders.BottomRight = tview.BoxDrawingsLightUpAndLeft
	tview.Borders.HorizontalFocus = tview.BoxDrawingsLightHorizontal
	tview.Borders.VerticalFocus = tview.BoxDrawingsLightVertical
	tview.Borders.TopLeftFocus = tview.BoxDrawingsLightDownAndRight
	tview.Borders.TopRightFocus = tview.BoxDrawingsLightDownAndLeft
	tview.Borders.BottomLeftFocus = tview.BoxDrawingsLightUpAndRight
	tview.Borders.BottomRightFocus = tview.BoxDrawingsLightUpAndLeft
}

func (ui *Ui) loadResponse(req *core.Request) {
	ui.ResponseBodyView.SetText(req.RespBody)
	ui.ResponseHeadersTable.updateTableFromData()
    ui.ParsedVariablesTable.updateTableFromData()
}

func (ui *Ui) loadRequest(req *core.Request) {
	ui.UrlInput.SetText(req.Url)
	ui.RequestHeadersTable.updateTableFromData()
	ui.RequestBodyInput.SetText(req.ReqBody, false)
}
