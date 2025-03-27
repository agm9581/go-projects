package main

import (
	"encoding/json"
	"fmt"
	"log"
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
	Id: 1, Title: "Short film", Length: 2, Director: &director1,
}, {Id: 2, Title: "Long film", Length: 1000, Director: &director1}}

func getMoviesHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		fmt.Fprintf(w, strconv.Itoa(http.StatusBadRequest), "Method not allowed")
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movieList)
}
func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", getMoviesHandler)

	fmt.Println("Server started in port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))

}
