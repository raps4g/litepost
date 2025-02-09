package ui

import (
	"net/url"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Ui struct {
    App *tview.Application
    Pages *tview.Pages
    MainGrid *tview.Grid
    RequestGrid *tview.Grid
    UrlInput *tview.TextArea
    RequestInput *tview.TextArea
    ResponseView *tview.TextView
    HistoryList *tview.List
    MethodList *tview.List
}

var methods [7]string = [7]string{"GET",
                                "POST",
                                "PUT",
                                "PATCH",
                                "DELETE",
                                "HEAD",
                                "OPTIONS"}

func InitUi() *Ui {

    defaultStyle := tcell.StyleDefault
    placeholderStyle := tcell.StyleDefault.Foreground(tcell.ColorGray)
    _, backgroundColor, _ := defaultStyle.Decompose()
    foregroundColor := tcell.ColorGray

    setCustomBorders()

    ui := &Ui{
        App: tview.NewApplication(),
        Pages: tview.NewPages(),
        MainGrid: tview.NewGrid(),
        RequestGrid: tview.NewGrid(),
        UrlInput: tview.NewTextArea(),
        RequestInput: tview.NewTextArea(),
        ResponseView: tview.NewTextView(),
        HistoryList: tview.NewList(),
        MethodList: tview.NewList(),
    } 

    ui.MethodList.SetMainTextStyle(defaultStyle).
        ShowSecondaryText(false).
        SetSelectedFocusOnly(true).
        SetHighlightFullLine(true).
        SetShortcutStyle(defaultStyle).
        SetTitle(" Method ").
        SetBorder(true).
        SetBackgroundColor(backgroundColor).
        SetBorderColor(foregroundColor).
        SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
            if event.Key() == tcell.KeyEnter {
                ui.Pages.RemovePage("method_modal")
            }    
            return event})

    for i, method := range methods {
        ui.MethodList.AddItem(method, "", rune('1'+i), nil)
    }

    ui.UrlInput.
        SetTextStyle(defaultStyle).
        SetLabelStyle(defaultStyle).
        SetPlaceholderStyle(placeholderStyle).
        SetPlaceholder("Enter URL [F2]").
        SetWrap(false).
        SetSelectedStyle(defaultStyle).
        SetDrawFunc(func(screen tcell.Screen, x int, y int, width int, height int) (int, int, int, int) {
            methodId := ui.MethodList.GetCurrentItem() 
            tview.Print(screen, methods[methodId], x+2, y+1, 8, 0, tcell.ColorWhite)
            screen.SetContent(x+9, y, tview.BoxDrawingsLightDownAndHorizontal, nil, tcell.StyleDefault.Foreground(ui.UrlInput.GetBorderColor()))
            screen.SetContent(x+9, y+1, tview.BoxDrawingsLightVertical, nil, tcell.StyleDefault.Foreground(ui.UrlInput.GetBorderColor()))
            screen.SetContent(x+9, y+2, tview.BoxDrawingsLightUpAndHorizontal, nil, tcell.StyleDefault.Foreground(ui.UrlInput.GetBorderColor()))
            return x+11,y+1,width-12,height-2
        }).
        SetBorderStyle(defaultStyle).
        SetBackgroundColor(backgroundColor).
        SetBorderColor(tcell.ColorRed).SetBorder(true).
        SetTitle(" [red]URL ").
        SetTitleAlign(0).
        SetBlurFunc(func () {
            ui.UrlInput.SetBorderColor(tcell.ColorGray)
            ui.UrlInput.SetPlaceholder("Enter URL [F2]")}).
        SetFocusFunc(func () {
            ui.UrlInput.SetPlaceholder("").
                SetBorderColor(tcell.ColorGray)
        })
    
    ui.RequestInput.SetPlaceholder("Enter request body [F3]").
        SetTextStyle(defaultStyle).
        SetLabelStyle(defaultStyle).
        SetPlaceholderStyle(placeholderStyle).
        SetOffset(0,10).
        SetBorderStyle(defaultStyle).
        SetTitle(" Request body ").
        SetTitleAlign(0).
        SetBorderColor(tcell.ColorGray).SetBorder(true).
        SetBlurFunc(func () {
            ui.RequestInput.SetPlaceholder("Enter request body [F3]")}).
        SetFocusFunc(func () {
            ui.RequestInput.SetPlaceholder("")})
    
    ui.ResponseView.
        SetToggleHighlights(true).
        SetTextStyle(defaultStyle).
        SetTitle(" Response ").
        SetTitleAlign(0).
        SetBorder(true).
        SetBackgroundColor(backgroundColor).
        SetBorderColor(foregroundColor)

    ui.HistoryList.
        SetMainTextStyle(defaultStyle).
        SetSecondaryTextStyle(defaultStyle).
        SetSelectedFocusOnly(true).
        SetTitle(" History ").
        SetTitleAlign(0).
        SetBorder(true).
        SetBackgroundColor(backgroundColor).
        SetBorderColor(foregroundColor)

    ui.RequestGrid.SetRows(3,0,0).
        AddItem(ui.UrlInput,0,0,1,1,1,0,true).
        AddItem(ui.RequestInput,1,0,1,1,0,0,false).
        AddItem(ui.ResponseView,2,0,1,1,1,2,false)

    ui.MainGrid.SetColumns(30,0).
        AddItem(ui.HistoryList,0,0,1,1,0,80,false).
        AddItem(ui.RequestGrid,0,1,1,1,0,80,true).
        AddItem(ui.RequestGrid,0,0,1,2,0,0,true)

    ui.Pages.AddPage("main_grid", ui.MainGrid, true, true)
    ui.App.SetRoot(ui.Pages, true)
    ui.App.SetInputCapture(ui.capture)
    return ui
}

