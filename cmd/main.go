package main

import "github.com/raps4g/litepost/internal/ui"

func main() {
    appContext := ui.InitAppContext()
    
    appContext.App.Run()
}
