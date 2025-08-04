package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// fake database :)
type Movie struct {
	Title string `json:"title"`
	ID    int    `json:"id"`
	Actor string `json:"actor"`
	Genre string `json:"genre"`
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

func AddMovies(c *gin.Context) {
	var mymovie Movie

	// trying to parse request json into struct

	if err := c.BindJSON(&mymovie); err != nil {
		// Send back a 400 Bad Request with an error message

		c.JSON(http.StatusBadRequest, gin.H{ // returning a map
			"error":   "Invalid JSON data",
			"details": err.Error(),
		})
		return
	}
	movies = append(movies, mymovie)
	c.IndentedJSON(http.StatusAccepted, movies)
}

func main() {
	router := gin.Default()
	router.GET("/getmovies", GetMovies)
	router.POST("postmovies", AddMovies)
	router.Run("localhost:5050")
}
