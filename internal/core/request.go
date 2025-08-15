package core

import (
	"bytes"
	"io"
	"net/http"
	"strconv"
	"time"
)

func SendHttpRequest(reqData *Request, variables *map[string]string) (error) {
	client := &http.Client{Timeout: 10 * time.Second}

	var body io.Reader
	if reqData.ReqBody != "" {
        bodyStr := ReplaceVariables(reqData.ReqBody, *variables)
		body = bytes.NewBufferString(bodyStr)
	}

	req, err := http.NewRequest(reqData.Methods[reqData.SelectedMethod], reqData.Url, body)
	if err != nil {
		return err
	}

	for key, value := range reqData.ReqHeaders {
        if (key != "" && value != "") {
            canonKey := http.CanonicalHeaderKey(key)
            value = ReplaceVariables(value, *variables)
            req.Header.Set(canonKey, value)
        }
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

    respBodyBytes, err := io.ReadAll(resp.Body)
    if err != nil {
        return err
    }

    if resp.Header.Get("Content-Type") == "application/json" {
        reqData.RespBody, err = PrettyPrintJSON(string(respBodyBytes))
    } else {
        reqData.RespBody = string(respBodyBytes)
    }
    reqData.Status = strconv.Itoa(resp.StatusCode) + " " + http.StatusText(resp.StatusCode)
    reqData.RespHeaders = httpRespToMap(&resp.Header)
    parsedVariables, err := ConvertToMap(reqData.RespBody)
    reqData.ParsedVariables = parsedVariables
	return nil
}

func httpRespToMap(header *http.Header) (map[string]string) {

    headerMap := map[string]string{}

    for key, vals := range *header {
        headerMap[key] = vals[0]
    }

    return headerMap
}
