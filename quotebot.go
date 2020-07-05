package main

import (
  "fmt"
)

// String appears in array
func Contains(a []string, x string) bool {
  for _, n := range a {
    if x == n {
      return true
    }
  }
  return false
}

type zone struct {
  Name      string
  Postcodes []string
}

type rate struct {
  RateID    string
  MaxWeight int
  Cost      float64 // I know this is bad. I hate myself.
  FromZone  string
  ToZone    string
}

type quoteBot struct {
  zones []zone
  rates []rate
}

/**
  Creator Functions
**/

func newZone(name string, postcodes []string) zone {
  z := zone{Name: name, Postcodes: postcodes}
  return z
}

func newRate(maxWeight int, cost float64, fromZone string, toZone string) rate {
  r := rate{MaxWeight: maxWeight, FromZone: fromZone, ToZone: toZone, Cost: cost}
  r.RateID = fmt.Sprintf("%s-%s-%dkg", fromZone, toZone, maxWeight)
  return r
}

func newQuoteBot() *quoteBot {
  q := quoteBot{}
  q.rates = make([]rate, 0)
  q.zones = make([]zone, 0)
  return &q
}

/**
  quoteBot methods
**/

func (qq *quoteBot) registerZone(name string, postcodes []string) {
  qq.zones = append(qq.zones, newZone(name, postcodes))
}

func (qq *quoteBot) registerRate(maxWeight int, cost float64, fromZone string, toZone string) {
  qq.rates = append(qq.rates, newRate(maxWeight, cost, fromZone, toZone))
}

func (qq *quoteBot) matchingRates(fromZone string, toZone string, weight int) []rate {
  rrf := make([]rate, 0)
  for _, r := range qq.rates {
    if r.FromZone == fromZone && r.ToZone == toZone && weight <= r.MaxWeight {
      rrf = append(rrf, r)
    }
  }
  return rrf
}

func (qq *quoteBot) zoneFromPostcode(postcode string) string {
  s := ""
  for _, elem := range qq.zones {
    if Contains(elem.Postcodes, postcode) {
      s = elem.Name
    }
  }
  return s
}

func (qq *quoteBot) quote(from string, to string, weight int) []rate {
  // Get the zones matching the postcodes. If we don't have one or other of these, return an empty array
  fz := qq.zoneFromPostcode(from)
  tz := qq.zoneFromPostcode(to)
  if fz == "" || tz == "" {
    return make([]rate, 0)
  }
  // Get and return the matching rates
  return qq.matchingRates(fz, tz, weight)
}
