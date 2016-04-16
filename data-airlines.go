package aviasales

type Airline struct {
	Name     string `json:"name" bson:"name"`
	Alias    string `json:"alias" bson:"alias"`
	Iata     string `json:"iata" bson:"iata"`
	Icao     string `json:"icao" bson:"icao"`
	Callsign string `json:"callsign" bson:"callsign"`
	Country  string `json:"country" bson:"country"`
	IsActive bool   `json:"is_active" bson:"is_active"`
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
