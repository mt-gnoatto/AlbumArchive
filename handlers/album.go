package handlers

import (
	"album-tracker/db"
	"album-tracker/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func RegisterAlbum(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	a := models.Album{}
	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		fmt.Println("server failed to handle", err)
		return
	}

	if a.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "The field name cannot be empty",
		})
		return
	}

	if a.Genre == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "The field genre cannot be empty",
		})
		return
	}

	_, err = db.DB.Exec("INSERT INTO Album (name, genre, score, liked, played) VALUES ($1, $2, $3, $4, $5)", a.Name, a.Genre, a.Score, a.Liked, a.Played)
	if err != nil {
		fmt.Println("server failed to handle", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// FindAlbum find in DB the registreded albuns
func FindAlbuns(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// Busca as informações na tabela do banco
	rows, err := db.DB.Query("SELECT * FROM Album")
	if err != nil {
		fmt.Println("server failed to handle", err)
		return
	}

	defer rows.Close()
	data := make([]models.Album, 0)

	// Scaneia as linhas da tabela e retorna os valores
	for rows.Next() {
		album := models.Album{}
		err := rows.Scan(&album.Name, &album.Genre, &album.Score, &album.Liked, &album.Played)
		if err != nil {
			fmt.Println("server failed to handle", err)
			return
		}

		data = append(data, album)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("server failed to handle", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func FindAlbum(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	nameParam := vars["name"]

	// Busca as informações na tabela do banco
	rows, err := db.DB.Query("SELECT name, genre, score, liked, played from album where name LIKE $1", nameParam)
	if err != nil {
		fmt.Println("server failed to handle", err)
		return
	}

	defer rows.Close()
	data := make([]models.Album, 0)

	// Scaneia as linhas da tabela e retorna os valores
	for rows.Next() {
		album := models.Album{}
		err := rows.Scan(&album.Name, &album.Genre, &album.Score, &album.Liked, &album.Played)
		if err != nil {
			fmt.Println("server failed to handle", err)
			return
		}

		data = append(data, album)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("server failed to handle", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

// FindForScore filtra a busca pela nota do album, informada pelo usuário
func FindForScore(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	scoreParam := vars["score"]

	score, err := strconv.ParseUint(scoreParam, 10, 64)
	if err != nil {
		http.Error(w, "Invalid score parameter", http.StatusBadRequest)
		return
	}

	rows, err := db.DB.Query("SELECT name, genre, score, liked, played FROM Album WHERE score = $1", score)
	if err != nil {
		fmt.Println("server failed to handle", err)
		return
	}

	defer rows.Close()
	data := make([]models.Album, 0)

	for rows.Next() {
		album := models.Album{}
		err := rows.Scan(&album.Name, &album.Genre, &album.Score, &album.Liked, &album.Played)
		if err != nil {
			fmt.Println("server failed to handle", err)
			return
		}

		data = append(data, album)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("server failed to handle", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func FindForGenre(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	genreParam := vars["genre"]

	rows, err := db.DB.Query("SELECT name, genre, score, liked, played FROM album WHERE genre = $1", genreParam)
	if err != nil {
		fmt.Println("server failed to handle", err)
		http.Error(w, "Failed to query the database", http.StatusInternalServerError)
		return
	}

	defer rows.Close()
	data := make([]models.Album, 0)

	for rows.Next() {
		album := models.Album{}
		err := rows.Scan(&album.Name, &album.Genre, &album.Score, &album.Liked, &album.Played)
		if err != nil {
			fmt.Println("server failed to handle", err)
			http.Error(w, "Failed to read query results", http.StatusInternalServerError)
			return
		}

		data = append(data, album)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("server failed to handle", err)
		http.Error(w, "Error during query iteration", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func FindLiked(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	liked := vars["bool"]

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.DB.Query("SELECT * FROM album where liked = $1", liked)
	if err != nil {
		fmt.Println("server failed to handle", err)
		return
	}

	defer rows.Close()
	data := make([]models.Album, 0)

	// Scaneia as linhas da tabela e retorna os valores
	for rows.Next() {
		album := models.Album{}
		err := rows.Scan(&album.Name, &album.Genre, &album.Score, &album.Liked, &album.Played)
		if err != nil {
			fmt.Println("server failed to handle", err)
			return
		}

		data = append(data, album)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("server failed to handle", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func FindPlayed(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	played := vars["bool"]

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.DB.Query("SELECT name, genre, score, liked, played FROM album WHERE played = $1", played)
	if err != nil {
		fmt.Println("server failed to handle", err)
		http.Error(w, "Failed to query the database", http.StatusInternalServerError)
		return
	}

	defer rows.Close()
	data := make([]models.Album, 0)

	for rows.Next() {
		album := models.Album{}
		err := rows.Scan(&album.Name, &album.Genre, &album.Score, &album.Liked, &album.Played)
		if err != nil {
			fmt.Println("server failed to handle", err)
			http.Error(w, "Failed to read query results", http.StatusInternalServerError)
			return
		}

		data = append(data, album)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("server failed to handle", err)
		http.Error(w, "Error during query iteration", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	nameParam := vars["name"]

	// Busca as informações na tabela do banco
	rows, err := db.DB.Query("delete from album where name = $1", nameParam)
	if err != nil {
		fmt.Println("server failed to handle", err)
		return
	}

	defer rows.Close()
	data := make([]models.Album, 0)

	// Scaneia as linhas da tabela e retorna os valores
	for rows.Next() {
		album := models.Album{}
		err := rows.Scan(&album.Name, &album.Genre, &album.Score, &album.Liked, &album.Played)
		if err != nil {
			fmt.Println("server failed to handle", err)
			return
		}

		data = append(data, album)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("server failed to handle", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func UpdateAlbum(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	nameParam := vars["name"]

	a := models.Album{}
	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		fmt.Println("server failed to handle", err)
		return
	}

	if a.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "The field name cannot be empty",
		})
		return
	}

	if a.Genre == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "The field genre cannot be empty",
		})
		return
	}

	result, err := db.DB.Exec("UPDATE Album SET name = $1, genre = $2, score = $3, liked = $4, played = $5 WHERE name = $6", a.Name, a.Genre, a.Score, a.Liked, a.Played, nameParam)
	if err != nil {
		fmt.Println("server failed to handle", err)
		return
	}

	// Verifica se o álbum foi encontrado e atualizado
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Failed to check the number of affected rows", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "No album found with the provided name", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
