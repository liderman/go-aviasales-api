package aviasales

type Country struct {
	Code             string            `json:"code"`
	Name             string            `json:"name"`
	Currency         string            `json:"currency"`
	NameTranslations map[string]string `json:"name_translations"`
}

// DataCountries returns a list of countries from the database.
//
// Documentation:
// https://support.travelpayouts.com/hc/ru/articles/203956163#09
//
// Example URI:
// http://api.travelpayouts.com/data/countries.json
func (a *AviasalesApi) DataCountries() (countries []Country, err error) {
	err = a.getJson("data/countries.json", map[string]string{}, &countries)
	return
}
