package aviasales

import (
	"encoding/xml"
)

type XMLSpecialOffers struct {
	XMLName xml.Name       `xml:"offers"`
	Offers  []SpecialOffer `xml:"offer"`
}

type SpecialOffer struct {
	Airline         string              `xml:"airline,attr" json:"airline" bson:"airline"`
	AirlineCode     string              `xml:"airline_code,attr" json:"airline_code" bson:"airline_code"`
	Title           string              `xml:"title,attr" json:"title" bson:"title"`
	Id              int                 `xml:"id,attr" json:"id" bson:"id"`
	Href            string              `xml:"href,attr" json:"href" bson:"href"`
	SaleDateBegin   int                 `xml:"sale_date_begin,attr" json:"sale_date_begin" bson:"sale_date_begin"`
	SaleDateEnd     int                 `xml:"sale_date_end,attr" json:"sale_date_end" bson:"sale_date_end"`
	FlightDateBegin int                 `xml:"flight_date_begin,attr" json:"flight_date_begin" bson:"flight_date_begin"`
	FlightDateEnd   int                 `xml:"flight_date_end,attr" json:"flight_date_end" bson:"flight_date_end"`
	Link            string              `xml:"link,attr" json:"link" bson:"link"`
	Description     string              `xml:"description" json:"description" bson:"description"`
	Conditions      string              `xml:"conditions" json:"conditions" bson:"conditions"`
	Routes          []SpecialOfferRoute `xml:"route" json:"routes" bson:"routes"`
}

type SpecialOfferRoute struct {
	FromIata       string `xml:"from_iata,attr" json:"from_iata" bson:"from_iata"`
	ToIata         string `xml:"to_iata,attr" json:"to_iata" bson:"to_iata"`
	FromName       string `xml:"from_name,attr" json:"from_name" bson:"from_name"`
	ToName         string `xml:"to_name,attr" json:"to_name" bson:"to_name"`
	Class          string `xml:"class,attr" json:"class" bson:"class"`
	OnewayPrice    string `xml:"oneway_price,attr" json:"oneway_price" bson:"oneway_price"`
	RoundtripPrice string `xml:"roundtrip_price,attr" json:"roundtrip_price" bson:"roundtrip_price"`
}

// SpecialOffers returns latest special offers from airlines.
//
// Documentation:
// https://support.travelpayouts.com/hc/ru/articles/203956163#18
//
// Example URI:
// http://api.travelpayouts.com/v2/prices/special-offers?token=YOUR_TOKEN
func (a *AviasalesApi) SpecialOffers() (specialOffers []SpecialOffer, err error) {
	data := XMLSpecialOffers{}
	err = a.getXML("v2/prices/special-offers", map[string]string{
		"token": a.token,
	}, &data)
	if err != nil {
		return
	}

	specialOffers = data.Offers
	return
}
