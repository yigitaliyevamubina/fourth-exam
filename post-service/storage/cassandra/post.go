package cassandra

import (
	"time"

	pb "template-post-service/genproto/post_service"

	"github.com/gocql/gocql"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
)

type PostRepo struct {
	session *gocql.Session
}

func NewPostRepo(cluster *gocql.ClusterConfig) *PostRepo {
	session, err := cluster.CreateSession()
	if err != nil {
		return nil
	}
	return &PostRepo{session: session}
}

func (p *PostRepo) Create(req *pb.Post) (*pb.Post, error) {
	if req.Id == "" {
		req.Id = uuid.NewString()
	}

	timestamp := time.Now().UnixNano() / int64(time.Millisecond)

	query := `INSERT INTO posts(
		id, 
		user_id, 
		content, 
		title,
		likes,
		dislikes,
		views,
		category,
		created_at
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	err := p.session.Query(query,
		req.Id,
		req.UserId,
		req.Content,
		req.Title,
		req.Likes,
		req.Dislikes,
		req.Views,
		req.Category,
		timestamp,
	).Exec()

	if err != nil {
		return nil, err
	}

	return req, nil
}

func (p *PostRepo) Get(id *pb.Id) (*pb.Post, error) {
	query := `SELECT id, 
		user_id, 
		content, 
		title,
		likes,
		dislikes,
		views,
		category,
		created_at
		FROM posts WHERE id = ?`

	var resp pb.Post
	var createdAt time.Time

	err := p.session.Query(query, id.PostId).Scan(&resp.Id,
		&resp.UserId,
		&resp.Content,
		&resp.Title,
		&resp.Likes,
		&resp.Dislikes,
		&resp.Views,
		&resp.Category,
		&createdAt)

	if err != nil {
		return nil, err
	}

	resp.CreatedAt = createdAt.Format("2006-01-02 15:04:05")

	return &resp, nil
}

func (p *PostRepo) Update(req *pb.Post) (*pb.Post, error) {
	query := `UPDATE posts 
        SET user_id = ?, 
            content = ?, 
            title = ?,
            likes = ?,
            dislikes = ?,
            views = ?,
            category = ?,
            updated_at = ?
            WHERE id = ?`

	timestamp := time.Now().UnixNano() / int64(time.Millisecond)

	err := p.session.Query(query,
		req.UserId,
		req.Content,
		req.Title,
		req.Likes,
		req.Dislikes,
		req.Views,
		req.Category,
		timestamp,
		req.Id).Exec()

	if err != nil {
		return nil, err
	}

	return req, nil
}

func (p *PostRepo) Delete(id *pb.Id) (*empty.Empty, error) {
	query := `DELETE FROM posts WHERE id = ?`

	err := p.session.Query(query, id.PostId).Exec()
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (p *PostRepo) List(req *pb.GetListFilter) (*pb.Posts, error) {
	// offset := (req.Page - 1) * req.Limit
	query := `SELECT id, 
		user_id, 
		content, 
		title,
		likes,
		dislikes,
		views,
		category,
		created_at
		FROM posts`

	var rows *gocql.Iter

	if req.UserId != "" {
		query += " WHERE user_id = ?"
		rows = p.session.Query(query, req.UserId).Iter()
	} else {
		rows = p.session.Query(query).Iter()
	}

	var posts pb.Posts

	for {
		var resp pb.Post
		var createdAt time.Time

		if !rows.Scan(&resp.Id,
			&resp.UserId,
			&resp.Content,
			&resp.Title,
			&resp.Likes,
			&resp.Dislikes,
			&resp.Views,
			&resp.Category,
			&createdAt) {
			break
		}

		resp.CreatedAt = createdAt.Format("2006-01-02 15:04:05")
		posts.Count++
		posts.Items = append(posts.Items, &resp)
	}

	return &posts, nil
}
