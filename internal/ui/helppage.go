package ui

import (
	"github.com/rivo/tview"
)

func (ui *Ui) NewHelpPage() *tview.TextView {

    HelpPage := tview.NewTextView()
    HelpPage.
        SetTextStyle(defaultStyle).
		SetDynamicColors(true).
		SetText(`
[green]LitePost[white] is a terminal-based API client for exploring, sending, and testing HTTP requests.

[green]Global Shortcuts
[yellow]Ctrl + ↑ / ↓ / ← / →   [white]Move focus between panes
[yellow]Ctrl + Space           [white]Send the current request
[yellow]Ctrl + T               [white]Open HTTP method list
[yellow]Ctrl + V               [white]Open the parsed variable list
[yellow]F1                     [white]Open this help page

[green]Tabs[white]
The Request and Response panes are divided into different tabs
[yellow][                      [white]Switch to previous tab
[yellow]]                      [white]Switch to next tab

[green]Key/Value Table[white]
Headers and variables are managed in interactive Key/Value tables.
They may be editable or read-only.

[green]Editable [white](request headers and variables)
[yellow]Enter                   [white]Edit key or value
[yellow]a                       [white]Add a new key
[yellow]d                       [white]Delete selected key

[green]Read-Only [white](response headers and parsed variables)
[yellow]Enter                   [white]Copy entry as a variable

`).
        SetTitle(" Help ").
        SetBorder(true).
        SetTitleColor(focusedTitleColor).
        SetBackgroundColor(backgroundColor).
        SetBorderStyle(borderStyle)

    return HelpPage
}
