package model

import (
	"log"
	"math"
	"strings"
	"time"
)

type PostMeta struct {
	ID          int
	Subject     string
	Preview     string
	CreatedTime time.Time
	UpdatedTime time.Time
}

func PostMetaFromFile(file File, preview string) PostMeta {
	res := PostMeta{}
	ct := time.Time{}

	if preview == "" {
		log.Println("no preview")
	}
	if !file.CreatedTime.Equal(time.Time{}) {
		ct = file.CreatedTime
	}

	fns := strings.Split(file.Name, "/")
	lastIdx := int(math.Max(float64(len(fns)-1), 0))

	res.Subject = fns[lastIdx]
	res.Preview = preview
	res.CreatedTime = ct

	return res
}
