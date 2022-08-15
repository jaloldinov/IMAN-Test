package posts

type Links struct {
	Previous string `json:"previous"`
	Current  string `json:"current"`
	Next     string `json:"next"`
}

type Pagination struct {
	Total int   `json:"total"`
	Pages int   `json:"pages"`
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
	Links Links `json:"links"`
}

type Data struct {
	ID     int    `db:"id" json:"id"`
	UserID int    `db:"user_id" json:"user_id"`
	Title  string `db:"title" json:"title"`
	Body   string `db:"body" json:"body"`
}

type Meta struct {
	Pagination Pagination `json:"pagination"`
	Data       []Data     `json:"data"`
}
