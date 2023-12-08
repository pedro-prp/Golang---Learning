package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Define a struct to hold album data.
type album struct {
	ID		 string `json:"id"`
	Title		 string `json:"title"`
	Artist		 string `json:"artist"`
	Price		 float64 `json:"price"`
}

// define varible to hold all albums
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// function to return a list of all albums in the collection

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}


func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// loop in albums for match
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJson to bind the receiced JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}