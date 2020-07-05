My starting point for this assignment is this:

- I'm new to Go, having done only the equivelent of Hello World
- The the README and
- THe weekend

The assumed scope for this, based on documentation and hints provided prior
- a Go runtime
- that takes the events defining zones and rates and somehow registers them
- then takes quote requests and returns 0-n matching rates

I've assumed part of what I'll need to solve as part of this task is:

- reading the events in
- data structures representing the zones and rates themselves, as well as storing the collections thereof
- An iteration and/or filtering algorithm for giving us the valid rates
- Reading in a quote request
- Respondign to the quote request with rates (or lack thereof)



----

So i've created structs representing Zones and Rates and the functions that create them

- the newZone() function simply passes in the values and returns the new struct
- the newRate() function does some magic to interpolate some of the values to create the ID and the converts the number to an int64 by multiplying x 100 and casting it from float64.
- the zone struct stores the postcodes as a slice of strings.


At this point, I've discovered

- The most common approach to handling decimals (currency) is to either store them as int64s and multiple by 100 (the approach i've taken), install the decimal pkg, or create your own type entirely.
- WHen choosing slices or arrays for postcodes, I decided on slices because I didn't know what length would be coming through ofr each of the zones. Given that a slice points to an underlying array anyway, for my use case, this is fine.


That appears to have worked well. So now it's time to see how we go creating slices of these structs

----

I've created two slices capable of storing the zone and rate structs and have iterated them using for loops. I can access the values, formatting them and logging them out.

Some concepts that I have started to grapple with:

- Pointers: My new{[}Struct} functions were all setup to return pointers and as such, the slices representing the collections contained the pointers to those objects. Iterating them and reading out values caused issues. `runtime error: invalid memory address or nil pointer dereference` I made them not return pointers and itall works but intuitively, this feels wrong. I need to do more reading to understand what this is meant to look like.
- Make/New: This relates to the above so I need work out what I'm going. I know that new returns pointers and make does not. THese conventions flow through to the names of constructor functions.


Next steps at this point are to introduce methods to start finding an appropriate means of matching a rate. Given that the input is a postcode, a lookup is a two-step process to first match to/from zones (Names), then match rates, based on both zones and weight.

My thinking is that the zone and rates collections are structs instead and the lookups are methods, along with the registration.

----

A quick refactor of that code as per latest thinking about collecitons as structs has paid off. Miraculously, it worked first go. Lol. Go
And it is much cleaner and much less repetitive. I might move the creation of rate/zone instances to into the register functions.

Now i've got structs for zones/rates, I'm going to start writing the query methods.

----

Query methods are now written and work correctly

- The Zone Collection struct has a lookup for returning a Zone Name for a given postcode. It's rudimentary with no early return as yet. It has two loops, the inner of which is broken out into a boolean-returning contains() func

- The Rates Collection struct has a lookup returning rates matching the given to/from postcodes and weight. This is a single loop that just asserts against each of the members of the rate.

Some musings around performance optimisation of these lookups:

- It would be interesting to benchmark different approaches to this to progressively reduce the set, rather than interrogating each property in every rate struct returned. An initial filter to match weight first might be good, especially if the weight is on the higher side. (fewer rates available) Then do the more expensive string matching operation.

- If you really wanted to optimise for speed, you could take an adaptive filtering approach by building an index representing how to, from and weight is distributed, recalculating with each registration of a rate. Ideally, you would want to optimise for a balance of least expensive comparison to eliminating the most options up front.


Thoughts at this stage:
- This file is gettign way big. it's time to start breaking some things out
- In tying this together, I've stepped well-clear of my comfort zone,
- I dread to think how many sacred Go cows I've managed to kill so far. Probably all of them.

Next steps:
- I'm thinking a struct that stores the rates and zones together and offers an interface for doing the combined lookup should work well, largely because:
  - It simplifies the interface to the rest of the app
  - It will reduce use of unnecesary global vars

----

The quoteBot struct is complete

- Stores the zoneCollection and rateCollection
- Offers registration functions that inline the newZone/newRate creation and trigger each structs own registration function (I'm thinking now this is probable an unneceassry duplication)
- Also contains the magic quote() function that triggers both the look up functions. It has an early return if either postcodes return no value.

Functionally, it is working, with the events and IO being called by way of functions for the quoteBot itself.

Worth mentioning, all of this is still in the main go file at the moment. :-S

I would have loved the opportunity to pick someone's brain on this as I went. My over-arching feeling is that "i'm a long way from home now."

Onwards.

Next: I'm going to try and build a small webserver to actually parse the events and I/O as requests. Failing that, I might actually just read in the JSON files.

----

OK. I've started implementing the JSON webserver and now I find this gem: "You need to export the fields in Structs by capitalizing the first letter in the field name." Blergh.

- /rates and /zones return JSON representations of their registered members. I set the headers to return them.
- The response type is also set correctly to application/json

I was looking at the code and realised with the quoteBot stuct now handling the bulk of things, there is no longer a need for separate structs for managing Zone and Rate collections. I'll be refactoring this to:

- Move the methods to the quoteBot struct
- Change quoteBot to have its members jsut be []rate and []zone
- There was a function to get both postcodes at once. I think this is some old scripted langauge bad habits coming through. Only one thign calls it. It's going away too.

Next, I'll allow the router to handle multiple methods so we can use POST calls to register new rates and zones.

--

Multiple methods enabled. and that refactoring has been done. Now I'm having issues with the tpying as the JSON is decoded to the rate struct. Changing the rate struct member Cost to be float64 feels dirty but I just don't have the time to do it, nor the knowledge to do it expediently. I'm doing it. Don't @ me.

Router is done and working.

New rates and zones can be registered and quote requests are done via GET requests.  This all appears to work correctly.

Unfortunately, it is all still in the one file. I need to find out how to break things down.

It would also be great to script requests to help scaffold the data in there.

---

