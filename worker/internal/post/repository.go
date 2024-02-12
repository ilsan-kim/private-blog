package post

import (
	"database/sql"
	"github.com/ilsan-kim/private-blog/worker/config"
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

func (r PGRepository) Insert(data PostMeta) error {
	//TODO implement me
	panic("implement me")
}

func (r PGRepository) Update(data PostMeta) error {
	//TODO implement me
	panic("implement me")
}

func (r PGRepository) Delete(pkey int) error {
	//TODO implement me
	panic("implement me")
}

func (r PGRepository) Get(pkey int) PostMeta {
	//TODO implement me
	panic("implement me")
}

func (r PGRepository) GetAll() []PostMeta {
	//TODO implement me
	panic("implement me")
}
