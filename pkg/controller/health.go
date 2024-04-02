package controller

import (
	"encoding/json"
	"net/http"
)

type healthCheck struct {
	Status string
}

func GetGealthCheck(w http.ResponseWriter, r *http.Request) {
	st, _ := json.Marshal(&healthCheck{Status: "OK"})
	w.Write(st)
}