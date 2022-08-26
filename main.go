package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprintf(writer, "string: %s", "123")
		check(err)
	})
	errHttp := http.ListenAndServe(":80", router)
	check(errHttp)
}
