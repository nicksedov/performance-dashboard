package controller

import (
	"io"
	"net/http"
)

func GetGealth(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "{\"Status\":\"OK\"}\n")
}