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

func myBooks(c echo.Context) error {

	username := c.Param("username")
	borrowed, err := bookColl.List(&ListOption{Name: username})
	if err != nil {
		return err
	}
	if borrowed == nil {
		borrowed = []*Book{}
	}

	donated, err := bookColl.List(&ListOption{Donator: username})
	if err != nil {
		return err
	}
	if donated == nil {
		donated = []*Book{}
	}

	mybooks := new(MyBooks)
	mybooks.Borrowed = borrowed
	mybooks.Donated = donated

	return c.JSON(http.StatusOK, mybooks)
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
