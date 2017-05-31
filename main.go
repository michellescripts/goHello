package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Fruit struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Color  string `json:"color,omitempty"`
	Rating int    `json:"rating,omitempty"`
}

var fruits []Fruit

func GetFruitsEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range fruits {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Fruit{})
}

func GetFruitEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(fruits)
}

func CreateFruitEndpoint(w http.ResponseWriter, req *http.Request) {
	var fruit Fruit
	_ = json.NewDecoder(req.Body).Decode(&fruit)
	fruits = append(fruits, fruit)
	json.NewEncoder(w).Encode(fruits)
}

func DeleteFruitEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range fruits {
		if item.ID == params["id"] {
			fruits = append(fruits[:index], fruits[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(fruits)
}

func UpdateFruitEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var fruit Fruit
	for index, item := range fruits {
		if item.ID == params["id"] {
			fruits = append(fruits[:index], fruits[index+1:]...)
			_ = json.NewDecoder(req.Body).Decode(&fruit)
			fruits = append(fruits, fruit)
			break
		}
	}
	json.NewEncoder(w).Encode(fruits)
}

func main() {
	router := mux.NewRouter()
	fruits = append(fruits, Fruit{ID: "1", Name: "Orange", Color: "Orange", Rating: 1})
	fruits = append(fruits, Fruit{ID: "2", Name: "Banana", Color: "Yellow", Rating: 5})
	router.HandleFunc("/fruits", GetFruitEndpoint).Methods("GET")
	router.HandleFunc("/fruits/{id}", GetFruitsEndpoint).Methods("GET")
	router.HandleFunc("/fruits", CreateFruitEndpoint).Methods("POST")
	router.HandleFunc("/fruits/{id}", DeleteFruitEndpoint).Methods("DELETE")
	router.HandleFunc("/fruits/{id}", UpdateFruitEndpoint).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", router))
}
