package persistence

import (
	"context"

	"github.com/duvansh91/xmen/internal/human/models"
	"github.com/duvansh91/xmen/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

// HumanMongoDBRepository groups structs needed to human repository.
type HumanMongoDBRepository struct {
	client *mongodb.MongoDB
}

// NewHumanMongoDBRepository creates a new instance of HumanMongoDBRepository.
func NewHumanMongoDBRepository(client *mongodb.MongoDB) *HumanMongoDBRepository {
	return &HumanMongoDBRepository{
		client: client,
	}
}

// Save stores a human in DB.
func (r *HumanMongoDBRepository) Save(human *models.Human) error {
	_, err := r.client.Collection.InsertOne(context.TODO(), human)
	if err != nil {
		return err
	}
	return nil
}

// FindAll searches all human verification documents.
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
