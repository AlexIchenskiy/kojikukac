package handlers

import (
	"bufio"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var n = 3
var data map[string]float64
var mutex sync.Mutex

func SaxpyHandler(w http.ResponseWriter, r *http.Request) {
	if data == nil {
		data = make(map[string]float64)
	}
	if r.Method == http.MethodPost {
		scanner := bufio.NewScanner(r.Body)
		defer r.Body.Close()

		mutex.Lock()
		for scanner.Scan() {
			line := scanner.Text()
			l := strings.Split(line, ":")
			i, err := strconv.ParseFloat(l[1], 64)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			data[l[0]] = i
		}
		mutex.Unlock()

		if err := scanner.Err(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Failed to read request body: %v", err)
			return
		}

		if len(data) == 2*n+1 { // ako su upisene sve vrijdnosti i a, predpostavlja se da korisnik neće unositi besmislene podatke
			_, err1 := w.Write([]byte(saxpy()))
			if err1 != nil {
				http.Error(w, "Failed to write response", http.StatusInternalServerError)
				return
			}
		}
	} else if r.Method == http.MethodGet {
		if len(data) == 2*n+1 { // ako su upisene sve vrijdnosti i a
			_, err1 := w.Write([]byte(saxpy()))
			if err1 != nil {
				http.Error(w, "Failed to write response", http.StatusInternalServerError)
				return
			}
		} else {
			_, err1 := w.Write([]byte("Nema dovoljno informacija"))
			if err1 != nil {
				http.Error(w, "Failed to write response", http.StatusInternalServerError)
				return
			}
		}
	} else if r.Method == http.MethodDelete {
		scanner := bufio.NewScanner(r.Body)
		defer r.Body.Close()

		for scanner.Scan() {
			line := scanner.Text()
			l := strings.Split(line, ":")
			delete(data, l[0])
		}
		if err := scanner.Err(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}

func saxpy() string {
	r := 0.0

	mutex.Lock()
	for i := 0; i < n; i++ {
		keyX := "x" + strconv.Itoa(i)
		keyY := "y" + strconv.Itoa(i)

		r = r + data[keyX]*data["a"] + data[keyY]
	}
	mutex.Unlock()
	str := strconv.FormatFloat(r, 'f', -1, 64)
	return str
}

/*
podaci se primju u obliku opis:podatak\n
	x3:1.0
	y0:2.0
	a:0.5
	x1:3.0
	y4:4.0
	x2:5.0 ...
*/
