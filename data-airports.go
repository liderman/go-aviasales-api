package aviasales

type Airport struct {
	Code             string            `json:"code"`
	Name             string            `json:"name"`
	Coordinates      Coordinates       `json:"coordinates"`
	TimeZone         string            `json:"time_zone"`
	CountryCode      string            `json:"country_code"`
	CityCode         string            `json:"city_code"`
	NameTranslations map[string]string `json:"name_translations"`
}

// DataAirports a list of airports in the database.
//
// Documentation:
// https://support.travelpayouts.com/hc/ru/articles/203956163#11
//
// Example URI:
// http://api.travelpayouts.com/data/airports.json
func (a *AviasalesApi) DataAirports() (airports []Airport, err error) {
	err = a.getJson("data/airports.json", map[string]string{}, &airports)
	return
}
