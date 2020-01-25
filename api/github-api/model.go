package main

type comment struct {
	ID    int     json:"id"
	OrgId  string  json:"name"
	Githubuser string json:"github_user"
	Comment string json:"comment"
	DeletedAt bool json"deleted"
}

type comments  comment[]

func (c *comment) getComment(db *sql.DB) error{

}