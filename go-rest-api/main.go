package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type event struct {
	ID          string
	Title       string
	Description string
}

type allEvent []event

var events = allEvent{
	{
		ID:          "1",
		Title:       "Introduction to GoLang",
		Description: "This is a basic thing important to know",
	},
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome home")
}

func createEvent(w http.ResponseWriter, req *http.Request) {
	var newEvent event
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "Please check data again")
	}

	json.Unmarshal(body, &newEvent)
	events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/event", createEvent).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
