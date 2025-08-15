package core

type Request struct {
    Url string
    SelectedMethod int
    ReqBody string
    ReqHeaders map[string]string
    Status string
    RespBody string
    RespHeaders map[string]string
    Methods []string
    ParsedVariables map[string]string
}
