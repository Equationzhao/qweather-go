package geo

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// SearchCity
// GET https://geoapi.qweather.com/v2/city/lookup?parameters&key=key
func SearchCity(para *Para, key string) (*SearchResponse, error) {
	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", "https://geoapi.qweather.com/v2/city/lookup", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("key", key)
	q.Add("location", para.Location)
	q.Add("adm", para.Adm)
	q.Add("lang", para.Lang)
	q.Add("number", strconv.Itoa(int(para.Number)))
	q.Add("range", para.Range)

	req.URL.RawQuery = q.Encode()
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var response SearchResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// HitCity
// Get https://geoapi.qweather.com/v2/city/top?number=5&range=cn&key=key
func HitCity(para *Para, key string) (*HitResponse, error) {
	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", "https://geoapi.qweather.com/v2/city/top", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("key", key)
	q.Add("number", strconv.Itoa(int(para.Number)))
	q.Add("range", para.Range)

	req.URL.RawQuery = q.Encode()
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var response HitResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// POI
// Get https://geoapi.qweather.com/v2/poi/lookup?type=scenic&location=jings&key=key
func POI(para *Para, key string) (*POIResponse, error) {
	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", "https://geoapi.qweather.com/v2/poi/lookup", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("key", key)
	q.Add("type", para.Type.String())
	q.Add("location", para.Location)
	q.Add("city", para.City)
	q.Add("number", strconv.Itoa(int(para.Number)))
	q.Add("lang", para.Lang)

	req.URL.RawQuery = q.Encode()
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var response POIResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// POIRange
// https://geoapi.qweather.com/v2/poi/range?location=116.40528,39.90498&type=scenic&radius=10&key=key
func POIRange(para *Para, key string) (*POIResponse, error) {
	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", "https://geoapi.qweather.com/v2/poi/lookup", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("key", key)
	q.Add("type", para.Type.String())
	q.Add("location", para.Location)
	q.Add("radius", strconv.Itoa(int(para.Radius)))
	q.Add("number", strconv.Itoa(int(para.Number)))
	q.Add("lang", para.Lang)

	req.URL.RawQuery = q.Encode()
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var response POIResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
