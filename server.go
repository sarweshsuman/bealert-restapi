package main

import (
	"net/http"
	"github.com/sarweshsuman/bealert-restapi/handlers"
	"log"
	"os"
	"fmt"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("In-correct options, please only provide configuration file..")
		os.Exit(1)
	}
	config := handlers.New(os.Args[1])
	router := config.GetRouter()
	log.Fatal(http.ListenAndServe(":8080",router))
}