package handlers

import (
	"io"
	"net/http"
	"os"
)

func MyJmbagHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}

		err = os.WriteFile("student.txt", body, 0644)
		if err != nil {
			http.Error(w, "Failed to save data", http.StatusInternalServerError)
			return
		}
	} else if r.Method == http.MethodGet {
		data, err := os.ReadFile("student.txt")
		if err != nil {
			http.Error(w, "Failed to read data", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		_, err = w.Write(data)
		if err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}
