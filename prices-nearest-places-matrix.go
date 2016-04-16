package aviasales

import "fmt"

type InputPricesNearestPlacesMatrix struct {
	Currency         string
	Origin           string
	Destination      string
	ShowToAffiliates bool
	DepartDate       string
	ReturnDate       string
}

type DataPriceNearest struct {
	Success bool           `json:"success"`
	Data    []PriceNearest `json:"data"`
}

type PriceNearest struct {
	Origins      []string `json:"origins"`
	Destinations []string `json:"destinations"`
	Prices       []Price  `json:"prices"`
}

// PricesNearestPlacesMatrix returns prices directions between nearest to the target cities.
// Note! If you do not specify a starting point and destination, the API returns a list of the cheapest tickets, which were found in the last 48 hours.
//
// Documentation:
// https://support.travelpayouts.com/hc/ru/articles/203956163#04
//
// Example URI:
// http://api.travelpayouts.com/v2/prices/nearest-places-matrix?currency=rub&origin=LED&destination=HKT&show_to_affiliates=true&token=YOUR_TOKEN
func (a *AviasalesApi) PricesNearestPlacesMatrix(input InputPricesNearestPlacesMatrix) (prices DataPriceNearest, err error) {
	err = a.getJson("v2/prices/nearest-places-matrix", map[string]string{
		"currency":           input.Currency,
		"origin":             input.Origin,
		"destination":        input.Destination,
		"show_to_affiliates": fmt.Sprint(input.ShowToAffiliates),
		"depart_date":        input.DepartDate,
		"return_date":        input.ReturnDate,
		"token":              a.token,
	}, &prices)
	return
}
