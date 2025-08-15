package ui

import (
	"reflect"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TabPage struct {
    PageTitle string
    *tview.Pages
    Tabs []tview.Primitive
    TabNames []string
}

func (ui *Ui) NewTabPages(pageTitle string) *TabPage {
    
    tabPage := &TabPage {
        PageTitle: pageTitle,
        Pages: tview.NewPages(),
        Tabs: []tview.Primitive{},
        TabNames: []string{},
    }

    tabPage.
        SetChangedFunc(func() {
            tabPage.SetTitle(tabPage.formatTitle())
        }).
        SetBorder(true).
        SetTitle(tabPage.formatTitle()).
        SetDrawFunc(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
            return x+2, y+1, width-3, height-2
        }).
        SetBackgroundColor(backgroundColor).
        SetTitleColor(bluredTitleColor).
        SetTitleAlign(0).
        SetBorderStyle(borderStyle)
    tabPage.SetTabInputCapture()

    return tabPage
}

func (tabPage *TabPage) SetTabInputCapture() {
    tabPage.Pages.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        if 
            event.Rune() == '[' {
            tabPage.cycleTabs(false)
            return nil
        } else if 
            event.Rune() == ']' {
            tabPage.cycleTabs(true)
            return nil
        }
        return event
    })
}

func (tabPage *TabPage) cycleTabs(direction bool) {
    frontPageName, _ := tabPage.Pages.GetFrontPage()
    var nextTabIndex int
    for i, tabName := range tabPage.TabNames {
        if frontPageName == tabName {
            nextTabIndex = i
            break
        }
    }
    if direction {
        nextTabIndex = (nextTabIndex + 1) % len(tabPage.TabNames)
    } else {
        nextTabIndex = (len(tabPage.TabNames) + nextTabIndex - 1) % len(tabPage.TabNames)
    }
    tabPage.Pages.SendToFront(tabPage.TabNames[nextTabIndex])
}

func (tabPage *TabPage) AddTab(name string, item tview.Primitive) {
    tabPage.Tabs = append([]tview.Primitive{item}, tabPage.Tabs...) 
    tabPage.TabNames = append([]string{name}, tabPage.TabNames...)
    tabPage.Pages.AddPage(name, item, true, true)
   
    tabPage.SetFocusAndBlurFunc(item)
}

func (tabPage *TabPage) SetFocusAndBlurFunc(item tview.Primitive) {
    if box := getEmbeddedBox(item); box != nil {
        box.SetFocusFunc(func() {
            tabPage.SetTitleColor(focusedTitleColor)
        }).
            SetBlurFunc(func() {
                tabPage.SetTitleColor(bluredTitleColor)
            })
    }
}

func getEmbeddedBox(item tview.Primitive) *tview.Box {

    if box, ok := item.(*tview.Box); ok {
        return box
    }

    val := reflect.ValueOf(item)
    
    if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
        return nil
    }

    val = val.Elem()

    for i := 0; i < val.NumField(); i++ {
        field := val.Field(i)

        if field.Type() == reflect.TypeOf((*tview.Box)(nil)) {
            return field.Interface().(*tview.Box)
        }
    }
    
    return nil
}

func (tabpage *TabPage) formatTitle() string {

    var formatted []string
    frontPageName, _ := tabpage.GetFrontPage()
    for _, tabName := range tabpage.TabNames {
        if tabName == frontPageName {
            formatted = append(formatted, tabName)    
        } else {
            formatted = append(formatted, "[::d]"+tabName+"[::D]")    
        }
    }
    return " " + tabpage.PageTitle + ": " + strings.Join(formatted, " | ") + " "
}
