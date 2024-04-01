package mongodb

import (
	pb "comment-service/genproto/comment_service"
	"comment-service/pkg/logger"
	"context"
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type commentRepo struct {
	collection *mongo.Collection
	log        logger.Logger
}

func NewCommentRepo(collection *mongo.Collection, log logger.Logger) *commentRepo {
	return &commentRepo{collection: collection, log: log}
}

func (p *commentRepo) Create(req *pb.Comment) (*pb.Comment, error) {
	if req.Id == "" {
		req.Id = uuid.NewString()
	}
	req.CreatedAt = time.Now().String()
	result, err := p.collection.InsertOne(context.Background(), req)
	if err != nil {
		return nil, err
	}

	var resp pb.Comment
	filter := bson.M{"_id": result.InsertedID}
	err = p.collection.FindOne(context.Background(), filter).Decode(&resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (p *commentRepo) Get(id *pb.Id) (*pb.Comment, error) {
	var (
		resp = pb.Comment{}
	)
	filter := bson.M{"id": id.CommentId}
	err := p.collection.FindOne(context.Background(), filter).Decode(&resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (p *commentRepo) Update(req *pb.Comment) (*pb.Comment, error) {
	filter := bson.M{"id": req.Id}

	updateReq := bson.M{
		"$set": bson.M{
			"postid":    req.PostId,
			"userid":    req.UserId,
			"content":   req.Content,
			"updatedat": time.Now().String(),
		},
	}

	result, err := p.collection.UpdateOne(context.Background(), filter, updateReq)
	if err != nil {
		return nil, err
	}

	if result.ModifiedCount == 0 {
		return nil, fmt.Errorf("comment not found")
	}

	return req, nil
}

func (p *commentRepo) Delete(id *pb.Id) (*empty.Empty, error) {
	filter := bson.M{"id": id.CommentId}
	_, err := p.collection.DeleteOne(context.Background(), filter)
	return &empty.Empty{}, err
}

func (p *commentRepo) List(req *pb.GetListFilter) (*pb.Comments, error) {
	var commments pb.Comments

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
			resp = pb.Comment{}
		)
		err = cursor.Decode(&resp)
		if err != nil {
			return nil, err
		}

		commments.Count++
		commments.Items = append(commments.Items, &resp)
	}

	return &commments, nil
}
