package main

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

var drinks = []Drinks{
	{
		ID: 0, Brand: "The Coca-Cola Company", Name: "Coke", Price: 56.99,
	},
	{
		ID: 1, Brand: "Sprite", Name: "Sprite", Price: 17.99,
	},
	{
		ID: 2, Brand: "A&W", Name: "Root Beer", Price: 39.99,
	},
}

func getdrinks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, drinks)
}

func postdrinks(c *gin.Context) {
	var newDrinks Drinks

	if err := c.BindJSON(&newDrinks); err != nil {
		return
	}

	drinks = append(drinks, newDrinks)
	c.IndentedJSON(http.StatusCreated, newDrinks)
}

func getDrinksByID(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)

	for _, a := range drinks {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Drinks not found"})
}

func main() {
	router := gin.Default()
	router.GET("/drinks", getdrinks)
	router.GET("/drinks/:id", getDrinksByID)
	router.POST("/drinks", postdrinks)

	router.Run("localhost:8080")
}
