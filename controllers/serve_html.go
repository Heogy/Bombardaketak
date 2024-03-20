package controllers

import (
	"bombardaketak/gramatika"
	"encoding/json"
	"html/template"
	"net/http"
)

func InitViews() {

	serveTemplates()
	println("templated files registered")
}

func serveTemplates() {
	http.Handle("/", http.FileServer(http.Dir("./views/statics")))
	http.Handle("/ariketa", http.HandlerFunc(ariketaHandler))
	http.Handle("/erantzun", http.HandlerFunc(erantzunHandler))
}

func erantzunHandler(w http.ResponseWriter, r *http.Request) {
	g := r.FormValue("galdera")
	era := r.FormValue("era")

	var guess gramatika.Guess
	err := json.Unmarshal([]byte(g), &guess)
	if err != nil {
		panic(err)
	}
	guess.Erantzuna = era
	defer r.Body.Close()
	validationError := validate.Struct(guess)
	if validationError != nil {

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"status"` + validationError.Error() + `}`))
		return
	}

	b, erantzuna, err := gramatika.Verify(guess)

	var tmplFile = "./views/templates/erantzuna.tmpl"
	tmpl, err := template.ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}
	type Erantzuna struct {
		Xuxen     string
		Erantzuna string
	}
	var msg string
	if b {
		msg = "Xuxen"
	} else {
		msg = "Oker"
	}

	var eran = Erantzuna{Xuxen: msg, Erantzuna: erantzuna}
	err = tmpl.Execute(w, eran)
	if err != nil {
		panic(err)
	}
	println("erantzunaHandler end")

	// if err != nil {
	// 	println(err.Error())
	// 	w.WriteHeader(http.StatusNotImplemented)
	// 	w.Write([]byte(`{"status": ` + err.Error() + `}`))
	// 	return
	// }
	// if !b {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte(`{"status": "error"}`))
	// 	return
	// }
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte(`{"status": "ok"}`))
}

func ariketaHandler(w http.ResponseWriter, r *http.Request) {
	println("ariketaHandler")
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"status": "error"}`))
		return
	}
	var tmplFile = "./views/templates/ariketa.tmpl"
	tmpl, err := template.ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}
	randomGaldera := gramatika.RandomGaldera()
	println(randomGaldera.RandomNor)
	err = tmpl.Execute(w, randomGaldera)
	if err != nil {
		panic(err)
	}
	println("ariketaHandler end")
}
