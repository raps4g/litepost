package core

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

func extractUrlData(URL string) (string, string) {

    u, err := url.Parse(URL)
    if err != nil {
        panic(err)
    }

    return u.Host, u.Path
}

func FlattenJSON(data interface{}, parentKey string, result map[string]string) {
	
    switch t := data.(type) {
    
    case map[string]interface{}:
        for key, value := range t {
            newKey := key
            if parentKey != "" {
                newKey = parentKey + "." + key
            }

            switch v := value.(type) {

            case map[string]interface{}:
                FlattenJSON(v, newKey, result)

            default:
                result[newKey] = fmt.Sprintf("%v", v)
            }
        }

    case []interface{}:
        for i, value := range t {
            indexKey := fmt.Sprintf("%s[%d]", parentKey, i)
                FlattenJSON(value, indexKey, result)
        }
    }

}

func ConvertToMap(jsonStr string) (map[string]string, error) {
	var jsonData interface{}
	if err := json.Unmarshal([]byte(jsonStr), &jsonData); err != nil {
		return nil, err
	}

	result := make(map[string]string)

    FlattenJSON(jsonData, "", result)

	return result, nil
}

func ReplaceVariables(input string, variables map[string]string) string {
	for key, value := range variables {
		placeholder := fmt.Sprintf("${%s}", key)
        input = strings.ReplaceAll(input, placeholder, value)
		placeholder = fmt.Sprintf("$%s", key)
        input = strings.ReplaceAll(input, placeholder, `"`+value+`"`)
	}
	return input
}


func PrettyPrintJSON(input string) (string, error) {
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
