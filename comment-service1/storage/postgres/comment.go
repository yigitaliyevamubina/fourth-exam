package postgres

import (
	pb "comment-service/genproto/comment_service"
	"database/sql"
	"time"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/google/uuid"
)

type commentRepo struct {
	db *sql.DB
}

func NewCommentRepo(db *sql.DB) *commentRepo {
	return &commentRepo{
		db: db,
	}
}

// rpc Create(Comment) returns (Comment);
// rpc Update(Comment) returns (Comment);
// rpc Get(Id) returns (Comment);
// rpc Delete(Id) returns (google.protobuf.Empty);
// rpc List(GetListFilter) returns (Comments);

func (c *commentRepo) Create(req *pb.Comment) (*pb.Comment, error) {
	if req.Id == "" {
		req.Id = uuid.NewString()
	}
	query := `INSERT INTO comments(
							id, 
							post_id,
							user_id,
							content,
							created_at) 
							VALUES($1, $2, $3, $4, $5) 
							RETURNING id, 
										post_id,
										user_id,
										content,
										created_at,
										updated_at`

	var (
		updatedAt sql.NullTime
	)
	rowComment := c.db.QueryRow(query,
		req.Id,
		req.PostId,
		req.UserId,
		req.Content,
		time.Now())
	var resp pb.Comment
	if err := rowComment.Scan(&resp.Id,
		&resp.PostId,
		&resp.UserId,
		&resp.Content,
		&resp.CreatedAt,
		&updatedAt); err != nil {
		return nil, err
	}

	if updatedAt.Valid {
		resp.UpdatedAt = updatedAt.Time.Format("2006-01-02 15:04:05")
	}

	return &resp, nil
}

func (c *commentRepo) Get(id *pb.Id) (*pb.Comment, error) {
	query := `SELECT id, 
					post_id,
					user_id,
					content,
					created_at,
					updated_at FROM comments WHERE id = $1`

	rowComment := c.db.QueryRow(query, id.CommentId)

	var (
		resp      = pb.Comment{}
		updatedAt sql.NullTime
	)

	if err := rowComment.Scan(&resp.Id,
		&resp.PostId,
		&resp.UserId,
		&resp.Content,
		&resp.CreatedAt,
		&updatedAt); err != nil {
		return nil, err
	}

	if updatedAt.Valid {
		resp.UpdatedAt = updatedAt.Time.Format("2006-01-02 15:04:05")
	}

	return &resp, nil
}

func (p *commentRepo) Update(req *pb.Comment) (*pb.Comment, error) {
	query := `UPDATE comments 
				SET post_id = $1, 
					user_id = $2, 
					content = $3,
					updated_at = $4
					WHERE id = $5
					RETURNING  id, 
								post_id,
								user_id,
								content,
								created_at,
								updated_at`

	rowComment := p.db.QueryRow(query,
		req.PostId,
		req.UserId,
		req.Content,
		time.Now(),
		req.Id)
	var resp pb.Comment
	var (
		updatedAt sql.NullTime
	)

	if err := rowComment.Scan(&resp.Id,
		&resp.PostId,
		&resp.UserId,
		&resp.Content,
		&resp.CreatedAt,
		&updatedAt); err != nil {
		return nil, err
	}

	if updatedAt.Valid {
		resp.UpdatedAt = updatedAt.Time.Format("2006-01-02 15:04:05")
	}

	return &resp, nil
}

func (p *commentRepo) Delete(id *pb.Id) (*empty.Empty, error) {
	query := `DELETE FROM comments WHERE id = $1`

	_, err := p.db.Exec(query, id.CommentId)
	return &empty.Empty{}, err
}

func (p *commentRepo) List(req *pb.GetListFilter) (*pb.Comments, error) {
	offset := (req.Page - 1) * req.Limit
	query := `SELECT id, 
					post_id,
					user_id,
					content,
					created_at,
					updated_at FROM comments`
	var args []interface{}

	if req.UserId != "" {
		query += " WHERE user_id = $1"
		args = append(args, req.UserId)
	}

	if req.PostId != "" {
		if req.UserId != "" {
			query += " AND"
			query += " post_id = $2"
		} else {
			query += " WHERE"
			query += " post_id = $1"
		}
		args = append(args, req.PostId)
	}

	if req.OrderBy != "" {
		query += " ORDER BY " + req.OrderBy
	}

	if req.PostId != "" {
		if req.UserId != "" {
			query += " LIMIT $3 OFFSET $4"
		} else {
			query += " LIMIT $2 OFFSET $3"
		}
	} else {
		if req.UserId != "" {
			query += " LIMIT $2 OFFSET $3"
		} else {
			query += " LIMIT $1 OFFSET $2"
		}
	}

	args = append(args, req.Limit, offset)

	rows, err := p.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments pb.Comments

	for rows.Next() {
		var (
			updatedAt sql.NullTime
			resp      = pb.Comment{}
		)
		if err := rows.Scan(&resp.Id,
			&resp.PostId,
			&resp.UserId,
			&resp.Content,
			&resp.CreatedAt,
			&updatedAt); err != nil {
			return nil, err
		}
		if updatedAt.Valid {
			resp.UpdatedAt = updatedAt.Time.Format("2006-01-02 15:04:05")
		}
		comments.Count++
		comments.Items = append(comments.Items, &resp)
	}

	return &comments, nil
}
