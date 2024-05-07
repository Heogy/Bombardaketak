package controllers

import (
	"bombardaketak/gramatika"
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func InitViews() {
	validate = validator.New(validator.WithRequiredStructEnabled())
	serveTemplates()
	println("templated files registered")
}

func serveTemplates() {
	http.Handle("/", http.FileServer(http.Dir("./views")))
	http.Handle("/ariketa", http.HandlerFunc(ariketaHandler))
	http.Handle("/nornori", http.HandlerFunc(norNoriHandler))
	http.Handle("/erantzun", http.HandlerFunc(erantzunHandler))
}

func erantzunHandler(w http.ResponseWriter, r *http.Request) {
	g := r.FormValue("galdera")
	era := r.FormValue("era")

	var guess gramatika.GuessNorNori
	err := json.Unmarshal([]byte(g), &guess)
	if err != nil {
		panic(err)
	}
	guess.Erantzuna = era
	println(guess.Erantzuna)
	println(guess.Denbora)
	println(guess.Nor)
	println(guess.Nori)
	println(guess.Nondik)
	defer r.Body.Close()
	println("erantzunaHandler")
	validationError := validate.Struct(guess)
	println("validationError")
	if validationError != nil {

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"status"` + validationError.Error() + `}`))
		return
	}

	b, erantzuna, err := gramatika.VerifyNorNori(guess)

	if err != nil {
		println("error verfifying")
		panic(err)
	}
	var tmplFile = "./views/erantzuna.html"
	tmpl, err := template.ParseFiles(tmplFile)
	if err != nil {
		println("error parsing template")
		panic(err)
	}
	type Erantzuna struct {
		Xuxen     string
		Erantzuna string
		Berritz   string
	}
	var msg string
	if b {
		msg = "Xuxen"
	} else {
		msg = "Oker"
	}

	var eran = Erantzuna{
		Xuxen:     msg,
		Erantzuna: erantzuna,
		Berritz:   guess.Nondik,
	}
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
	var tmplFile = "./views/ariketa.html"
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

func norNoriHandler(w http.ResponseWriter, r *http.Request) {
	println("norNoriHandler")
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"status": "error"}`))
		return
	}
	var tmplFile = "./views/nornori.html"
	tmpl, err := template.ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}
	randomGaldera := gramatika.RandomNorNori()
	println(randomGaldera.RandomNor)
	err = tmpl.Execute(w, randomGaldera)
	if err != nil {
		panic(err)
	}
	println("norNoriHandler end")
}
