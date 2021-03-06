package main

import (
	"errors"
	"fmt"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Book struct
type Book struct {
	// ID is auto generated by mgo, ready only
	ID         bson.ObjectId `bson:"_id"                json:"id"`
	Author     string        `bson:"author"             json:"author"`
	Name       string        `bson:"name"               json:"name"`
	Donator    string        `bson:"donator"            json:"donator"`
	UpdateTime int64         `bson:"update_time"        json:"update_time"`
	Borrower   string        `bson:"current_borrower"   json:"current_borrower"`
	Image      string        `bson:"image_url"          json:"image_url"`
}

// BookOpt ...
type BookOpt struct {
	ID       bson.ObjectId `json:"id"`
	Borrower string        `json:"borrower"`
}

// MyBooks struct
type MyBooks struct {
	Borrowed []*Book `bson:"borrowed"             json:"borrowed"`
	Donated  []*Book `bson:"donated"              json:"donated"`
}

// BookColl struct
type BookColl struct {
	coll *mgo.Collection
}

// NewBookColl constructor
func NewBookColl(cfg *MgoConfig) (*BookColl, error) {
	session := Open(cfg)

	err := session.Coll.EnsureIndex(mgo.Index{Key: []string{"name", "donator"}, Unique: true})
	if err != nil {
		return nil, fmt.Errorf("EnsureIndex error: %v", err)
	}

	return &BookColl{coll: session.Coll}, nil
}

// Create ...
func (c *BookColl) Create(book *Book) error {
	coll := FastCopyCollection(c.coll)
	defer CloseCollection(coll)

	if book == nil {
		return errors.New("nil Favorite args")
	}

	book.ID = bson.NewObjectId()
	book.UpdateTime = time.Now().Unix()

	return coll.Insert(book)
}

// Borrow ...
func (c *BookColl) Borrow(id bson.ObjectId, borrower string) error {
	coll := FastCopyCollection(c.coll)
	defer CloseCollection(coll)

	var resp *Book
	if err := coll.FindId(id).One(&resp); err != nil {
		return err
	}

	if resp.Borrower != "" {
		return errors.New("already borrowed")
	}

	change := bson.M{"$set": bson.M{
		"update_time":      time.Now().Unix(),
		"current_borrower": borrower,
	}}
	return coll.UpdateId(id, change)
}

// Return ...
func (c *BookColl) Return(id bson.ObjectId, borrower string) error {
	coll := FastCopyCollection(c.coll)
	defer CloseCollection(coll)

	var resp *Book
	if err := coll.FindId(id).One(&resp); err != nil {
		return err
	}

	if resp.Borrower != borrower {
		return errors.New("cannot return")
	}
	change := bson.M{"$set": bson.M{
		"update_time":      time.Now().Unix(),
		"current_borrower": "",
	}}
	return coll.UpdateId(id, change)
}

// ListOption ...
type ListOption struct {
	Name     string
	Author   string
	Borrower string
	Donator  string
}

// List ...
func (c *BookColl) List(opt *ListOption) ([]*Book, error) {
	coll := FastCopyCollection(c.coll)
	defer CloseCollection(coll)

	if opt == nil {
		return nil, errors.New("nil ListOption")
	}

	query := bson.M{}

	if len(opt.Name) != 0 {
		query["name"] = opt.Name
	}

	if len(opt.Author) != 0 {
		query["author"] = opt.Author
	}

	if len(opt.Donator) != 0 {
		query["donator"] = opt.Donator
	}

	if len(opt.Borrower) != 0 {
		query["current_borrower"] = opt.Borrower
	}

	var resp []*Book
	err := coll.Find(query).All(&resp)
	return resp, err
}
