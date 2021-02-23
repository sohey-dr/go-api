package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	_"github.com/go-sql-driver/mysql"
)

func errorInResponse(w http.ResponseWriter, status int, error Error) {
	w.WriteHeader(status) // 400 とか 500 などの HTTP status コードが入る
	json.NewEncoder(w).Encode(error)
	return
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successfully called createUser"))
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successfully called getUser"))
}

var Db *sql.DB
func main() {

    var err error
    Db, err = sql.Open("mysql", "username:password@/techtrain_go_development")
    if err != nil {
        log.Fatal("DBエラー")
    }

    err = Db.Ping()

    if err != nil {
        log.Fatal(err)
    }

    router := mux.NewRouter()

    router.HandleFunc("/user/create", createUser).Methods("POST")
    router.HandleFunc("/user/get", getUser).Methods("POST")

		log.Println("サーバー起動 : 8080 port で受信")

    log.Fatal(http.ListenAndServe(":8080", router))
}

type User struct {
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
