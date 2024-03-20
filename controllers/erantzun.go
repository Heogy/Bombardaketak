package controllers

import (
	"bombardaketak/gramatika"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func InitAPI() {
	validate = validator.New(validator.WithRequiredStructEnabled())
	http.Handle("/api/erantzun", http.HandlerFunc(erantzunHandlerApi))
	http.Handle("/api/ariketa", http.HandlerFunc(ariketaHandler))
	println("endpoint /erantzun registered")
}

func ariketaHandlerApi(w http.ResponseWriter, r *http.Request) {
	println("ariketaHandler")
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"status": "error"}`))
		return
	}
	randomGaldera := gramatika.RandomGaldera()
	println(randomGaldera.RandomNor)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(randomGaldera)
	println("ariketaHandler end")
}

func erantzunHandlerApi(w http.ResponseWriter, r *http.Request) {

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

	b, _, err := gramatika.Verify(guess)
	if err != nil {
		println(err.Error())
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(`{"status": ` + err.Error() + `}`))
		return
	}
	if !b {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"status": "error"}`))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "ok"}`))

}
