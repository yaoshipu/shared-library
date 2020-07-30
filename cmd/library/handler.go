package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handler
func ping(c echo.Context) error {
	return c.String(http.StatusOK, "Pong")
}

func myBooks(c echo.Context) error {
	mybooks := new(MyBooks)
	mybooks.Donated = []*Book{}
	mybooks.Borrowed = []*Book{}

	username := c.Param("username")
	borrowed, err := bookColl.List(&ListOption{Borrower: username})
	if err != nil {
		return c.JSON(http.StatusBadRequest, mybooks)
	}
	mybooks.Borrowed = borrowed

	donated, err := bookColl.List(&ListOption{Donator: username})
	if err != nil {
		return c.JSON(http.StatusBadRequest, mybooks)
	}
	mybooks.Donated = donated

	return c.JSON(http.StatusOK, mybooks)
}

func listBooks(c echo.Context) error {
	opt := ListOption{}

	list, err := bookColl.List(&opt)

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
	if err := bookColl.Create(newBook); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, newBook)
}

func borrowBook(c echo.Context) error {
	opt := new(BookOpt)
	if err := c.Bind(opt); err != nil {
		return err
	}
	if err := bookColl.Borrow(opt.ID, opt.Borrower); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, opt)
}

func returnBook(c echo.Context) error {
	opt := new(BookOpt)
	if err := c.Bind(opt); err != nil {
		return err
	}

	if err := bookColl.Return(opt.ID, opt.Borrower); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, opt)
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
