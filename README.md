# go-rate-bot

> A code challenge homework task in Go.

## Introduction

Thank you for the opportunity to partcipate in this code challenge.

This is my first time doing Go. My ramblings on this journey from start to finish are in the [DevNotes.md](https://github.com/delxa/go-rate-bot/blob/master/DevNotes.md "Might want to find a comfortable chair.") file. This is in lieu of not keeping a running commit history of progress.

[BinnedCode.md](https://github.com/delxa/go-rate-bot/blob/master/BinnedCode.md "No promises that it won't make you stab your own eyes.") contains some of the code that was ultimately discarded or otherwise refactored as the project came together.


## What I've done

The app I've written basically consists of:

- structs representing the zones and rates `rate` and `quote`
- A struct representing the main mechanism of registering these and offering methods to query them `quoteBot`
- A webserveer to register zones and rates and then request a quote via REST calls.

Registered zones and rates are persisted in memory. When the app dies, so do they.

## What I've not done

Being limited in both time and experience with Go, there were some compromises to be made.

- Used any external libraries or types (decimal. mux, etc)
- Shaped the rates object returned by the quotes endpoint
- Tests: I tried testing something basic but I ran out of time to dig into why it wasn't working.
- A good job at managing pointers: I need to do more reading and more experimentation
- A good job managing types: Cost was never intended to be a Float64 but in the end, I ran out of time to optimise for a more appropriate type.
- Error handling: Because no one ever makes mistakes.
- Setup to handle weights under 1kg (again, because types) 

## How to use

### Installation and running

Assumes you've got Go installed

1. Clone this repo into your GOPATH (I'll be honest, I'm still working this part out)
2. `go build` to compile the binary
3. `chmod +x go-rate-bot` to give it execute permissions
4. `./go-rate-bot` to fire it up
5. The webserver will be running on `localhost:8000`

*Note:* Ordinarily, you would start with a blank slate and have to create the data yourself via the POST requests detailed below. To save time, if you start the app with the `-scaffold` flag, it will create the data from the sample events when the app boots. 

### Register a zone

`POST /zones`

```json
{
  "Name": "MEL",
  "Postcodes": ["3000", "3001"]
}
```

You can then `GET /zones` to observe the registered zones


### Register a rate

`POST /rates`

```json
{
  "MaxWeight": 7,
  "Cost": 5,
  "FromZone": "MEL",
  "ToZone": "SYD"
}
```

You can then `GET /rates` to observe the registered rates

*Note:* RateIDs are concatenated from To, From and weight inputs.


### Get a quote

`GET /quotes/?weight={weight}&from={from}&to={to}`

where:

- `{weight}` is the shipping weight
- '{from}' is the Postcode of the shipper
- '{to}' is the Postcode of the recipient

Hitting `/quotes/?weight=2&from=MEL&to=SYD` will yield:

```json
[
  {
    "RateID": "MEL-SYD-2kg",
    "MaxWeight": 2,
    "Cost": 2.5,
    "FromZone": "MEL",
    "ToZone": "SYD"
  },
  {
    "RateID": "MEL-SYD-7kg",
    "MaxWeight": 7,
    "Cost": 5,
    "FromZone": "MEL",
    "ToZone": "SYD"
  }
]
```

## Conclusions

Here are some parting thoughts on pulling this together.

- Due to the newness to the language, I've delivered very literally against the requirements, rather than trying to reason about a more efficient approach to achieving the business/experience outcomes. Realistically, I'm sure there were much more efficient means of storing the data, enabling even faster querying.
- Dealing with mainly scripted and non-typed languages before, these topics were most challenging. I do need to spend more time reading and experimenting with more bad code to understand this in more detail through real-world consequences.
- I feel as though i've approached this as if I were writing it in a scripted language. I did catch myself on a few things, particularly unnecessarily breaking out "convenience" functions that should have been inlined.
- The router is verbose and repetitive as hell. Using MUX would probably simplify thing. I ran out of time to refactor.
- This was a huge amount of fun and I'm super excited to try and write more and better Go in the future.