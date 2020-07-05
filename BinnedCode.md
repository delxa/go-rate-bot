# Binned Code

I've kept these here to show you more specifically some of the earlier implementations.


These were used as the initial mocks for firing the events against the quoteBot to register rates and zones and then retrieve quotes.
Ultimately, these interfaces were accessed on webserver events


```go

  // Events
  qb.registerZone("MEL", []string{"3000", "3001"})
  qb.registerZone("SYD", []string{"2000", "2001"})
  qb.registerZone("BRI", []string{"4000"})
  
  qb.registerRate(2, 2.5, "MEL", "SYD")
  qb.registerRate(7, 5, "MEL", "SYD")
  qb.registerRate(4, 3.55, "MEL", "BRI")

  // Input/Output
  fmt.Println(qb.quote("3000", "2000", 2))
  fmt.Println(qb.quote("20", "2000", 2))
  fmt.Println(qb.quote("3000", "2000", 5))
  fmt.Println(qb.quote("3000", "4000", 1))
  fmt.Println(qb.quote("4000", "3000", 1))


  //qb.zoneCollection.printOut()
  //qb.rateCollection.printOut()

```



Before the existence of the quoteBot, these were used to maintain the collections and had their own registration functions.
After this point, they simply became useless boilerplate as the quoteBot was more than capable of directly storing and querying the data.
The query functions were kept and the rest discarded.


```go
/* Rate Collecton */

type rateCollection struct {
  rates []rate
}

func (rr *rateCollection) register(nr rate) {
  rr.rates = append(rr.rates, nr)
}

func (rr *rateCollection) printOut() {
  for index, elem := range rr.rates {
    fmt.Println(fmt.Sprintf("%d - %s (%d)", index, elem.RateID, elem.Cost))
  }
}

func (rr *rateCollection) matchingRates(fromZone string, toZone string, weight int) []rate {
  rrf := make([]rate, 0)
  for _, r := range rr.rates {
    if r.FromZone == fromZone && r.ToZone == toZone && weight <= r.MaxWeight {
      rrf = append(rrf, r)
    }
  }
  return rrf
}


  /* Zone Collecton */
type zoneCollection struct {
  zones []zone
}

func (zz *zoneCollection) register(nz zone) {
  zz.zones = append(zz.zones, nz)
}

func (zz *zoneCollection) printOut() {
  for index, elem := range zz.zones {
    fmt.Println(fmt.Sprintf("%d - %s (%s)", index, elem.Name, elem.Postcodes))
  }
}

func (zz *zoneCollection) zoneFromPostcode(postcode string) string {
  s := "" 
  for _, elem := range zz.zones {
    if contains(elem.Postcodes, postcode) {
      s = elem.Name
    }
  }
  return s
}

func (zz *zoneCollection) zonesFromPostcodes(fromPostcode string, toPostcode string) (string, string) {
  return zz.zoneFromPostcode(fromPostcode), zz.zoneFromPostcode(toPostcode)
}

```


Before it was refactored, you can see the level of duplication in this particular implementation. Disgusting, the things we do.

```go

/* QuoteBot */

type quoteBot struct {
  zoneCollection zoneCollection
  rateCollection rateCollection
}

func newQuoteBot () *quoteBot {
  q := quoteBot{}
  q.rateCollection = rateCollection{}
  q.zoneCollection = zoneCollection{}  
  return &q
}

func (qq *quoteBot) registerZone(name string, postcodes []string) {
  qq.zoneCollection.register(newZone(name, postcodes))
}

func (qq *quoteBot) registerRate(maxWeight int, cost float64, fromZone string, toZone string) {
  qq.rateCollection.register(newRate(maxWeight, cost, fromZone, toZone))
}

func (qq *quoteBot) quote(from string, to string, weight int) []rate {
  // Get the zones matching the postcodes. If we don't have one or other of these, return an empty array
  fz, tz := qq.zoneCollection.zonesFromPostcodes(from, to)
  if fz == "" || tz == "" {
    return make([]rate, 0)
  }
  // Get and return the matching rates
  return qq.rateCollection.matchingRates(fz, tz, weight)

}
```