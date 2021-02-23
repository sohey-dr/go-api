package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/go-sql-driver/mysql"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successfully called createUser"))
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successfully called getUser"))
}

var db *sql.DB
func main() {
    // parmas.go から DB の URL を取得
    i := tool.Info{}

    // Convert
    // https://github.com/lib/pq/blob/master/url.go
    // "postgres://bob:secret@1.2.3.4:5432/mydb?sslmode=verify-full"
    // ->　"user=bob password=secret host=1.2.3.4 port=5432 dbname=mydb sslmode=verify-full"
    pgUrl, err := pq.ParseURL(i.GetDBUrl())

    // 戻り値に err を返してくるので、チェック
    if err != nil {
        // エラーの場合、処理を停止する
        log.Fatal()
    }

    // DB 接続
    db, err = sql.Open("postgres", pgUrl)
    if err != nil {
        log.Fatal(err)
    }

    // DB 疎通確認
    err = db.Ping()

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
