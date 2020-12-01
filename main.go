package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type (
	// User ...
	User struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

var (
	users = map[int]*User{}
	seq   = 1
)

//----------
// Handlers
//----------

// CreateUser ...
func CreateUser(c echo.Context) error {
	u := &User{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	users[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

// GetUser ...
func GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, users[id])
}

// UpdateUser ...
func UpdateUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	users[id].Name = u.Name
	return c.JSON(http.StatusOK, users[id])
}

// DeleteUser ...
func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}

// GetAllUsers ...
func GetAllUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func main() {
	e := echo.New()

	// Routes
	e.GET("/users", GetAllUsers)
	e.POST("/users", CreateUser)
	e.GET("/users/:id", GetUser)
	e.PUT("/users/:id", UpdateUser)
	e.DELETE("/users/:id", DeleteUser)

	// Start server
	e.Logger.Fatal(e.Start(":9090"))
}
