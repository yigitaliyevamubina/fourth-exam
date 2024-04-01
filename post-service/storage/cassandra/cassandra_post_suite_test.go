package cassandra

import (
	"template-post-service/config"
	pb "template-post-service/genproto/post_service"
	"template-post-service/storage/repo"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type PostTestSuite struct {
	suite.Suite
	Repository repo.PostStorageI
}

func (p *PostTestSuite) SetupSuite() {
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

	// Create posts table
	if err := session.Query(`
        CREATE TABLE IF NOT EXISTS posts (
            id UUID PRIMARY KEY,
            user_id UUID,
            content TEXT,
            title TEXT,
            likes BIGINT,
            dislikes BIGINT,
            views BIGINT,
            category TEXT,
            created_at TIMESTAMP,
            updated_at TIMESTAMP,
            deleted_at TIMESTAMP
        );
    `).Exec(); err != nil {
		panic(err)
	}
	p.Repository = NewPostRepo(cluster)
}

func (p *PostTestSuite) TestPostCRUD() {
	//Create Post (make sure you already user in the users table before creating a post)
	id := uuid.NewString()
	post := &pb.Post{
		Id:       id,
		UserId:   "f8e1c19a-6eca-411d-8cd0-30d9716db606",
		Content:  gofakeit.Sentence(2),
		Title:    gofakeit.Book().Title,
		Likes:    10,
		Dislikes: 5,
		Views:    50,
		Category: gofakeit.Noun(),
	}

	createResp, err := p.Repository.Create(post)
	p.Suite.NoError(err)
	p.Suite.NotNil(createResp)

	//Get post
	postId := &pb.Id{PostId: post.Id}
	getResp, err := p.Repository.Get(postId)
	p.Suite.NoError(err)
	p.Suite.NotNil(getResp)
	p.Suite.Equal(createResp.Id, getResp.Id)
	p.Suite.Equal(createResp.UserId, getResp.UserId)
	p.Suite.Equal(createResp.Content, getResp.Content)
	p.Suite.Equal(createResp.Title, getResp.Title)
	p.Suite.Equal(createResp.Likes, getResp.Likes)
	p.Suite.Equal(createResp.Dislikes, getResp.Dislikes)
	p.Suite.Equal(createResp.Views, getResp.Views)
	p.Suite.Equal(createResp.Category, getResp.Category)

	//List posts
	listResp, err := p.Repository.List(&pb.GetListFilter{Page: 1, Limit: 10, OrderBy: "created_at", UserId: "f8e1c19a-6eca-411d-8cd0-30d9716db606"})
	p.Suite.NoError(err)
	p.Suite.NotNil(listResp)

	//Update post
	updatedTitle := gofakeit.Book().Title
	post.Title = updatedTitle
	updatedContent := gofakeit.Sentence(3)
	post.Content = updatedContent
	updateResp, err := p.Repository.Update(post)
	p.Suite.NoError(err)
	p.Suite.NotNil(updateResp)
	p.Suite.Equal(post.Title, updateResp.Title)
	p.Suite.Equal(post.Content, updateResp.Content)

	//Delete post
	_, err = p.Repository.Delete(postId)
	p.Suite.NoError(err)
}

func TestPostRepository(t *testing.T) {
	suite.Run(t, new(PostTestSuite))
}
