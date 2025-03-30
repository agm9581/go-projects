package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movies struct {
	Id       int       `json:"id"`
	Title    string    `json:"title"`
	Length   int       `json:"length"`
	Director *Director `json:"director"`
}
type Director struct {
	Name          string `json:"name"`
	Surname       string `json:"surname"`
	AverageRating int    `json:"averageRating"`
}

var director1 = Director{Name: "Antonio", Surname: "Gonzalez", AverageRating: 10}
var movieList []Movies = []Movies{{
	Id: 1, Title: "Short film", Length: 2, Director: &director1},
	{Id: 2, Title: "Long film", Length: 1000, Director: &director1},
	{Id: 3, Title: "Long film", Length: 1000, Director: &director1},
	{Id: 4, Title: "Long film", Length: 1000, Director: &director1}}

func getMoviesHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		fmt.Fprintf(w, strconv.Itoa(http.StatusBadRequest), "Method not allowed")
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movieList)

}

func getOneMovieHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieId := vars["id"]
	for _, movie := range movieList {
		if value, err := strconv.Atoi(movieId); err == nil && value == movie.Id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

	fmt.Fprintf(w, "Movie with id %s not found", movieId)

}

func deleteOneMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieId := vars["id"]
	for index, movie := range movieList {
		if value, err := strconv.Atoi(movieId); err == nil && value == movie.Id {
			w.Header().Set("Content-Type", "application/json")
			movieList = append(movieList[:index], movieList[index+1:]...)
			json.NewEncoder(w).Encode(movieList)
			return
		}
	}

	fmt.Fprintf(w, "Movie with id %s not found", movieId)

}

func createOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movies
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = rand.Intn(100000)
	movieList = append(movieList, movie)
	json.NewEncoder(w).Encode(movieList)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	updateMovieId := vars["id"]
	w.Header().Set("Content-Type", "application/json")
	var updateMovie Movies
	_ = json.NewDecoder(r.Body).Decode(&updateMovie)
	for index, movie := range movieList {
		if movieId, err := strconv.Atoi(updateMovieId); err == nil && movieId == movie.Id {
			updateMovie.Id = movieId
			movieList = append(movieList[:index], movieList[index+1:]...)
			movieList = append(movieList, updateMovie)
			json.NewEncoder(w).Encode(movieList)
			return
		}
	}
	fmt.Fprintf(w, "Movie with id %s not found", updateMovieId)
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", getMoviesHandler).Methods("GET")
	router.HandleFunc("/", createOneMovie).Methods("POST")
	router.HandleFunc("/{id}", getOneMovieHandler).Methods("GET")
	router.HandleFunc("/{id}", deleteOneMovie).Methods("DELETE")
	router.HandleFunc("/{id}", updateMovie).Methods("PUT")

	fmt.Println("Server started in port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))

}
