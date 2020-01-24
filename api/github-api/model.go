package main

type comment struct {
	ID    int     json:"id"
	OrgId  string  json:"name"
	Comment strimg json:"comment"
}

func (c *comment) getComment(db *sql.DB) error{

}