package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Define data schema
type user struct {
	ID             string `json:"id"`
	First_Name     string `json:"first_name"`
	Last_Name      string `json:"last_name"`
	Preffered_Name string `json:"nick_name"`
	Email_Address  string `json:"email_address"`
}

// seeding mock data
var users = []user{
	{ID: "1", First_Name: "James", Last_Name: "Smith", Preffered_Name: "", Email_Address: "james_smith@gmail.com"},
	{ID: "2", First_Name: "Jocelyn", Last_Name: "Spencer", Preffered_Name: "Joyce", Email_Address: "jc_spencer@gmail.com"},
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func createUser(c *gin.Context) {
	var newUser user

	// Context.BindJSON binds the request body to newUser
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	// appending the newUser to the users slices(in-memory)
	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func getUserByID(c *gin.Context) {
	id := c.Param("id")

	// using underscore to ignore the index value and not trip the linter
	for _, userInstance := range users {
		if userInstance.ID == id {
			c.IndentedJSON(http.StatusOK, userInstance)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func updateUser(c *gin.Context) {
	var updatedUser user
	if err := c.BindJSON(&updatedUser); err != nil {
		return
	}
	id := c.Param("id")

	// using underscore to ignore the index value and not trip the linter
	for index, userInstance := range users {
		if userInstance.ID == id {
			users[index] = updatedUser
			c.IndentedJSON(http.StatusOK, updatedUser)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")

	// using underscore to ignore the index value and not trip the linter
	for index, userInstance := range users {
		if userInstance.ID == id {
			users = append(users[:index], users[index+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "user deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUserByID)
	router.POST("/users", createUser)
	router.PATCH("/users/:id", updateUser)
	router.DELETE("/users/:id", deleteUser)

	router.Run("localhost:8080")
}
