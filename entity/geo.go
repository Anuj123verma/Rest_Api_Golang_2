package entity

// entity that made on the basis of consumed api data
type Geo struct {
	Items []struct {
		Title           string `json:"title"`
		ID              string `json:"id"`
		ResultType      string `json:"resultType"`
		HouseNumberType string `json:"houseNumberType"`
		Address         struct {
			Label       string `json:"label"`
			CountryCode string `json:"countryCode"`
			CountryName string `json:"countryName"`
			State       string `json:"state"`
			CountyCode  string `json:"countyCode"`
			County      string `json:"county"`
			City        string `json:"city"`
			District    string `json:"district"`
			Street      string `json:"street"`
			PostalCode  string `json:"postalCode"`
			HouseNumber string `json:"houseNumber"`
		} `json:"address"`
		Position struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"position"`
		Access []struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"access"`
		Distance int `json:"distance"`
		MapView  struct {
			West  float64 `json:"west"`
			South float64 `json:"south"`
			East  float64 `json:"east"`
			North float64 `json:"north"`
		} `json:"mapView"`
	} `json:"items"`
}
