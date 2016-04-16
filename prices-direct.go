package aviasales

import "fmt"

type InputPricesCheap struct {
	Origin      string
	Destination string
	DepartDate  string
	ReturnDate  string
	Currency    string
}

type DataFlight struct {
	Success bool                         `json:"success"`
	Data    map[string]map[string]Flight `json:"data"`
}

type Flight struct {
	Price        int    `json:"price"`
	Airline      string `json:"airline"`
	FlightNumber int    `json:"flight_number"`
	DepartureAt  string `json:"departure_at"`
	ReturnAt     string `json:"return_at"`
	ExpiresAt    string `json:"expires_at"`
}

// PricesDirect returns the cheapest ticket direct to the chosen direction with filters by date of departure and return.
//
// Documentation:
// https://support.travelpayouts.com/hc/ru/articles/203956163#06
//
// Example URI:
// http://api.travelpayouts.com/v1/prices/direct?origin=MOW&destination=LED&depart_date=2016-11&return_date=2016-12&token=YOUR_TOKEN
func (a *AviasalesApi) PricesDirect(input InputPricesCheap) (flights DataFlight, err error) {
	err = a.getJson("v1/prices/direct", map[string]string{
		"currency":    input.Currency,
		"origin":      input.Origin,
		"destination": input.Destination,
		"depart_date": input.DepartDate,
		"return_date": input.ReturnDate,
		"token":       a.token,
	}, &flights)
	return
}
