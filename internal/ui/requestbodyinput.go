package ui

import (
	"github.com/raps4g/litepost/internal/core"
	"github.com/rivo/tview"
)

func (ui *Ui) NewRequestBodyInput(req *core.Request) *tview.TextArea {

    bodyPlaceholder := "Enter request body"
    RequestBodyInput := tview.NewTextArea()
    RequestBodyInput.
        SetPlaceholder(bodyPlaceholder).
        SetTextStyle(defaultStyle).
        SetLabelStyle(defaultStyle).
        SetPlaceholderStyle(placeholderStyle).
        SetChangedFunc(func() {
            req.ReqBody = RequestBodyInput.GetText()  
        })
    return RequestBodyInput
}
