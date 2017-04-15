package aviasales

import (
	"compress/gzip"
	"encoding/json"
	"encoding/xml"
	"net/http"
	"net/url"
)

type AviasalesApi struct {
	token string
	log   LoggerInterface
}

type LoggerInterface interface {
	Debug(...interface{})
}

// NewAviasalesApi creates a new instance AviasalesApi.
func NewAviasalesApi(token string) *AviasalesApi {
	return &AviasalesApi{
		token: token,
	}
}

type Coordinates struct {
	Lon float64 `json:"lon" bson:"lon"`
	Lan float64 `json:"lat" bson:"lat"`
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
		if v == "" {
			continue
		}
		params.Add(k, v)
	}

	apiUrl.RawQuery = params.Encode()

	if a.log != nil {
		a.log.Debug("API Send: " + apiUrl.String())
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", apiUrl.String(), nil)
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.Header.Get("Content-Encoding") == "gzip" {
		// decompress data
		reader, err := gzip.NewReader(res.Body)
		if err != nil {
			return err
		}
		defer reader.Close()
		return json.NewDecoder(reader).Decode(&v)
	}
	return json.NewDecoder(res.Body).Decode(&v)
}

func (a *AviasalesApi) getXML(path string, args map[string]string, v interface{}) error {
	apiUrl, err := url.Parse("http://api.travelpayouts.com/" + path)
	if err != nil {
		return err
	}
	params := url.Values{}
	for k, v := range args {
		if v == "" {
			continue
		}
		params.Add(k, v)
	}

	apiUrl.RawQuery = params.Encode()

	if a.log != nil {
		a.log.Debug("API Send: " + apiUrl.String())
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", apiUrl.String(), nil)
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.Header.Get("Content-Encoding") == "gzip" {
		// decompress data
		reader, err := gzip.NewReader(res.Body)
		if err != nil {
			return err
		}

		defer reader.Close()
		return xml.NewDecoder(reader).Decode(&v)
	}
	return xml.NewDecoder(res.Body).Decode(&v)
}
