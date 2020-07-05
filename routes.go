package main

import (
  "encoding/json"
  "fmt"
  "net/http"
  "strconv"
)

/**
  Route handlers
**/

func ratesEndpoint(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Endpoint Hit: /rates")
  switch r.Method {
  case "GET":
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(qb.rates)
  case "POST":
    d := json.NewDecoder(r.Body)
    rt := &rate{}
    err := d.Decode(rt)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    qb.registerRate(rt.MaxWeight, rt.Cost, rt.FromZone, rt.ToZone)
    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "New rate registered.")
  default:
    w.WriteHeader(http.StatusMethodNotAllowed)
    fmt.Fprintf(w, "Method not allowed on this endpoint.")
  }
}

func zonesEndpoint(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Endpoint Hit: /zones")
  switch r.Method {
  case "GET":
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(qb.zones)
  case "POST":
    d := json.NewDecoder(r.Body)
    z := &zone{}
    err := d.Decode(z)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    qb.registerZone(z.Name, z.Postcodes)
    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "New zone registered.")
  default:
    w.WriteHeader(http.StatusMethodNotAllowed)
    fmt.Fprintf(w, "Method not allowed on this endpoint.")
  }
}

func quotesEndpoint(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Endpoint Hit: /quotes")
  switch r.Method {
  case "GET":
    w.Header().Set("Content-Type", "application/json")
    q := r.URL.Query()
    wt, err := strconv.Atoi(q.Get("weight"))
    if err != nil {
      http.Error(w, err.Error(), http.StatusBadRequest)
    }
    f := q.Get("from")
    t := q.Get("to")
    json.NewEncoder(w).Encode(qb.matchingRates(f, t, wt))
  default:
    w.WriteHeader(http.StatusMethodNotAllowed)
    fmt.Fprintf(w, "Method not allowed on this endpoint.")
  }
}
