package post

import (
	"github.com/ilsan-kim/private-blog/worker/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testDBConf = config.DBConfig{
	Host:     "localhost",
	User:     "postgres",
	Password: "postgres",
	DB:       "blog_test",
}
var pgRepo *PGRepository

func setup() {
	pgRepo = NewPGRepository(testDBConf)

	pgRepo.db.Exec(`
		create table posts (
			id          bigserial	primary key,
			subject     varchar(255),
			preview     varchar(255),
			thumbnail   varchar(255),
			file_path   varchar(255),
			inserted_at timestamp(0) not null,
			updated_at  timestamp(0) not null
		);
	`)
}

func teardown() {
	pgRepo.db.Exec("truncate posts")
	defer pgRepo.db.Close()
}

func TestNewPGRepository(t *testing.T) {
	// TODO 테스트 고도화
	setup()
	defer teardown()

	t.Run("insert / get / update / delete", func(t *testing.T) {
		meta := PostMeta{
			Subject:   "test",
			Preview:   "test",
			Thumbnail: "test",
			FilePath:  "test",
		}

		err := pgRepo.Insert(meta)
		assert.NoError(t, err)

		//read, err := pgRepo.Get(1)
		//assert.NoError(t, err)
		//assert.Equal(t, "test", read.Subject)
	})
}
