package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/Aarav238/mongoapi/router"
)

func main () {
	fmt.Println("MongoAPi")
	r:= router.Router()
	fmt.Println("server is getting started")
	log.Fatal(http.ListenAndServe(":4000",r))
	fmt.Println("listening at port 4000")

}