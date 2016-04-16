package aviasales

import "fmt"

type DataAirlineDirections struct {
	Success  bool           `json:"success"`
	Data     map[string]int `json:"data"`
	Error    string         `json:"error"`
	Currency string         `json:"currency"`
}

// AirlineDirections returns the ways in which the airline operates flights sorted by popularity.
//
// Documentation:
// https://support.travelpayouts.com/hc/ru/articles/203956163#08
//
// Example URI:
// http://api.travelpayouts.com/v1/airline-directions?airline_code=SU&limit=10&token=YOUR_TOKEN
func (a *AviasalesApi) AirlineDirections(airlineCode string, limit int) (airlineDirections DataAirlineDirections, err error) {
	err = a.getJson("v1/airline-directions", map[string]string{
		"airline_code": airlineCode,
		"limit":        fmt.Sprintf("%d", limit),
		"token":        a.token,
	}, &airlineDirections)
	return
}
