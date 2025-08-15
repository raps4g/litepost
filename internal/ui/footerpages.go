package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (ui *Ui) newFooter(FootersPages *tview.Pages, name, text string, style tcell.Style) {
    newFooter := tview.NewTextView()
    newFooter.SetBackgroundColor(backgroundColor)
    newFooter.SetTextStyle(style)
    newFooter.SetText(text)
    FootersPages.AddPage(name, newFooter, true, true)
    FootersPages.SendToFront(name)
}

func (ui *Ui) newFooterInput(label, text string, doneFunc func(string)) string{

    var input string

    inputField := tview.NewInputField()
    inputField.SetFormAttributes(0, tcell.ColorRed, backgroundColor, foregroundColor, backgroundColor)

    inputField.SetLabel(label)
    inputField.SetText(text)
    ui.FooterPages.AddPage("footer_input", inputField, true, true)
    ui.App.SetFocus(inputField)

    inputField.
        SetDoneFunc(func(key tcell.Key) {
            if key == tcell.KeyEnter {
                input = inputField.GetText()
                doneFunc(input)
            } 
            ui.FooterPages.RemovePage("footer_input")
            if ui.FocusLayout.OpenDialog != nil {
                ui.App.SetFocus(ui.FocusLayout.OpenDialog)
            } else {
                ui.App.SetFocus(ui.FocusLayout.OnFocus)
            }
        })

    return input

}

