package meander

import "strings"

type j struct {
	Name       string
	PlaceTypes []string
}

// Journeys lol
var Journeys = []interface{}{
	j{Name: "Romantic", PlaceTypes: []string{"park", "bar",
		"movie_theater", "restaurant", "florist", "taxi_stand"}},
	j{Name: "Shopping", PlaceTypes: []string{"department_store", "cafe",
		"clothing_store", "jewelry_store", "shoe_store"}},
	j{Name: "Night out", PlaceTypes: []string{"bar", "casino", "food",
		"bar", "night_club", "bar", "bar", "hospital"}},
	j{Name: "Culture", PlaceTypes: []string{"museum", "cafe", "cemetery",
		"library", "art_gallery"}},
	j{Name: "Pamper", PlaceTypes: []string{"hair_care", "beauty_salon",
		"cafe", "spa"}},
}

// Public satisfy Facade interface
func (j j) Public() interface{} {
	return map[string]interface{}{
		"name":    j.Name,
		"journey": strings.Join(j.PlaceTypes, "|"),
	}
}
