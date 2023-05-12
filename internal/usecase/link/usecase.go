package link

import (
	"context"
	linkModel "github.com/aryahmph/url-shortener/internal/model/link"
)

type Usecase interface {
	CreateLink(ctx context.Context, arg linkModel.Link) (string, error)
	GetLink(ctx context.Context, id string) (linkModel.Link, error)
}
