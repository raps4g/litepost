package ui

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (s *State) SetHeaderTable() *State {
	s.HeaderTable = tview.NewTable()
	s.HeaderTable.
		SetSelectable(true, true).
		SetFixed(10, 2).
		SetTitle(" Headers ").
		SetTitleColor(focusedTitleColor).
		SetBorder(true).
		SetBorderStyle(borderStyle).
		SetBackgroundColor(backgroundColor).
		SetDrawFunc(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
			for i := range len(s.Headers) {
				tview.Print(screen, fmt.Sprintf("(%d)", i+1), x+1, y+1+i, 3, 0, placeholderColor)
			}
			return x + 5, y + 1, width - 4, height
		}).
        SetInputCapture(s.headerTaableInputCapture)

	for key, value := range s.Headers {
		s.addHeaderCell(key, value)
	}

    return s
}

func (s *State) headerTaableInputCapture (event *tcell.EventKey) *tcell.EventKey {
    row, col := s.HeaderTable.GetSelection()

    oldText := s.HeaderTable.GetCell(row, col).Text
    oldText = strings.TrimSpace(oldText)

    switch {

    case event.Key() == tcell.KeyEnter:
        var label string
        if col == 0 {
            label = " Rename header key: "
        } else if col == 1 {
            label = " Edit header value: "
        }
        s.createFooterInput(label, oldText, s.HeaderTable, func(input string) {s.modifyHeaderCell(row, col, input)})
    
    case event.Key() == tcell.KeyEsc:
        s.removeModal("header_table")
        s.removeFooter("header_footer")

    case event.Rune() == 'a':
        s.Headers["New header"] = ""
        newRow := s.addHeaderCell("New header", "")
        s.HeaderTable.Select(newRow, 0)
        s.createFooterInput("New header key: ", "", s.HeaderTable, func(input string) {s.modifyHeaderCell(newRow, 0, input)})
    
    case event.Rune() == 'd':
        key := s.HeaderTable.GetCell(row, 0).Text
        key = strings.TrimSpace(key)
        delete(s.Headers, key)
        s.HeaderTable.RemoveRow(row)
        s.HeaderTable.Select(row-1, 0)
    }
    return event
}

func (s *State) addHeaderCell(key, value string) int {


    keyLen := utf8.RuneCountInString(key)
    valueLen := utf8.RuneCountInString(value)
    if keyLen < 20 {
        key += strings.Repeat(" ", 20 - keyLen)
    }
    if valueLen < 31 {
        value += strings.Repeat(" ", 31 - valueLen)
    }

    row := s.HeaderTable.GetRowCount()

    s.HeaderTable.SetCell(row, 0, tview.NewTableCell(key).SetAlign(tview.AlignLeft).SetMaxWidth(20))
    s.HeaderTable.SetCell(row, 1, tview.NewTableCell(value).SetAlign(tview.AlignLeft).SetMaxWidth(31))

    return row
}

func (s *State) modifyHeaderCell(row, col int, text string) {
    
    oldKey := s.HeaderTable.GetCell(row, 0).Text
    oldKey = strings.TrimSpace(oldKey)
    
    textLen := utf8.RuneCountInString(text)

    if col == 0 {
        value := s.Headers[oldKey]
        delete(s.Headers, oldKey)
        s.Headers[text] = value
        if textLen < 20 {
            text += strings.Repeat(" ", 20 - textLen)
        }
        s.HeaderTable.SetCell(row, 0, tview.NewTableCell(text).SetAlign(tview.AlignLeft).SetMaxWidth(20))
    } else if col == 1 {
        s.Headers[oldKey] = text
        if textLen < 31 {
            text += strings.Repeat(" ", 31 - textLen)
        }
        s.HeaderTable.SetCell(row, 1, tview.NewTableCell(text).SetAlign(tview.AlignLeft).SetMaxWidth(31))
    }
}
