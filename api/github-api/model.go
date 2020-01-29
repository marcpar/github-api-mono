package main

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type Comment struct {
	Id         int `gorm:"primary_key"`
	OrgId      string
	Githubuser string
	Comment    string
	DeletedAt  *time.Time
	CreatedAt  time.Time
}

type Comments []Comment

//Create automigrate this creates the Comment model without
func IntializeModel(db *gorm.DB) {

	db.AutoMigrate(&Comment{})
}

//Create comment
func (c *Comment) CreateComment(db *gorm.DB) error {
	return errors.New("Not implemented")
}

//delete comment softdelete
func (c *Comment) DeleteComment(db *gorm.DB) error {
	return errors.New("Not implemented")
}
