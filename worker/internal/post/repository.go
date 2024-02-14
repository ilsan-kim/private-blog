package post

import (
	"database/sql"
	"github.com/ilsan-kim/private-blog/worker/config"
	"github.com/ilsan-kim/private-blog/worker/internal/model"
	_ "github.com/jackc/pgx/stdlib"
	"log"
)

type PGRepository struct {
	db *sql.DB
}

func NewPGRepository(dbConfig config.DBConfig) *PGRepository {
	log.Println(dbConfig.ConnectionString())
	db, err := sql.Open("pgx", dbConfig.ConnectionString())
	if err != nil {
		log.Panicf("can not connect to postgresql %s, err %v", dbConfig.ConnectionString(), err)
		return nil
	}
	return &PGRepository{db: db}
}

func (r PGRepository) Insert(data model.PostMeta) error {
	_, err := r.db.Exec("insert into posts (subject, preview, thumbnail, file_path, inserted_at, updated_at) values ($1, $2, $3, $4, now(), now())",
		data.Subject, data.Preview, data.Thumbnail, data.FilePath)
	return err
}

func (r PGRepository) Update(data model.PostMeta) error {
	_, err := r.db.Exec("update posts set subject = $1, preview = $2, thumbnail = $3, file_path = $4, updated_at = now() where id = $5",
		data.Subject, data.Preview, data.Thumbnail, data.FilePath, data.ID)
	return err
}

func (r PGRepository) Delete(id int) error {
	_, err := r.db.Exec("delete from posts where id = $1", id)
	return err
}

func (r PGRepository) Get(id int) (model.PostMeta, error) {
	res := model.PostMeta{}
	err := r.db.QueryRow("select id, subject, preview, thumbnail, file_path, inserted_at, updated_at from posts where id = $1", id).Scan(
		&res.ID, &res.Subject, &res.Preview, &res.Thumbnail, &res.FilePath, &res.CreatedTime, &res.UpdatedTime)
	return res, err
}

func (r PGRepository) GetAll() ([]model.PostMeta, error) {
	var res []model.PostMeta
	rows, err := r.db.Query("select id, subject, preview, thumbnail, file_path, inserted_at, updated_at from posts")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		r := model.PostMeta{}
		err = rows.Scan(&r.ID, &r.Subject, &r.Preview, &r.Thumbnail, &r.FilePath, &r.CreatedTime, &r.UpdatedTime)
		if err != nil {
			return nil, err
		}

		res = append(res, r)
	}

	return res, nil
}
