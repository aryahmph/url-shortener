package link

import (
	"context"
	hashCommon "github.com/aryahmph/url-shortener/common/hash"
	linkModel "github.com/aryahmph/url-shortener/internal/model/link"
	counterRepo "github.com/aryahmph/url-shortener/internal/repository/counter"
	linkRepo "github.com/aryahmph/url-shortener/internal/repository/link"
	"time"
)

type linkUsecase struct {
	linkRepository    linkRepo.Repository
	counterRepository counterRepo.Repository

	hashManager hashCommon.Hash
}

func NewLinkUsecase(
	linkRepository linkRepo.Repository,
	counterRepository counterRepo.Repository,
	hashManager hashCommon.Hash,
) *linkUsecase {
	return &linkUsecase{
		linkRepository:    linkRepository,
		counterRepository: counterRepository,
		hashManager:       hashManager,
	}
}

func (u linkUsecase) CreateLink(ctx context.Context, arg linkModel.Link) (string, error) {
	// Get counter value from zookeeper
	counter, err := u.counterRepository.GetCurrentCounter()
	if err != nil {
		return "", err
	}

	// Create link
	hash := u.hashManager.Hash(counter)
	now := time.Now()
	err = u.linkRepository.CreateLink(ctx, linkModel.Link{
		ID:          hash,
		OriginalURL: arg.OriginalURL,
		ExpireAt:    now.Add(365 * 24 * time.Hour),
		CreatedAt:   now,
	})
	return hash, err
}

func (u linkUsecase) GetLink(ctx context.Context, id string) (linkModel.Link, error) {
	link, err := u.linkRepository.GetLink(ctx, id)
	if err != nil {
		return linkModel.Link{}, err
	}
	return link, nil
}
