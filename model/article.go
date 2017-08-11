package model

// Article domain object
type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}

// Articles - is the list of domain objects
type Articles []Article
