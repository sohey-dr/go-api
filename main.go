package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func errorInResponse(w http.ResponseWriter, status int, error Error) {
	w.WriteHeader(status) // 400 とか 500 などの HTTP status コードが入る
	json.NewEncoder(w).Encode(error)
	return
}

func createToken(user User) (string, error) {
	var err error

	// 鍵となる文字列(多分なんでもいい)
	secret := "secret"

	// Token を作成
	// jwt -> JSON Web Token - JSON をセキュアにやり取りするための仕様
	// jwtの構造 -> {Base64 encoded Header}.{Base64 encoded Payload}.{Signature}
	// HS254 -> 証明生成用(https://ja.wikipedia.org/wiki/JSON_Web_Token)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": user.Name,
		"iss":  "__init__", // JWT の発行者が入る(文字列(__init__)は任意)
	})

	//Dumpを吐く
	spew.Dump(token)

	tokenString, err := token.SignedString([]byte(secret))

	fmt.Println("-----------------------------")
	fmt.Println("tokenString:", tokenString)

	if err != nil {
		log.Fatal(err)
	}

	return tokenString, nil
}

func createUser(w http.ResponseWriter, r *http.Request) {

	// r.body に何が帰ってくるか確認
	fmt.Println(r.Body)

	var user User
	json.NewDecoder(r.Body).Decode(&user)

	if user.Name == "" {
		var error Error
		error.Message = "Name は必須です。"
		errorInResponse(w, http.StatusBadRequest, error)
		return
	}

	// if user.Password == "" {
	// 	var error Error
	// 	error.Message = "パスワードは必須です。"
	// 	errorInResponse(w, http.StatusBadRequest, error)
	// 	return
	// }

	// user に何が格納されているのか
	fmt.Println(user)

	// dump も出せる
	fmt.Println("---------------------")
	spew.Dump(user)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	var user User
	var error Error

	json.NewDecoder(r.Body).Decode(&user)

	if user.Name == "" {
		error.Message = "Email は必須です。"
		errorInResponse(w, http.StatusBadRequest, error)
		return
	}

	// if user.Password == "" {
	// 	error.Message = "パスワードは、必須です。"
	// 	errorInResponse(w, http.StatusBadRequest, error)
	// }

	// 認証キー(Emal)のユーザー情報をDBから取得
	row := Db.QueryRow("SELECT * FROM USERS WHERE email=$1;", user.Name)
	err := row.Scan(&user.ID, &user.Name)

	if err != nil {
		if err == sql.ErrNoRows { // https://golang.org/pkg/database/sql/#pkg-variables
			error.Message = "ユーザが存在しません。"
			errorInResponse(w, http.StatusBadRequest, error)
		} else {
			log.Fatal(err)
		}
	}
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
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type JWT struct {
	Token string `json:"token"`
}

type Error struct {
	Message string `json:"message"`
}
