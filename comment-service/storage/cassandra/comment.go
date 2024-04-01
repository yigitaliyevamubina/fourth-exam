package cassandra

import (
	"time"

	pb "comment-service/genproto/comment_service"

	"github.com/gocql/gocql"
	"github.com/golang/protobuf/ptypes/empty"
)

type CommentRepo struct {
	session *gocql.Session
}

func NewCommentRepo(cluster *gocql.ClusterConfig) *CommentRepo {
	session, err := cluster.CreateSession()
	if err != nil {
		return nil
	}
	return &CommentRepo{session: session}
}

func (c *CommentRepo) Close() {
	c.session.Close()
}

func (c *CommentRepo) Create(req *pb.Comment) (*pb.Comment, error) {
	if req.Id == "" {
		req.Id = gocql.TimeUUID().String()
	}

	timestamp := time.Now().UnixNano() / int64(time.Millisecond)

	query := `INSERT INTO comments (
		id, 
		post_id,
		user_id,
		content,
		created_at
	) VALUES (?, ?, ?, ?, ?)`

	err := c.session.Query(query,
		req.Id,
		req.PostId,
		req.UserId,
		req.Content,
		timestamp,
	).Exec()

	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *CommentRepo) Get(req *pb.Id) (*pb.Comment, error) {
	query := `SELECT id, 
		post_id,
		user_id,
		content,
		created_at
		FROM comments WHERE id = ?`

	var (
		resp       pb.Comment
		created_at time.Time
	)

	err := c.session.Query(query, req.CommentId).Scan(&resp.Id, &resp.PostId, &resp.UserId, &resp.Content, &created_at)
	if err != nil {
		return nil, err
	}
	resp.CreatedAt = created_at.Format("2006-01-02 15:04:05")

	return &resp, nil
}

func (c *CommentRepo) Update(req *pb.Comment) (*pb.Comment, error) {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)

	query := `UPDATE comments 
	SET content = ?,
		updated_at = ?
	WHERE id = ? AND post_id = ? AND user_id = ?`

	err := c.session.Query(query,
		req.Content,
		timestamp,
		req.Id, req.PostId, req.UserId).Exec()

	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *CommentRepo) Delete(req *pb.Id) (*empty.Empty, error) {
	query := `DELETE FROM comments WHERE id = ?`

	err := c.session.Query(query, req.CommentId).Exec()
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (c *CommentRepo) List(req *pb.GetListFilter) (*pb.Comments, error) {
	// offset := (req.Page - 1) * req.Limit
	query := `SELECT id, 
		post_id,
		user_id,
		content,
		created_at
		FROM comments`
	var args []interface{}

	if req.UserId != "" {
		query += " WHERE user_id = ?"
		args = append(args, req.UserId)
	}

	if req.PostId != "" {
		if req.UserId != "" {
			query += " AND"
			query += " post_id = ?"
		} else {
			query += " WHERE"
			query += " post_id = ?"
		}
		args = append(args, req.PostId)
	}

	if req.OrderBy != "" {
		query += " ORDER BY token(id)"
	}

	query += " LIMIT ? ALLOW FILTERING"
	args = append(args, req.Limit)

	iter := c.session.Query(query, args...).Iter()

	var comments pb.Comments
	for {
		var (
			resp       pb.Comment
			created_at time.Time
		)
		if !iter.Scan(&resp.Id, &resp.PostId, &resp.UserId, &resp.Content, &created_at) {
			break
		}
		resp.CreatedAt = created_at.Format("2006-01-02 15:04:05")
		comments.Count++
		comments.Items = append(comments.Items, &resp)
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return &comments, nil
}
