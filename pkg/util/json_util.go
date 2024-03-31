package util

import (
	"bytes"
	"encoding/json"
	"log"
)

func PrettyPrintJSON(body []byte) {
	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, body, "", "  ")
	if error != nil {
		return
	}
	output := prettyJSON.String()
	log.Printf("JSON data:\n%s\n", output)
}
