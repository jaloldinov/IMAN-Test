package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
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
			return "failed insert data", err
		}
	}

	return "success insert data", nil
}
