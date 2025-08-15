package ui

import (
	"errors"
	"strings"
	"unicode/utf8"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type KeyValueTable struct {
    *tview.Table
    Data *map[string]string
    Category string
}

func (ui *Ui) NewKeyValueTable(keyValue *map[string]string, category string, edtiable bool) *KeyValueTable {
    table := &KeyValueTable{
        Table: tview.NewTable(),
        Data: keyValue,
        Category: category,
    }
	table.
		SetSelectable(true, true).
		SetFixed(0, 2).
		SetBackgroundColor(backgroundColor)
    
    table.updateTableFromData()

    if edtiable {
        table.SetEditableInputCapture(ui)
    } else {
        table.SetReadOnlyInputCapture(ui)
    }

    return table
}

func (table *KeyValueTable) addKeyValueCell(key, value string, defineInMap bool) int {

    if defineInMap {
        (*table.Data)[key] = value
    }

    keyLen := utf8.RuneCountInString(key)
    paddedKey := key
    if keyLen < 20 {
        paddedKey = key + strings.Repeat(" ", 20 - keyLen)
    }

    row := table.GetRowCount()

    table.SetCell(row, 0, tview.NewTableCell(paddedKey).SetAlign(tview.AlignLeft).SetMaxWidth(25).SetReference(key))
    table.SetCell(row, 1, tview.NewTableCell(value).SetAlign(tview.AlignLeft).SetReference(value).SetExpansion(1))

    return row
}

func (table *KeyValueTable) updateTableFromData() {
    table.Clear()
    for key, value := range *table.Data {
        table.addKeyValueCell(key, value, false)
    }
}

func (table *KeyValueTable) modifyKeyValueCell(row, col int, text string, redefineMap bool) {
  
    if redefineMap {
        key, err := table.getTextFromCell(row, 0)
        if err != nil { panic(err) }
        if col == 0 {
            value := (*table.Data)[key]
            delete(*table.Data, key)
            (*table.Data)[text] = value
        } else {
            (*table.Data)[key] = text
        }
    }
    textLen := utf8.RuneCountInString(text)

    if col == 0 && textLen < 20 {
        paddedText := text + strings.Repeat(" ", 20 - textLen)
        table.SetCell(row, col, tview.NewTableCell(paddedText).SetAlign(tview.AlignLeft).SetMaxWidth(25).SetReference(text))
    } else {
        table.SetCell(row, col, tview.NewTableCell(text).SetAlign(tview.AlignLeft).SetReference(text).SetExpansion(1))
    } 
}

func (table *KeyValueTable) getTextFromCell(row, col int) (string, error) {
    cell := table.GetCell(row, col)
    if ref, ok := cell.GetReference().(string); ok {
        return ref, nil
    }
    return "", errors.New("Cell reference is not a string.")
}

func (table *KeyValueTable) SetEditableInputCapture(ui *Ui) {
    table.
        SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
            row, col := table.GetSelection() 
            switch {
            
            case event.Key() == tcell.KeyEnter:
                table.ModifyKeyValueHandler(row, col, ui)
            case event.Rune() == 'a':
                table.AddKeyHandler(ui)
            case event.Rune() == 'd':
                table.DeleteKey(row, col)
            case event.Rune() == 'p':
                table.AddAsVariableHandler(row, col, ui)
            }

            return event
        })
}

func (table *KeyValueTable) SetReadOnlyInputCapture(ui *Ui) {
    table.
        SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
            row, col := table.GetSelection() 
            switch {
            
            case event.Key() == tcell.KeyEnter:
                table.AddAsVariableHandler(row, col, ui)
            }

            return event
        })
}
func (table *KeyValueTable) DeleteKey(row, col int) {
    key, err := table.getTextFromCell(row, 0)
    if err != nil {
        panic(err)
    }
    delete(*table.Data, key)
    table.RemoveRow(row)

}

func (table *KeyValueTable) ModifyKeyValueHandler(row, col int, ui *Ui) {
    oldText, err := table.getTextFromCell(row, col)
    if err != nil {
        panic(err)
    }

    prompt := "Edit " + table.Category + " value: "
    if col == 0 {
        prompt = "Rename " + table.Category + ": "
    }

    ui.newFooterInput(prompt, oldText, func(input string) {
        table.modifyKeyValueCell(row, col, input, true)
    })
}

func (table *KeyValueTable) AddAsVariableHandler(row, col int, ui *Ui) {
    key, err := table.getTextFromCell(row, 0)
    if err != nil {
        panic(err)
    }

    value, err := table.getTextFromCell(row, 1)
    if err != nil {
        panic(err)
    }

    prompt := "Add variable: "

    ui.newFooterInput(prompt, key, func(input string) {
        ui.VariablesTable.addKeyValueCell(input, value, true)
    })
}

func (table *KeyValueTable) AddKeyHandler(ui *Ui) {
    ui.newFooterInput("New " + table.Category + ": ", "", func(input string) {
        table.addKeyValueCell(input, "", true)
    })
}
