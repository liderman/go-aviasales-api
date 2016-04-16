package aviasales

type Route struct {
	AirlineIata          string   `json:"airline_iata"`
	AirlineIcao          string   `json:"airline_icao"`
	DepartureAirportIata string   `json:"departure_airport_iata"`
	DepartureAirportIcao string   `json:"departure_airport_icao"`
	ArrivalAirportIata   string   `json:"arrival_airport_iata"`
	ArrivalAirportIcao   string   `json:"arrival_airport_icao"`
	Codeshare            bool     `json:"codeshare"`
	transfers            int      `json:"transfers"`
	planes               []string `json:"planes"`
}

// DataRoutes a list of routes from the database.
//
// Documentation:
// https://support.travelpayouts.com/hc/ru/articles/203956163#15
//
// Example URI:
// http://api.travelpayouts.com/data/routes.json
func (a *AviasalesApi) DataRoutes() (routes []Route, err error) {
	err = a.getJson("data/routes.json", map[string]string{}, &routes)
	return
}
