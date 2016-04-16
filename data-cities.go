package aviasales

type City struct {
	Code             string            `json:"code"`
	Name             string            `json:"name"`
	Coordinates      Coordinates       `json:"coordinates"`
	TimeZone         string            `json:"time_zone"`
	CountryCode      string            `json:"country_code"`
	NameTranslations map[string]string `json:"name_translations"`
}

// DataCities returns a list of cities in the database.
//
// Documentation:
// https://support.travelpayouts.com/hc/ru/articles/203956163#10
//
// Example URI:
// http://api.travelpayouts.com/data/cities.json
func (a *AviasalesApi) DataCities() (cities []City, err error) {
	err = a.getJson("data/cities.json", map[string]string{}, &cities)
	return
}
