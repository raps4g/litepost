package ui

import (
	"github.com/raps4g/litepost/internal/core"
	"github.com/rivo/tview"
)

type AppContext struct {
    Request core.Request
    History []core.Request
    Variables map[string]string
    //FocusLayout
    App *tview.Application
}

func InitAppContext() *AppContext {

    appContext := AppContext{}
    appContext.Request = core.Request{
        Url: "",
        ReqBody: "",
        ReqHeaders: make(map[string]string),
        RespBody: "",
        RespHeaders: make(map[string]string),
        ParsedVariables: make(map[string]string),
        Status: "",
        SelectedMethod: 0,
        Methods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
    }

    appContext.Request.ReqHeaders["Content-Type"] = "application/json"
    appContext.History = []core.Request{}
    appContext.Variables = make(map[string]string)

    appContext.App = SetApp(&appContext.Request, &appContext.Variables, &appContext.History)

    return &appContext
}

