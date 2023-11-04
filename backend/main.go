package main

import (
	"fmt"
	"net/http"

	"Projekt/handlers"
	"Projekt/login"
)

func main() {

	http.HandleFunc("/api/reservation/add", handlers.AddReservationHandler)
	http.HandleFunc("/api/reservation/delete", handlers.DelReservationHandler)
	http.HandleFunc("/api/search", handlers.SearchHandler)
	http.HandleFunc("/api/get/status", handlers.StatusHandler)
	http.HandleFunc("/api/login", login.Authenticate(handlers.SumHandler))

	/*	http.HandleFunc("/jmbag", handlers.JmbagHandler)
		http.HandleFunc("/sum", login.Authenticate(handlers.SumHandler))
		http.HandleFunc("/fetch", handlers.FetchHandler)
		http.HandleFunc("/0036540336", login.Authenticate(handlers.MyJmbagHandler))
		http.HandleFunc("/saxpy", login.Authenticate(handlers.SaxpyHandler))
	*/
	err1 := http.ListenAndServe("127.0.0.1:3000", nil)
	if err1 != nil {
		fmt.Println("Server error:", err1)
	}

}