func (ui *Ui) capture(event *tcell.EventKey) *tcell.EventKey {
    if event.Key() == tcell.KeyF2 {
        ui.App.SetFocus(ui.UrlInput)
    } else if event.Key() == tcell.KeyF3 {
        ui.App.SetFocus(ui.RequestInput)
    } else if event.Key() == tcell.KeyF6 {
        ui.Pages.AddPage("method_modal", createModal(ui.MethodList, 40, 9), true, true)
    } else if event.Key() == tcell.KeyF4 {
        methodId := ui.MethodList.GetCurrentItem()
        u, _ := url.Parse(ui.UrlInput.GetText())
        ui.HistoryList.AddItem(methods[methodId] + " - " + u.Host, string('\u2514')+" "+u.Path, 0, nil)
    } 
    return event
}

func createModal(p tview.Primitive, width, height int) tview.Primitive {
	return tview.NewGrid().
		SetColumns(0, width, 0).
		SetRows(10, height, 0).
		AddItem(p, 1, 1, 1, 1, 0, 0, true)
}
func setCustomBorders() {
    tview.Borders.TopLeft = tview.BoxDrawingsLightArcDownAndRight
    tview.Borders.TopRight = tview.BoxDrawingsLightArcDownAndLeft
    tview.Borders.BottomLeft = tview.BoxDrawingsLightArcUpAndRight
    tview.Borders.BottomRight = tview.BoxDrawingsLightArcUpAndLeft
    tview.Borders.HorizontalFocus = tview.BoxDrawingsLightHorizontal
    tview.Borders.VerticalFocus = tview.BoxDrawingsLightVertical
    tview.Borders.TopLeftFocus = tview.BoxDrawingsLightArcDownAndRight
    tview.Borders.TopRightFocus = tview.BoxDrawingsLightArcDownAndLeft
    tview.Borders.BottomLeftFocus = tview.BoxDrawingsLightArcUpAndRight
    tview.Borders.BottomRightFocus = tview.BoxDrawingsLightArcUpAndLeft
}


