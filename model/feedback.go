package model

type Feedback struct {
	ID           string  `json:"id" bson:"id"`
	Name         string  `json:"name" bson:"name"`
	Address      string  `json:"address" bson:"address"`
	Job          string  `json:"job" bson:"job"`
	Feedback     string  `json:"feedback" bson:"feedback"`
	NextMaterial string  `json:"next_material" bson:"next_material"`
	Rating       float64 `json:"rating" bson:"rating"`
}
