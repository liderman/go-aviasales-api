package aviasales

type Airline struct {
	Name     string `json:"name"`
	Alias    string `json:"alias"`
	Iata     string `json:"iata"`
	Icao     string `json:"icao"`
	Callsign string `json:"callsign"`
	Country  string `json:"country"`
	IsActive bool   `json:"is_active"`
}

// DataAirlines a list of airlines from the database.
//
// Documentation:
// https://support.travelpayouts.com/hc/ru/articles/203956163#12
//
// Example URI:
// http://api.travelpayouts.com/data/airlines.json
func (a *AviasalesApi) DataAirlines() (airlines []Airline, err error) {
	err = a.getJson("data/airlines.json", map[string]string{}, &airlines)
	return
}
