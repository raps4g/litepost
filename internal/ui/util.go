package ui

import (
	"encoding/json"
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

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


func prettyPrintJSON(input string) (string, error) {
	var parsed interface{}

	if err := json.Unmarshal([]byte(input), &parsed); err != nil {
		return "", fmt.Errorf("invalid JSON: %v", err)
	}

	prettyJSON, err := json.MarshalIndent(parsed, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to format JSON: %v", err)
	}

	return string(prettyJSON), nil
}

func (s *State) createFooterInput(label, text string, caller tview.Primitive, doneFunc func(string)) string{

    var input string

    inputField := tview.NewInputField()
    inputField.SetFormAttributes(0, tcell.ColorRed, backgroundColor, foregroundColor, backgroundColor)

    inputField.SetLabel(label)
    inputField.SetText(text)
    inputField.
        SetDoneFunc(func(key tcell.Key) {
            if key == tcell.KeyEnter {
                input = inputField.GetText()
                doneFunc(input)
            } 
            s.Footers.RemovePage("footer_input")
            s.App.SetFocus(caller)
        })

    s.Footers.AddPage("footer_input", inputField, true, true)
    s.App.SetFocus(inputField)

    return input

}

