package routes

import (
	"github.com/Berkemah-dev/be-quisioner/controller"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/feedbacks", controller.CreateFeedback).Methods("POST")
	router.HandleFunc("/feedbacks", controller.GetFeedbacks).Methods("GET")
	// Tambahkan rute untuk update dan delete
	return router
}
