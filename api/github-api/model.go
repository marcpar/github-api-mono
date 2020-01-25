package main

import "time"

type Comment struct {
	Id         int `gorm:"primary_key"`
	OrgId      string
	Githubuser string
	Comment    string
	DeletedAt  *time.Time
	CreatedAt  time.Time
}

type Comments []Comment

func (a *App) IntializeModel() {
	db := a.DB
	db.AutoMigrate(&Comment{})
}
