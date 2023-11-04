package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type SumResponse struct {
	A      int `json:"a"`
	B      int `json:"b"`
	Result int `json:"result"`
}

func SumHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	aStr := queryParams.Get("a")
	bStr := queryParams.Get("b")

	a, err := strconv.Atoi(aStr)
	if err != nil {
		http.Error(w, "Invalid value for 'a'", http.StatusBadRequest)
		return
	}
	b, err := strconv.Atoi(bStr)
	if err != nil {
		http.Error(w, "Invalid value for 'b'", http.StatusBadRequest)
		return
	}
	sum := a + b

	response := SumResponse{
		A:      a,
		B:      b,
		Result: sum,
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}
