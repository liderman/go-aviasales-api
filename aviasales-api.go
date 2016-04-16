package aviasales

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type AviasalesApi struct {
	token string
	log   LoggerInterface
}

type LoggerInterface interface {
	Debug(interface{})
}

// NewAviasalesApi creates a new instance AviasalesApi.
func NewAviasalesApi(token string) *AviasalesApi {
	return &AviasalesApi{
		token: token,
	}
}

type Coordinates struct {
	Lon float64 `json:"lon"`
	Lan float64 `json:"lat"`
}

func (a *AviasalesApi) SetLogger(logger LoggerInterface) {
	a.log = logger
}

func (a *AviasalesApi) getJson(path string, args map[string]string, v interface{}) error {
	apiUrl, err := url.Parse("http://api.travelpayouts.com/" + path)
	if err != nil {
		return err
	}
	params := url.Values{}
	for k, v := range args {
		if args == "" {
			continue
		}
		params.Add(k, v)
	}

	apiUrl.RawQuery = params.Encode()

	a.log.Debug("API Send: " + apiUrl.String())
	res, err := http.Get(apiUrl.String())
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(&v)
}
