package server

import (
	"io/ioutil"
	"log"
	"net/http"
)

const (
	country   = "assets/country.json"
	regions   = "assets/regions.json"
	districts = "assets/districts.json"
	communes  = "assets/communes.json"
	fokontany = "assets/fokontany.json"
)

func StartHTTPServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/country", func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadFile(country)
		if err != nil {
			log.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	})
	mux.HandleFunc("/regions", func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadFile(regions)
		if err != nil {
			log.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	})
	mux.HandleFunc("/districts", func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadFile(districts)
		if err != nil {
			log.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	})
	mux.HandleFunc("/communes", func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadFile(communes)
		if err != nil {
			log.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	})
	mux.HandleFunc("/fokontany", func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadFile(fokontany)
		if err != nil {
			log.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	})

	s := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	log.Fatal(s.ListenAndServe())
}
