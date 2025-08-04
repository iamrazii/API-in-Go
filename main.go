package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// fake database :)
type Movie struct {
	Title string `json:"title"`
	ID    int    `json:"id"`
	Actor string `json:"author"`
	Genre string `json:"quantity"`
}

var movies = []Movie{
	{Title: "Oppenheimer", ID: 1, Actor: "Cillian Murphy", Genre: "History"},
	{Title: "The Dark Knight", ID: 2, Actor: "Christian Bale", Genre: "Action"},
	{Title: "Ford v Ferrari", ID: 3, Actor: "Christian Bale", Genre: "Sports"},
	{Title: "Man of Steel", ID: 4, Actor: "Henry Cavill", Genre: "Fiction"},
}

func GetMovies(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, movies)
}
func main() {
	router := gin.Default()
	router.GET("/getmovies", GetMovies)
	fmt.Println(movies)
	router.Run("localhost:5050")
}
