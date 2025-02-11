package core

type Request struct {
    Url string
    Body string
    Response string
    Status string
    SelectedMethod int
    Headers map[string]string
    Methods []string
}
