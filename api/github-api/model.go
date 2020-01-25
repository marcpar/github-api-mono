package main

type Comment struct {
	Id         int `gorm:"primary_key"`
	OrgId      string
	Githubuser string
	Comment    string
	DeletedAt  bool
}

type Comments []Comment
