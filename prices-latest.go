package aviasales

import "fmt"

type InputPricesLatest struct {
	Currency          string
	Origin            string
	Destination       string
	BeginningOfPeriod string
	PeriodType        string
	OneWay            bool
	Page              int
	Limit             int
	ShowToAffiliates  bool
	Sorting           string
	TripClass         int
	TripDuration      string
}

type DataPrice struct {
	Success bool    `json:"success"`
	Data    []Price `json:"data"`
}

type Price struct {
	Actual           bool   `json:"actual"`
	DepartDate       string `json:"depart_date"`
	Destination      string `json:"destination"`
	Distance         int    `json:"distance"`
	FoundAt          string `json:"found_at"`
	NumberOfChanges  int    `json:"number_of_changes"`
	Origin           string `json:"origin"`
	ReturnDate       string `json:"return_date"`
	ShowToAffiliates bool   `json:"show_to_affiliates"`
	TripClass        int    `json:"trip_class"`
	Value            int    `json:"value"`
}

// PricesLatest returns a list of the prices found aviasales users in the last 48 hours, in accordance with the exposed filters.
// Note! If you do not specify a starting point and destination, the API returns 30 for the cheapest tickets, which were found in the last 48 hours.
//
// Documentation:
// https://support.travelpayouts.com/hc/ru/articles/203956163#02
//
// Example URI:
// http://api.travelpayouts.com/v2/prices/latest?currency=rub&period_type=year&page=1&limit=30&show_to_affiliates=true&sorting=price&trip_class=0&token=YOUR_TOKEN
func (a *AviasalesApi) PricesLatest(input InputPricesLatest) (prices DataPrice, err error) {
	query := map[string]string{
		"currency":            input.Currency,
		"origin":              input.Origin,
		"destination":         input.Destination,
		"beginning_of_period": input.BeginningOfPeriod,
		"period_type":         input.PeriodType,
		"one_way":             fmt.Sprint(input.OneWay),
		"show_to_affiliates":  fmt.Sprint(input.ShowToAffiliates),
		"sorting":             input.Sorting,
		"trip_duration":       input.TripDuration,
		"token":               a.token,
	}
	if input.Page > 0 {
		query["page"] = fmt.Sprintf("%d", input.Page)
	}
	if input.Limit > 0 {
		query["limit"] = fmt.Sprintf("%d", input.Limit)
	}
	if input.TripClass != -1 {
		query["trip_class"] = fmt.Sprintf("%d", input.TripClass)
	}
	err = a.getJson("v2/prices/latest", query, &prices)
	return
}
