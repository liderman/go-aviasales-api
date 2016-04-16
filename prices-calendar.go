package aviasales

import "fmt"

type InputPricesCalendar struct {
	Origin       string
	Destination  string
	DepartDate   string
	ReturnDate   string
	CalendarType string
	TripDuration int
	Currency     string
}

type DataFlightCalendar struct {
	Success bool                      `json:"success"`
	Data    map[string]FlightCalendar `json:"data"`
}

type FlightCalendar struct {
	Price        int    `json:"price"`
	Airline      string `json:"airline"`
	FlightNumber int    `json:"flight_number"`
	DepartureAt  string `json:"departure_at"`
	ReturnAt     string `json:"return_at"`
	ExpiresAt    string `json:"expires_at"`
	Origin       string `json:"origin"`
	Destination  string `json:"destination"`
	Transfers    int    `json:"transfers"`
}

// PricesCalendar returns the cheapest ticket (without transfer, with one or two changes) for the specified direction for each day of the selected month.
//
// Documentation:
// https://support.travelpayouts.com/hc/ru/articles/203956163#07
//
// Example URI:
// http://api.travelpayouts.com/v1/prices/calendar?depart_date=2016-11&origin=MOW&destination=BCN&calendar_type=departure_date&token=YOUR_TOKEN
func (a *AviasalesApi) PricesCalendar(input InputPricesCalendar) (flights DataFlightCalendar, err error) {
	query := map[string]string{
		"currency":      input.Currency,
		"origin":        input.Origin,
		"destination":   input.Destination,
		"depart_date":   input.DepartDate,
		"return_date":   input.ReturnDate,
		"calendar_type": input.CalendarType,
		"token":         a.token,
	}
	if input.TripDuration > 0 {
		query["trip_duration"] = fmt.Sprintf("%d", input.TripDuration)
	}
	err = a.getJson("v1/prices/calendar", query, &flights)
	return
}
