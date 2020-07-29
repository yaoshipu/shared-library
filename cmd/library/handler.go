package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handler
func ping(c echo.Context) error {
	return c.String(http.StatusOK, "Pong")
}

func listBooks(c echo.Context) error {

	opt := ListOption{}
	list, err := bookColl.List(&opt)

	fmt.Println(list)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, list)
}

func createBook(c echo.Context) error {
	newBook := new(Book)
	if err := c.Bind(newBook); err != nil {
		return err
	}
	fmt.Println(newBook.Name)
	if err := bookColl.Create(newBook); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, newBook)
}

// User ...
type User struct {
	Code int `bson:"code"	json:"code"`
	Data struct {
		Token string `bson:"token"     json:"token"`
	} `bson:"data"               json:"data"`
}

func userInfo(c echo.Context) error {
	var user User
	user.Code = 20000
	user.Data.Token = "Admin Token"

	return c.JSON(http.StatusOK, user)
}

func userLogin(c echo.Context) error {

	var user User
	user.Code = 20000
	user.Data.Token = "Admin Token"

	return c.JSON(http.StatusOK, user)
}
