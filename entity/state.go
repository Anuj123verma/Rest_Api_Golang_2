package entity

// state entity that we want to show as final output
type State struct {
	State     string `json:"state" bson:"state"`
	Confirmed int    `json:"confirmed" bson:"confirmed cases"`
	Recovered int    `json:"recovered" bson:"recovered cases"`
	Deaths    int    `json:"deaths" bson:"deaths"`
	Active    int    `json:"active" bson:"active cases"`
}
