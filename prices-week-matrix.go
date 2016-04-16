package aviasales

import "fmt"

type InputPricesWeekMatrix struct {
	Currency         string
	Origin           string
	Destination      string
	ShowToAffiliates bool
	DepartDate       string
	ReturnDate       string
}

// PricesWeekMatrix calendar of prices for a week.
// Returns  prices nearest the target dates.
//
// Documentation:
// https://support.travelpayouts.com/hc/ru/articles/203956163#19
//
// Example URI:
// http://api.travelpayouts.com/v2/prices/week-matrix?currency=rub&origin=LED&destination=HKT&show_to_affiliates=true&depart_date=2016-03-04&return_date=2016-03-18&token=YOUR_TOKEN
func (a *AviasalesApi) PricesWeekMatrix(input InputPricesWeekMatrix) (prices DataPrice, err error) {
	err = a.getJson("v2/prices/week-matrix", map[string]string{
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
