package aviasales

type Alliance struct {
	Name     string   `json:"name" bson:"name"`
	Airlines []string `json:"alias" bson:"alias"`
}

// DataAirlinesAlliances a list of alliances from the database.
//
// Documentation:
// https://support.travelpayouts.com/hc/ru/articles/203956163#13
//
// Example URI:
// http://api.travelpayouts.com/data/airlines_alliances.json
func (a *AviasalesApi) DataAirlinesAlliances() (airlinesAlliances []Alliance, err error) {
	err = a.getJson("data/airlines_alliances.json", map[string]string{}, &airlinesAlliances)
	return
}
