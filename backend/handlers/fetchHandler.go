package handlers

import (
	"encoding/json"
	"net/http"
)

type RequestBody struct {
	URL string `json:"url"`
}
type ResponseBody struct {
	Headers http.Header `json:"headers"`
}

func FetchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	resp, err := http.Get(requestBody.URL)
	if err != nil {
		http.Error(w, "Failed to perform GET request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	headers := resp.Header

	responseBody := ResponseBody{
		Headers: headers,
	}

	jsonData, err := json.Marshal(responseBody)
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
