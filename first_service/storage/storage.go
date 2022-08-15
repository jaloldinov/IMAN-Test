package storage

import "github.com/jaloldinov/IMAN-Updated/first_service/genproto/first_service"

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
	CheckPosts() (*first_service.CheckPostsResponse, error)
}
