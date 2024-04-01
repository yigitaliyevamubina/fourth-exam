package cassandra

import (
	"comment-service/config"
	pb "comment-service/genproto/comment_service"
	"comment-service/storage/repo"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gocql/gocql"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type CommentTestSuite struct {
	suite.Suite
	Repository repo.CommentStorageI
}

func (p *CommentTestSuite) SetupSuite() {
	cfg := config.Load()
	cluster := gocql.NewCluster(cfg.CassandraCluster)
	cluster.Keyspace = cfg.CassandraKeyspaceName
	cluster.Consistency = gocql.Quorum

	// Create a session
	session, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Create comments table
	if err := session.Query(`
        CREATE TABLE IF NOT EXISTS comments (
            id UUID,
            post_id UUID,
            user_id UUID,
            content TEXT,
            created_at TIMESTAMP,
            updated_at TIMESTAMP,
            deleted_at TIMESTAMP,
            PRIMARY KEY (id, post_id, user_id)
        );
    `).Exec(); err != nil {
		panic(err)
	}

	p.Repository = NewCommentRepo(cluster)
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
