package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

var APP_PASSWORD = os.Getenv("app_password")
var APP_LOGIN = os.Getenv("app_login")
var APP_IP = os.Getenv("app_ip")
var APP_PORT = os.Getenv("app_port")
var APP_DBNAME = os.Getenv("app_dbname")

var db_path = fmt.Sprintf("%s:%s@(%s:%s)/%s", APP_LOGIN, APP_PASSWORD, APP_IP, APP_PORT, APP_DBNAME)

func main() {
	fmt.Printf("appLogin: %s\nappPassword: %s\nappIp: %s\nappPort: %s\nappDBname: %s\n", APP_LOGIN, APP_PASSWORD, APP_IP, APP_PORT, APP_DBNAME)
	db, err := sql.Open("mysql", db_path)
	check(err)
	check(db.Ping())
	fmt.Println("DataBase ping is OK")
}
