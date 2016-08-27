# go-aviasales-api
Golang implementation Aviasales API for data access

[![GoDoc](https://godoc.org/github.com/liderman/go-aviasales-api?status.svg)](https://godoc.org/github.com/liderman/go-aviasales-api)

Installation
-----------
	go get github.com/liderman/go-aviasales-api

Usage
-----------
Creates a new instance AviasalesApi:
```go
aviaApi := aviasales.NewAviasalesApi("YOUR_TOKEN")
```

Getting directions in which the airline operates flights sorted by popularity:
```go
airlineDirections, err := aviaApi.AirlineDirections("SU", 10)
```

Requirements
-----------

* Need at least `go1.5` or newer.

Documentation
-----------

You can read package documentation [here](http:godoc.org/github.com/liderman/go-aviasales-api).
