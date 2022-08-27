package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func pingDB() error {
	var APP_LOGIN = os.Getenv("app_login")
	var APP_PASSWORD = os.Getenv("app_password")
	var APP_IP = os.Getenv("app_ip")
	var APP_PORT = os.Getenv("app_port")
	var APP_DBNAME = os.Getenv("app_dbname")

	var db_path = fmt.Sprintf("%s:%s@(%s:%s)/%s", APP_LOGIN, APP_PASSWORD, APP_IP, APP_PORT, APP_DBNAME)

	fmt.Printf("appLogin: %s\nappPassword: %s\nappIp: %s\nappPort: %s\nappDBname: %s\n", APP_LOGIN, APP_PASSWORD, APP_IP, APP_PORT, APP_DBNAME)
	db, err := sql.Open("mysql", db_path)
	check(err)
	if db.Ping() == nil {
		return nil
	}
	fmt.Println("DataBase ping is OK")
	return nil
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprintf(writer, "string: %s", "Hello")
		check(err)
	})
	router.HandleFunc("/func", func(writer http.ResponseWriter, request *http.Request) {
		err := pingDB()
		check(err)

	})
	router.HandleFunc("/directly", func(writer http.ResponseWriter, request *http.Request) {

	})
	errHttp := http.ListenAndServe(":80", router)
	check(errHttp)
}
