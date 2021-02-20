package main

import "github.com/gorilla/mux"

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/user/create", createUser).Methods("POST")
    router.HandleFunc("/user/get", getUser).Methods("POST")
}