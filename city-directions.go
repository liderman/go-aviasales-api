package aviasales

type DataCityDirectionsFlight struct {
	Success bool              `json:"success" bson:"success"`
	Data    map[string]Flight `json:"data" bson:"data"`
}

// CityDirections returns most popular destinations from the specified city.
//
// Documentation:
// https://support.travelpayouts.com/hc/ru/articles/203956163#20
//
// Example URI:
// http://api.travelpayouts.com/v1/city-directions?origin=MOW&token=YOUR_TOKEN
func (a *AviasalesApi) CityDirections(origin string) (flights DataCityDirectionsFlight, err error) {
	err = a.getJson("v1/city-directions", map[string]string{
		"origin": origin,
		"token":  a.token,
	}, &flights)
	return
}
