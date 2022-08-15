package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jaloldinov/IMAN-Updated/post_service/genproto/second_service"
	"github.com/jaloldinov/IMAN-Updated/post_service/pkg/helper"
	"github.com/jaloldinov/IMAN-Updated/post_service/storage"
)

type postRepo struct {
	db *pgxpool.Pool
}

func NewPostRepo(db *pgxpool.Pool) storage.PostI {
	return &postRepo{
		db: db,
	}
}

// GetAll returns all posts with given offset and limit
func (r *postRepo) ListPosts(limit uint32, offset uint32, search string) (*second_service.ListPostsResponse, error) {
	var (
		resp   second_service.ListPostsResponse
		err    error
		filter string
		params = make(map[string]interface{})
	)

	if len(search) > 0 {
		params["search"] = search
		filter += " AND ((title ILIKE '%' || :search || '%') OR (body ILIKE '%' || :search || '%'))"
	}
	countQuery := `SELECT count(1) FROM posts WHERE deleted_at is NULL and true ` + filter

	q, arr := helper.ReplaceQueryParams(countQuery, params)
	err = r.db.QueryRow(context.Background(), q, arr...).Scan(
		&resp.Count,
	)

	if err != nil {
		return nil, fmt.Errorf("error while scanning count %w", err)
	}

	query := `SELECT
		id,
		user_id,
		title,
		body
	FROM posts
	WHERE deleted_at is NULL and true` + filter

	query += " LIMIT :limit OFFSET :offset"
	params["limit"] = limit
	params["offset"] = offset

	q, arr = helper.ReplaceQueryParams(query, params)
	rows, err := r.db.Query(context.Background(), q, arr...)
	if err != nil {
		return nil, fmt.Errorf("error while getting rows %w", err)
	}

	for rows.Next() {
		var post second_service.Post

		err = rows.Scan(
			&post.PostId,
			&post.UserId,
			&post.Title,
			&post.Body,
		)

		if err != nil {
			return nil, fmt.Errorf("error  scanning row %w", err)
		}

		resp.Results = append(resp.Results, &post)
	}

	return &resp, nil
}

// GetPost returns post by id
func (r *postRepo) GetPost(id int) (*second_service.Post, error) {
	var (
		post second_service.Post
		err  error
	)

	query := `SELECT
		id,
		user_id,
		title,
		body
	FROM posts
	WHERE id = $1 AND deleted_at is NULL`

	err = r.db.QueryRow(context.Background(), query, id).Scan(
		&post.PostId,
		&post.UserId,
		&post.Title,
		&post.Body,
	)

	if err != nil {
		return nil, fmt.Errorf("not found %w", err)
	}

	return &post, nil
}

// UpdatePost updates post by id
func (r *postRepo) UpdatePost(req *second_service.UpdatePostRequest) (result *second_service.Result, err error) {
	query := `UPDATE posts SET user_id = $1, title = $2, body = $3, updated_at = now() WHERE deleted_at is NULL AND id = $4`

	res, err := r.db.Exec(context.Background(), query, req.UserId, req.Title, req.Body, req.PostId)

	if err != nil {
		return &second_service.Result{Message: "NOT UPDATED"}, fmt.Errorf("error while updating post %w", err)
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return &second_service.Result{Message: "post id not found"}, fmt.Errorf("post not found")
	}

	return &second_service.Result{
		Message: "Updated",
	}, nil
}

// DeletePost deletes port given id
func (r *postRepo) DeletePost(postId int) (*second_service.Result, error) {
	query := `
	UPDATE posts SET deleted_at = NOW()
	WHERE deleted_at is NULL AND  id = $1
	`
	res, err := r.db.Exec(context.Background(), query, postId)
	if err != nil {
		return &second_service.Result{Message: "NOT DELETED"}, fmt.Errorf("error while deleting post %w", err)
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return &second_service.Result{Message: "post id not found"}, fmt.Errorf("post not found")
	}

	return &second_service.Result{
		Message: "Deleted",
	}, nil
}
