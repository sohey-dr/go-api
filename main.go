package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successfully called createUser"))
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successfully called getUser"))
}

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/user/create", createUser).Methods("POST")
    router.HandleFunc("/user/get", getUser).Methods("POST")

		log.Println("サーバー起動 : 8080 port で受信")

    log.Fatal(http.ListenAndServe(":8080", router))
}

type User struct {
	// 大文字だと Public 扱い
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWT struct {
	Token string `json:"token"`
}

type Error struct {
	Message string `json:"message"`
}
