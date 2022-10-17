package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	// "math/rand"
	// "encoding/json"
)


func main(){
router := mux.NewRouter()

movies= append(movies,Movie{ID: "1", Isbn: "43448434653", Title: "True Grit", Director: &Director{Firstname: "Ethan", Lastname: " Coen"}}, )
movies= append(movies,Movie{ID: "3", Isbn: "434484344", Title: "Pulp Fiction", Director: &Director{Firstname: "Quentin", Lastname: "Tarantino"}},)
movies= append(movies,Movie{ID: "2", Isbn: "434448464", Title: "The Guard", Director: &Director{Firstname: "John", Lastname: " Michael"}},)
movies= append(movies,Movie{ID: "4", Isbn: "434998464", Title: "Gravity", Director: &Director{Firstname: "Alfonso", Lastname: "cauron"}},)
movies= append(movies,Movie{ID: "5", Isbn: "4344390464", Title: "Snowpiercer", Director: &Director{Firstname: "Joon-ho", Lastname: "Bong"}})
// handler functions

router.HandleFunc("/text", homepage).Methods("GET")
router.HandleFunc("/movies", getMovies).Methods("GET")
router.HandleFunc("/movie/{id}", getMovie).Methods("GET")
router.HandleFunc("/movies", createMovie).Methods("POST")
router.HandleFunc("/movie/{id}", updateMovie).Methods("PUT")
router.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE")
fmt.Printf("start at server at port :8080\n")
log.Fatal(http.ListenAndServe(":8080", router))

}



// creating movies struct
type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
	
	
}
// creating Director struct
type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}
// slice of movies to store the list of movies since we are not using database
 var movies  []Movie



// fmt.Println(movies)
// creating handler functions
func handlerRouter(){
	
}

// router funcutions
func homepage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello server")
}
//   getMovies Endpoint
func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

//  delete movie endpoint
func deleteMovie(w http.ResponseWriter, r *http.Request){
		w.Header().Set("content-Type", "application/json")
		param:= mux.Vars(r)
		for index,item := range movies{
			if item.ID == param["id"]{
				movies= append(movies[:index], movies[index+1:]...)
				break
			}
		}
		json.NewEncoder(w).Encode(movies)
		
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-Type", "application/json")
		param:= mux.Vars(r)
		for _,item := range movies{
			if item.ID == param["id"]{
				json.NewEncoder(w).Encode(item)
				return
			}
		}
}


func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-Type", "application/json")
	var movie Movie
	_= json.NewDecoder(r.Body).Decode(&movie)
		movie.ID= strconv.Itoa(rand.Intn(10000000))
	movies =append(movies, movie)
	json.NewEncoder(w).Encode(movie)	
}


func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-Type", "application/json")
	param:= mux.Vars(r)
	for index, item:= range movies{
		if item.ID==param["id"]{
			movies= append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID=param["id"]
			movies =append(movies, movie)
			json.NewEncoder(w).Encode(movie)	
			return
		}
	}

}



