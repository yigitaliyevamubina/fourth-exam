package mongodb

import (
	"context"
	"errors"
	"exam/api-gateway/api/handlers/models"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type adminRepo struct {
	collection *mongo.Collection
}

func NewAdminRepo(collection *mongo.Collection) *adminRepo {
	return &adminRepo{collection: collection}
}

func (r *adminRepo) Create(admin *models.AdminResp) error {
	_, err := r.collection.InsertOne(context.Background(), admin)
	return err
}

func (r *adminRepo) Delete(userName, password string) error {
	filter := bson.M{"username": userName, "password": password}
	result, err := r.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		fmt.Println("error")
		return errors.New("no documents were deleted")
	}

	return nil
}

func (r *adminRepo) Get(userName string) (string, string, bool, error) {
	filter := bson.M{"username": userName}
	var admin models.AdminResp
	err := r.collection.FindOne(context.Background(), filter).Decode(&admin)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", "", false, nil
		}
		return "", "", false, err
	}

	return admin.Role, admin.Password, true, nil
}
