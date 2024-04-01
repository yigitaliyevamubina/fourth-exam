package postgres

import (
	"database/sql"
	pb "template-post-service/genproto/post_service"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
)

type postRepo struct {
	db *sql.DB
}

// NewPostRepo
func NewPostRepo(db *sql.DB) *postRepo {
	return &postRepo{db: db}
}

func (p *postRepo) Create(req *pb.Post) (*pb.Post, error) {
	if req.Id == "" {
		req.Id = uuid.NewString()
	}
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
					) 
					VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) 
					RETURNING id, 
							user_id, 
							content, 
							title,
							likes,
							dislikes,
							views,
							category,
							created_at,
							updated_at`

	var resp pb.Post
	rowComment := p.db.QueryRow(
		query,
		req.Id,
		req.UserId,
		req.Content,
		req.Title,
		req.Likes,
		req.Dislikes,
		req.Views,
		req.Category,
		time.Now())

	var (
		updatedAt sql.NullTime
	)
	if err := rowComment.Scan(
		&resp.Id,
		&resp.UserId,
		&resp.Content,
		&resp.Title,
		&resp.Likes,
		&resp.Dislikes,
		&resp.Views,
		&resp.Category,
		&resp.CreatedAt,
		&updatedAt); err != nil {
		return nil, err
	}

	if updatedAt.Valid {
		resp.UpdatedAt = updatedAt.Time.Format("2006-01-02 15:04:05")
	}

	return &resp, nil
}

func (p *postRepo) Get(id *pb.Id) (*pb.Post, error) {
	query := `SELECT id, 
					user_id, 
					content, 
					title,
					likes,
					dislikes,
					views,
					category,
					created_at,
					updated_at FROM posts WHERE id = $1`

	rowPost := p.db.QueryRow(query, id.PostId)

	var (
		updatedAt sql.NullTime
		resp      = pb.Post{}
	)

	if err := rowPost.Scan(&resp.Id,
		&resp.UserId,
		&resp.Content,
		&resp.Title,
		&resp.Likes,
		&resp.Dislikes,
		&resp.Views,
		&resp.Category,
		&resp.CreatedAt,
		&updatedAt); err != nil {
		return nil, err
	}

	if updatedAt.Valid {
		resp.UpdatedAt = updatedAt.Time.Format("2006-01-02 15:04:05")
	}

	return &resp, nil
}

func (p *postRepo) Update(req *pb.Post) (*pb.Post, error) {
	query := `UPDATE posts 
				SET user_id = $1, 
					content = $2, 
					title = $3,
					likes = $4,
					dislikes = $5,
					views = $6,
					category = $7,
					updated_at = $8
					WHERE id = $9
					RETURNING id, 
							user_id, 
							content, 
							title,
							likes,
							dislikes,
							views,
							category,
							created_at,
							updated_at`

	rowPost := p.db.QueryRow(query,
		req.UserId,
		req.Content,
		req.Title,
		req.Likes,
		req.Dislikes,
		req.Views,
		req.Category,
		time.Now(),
		req.Id)
	var resp pb.Post
	var (
		updatedAt sql.NullTime
	)

	if err := rowPost.Scan(&resp.Id,
		&resp.UserId,
		&resp.Content,
		&resp.Title,
		&resp.Likes,
		&resp.Dislikes,
		&resp.Views,
		&resp.Category,
		&resp.CreatedAt,
		&updatedAt); err != nil {
		return nil, err
	}

	if updatedAt.Valid {
		resp.UpdatedAt = updatedAt.Time.Format("2006-01-02 15:04:05")
	}

	return &resp, nil
}

func (p *postRepo) Delete(id *pb.Id) (*empty.Empty, error) {
	query := `DELETE FROM posts WHERE id = $1`

	_, err := p.db.Exec(query, id.PostId)
	return &empty.Empty{}, err
}

func (p *postRepo) List(req *pb.GetListFilter) (*pb.Posts, error) {
	offset := (req.Page - 1) * req.Limit
	query := `SELECT id, 
					user_id, 
					content, 
					title,
					likes,
					dislikes,
					views,
					category,
					created_at,
					updated_at
					FROM posts`
	if req.UserId != "" {
		query += " WHERE user_id = $1"
	}

	if req.OrderBy != "" {
		query += " ORDER BY " + req.OrderBy
	}

	if req.UserId != "" {
		query += " LIMIT $2 OFFSET $3"
	} else {
		query += " LIMIT $1 OFFSET $2"
	}

	var (
		err  error
		rows *sql.Rows
	)
	if req.UserId != "" {
		rows, err = p.db.Query(query, req.UserId, req.Limit, offset)
	} else {
		rows, err = p.db.Query(query, req.Limit, offset)
	}

	if err != nil {
		return nil, err
	}

	var posts pb.Posts

	for rows.Next() {
		var (
			updatedAt sql.NullTime
			resp      = pb.Post{}
		)
		if err := rows.Scan(&resp.Id,
			&resp.UserId,
			&resp.Content,
			&resp.Title,
			&resp.Likes,
			&resp.Dislikes,
			&resp.Views,
			&resp.Category,
			&resp.CreatedAt,
			&updatedAt); err != nil {
			return nil, err
		}
		if updatedAt.Valid {
			resp.UpdatedAt = updatedAt.Time.Format("2006-01-02 15:04:05")
		}
		posts.Count++
		posts.Items = append(posts.Items, &resp)
	}

	return &posts, nil
}
