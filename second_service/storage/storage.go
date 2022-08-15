package storage

import "github.com/jaloldinov/IMAN-Updated/post_service/genproto/second_service"

type StorageI interface {
	Post() PostI
}

type PostI interface {
	ListPosts(limit, page uint32, search string) (*second_service.ListPostsResponse, error)
	GetPost(postId int) (*second_service.Post, error)
	UpdatePost(*second_service.UpdatePostRequest) (*second_service.Result, error)
	DeletePost(postID int) (*second_service.Result, error)
}
