package aviasales

type Plane struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// DataPlanes a list of planes from the database.
//
// Documentation:
// https://support.travelpayouts.com/hc/ru/articles/203956163#14
//
// Example URI:
// http://api.travelpayouts.com/data/planes.json
func (a *AviasalesApi) DataPlanes() (planes []Plane, err error) {
	err = a.getJson("data/planes.json", map[string]string{}, &planes)
	return
}
