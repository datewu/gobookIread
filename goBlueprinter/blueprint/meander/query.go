package meander

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// APIKey Google Places API
var APIKey string

var proxyURL, _ = url.Parse("https://192.168.2.112:7070")

// Place lol
type Place struct {
	*googleGeometry `json:"geometry"`
	Name            string         `json:"name"`
	Icon            string         `json:"icon"`
	Photos          []*googlePhoto `json:"photos"`
	Vicinity        string         `json:"vicinity"`
}

type googleResponse struct {
	Results []*Place `json:"results"`
}

type googleGeometry struct {
	*googleLocation `json:"location"`
}

type googleLocation struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type googlePhoto struct {
	PhotoRef string `json:"photo_reference"`
	URL      string `json:"url"`
}

// Public satify Facade interface
func (p *Place) Public() interface{} {
	return map[string]interface{}{
		"name":     p.Name,
		"icon":     p.Icon,
		"photos":   p.Photos,
		"vicinity": p.Vicinity,
		"lat":      p.Lat,
		"lng":      p.Lng,
	}
}

// Query interupt with google API service
type Query struct {
	Lat          float64
	Lng          float64
	Journey      []string
	Radius       int
	CostRangeStr string
}

func (q *Query) find(types string) (*googleResponse, error) {
	vals := make(url.Values)
	vals.Set("location", fmt.Sprintf("%g,%g", q.Lat, q.Lng))
	vals.Set("radius", fmt.Sprintf("%d", q.Radius))
	vals.Set("types", types)
	vals.Set("key", APIKey)
	if len(q.CostRangeStr) > 0 {
		r, err := ParseCostRange(q.CostRangeStr)
		if err != nil {
			log.Println("wrong format", err)
			return nil, err
		}
		vals.Set("minprice", fmt.Sprintf("%d", int(r.From)-1))
		vals.Set("maxprice", fmt.Sprintf("%d", int(r.To)-1))
	}

	u := "https://maps.googleapis.com/maps/api/place/nearbysearch/json"
	req, err := http.NewRequest("GET", u+"?"+vals.Encode(), nil)
	if err != nil {
		log.Println("create request", err)
		return nil, err
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Println("Get", err)
		return nil, err
	}
	defer resp.Body.Close()
	var response googleResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Println("decoding", err)
		return nil, err
	}
	return &response, nil
}

// Run runs the query concurrently, and returns the results.
func (q *Query) Run() []interface{} {
	rand.Seed(time.Now().UnixNano())
	var w sync.WaitGroup
	var mux sync.Mutex
	places := make([]interface{}, len(q.Journey))
	w.Add(len(q.Journey))
	for i, t := range q.Journey {
		go func(types string, i int) {
			defer w.Done()
			response, err := q.find(types)
			if err != nil {
				log.Println("Failed to find places:", err)
				return
			}
			if len(response.Results) == 0 {
				log.Println("No places found for", types)
				return
			}
			for _, result := range response.Results {
				for _, photo := range result.Photos {
					photo.URL =
						"https://maps.googleapis.com/maps/api/place/photo?" +
							"maxwidth=1000&photoreference=" + photo.PhotoRef + "&key=" +
							APIKey
				}
			}
			randI := rand.Intn(len(response.Results))
			mux.Lock()
			places[i] = response.Results[randI]
			mux.Unlock()
		}(t, i)
	}
	w.Wait()
	return places
}
