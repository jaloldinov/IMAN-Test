package storage

type Post struct {
	ID     int
	UserID int
	Title  string
	Body   string
}

type StorageI interface {
	Data() DataI
}

type DataI interface {
	InsertPosts() (string, error)
}
