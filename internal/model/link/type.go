package link

import (
	"time"
)

type (
	Link struct {
		ID          string    `bson:"_id"`
		OriginalURL string    `bson:"original_url"`
		Visits      int32     `bson:"visits"`
		ExpireAt    time.Time `bson:"expire_at"`
		CreatedAt   time.Time `bson:"created_at"`
	}
)
