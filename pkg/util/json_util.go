package util

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func PrettyPrintJSON(body []byte) {
	var prettyJSON bytes.Buffer
    error := json.Indent(&prettyJSON, body, "", "  ")
    if error != nil {
        return
    }
	output := prettyJSON.String()
	fmt.Printf("Project roles:\n%s\n", output)
}