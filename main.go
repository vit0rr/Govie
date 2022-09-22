package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type movie struct {
	ID    string  `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

var movies = []movie{
	{ID: "1", Title: "Orphan", Price: 56.99},
	{ID: "2", Title: "Avatar", Price: 17.99},
	{ID: "3", Title: "Spider-Man", Price: 39.99},
}

func getMovies(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, movies)
}

func getMovieByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range movies {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "movie not found"})
}

func postMovies(c *gin.Context) {
	var newMovie movie

	if err := c.BindJSON(&newMovie); err != nil {
		return
	}

	movies = append(movies, newMovie)
	c.IndentedJSON(http.StatusCreated, newMovie)
}

func delMovie(c *gin.Context) {
	id := c.Param("id")

	for i, a := range movies {
		if a.ID == id {
			movies = append(movies[:i], movies[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "movie deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "movie not found"})
}

func main() {
	router := gin.Default()
	router.GET("/movies", getMovies)
	router.GET("/movies/:id", getMovieByID)
	router.POST("/movies", postMovies)
	router.DELETE("/movies/:id", delMovie)

	router.Run("localhost:8080")
}
