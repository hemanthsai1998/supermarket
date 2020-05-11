package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	super "supermarket/market"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.HandleFunc("/super_market", itemPost).Methods("POST")
	router.HandleFunc("/super_market", itemGetAll).Methods("GET")
	router.HandleFunc("/super_market/{item}", itemGet).Methods("GET")
	router.HandleFunc("/super_market", itemUpdate).Methods("PUT")
	router.HandleFunc("/super_market/{item}", itemDelete).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to SuperMarket")
}

func itemPost(w http.ResponseWriter, r *http.Request) {
	var temp map[string]interface{}
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "enter data in correct format")
	}
	json.Unmarshal(req, &temp)
	for item, value := range temp {
		err := super.Post(item, value)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			//fmt.Fprintf(w, "Item already present")
		} else {
			fmt.Fprintf(w, "Item entered successfully")
			//fmt.Fprintln(w, super.Print())
		}
	}
}
func itemUpdate(w http.ResponseWriter, r *http.Request) {
	var temp map[string]interface{}
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "enter data in correct format")
	}
	json.Unmarshal(req, &temp)
	for item, value := range temp {
		err := super.Put(item, value)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			//fmt.Fprintf(w, "Item already present")
		} else {
			fmt.Fprintf(w, "Item updated successfully")
			//fmt.Fprintln(w, super.Print())
		}
	}
}
func itemGetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, super.Print())
}

func itemGet(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["item"]
	val, err := super.Get(name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		fmt.Fprintln(w, val)
		//fmt.Fprintln(w, super.Print())
	}
}
func itemGet(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["item"]
	err := super.Delete(name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		fmt.Fprintf(w, "Item deleted successfully")
	}
}
