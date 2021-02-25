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

	// r.body に何が帰ってくるか確認
	fmt.Println(r.Body)

	var user User
	json.NewDecoder(r.Body).Decode(&user)

	if user.Email == "" {
		var error Error
		error.Message = "Email は必須です。"
		errorInResponse(w, http.StatusBadRequest, error)
		return
	}

	if user.Password == "" {
		var error Error
		error.Message = "パスワードは必須です。"
		errorInResponse(w, http.StatusBadRequest, error)
		return
	}

	// user に何が格納されているのか
	fmt.Println(user)

	// dump も出せる
	fmt.Println("---------------------")
	spew.Dump(user)
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
