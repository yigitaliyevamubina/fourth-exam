package mongodb

import (
	"context"
	pb "exam/user-service/genproto/user_service"
	"exam/user-service/pkg/logger"
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepo struct {
	collection *mongo.Collection
	log        logger.Logger
}

func NewUserRepo(collection *mongo.Collection, log logger.Logger) *userRepo {
	return &userRepo{collection: collection, log: log}
}

func (u *userRepo) Create(ctx context.Context, user *pb.User) (*pb.User, error) {
	if user.Id == "" {
		user.Id = uuid.NewString()
	}
	user.CreatedAt = time.Now().String()
	user.IsActive = true
	result, err := u.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	var resp pb.User
	filter := bson.M{"_id": result.InsertedID}
	err = u.collection.FindOne(ctx, filter).Decode(&resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (u *userRepo) Get(ctx context.Context, req *pb.GetRequest) (*pb.UserModel, error) {
	var (
		resp   = pb.UserModel{Posts: []*pb.Post{}}
		filter bson.M
	)
	if req.UserId != "" {
		filter = bson.M{"id": req.UserId, "isactive": true}
	} else if req.Email != "" {
		filter = bson.M{"email": req.Email, "isactive": true}
	} else if req.Username != "" {
		filter = bson.M{"username": req.Username, "isactive": true}
	} else {
		return nil, fmt.Errorf("id/email/username, one of them is required")
	}

	err := u.collection.FindOne(ctx, filter).Decode(&resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (u *userRepo) Update(ctx context.Context, req *pb.User) (*pb.User, error) {
	filter := bson.M{"id": req.Id, "isactive": true}

	updateReq := bson.M{
		"$set": bson.M{
			"firstname": req.FirstName,
			"lastname":  req.LastName,
			"username":  req.Username,
			"bio":       req.Bio,
			"website":   req.Website,
			"isactive":  req.IsActive,
			"updatedat": time.Now().String(),
		},
	}

	result, err := u.collection.UpdateOne(ctx, filter, updateReq)
	if err != nil {
		return nil, err
	}

	if result.ModifiedCount == 0 {
		return nil, fmt.Errorf("user not found or already deleted")
	}

	return req, nil
}

func (u *userRepo) Delete(ctx context.Context, req *pb.GetRequest) (*empty.Empty, error) {
	var (
		filter bson.M
	)

	if req.UserId != "" {
		filter = bson.M{"id": req.UserId, "isactive": true}
	} else if req.Email != "" {
		filter = bson.M{"email": req.Email, "isactive": true}
	} else if req.Username != "" {
		filter = bson.M{"username": req.Username, "isactive": true}
	} else {
		return nil, fmt.Errorf("id/email/username, one of them is required")
	}

	updateReq := bson.M{
		"$set": bson.M{
			"deletedat": time.Now().String(),
			"isactive":  false,
		},
	}

	_, err := u.collection.UpdateOne(ctx, filter, updateReq)
	return &empty.Empty{}, err
}

func (u *userRepo) List(ctx context.Context, req *pb.GetListFilter) (*pb.Users, error) {
	var users pb.Users

	reqOptions := options.Find()

	reqOptions.SetSkip(int64((req.Page - 1) * req.Limit))
	reqOptions.SetLimit(int64(req.Limit))

	if req.OrderBy != "" {
		order := bson.D{{Key: req.OrderBy, Value: 1}}
		reqOptions.SetSort(order)
	}

	cursor, err := u.collection.Find(ctx, bson.M{"isactive": req.IsActive}, reqOptions)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var (
			resp = pb.UserModel{Posts: []*pb.Post{}}
		)
		err = cursor.Decode(&resp)
		if err != nil {
			return nil, err
		}

		users.Count++
		users.Users = append(users.Users, &resp)
	}

	return &users, nil
}

func (u *userRepo) CheckField(ctx context.Context, req *pb.CheckFieldReq) (*pb.Status, error) {
	filter := bson.M{req.Field: req.Value, "isactive": true}
	var res pb.User
	err := u.collection.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return &pb.Status{Status: false}, nil
	}
	return &pb.Status{Status: true}, nil
}

func (u *userRepo) UpdateRefresh(ctx context.Context, req *pb.UpdateRefreshReq) (*pb.User, error) {
	filter := bson.M{"id": req.UserId, "isactive": true}

	updateReq := bson.M{
		"$set": bson.M{
			"refreshtoken": req.RefreshToken,
			"updatedat":    time.Now().String(),
		},
	}

	result, err := u.collection.UpdateOne(ctx, filter, updateReq)
	if err != nil {
		return nil, err
	}

	if result.ModifiedCount == 0 {
		return nil, fmt.Errorf("user not found or already deleted")
	}

	var (
		resp pb.User
	)

	filter = bson.M{"id": req.UserId, "isactive": true}

	err = u.collection.FindOne(ctx, filter).Decode(&resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil

}
