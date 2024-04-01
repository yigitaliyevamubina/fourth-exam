package mongodb

import (
	"context"
	"fmt"
	"template-post-service/pkg/logger"
	"time"

	pb "template-post-service/genproto/post_service"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type postRepo struct {
	collection *mongo.Collection
	log        logger.Logger
}

func NewPostRepo(collection *mongo.Collection, log logger.Logger) *postRepo {
	return &postRepo{collection: collection, log: log}
}

// Create(*pb.Post) (*pb.Post, error)
// Get(*pb.Id) (*pb.Post, error)
// Update(*pb.Post) (*pb.Post, error)
// Delete(*pb.Id) (*empty.Empty, error)
// List(*pb.GetListFilter) (*pb.Posts, error)

func (p *postRepo) Create(req *pb.Post) (*pb.Post, error) {
	if req.Id == "" {
		req.Id = uuid.NewString()
	}
	req.CreatedAt = time.Now().String()
	result, err := p.collection.InsertOne(context.Background(), req)
	if err != nil {
		return nil, err
	}

	var resp pb.Post
	filter := bson.M{"_id": result.InsertedID}
	err = p.collection.FindOne(context.Background(), filter).Decode(&resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (p *postRepo) Get(id *pb.Id) (*pb.Post, error) {
	var (
		resp = pb.Post{}
	)
	filter := bson.M{"id": id.PostId}
	err := p.collection.FindOne(context.Background(), filter).Decode(&resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (p *postRepo) Update(req *pb.Post) (*pb.Post, error) {
	filter := bson.M{"id": req.Id}

	updateReq := bson.M{
		"$set": bson.M{
			"userid":    req.UserId,
			"content":   req.Content,
			"title":     req.Title,
			"likes":     req.Likes,
			"dislikes":  req.Dislikes,
			"views":     req.Views,
			"category":  req.Category,
			"updatedat": time.Now().String(),
		},
	}

	result, err := p.collection.UpdateOne(context.Background(), filter, updateReq)
	if err != nil {
		return nil, err
	}

	if result.ModifiedCount == 0 {
		return nil, fmt.Errorf("post not found")
	}

	return req, nil
}

func (p *postRepo) Delete(id *pb.Id) (*empty.Empty, error) {
	filter := bson.M{"id": id.PostId}
	_, err := p.collection.DeleteOne(context.Background(), filter)
	return &empty.Empty{}, err
}

func (p *postRepo) List(req *pb.GetListFilter) (*pb.Posts, error) {
	var posts pb.Posts

	reqOptions := options.Find()

	reqOptions.SetSkip(int64((req.Page - 1) * req.Limit))
	reqOptions.SetLimit(int64(req.Limit))

	if req.OrderBy != "" {
		order := bson.D{{Key: req.OrderBy, Value: 1}}
		reqOptions.SetSort(order)
	}

	cursor, err := p.collection.Find(context.Background(), bson.M{}, reqOptions)
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var (
			resp = pb.Post{}
		)
		err = cursor.Decode(&resp)
		if err != nil {
			return nil, err
		}

		posts.Count++
		posts.Items = append(posts.Items, &resp)
	}

	return &posts, nil
}
