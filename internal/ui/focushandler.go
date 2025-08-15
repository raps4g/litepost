package ui

import (
    "github.com/rivo/tview"
)

func (ui *Ui) SetFocusHandlers() {

    SetFocusTitleColorChange(ui.UrlInput.Box, ui.UrlInput.Box)
    // SetFocusTitleColorChange(ui.RequestBodyInput.Box, ui.RequestPages.Box)
    // SetFocusTitleColorChange(ui.RequestHeadersTable.Box, ui.RequestPages.Box)
    // SetFocusTitleColorChange(ui.VariablesTable.Box, ui.RequestPages.Box)
    // SetFocusTitleColorChange(ui.ResponseBodyView.Box, ui.ResponsePages.Box)
    // SetFocusTitleColorChange(ui.ResponseHeadersTable.Box, ui.ResponsePages.Box)
    // SetFocusTitleColorChange(ui.VariablesTable.Box, ui.RequestPages.Box)

}

func SetFocusTitleColorChange(b1, b2 *tview.Box) {
    b1.SetBlurFunc(func () {
        b2.SetTitleColor(bluredTitleColor)
    }).
        SetFocusFunc(func () {
            b2.SetTitleColor(focusedTitleColor)
        })
}
