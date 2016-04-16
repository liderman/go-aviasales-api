package aviasales

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

// PricesCheap returns the cheapest direct, as well as 1 and 2 transfers to the chosen direction with filters by date of departure and return. Tickets return flight there and back.
// Note! If the request specifies the old date, as a result of his work is not an error, but no data will not be back.
//
// Documentation:
// https://support.travelpayouts.com/hc/ru/articles/203956163#05
//
// Example URI:
// http://api.travelpayouts.com/v1/prices/cheap?origin=MOW&destination=HKT&depart_date=2016-11&return_date=2016-12&token=YOUR_TOKEN
func (a *AviasalesApi) PricesCheap(input InputPricesCheap) (flights DataFlight, err error) {
	err = a.getJson("v1/prices/cheap", map[string]string{
		"currency":    input.Currency,
		"origin":      input.Origin,
		"destination": input.Destination,
		"depart_date": input.DepartDate,
		"return_date": input.ReturnDate,
		"token":       a.token,
	}, &flights)
	return
}
