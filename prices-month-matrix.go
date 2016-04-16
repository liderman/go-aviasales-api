package aviasales

import "fmt"

type InputPricesMonthMatrix struct {
	Currency         string
	Origin           string
	Destination      string
	ShowToAffiliates bool
	Month            string
}

// PricesMonthMatrix returns the price for each day of the month, grouped by the number of transfer.
// Note! If you do not specify a starting point and destination, the API returns a list of the cheapest tickets, which were found in the last 48 hours.
//
// Documentation:
// https://support.travelpayouts.com/hc/ru/articles/203956163#03
//
// Example URI:
// http://api.travelpayouts.com/v2/prices/month-matrix?currency=rub&origin=LED&destination=HKT&show_to_affiliates=true&token=YOUR_TOKEN
func (a *AviasalesApi) PricesMonthMatrix(input InputPricesMonthMatrix) (prices DataPrice, err error) {
	err = a.getJson("v2/prices/month-matrix", map[string]string{
		"currency":           input.Currency,
		"origin":             input.Origin,
		"destination":        input.Destination,
		"show_to_affiliates": fmt.Sprint(input.ShowToAffiliates),
		"month":              input.Month,
		"token":              a.token,
	}, &prices)
	return
}
