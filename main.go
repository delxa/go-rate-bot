package main

import (
  "flag"
	"fmt"
  "log"
	"net/http"
)

// Probably also should have gone in routes.go
func handleRequests() {
	http.HandleFunc("/rates", ratesEndpoint)
  http.HandleFunc("/zones", zonesEndpoint)
  http.HandleFunc("/quotes", quotesEndpoint)
	fmt.Println("Server up on localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// Made this a global variable to greatly simplify playing with the webserver.
var qb = newQuoteBot()

/* The main() event */
func main() {

  scaffold := flag.Bool("scaffold", false, "Create some default zones and rates")
  flag.Parse()

	// These mimic the events that come through the webserver. use -scaffold in the command line to have at them.
	if *scaffold == true {
    qb.registerZone("MEL", []string{"3000", "3001"})
  	qb.registerZone("SYD", []string{"2000", "2001"})
  	qb.registerZone("BRI", []string{"4000"})

  	qb.registerRate(2, 2.5, "MEL", "SYD")
  	qb.registerRate(7, 5, "MEL", "SYD")
  	qb.registerRate(4, 3.55, "MEL", "BRI")
    fmt.Println("Scaffolded demo data.")
  }
	handleRequests()
}
