package mongodb

import (
	"context"
	errorCommon "github.com/aryahmph/url-shortener/common/error"
	linkModel "github.com/aryahmph/url-shortener/internal/model/link"
	linkRepo "github.com/aryahmph/url-shortener/internal/repository/link"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type mongodbLinkRepository struct {
	collection *mongo.Collection
}

func NewMongodbLinkRepository(collection *mongo.Collection) linkRepo.Repository {
	return &mongodbLinkRepository{collection: collection}
}

func (r mongodbLinkRepository) CreateLink(ctx context.Context, arg linkModel.Link) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	_, err := r.collection.InsertOne(ctx, linkModel.Link{
		ID:          arg.ID,
		OriginalURL: arg.OriginalURL,
		ExpireAt:    arg.ExpireAt,
		CreatedAt:   time.Now(),
	})
	return err
}

func (r mongodbLinkRepository) GetLink(ctx context.Context, id string) (linkModel.Link, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var link linkModel.Link
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&link)
	if err == mongo.ErrNoDocuments {
		return linkModel.Link{}, errorCommon.NewNotFoundError("Link not found")
	} else if err != nil {
		return linkModel.Link{}, err
	}
	return link, nil
}
