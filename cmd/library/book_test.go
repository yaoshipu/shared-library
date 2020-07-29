package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
)

func TestBookCRUD(t *testing.T) {
	assert := assert.New(t)
	var err error

	cfg := MgoConfig{Host: "40.112.174.85:27017", DB: "sharedlibrary_test", Coll: "book", Mode: "strong"}

	bookColl, err := NewBookColl(&cfg)
	if err != nil {
		assert.Fail(err.Error())
	}

	bookColl.coll.RemoveAll(bson.M{})
	// defer func() {
	// 	bookColl.coll.RemoveAll(bson.M{})
	// }()

	err = bookColl.Create(nil)
	assert.NotNil(err, "should not nil")

	book1 := &Book{Name: "Site Reliability Engineering", Author: "Chris Jones", Donator: "Spark", Image: "https://images-na.ssl-images-amazon.com/images/I/51XswOmuLqL._SX379_BO1,204,203,200_.jpg"}
	err = bookColl.Create(book1)
	assert.Nil(err)

	book2 := &Book{Name: "Effeictive Java", Author: "Joshua Bloch", Donator: "Steven", Image: "https://images-na.ssl-images-amazon.com/images/I/41JLgmt8MlL._SX402_BO1,204,203,200_.jpg"}
	err = bookColl.Create(book2)
	assert.Nil(err)

	opt := ListOption{}
	list, err := bookColl.List(&opt)
	assert.Nil(err)
	assert.Len(list, 2)

}
