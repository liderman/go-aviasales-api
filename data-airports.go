package aviasales

type Airport struct {
	Code             string            `json:"code" bson:"code"`
	Name             string            `json:"name" bson:"name"`
	Coordinates      Coordinates       `json:"coordinates" bson:"coordinates"`
	TimeZone         string            `json:"time_zone" bson:"time_zone"`
	CountryCode      string            `json:"country_code" bson:"country_code"`
	CityCode         string            `json:"city_code" bson:"city_code"`
	NameTranslations map[string]string `json:"name_translations" bson:"name_translations"`
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
