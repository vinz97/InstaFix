package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func startServer() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/newticket", newticket)
	http.HandleFunc("/getnome", getnome)
	http.HandleFunc("/getickets", getickets)
	http.ListenAndServe(":8000", nil)

}

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
	fmt.Println("Hello path raggiunto")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
	fmt.Println("Stampa degli headers")
}

func upload(w http.ResponseWriter, req *http.Request) {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(b))

}

func register(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	name := req.Form.Get("nome")
	surname := req.Form.Get("cognome")
	city := req.Form.Get("citta")
	address := req.Form.Get("indirizzo")
	email := req.Form.Get("email")
	number := req.Form.Get("telefono")
	passw := req.Form.Get("password")
	resp := registerQuery(sqlDB, name, surname, city, address, email, number, passw)
	fmt.Println(resp)
	fmt.Fprintf(w, resp)
}

func login(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	email := req.Form.Get("email")
	passw := req.Form.Get("password")
	resp := loginQuery(sqlDB, email, passw)
	fmt.Println(resp)
	fmt.Fprintf(w, resp)

}

func newticket(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	title := req.Form.Get("titolo")
	category := req.Form.Get("categoria")
	description := req.Form.Get("descrizione")
	email := req.Form.Get("email")
    resp := newticketQuery(sqlDB, title, category, description, email)
	fmt.Println(resp)
	fmt.Fprintf(w, resp)
}

func getnome(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	mail := req.Form.Get("email")
	resp := getnomeQuery(sqlDB, mail)
	fmt.Println(resp)
	fmt.Fprintf(w, resp)
}

func getickets (w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	mail := req.Form.Get("email")
	resp := getTickets(sqlDB, mail)
	fmt.Println(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

}
