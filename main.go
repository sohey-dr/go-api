package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("createUser 関数実行")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getUser 関数実行")
}

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/user/create", createUser).Methods("POST")
    router.HandleFunc("/user/get", getUser).Methods("POST")

		log.Println("サーバー起動 : 8080 port で受信")

    log.Fatal(http.ListenAndServe(":8080", router))
}