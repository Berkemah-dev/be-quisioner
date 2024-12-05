package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Berkemah-dev/be-quisioner/config"
	"github.com/Berkemah-dev/be-quisioner/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Fungsi untuk membuat feedback
func CreateFeedback(w http.ResponseWriter, r *http.Request) {
	var feedback model.Feedback
	if err := json.NewDecoder(r.Body).Decode(&feedback); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	collection := config.Mongoconn.Database("db_berkemah").Collection("feedbacks")
	feedback.ID = primitive.NewObjectID() // Menghasilkan ID baru
	_, err := collection.InsertOne(context.TODO(), feedback)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Membuat respons sukses
	response := model.SuccessResponse{
		Message:  "Feedback berhasil ditambahkan!",
		Feedback: feedback,
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Fungsi untuk mendapatkan semua feedback
func GetFeedbacks(w http.ResponseWriter, r *http.Request) {
	var feedbacks []model.Feedback
	collection := config.Mongoconn.Database("db_berkemah").Collection("feedbacks")

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
