package mongodb

import (
	"comment-service/config"
	pb "comment-service/genproto/comment_service"
	"comment-service/pkg/logger"
	"comment-service/storage/repo"
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type CommentTestSuite struct {
	suite.Suite
	Repository repo.CommentStorageI
}

func (p *CommentTestSuite) SetupSuite() {
	cfg := config.Load()
	l := logger.New("", "")
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", cfg.MongoDBhost, cfg.MongoDBport))
	clientMongo, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		l.Error("failed to connect to mongodb: %v", logger.Error(err))
	}

	collection := clientMongo.Database(cfg.MongoDBdatabase).Collection(cfg.MongoDBCollection)
	p.Repository = NewCommentRepo(collection, l)
}

func (c *CommentTestSuite) TestCommentCRUD() {
	//You should have user and post to create a comment
	id := uuid.NewString()
	comment := &pb.Comment{
		Id:      id,
		PostId:  "e292ca9d-d202-4aa2-a7de-487158b02dd4",
		UserId:  "f8e1c19a-6eca-411d-8cd0-30d9716db606",
		Content: gofakeit.Sentence(5),
	}

	createResp, err := c.Repository.Create(comment)
	c.Suite.NoError(err)
	c.Suite.NotNil(createResp)

	//Get coment
	commentID := &pb.Id{CommentId: comment.Id}
	getResp, err := c.Repository.Get(commentID)
	c.Suite.NoError(err)
	c.Suite.NotNil(getResp)
	c.Suite.Equal(createResp.Id, getResp.Id)
	c.Suite.Equal(createResp.PostId, getResp.PostId)
	c.Suite.Equal(createResp.UserId, getResp.UserId)
	c.Suite.Equal(createResp.Content, getResp.Content)

	//List comments
	listResp, err := c.Repository.List(&pb.GetListFilter{Page: 1, Limit: 10, UserId: comment.UserId, PostId: comment.PostId})
	c.Suite.NoError(err)
	c.Suite.NotNil(listResp)

	//Update comment
	updatedContent := gofakeit.Sentence(5)
	comment.Content = updatedContent
	updateResp, err := c.Repository.Update(comment)
	c.Suite.NoError(err)
	c.Suite.NotNil(updateResp)
	c.Suite.Equal(comment.Id, updateResp.Id)
	c.Suite.Equal(comment.Content, updateResp.Content)

	//Delete comment
	_, err = c.Repository.Delete(commentID)
	c.Suite.NoError(err)
}

func TestCommentRepository(t *testing.T) {
	suite.Run(t, new(CommentTestSuite))
}
