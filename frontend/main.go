package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/campagnes", CampagneHandler)

	// static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	// insersion des routes dans l'objet http
	// et lancement du serveur
	http.Handle("/", r)
	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}

func CampagneHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("CampagneHandler appelé")
	var c Campagne
	c.Title = "Titre"

	b, err := json.Marshal(c)
	if err != nil {
		// envoie un message d'erreur
		fmt.Fprint(w, "error")
	}

	w.Header().Add("Accept-Charset", "utf-8")
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, string(b))
}

type Campagnes struct {
	Campagnes []Campagne
}

type Campagne struct {
	Title       string
	Description string
	Image       string
	Categorie   string
	Agence      string
	Annonceur   string
	Date        string
	Video       []Video
	Site        []Site
	Annonce     []Annonce
	Audio       []Audio
}

type Video struct {
	Title string
	File  string
}
type Site struct {
	Title string
	File  string
}
type Annonce struct {
	Title string
	File  string
}
type Audio struct {
	Title string
	File  string
	Image string
}
