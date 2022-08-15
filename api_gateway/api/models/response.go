package models

// ResponseModel ...
type ResponseModel struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

type Status struct {
	Status string `json:"status"`
}

// post models
type Post struct {
	ID     int64  `json:"id" db:"id"`
	UserId int64  `json:"user_id" db:"user_id"`
	Title  string `json:"title" db:"title"`
	Body   string `json:"body" db:"body"`
}

// Checking
type Check struct {
	Message string `json:"message" db:"message"`
	Error   string `json:"error" db:"error"`
}
