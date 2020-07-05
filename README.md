# go-rate-bot

> A code challenge homework task in Go.

## Introduction

This is my first time doing Go. My ramblings on this journey from start to finish are in the DevNotes.md file. This is in lieu of not keeping a running commit history of progress.

BinnedCode.md contains some of the code that was ultimately discarded or otherwise refactored as the project came together.


## What I've done

The app I've written basically consists of:

- structs representing the zones and rates `rate` and `quote`
- A struct representing the main mechanism of registering these and offering methods to query them `quoteBot`
- A webserveer to register zones and rates and then request a quote via REST calls.

## What I've not done

- Used any external libraries or types (decimal. mux, etc)
- Shaped the rates object returned by the quotes endpoint
- Tests
- A good job at managing pointers: I need to do more reading and more experimentation
- A good job managing types: Cost was never intended to be a Float64 but in the end, I ran out of time to optimise for a more appropriate type.
- Error handling: Because no one ever makes mistakes.
- Setup to handle weights under 1kg (again, because types) 

## How to use

### Installation and running

Assumes you've got Go installed

1. Clone this repo into your GOPATH
2. `go build` to compile the binary
3. `chmod +x go-rate-bot` to give it execute permissions
4. `./go-rate-bot` to fire it up
5. The webserver will be running on `localhost:8000`


### Register a zone

`POST /zones`

```json
{
  "Name": "MEL",
  "Postcodes": ["3000", "3001"]
}
```

You can then GET /zones to observe the registered zones


### Register a rate

`POST /rates`

```json
{
  "Name": "MEL",
  "Postcodes": ["3000", "3001"]
}
```

*Note:* RateIDs are concatenated from To, From and weight inputs.

You can then `GET /rates` to observe the registered rates


### Get a quote

`GET /quotes/?weight={weight}&from={from}&to={to}`

where:

- `{weight}` is the shipping weight
- '{from}' is the Postcode of the shipper
- '{to}' is the Postcode of the recipient

## Conclusions

Here are some parting thoughts on pulling this together.

- Due to the newness to the language, I've delivered very literally against the requirements, rather than trying to reason about a more efficient approach to achieving the business/experience outcomes. Realistically, I'm sure there were much more efficient means of storing the data, enabling even faster querying.
- Dealing with mainly scripted and non-typed languages before, these topics were most challenging. I do need to spend more time reading and experimenting with more bad code to understand this in more detail through real-world consequences.
- I feel as though i've approached this as if I were writing it in a scripted language. I did catch myself on a few things, particularly unnecessarily breaking out "convenience" functions that should have been inlined.
- This was a huge amount of fun and I'm super excited to try and write more and better Go in the future.