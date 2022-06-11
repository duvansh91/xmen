package persistence

import (
	"context"

	"github.com/duvansh91/xmen/internal/human/models"
	"github.com/duvansh91/xmen/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

type HumanMongoDBRepository struct {
	client *mongodb.MongoDB
}

func (r *HumanMongoDBRepository) Save(human *models.Human) error {
	_, err := r.client.Collection.InsertOne(context.TODO(), human)
	if err != nil {
		return err
	}
	return nil
}

func (r *HumanMongoDBRepository) FindAll() ([]*models.Human, error) {
	cursor, err := r.client.Collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	results := []*models.Human{}

	err = cursor.All(context.TODO(), &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func NewHumanMongoDBRepository(client *mongodb.MongoDB) *HumanMongoDBRepository {
	return &HumanMongoDBRepository{
		client: client,
	}
}
