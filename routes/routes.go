package routes

import (
	"album-tracker/db"
	"album-tracker/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	db.Init() // Inicia o banco de dados

	r.HandleFunc("/albuns/register", handlers.RegisterAlbum).Methods("POST")
	r.HandleFunc("/albuns", handlers.FindAlbuns).Methods("GET")
	r.HandleFunc("/albuns/{name}", handlers.FindAlbum).Methods("GET")
	r.HandleFunc("/albuns/{name}", handlers.DeleteAlbum).Methods("DELETE")
	r.HandleFunc("/albuns/{name}", handlers.UpdateAlbum).Methods("PUT")
	r.HandleFunc("/albuns/score/{score}", handlers.FindForScore).Methods("GET")
	r.HandleFunc("/albuns/genre/{genre}", handlers.FindForGenre).Methods("GET")
	r.HandleFunc("/albuns/liked/{bool}", handlers.FindLiked).Methods("GET")
	r.HandleFunc("/albuns/played/{bool}", handlers.FindPlayed).Methods("GET")

	return r
}
