package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)

	router.Run("localhost:8080")
}

// album represent data about a record album
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrone", Price: 56.98},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligant", Price: 18.99},
	{ID: "3", Title: "Sarah Vaughan and Friend", Artist: "Sarah Vaughan", Price: 12.99},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums add a data album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbums album

	// Call BindJSON to bind the received JSON to
	// newAlbum.

	if err := c.BindJSON(&newAlbums); err != nil {
		return
	}

	albums = append(albums, newAlbums)
	c.IndentedJSON(http.StatusCreated, newAlbums)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Album Not found"})
}
