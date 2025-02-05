package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID		string `json:"id"`
	Title 	string `json:"title"`
	Artist 	string `json:"artist"`
	Price 	float64 `json:"price"`
}

var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main(){
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", PostAlbums)
	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK, albums )
}

func PostAlbums(c *gin.Context){
	var NewAlbum album
	if err := c.BindJSON(&NewAlbum); err != nil {
		return
	}
	albums = append(albums, NewAlbum)
	c.IndentedJSON(http.StatusCreated, NewAlbum)
}
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
