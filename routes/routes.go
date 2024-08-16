package routes

import (
	"album-tracker/db"
	"album-tracker/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	db.Init() // Inicia o banco de dados

	r.HandleFunc("/albums/register", handlers.RegisterAlbum).Methods("POST")
	r.HandleFunc("/albums", handlers.FindAlbums).Methods("GET")
	r.HandleFunc("/albums/{name}", handlers.FindAlbum).Methods("GET")
	r.HandleFunc("/albums/{name}", handlers.DeleteAlbum).Methods("DELETE")
	r.HandleFunc("/albums/artist/{artist}", handlers.FindForArtist).Methods("GET")
	r.HandleFunc("/albums/{name}", handlers.UpdateAlbum).Methods("PUT")
	r.HandleFunc("/albums/score/{score}", handlers.FindForScore).Methods("GET")
	r.HandleFunc("/albums/genre/{genre}", handlers.FindForGenre).Methods("GET")
	r.HandleFunc("/albums/liked/{bool}", handlers.FindLiked).Methods("GET")
	r.HandleFunc("/albums/played/{bool}", handlers.FindPlayed).Methods("GET")

	return r
}
