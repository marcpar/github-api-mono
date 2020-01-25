package main

type Comment struct {
	Id         int `gorm:"primary_key"`
	OrgId      string
	Githubuser string
	Comment    string
	DeletedAt  bool
}

type Comments []Comment

func (a *App) intializeModel() {
	db := a.DB
	db.CreateTable(&Comment{})
	// Also some useful functions
	db.HasTable(&Comment{})          // =>;; true
	db.DropTableIfExists(&Comment{}) //Drops the table if already exists
}
