package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jaloldinov/IMAN-Updated/first_service/genproto/first_service"
	"github.com/jaloldinov/IMAN-Updated/first_service/storage"
	"github.com/jaloldinov/IMAN-Updated/first_service/storage/posts"
)

type firstRepo struct {
	db *pgxpool.Pool
}

func NewDataRepo(db *pgxpool.Pool) storage.DataI {
	return &firstRepo{
		db: db,
	}
}

func (r *firstRepo) InsertPosts() (string, error) {

	query := `
		INSERT INTO posts 
			(id, user_id, title, body, created_at)
		VALUES
			($1, $2, $3, $4, $5)
	`

	queryForError := `
		INSERT INTO check_post 
		 	(message, error, created_at)
		VALUES
			($1, $2, $3)
	`

	firsts := posts.GetPosts()
	postDatas := []posts.Data{}

	// pretty.Print(firsts)

	for _, first := range firsts {
		postDatas = append(postDatas, first.Data...)
	}

	for k := range postDatas {
		timeNow := time.Now()
		_, err := r.db.Exec(
			context.Background(),
			query,
			postDatas[k].ID,
			postDatas[k].UserID,
			postDatas[k].Title,
			postDatas[k].Body,
			timeNow,
		)
		if err != nil {

			r.db.Exec(
				context.Background(),
				queryForError,
				"FAILED SAVING DATA",
				"true",
				timeNow,
			)

			return "failed saving data", err
		}
	}

	r.db.Exec(
		context.Background(),
		queryForError,
		"SUCCESSFULLY SAVED DATA",
		"false",
		time.Now(),
	)

	return "successfully saved data", nil
}

func (r *firstRepo) CheckPosts() (*first_service.CheckPostsResponse, error) {

	var checking first_service.CheckPostsResponse

	query := `
	select message, error from check_post ORDER by created_at DESC LIMIT 1;
	`
	err := r.db.QueryRow(context.Background(), query).Scan(
		&checking.Message,
		&checking.Error,
	)

	if err != nil {
		return nil, fmt.Errorf("faild scaning row %w", err)
	}

	return &first_service.CheckPostsResponse{
		Message: checking.Message,
		Error:   checking.Error,
	}, nil
}
