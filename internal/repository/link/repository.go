package link

import (
	"context"
	linkModel "github.com/aryahmph/url-shortener/internal/model/link"
)

//go:generate mockery --name=Repository
type Repository interface {
	CreateLink(ctx context.Context, arg linkModel.Link) error
	GetLink(ctx context.Context, id string) (linkModel.Link, error)
}
