package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Struktur data untuk feedback
type Feedback struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name         string             `bson:"name" json:"name"`
	Address      string             `bson:"address" json:"address"`
	Job          string             `bson:"job" json:"job"`
	Feedback     string             `bson:"feedback" json:"feedback"`
	NextMaterial string             `bson:"next_material" json:"next_material"`
	Rating       float64            `bson:"rating" json:"rating"`
}

// Struktur untuk respons sukses
type SuccessResponse struct {
	Message  string   `json:"message"`
	Feedback Feedback `json:"feedback"`
}
