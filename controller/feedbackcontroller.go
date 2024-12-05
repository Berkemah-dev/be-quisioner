package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Berkemah-dev/be-quisioner/config"
	"github.com/Berkemah-dev/be-quisioner/model"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateFeedback(w http.ResponseWriter, r *http.Request) {
	var feedback model.Feedback
	if err := json.NewDecoder(r.Body).Decode(&feedback); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	collection := config.Mongoconn.Database("berkemah_quisioner").Collection("feedbacks")
	_, err := collection.InsertOne(context.TODO(), feedback)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(feedback)
}

func GetFeedbacks(w http.ResponseWriter, r *http.Request) {
	var feedbacks []model.Feedback
	collection := config.Mongoconn.Database("berkemah_quisioner").Collection("feedbacks")

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var feedback model.Feedback
		if err := cursor.Decode(&feedback); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		feedbacks = append(feedbacks, feedback)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(feedbacks)
}

func UpdateFeedback(w http.ResponseWriter, r *http.Request) {
	// Logika untuk memperbarui feedback berdasarkan ID
}

func DeleteFeedback(w http.ResponseWriter, r *http.Request) {
	// Logika untuk menghapus feedback berdasarkan ID
}
