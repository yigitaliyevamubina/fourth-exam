package postgres

import (
	"context"
	"database/sql"
	pb "exam/user-service/genproto/user_service"
	"exam/user-service/pkg/db"
	"exam/user-service/pkg/logger"
	"exam/user-service/storage/repo"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
)

type userRepo struct {
	db  *db.Postgres
	log logger.Logger
}

// Constructor
func NewUserRepo(db *db.Postgres, log logger.Logger) repo.UserServiceI {
	return &userRepo{
		db:  db,
		log: log,
	}
}

func (u *userRepo) Create(ctx context.Context, req *pb.User) (*pb.User, error) {
	if req.Id == "" {
		req.Id = uuid.NewString()
	}
	query := u.db.Builder.Insert("users").
		Columns(`
				id, 
				username,
				email,
				password,
				first_name,
				last_name,
				bio,
				website,
				is_active,
				refresh_token,
				created_at
			`).
		Values(
			req.Id, req.Username, req.Email, req.Password,
			req.FirstName, req.LastName, req.Bio, req.Website,
			true, req.RefreshToken, time.Now(),
		).
		Suffix("RETURNING created_at, updated_at")
	var (
		updatedAt sql.NullTime
	)
	err := query.RunWith(u.db.DB).QueryRow().Scan(&req.CreatedAt, &updatedAt)
	if err != nil {
		return nil, err
	}

	if updatedAt.Valid {
		req.UpdatedAt = updatedAt.Time.Format("2006-01-02 15:04:05")
	}
	return req, nil
}

func (u *userRepo) Get(ctx context.Context, req *pb.GetRequest) (*pb.UserModel, error) {
	resp := &pb.UserModel{Posts: []*pb.Post{}}

	query := u.db.Builder.Select(`
						id, 
						username,
						email,
						password,
						first_name,
						last_name,
						bio,
						website,
						is_active,
						refresh_token,
						created_at,
						updated_at
					`).From("users")

	if req.UserId != "" {
		query = query.Where(squirrel.Eq{"id": req.UserId})
	} else if req.Email != "" {
		query = query.Where(squirrel.Eq{"email": req.Email})
	} else if req.Username != "" {
		query = query.Where(squirrel.Eq{"username": req.Username})
	} else {
		return nil, fmt.Errorf("id/email/username, one of them is required")
	}

	var (
		updatedAt sql.NullTime
	)
	err := query.RunWith(u.db.DB).QueryRow().Scan(
		&resp.Id,
		&resp.Username,
		&resp.Email,
		&resp.Password,
		&resp.FirstName,
		&resp.LastName,
		&resp.Bio,
		&resp.Website,
		&resp.IsActive,
		&resp.RefreshToken,
		&resp.CreatedAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}

	if updatedAt.Valid {
		resp.UpdatedAt = updatedAt.Time.Format("2006-01-02 15:04:05")
	}

	return resp, nil
}

func (u *userRepo) Update(ctx context.Context, req *pb.User) (*pb.User, error) {
	var (
		updateMap = make(map[string]interface{})
		where     = squirrel.And{squirrel.Eq{"id": req.Id}}
	)

	updateMap["first_name"] = req.FirstName
	updateMap["last_name"] = req.LastName
	updateMap["username"] = req.Username
	updateMap["bio"] = req.Bio
	updateMap["website"] = req.Website
	updateMap["is_active"] = req.IsActive
	updateMap["updated_at"] = time.Now()

	query := u.db.Builder.Update("users").SetMap(updateMap).
		Where(where).
		Suffix("RETURNING created_at, updated_at")

	var (
		updatedAt sql.NullTime
	)
	err := query.RunWith(u.db.DB).QueryRow().Scan(
		&req.CreatedAt, &updatedAt,
	)
	if err != nil {
		return nil, err
	}

	if updatedAt.Valid {
		req.UpdatedAt = updatedAt.Time.Format("2006-01-02 15:04:05")
	}

	return req, nil
}

func (u *userRepo) Delete(ctx context.Context, req *pb.GetRequest) (*empty.Empty, error) {
	query := u.db.Builder.Delete("users")
	if req.UserId != "" {
		query = query.Where(squirrel.Eq{"id": req.UserId})
	} else if req.Email != "" {
		query = query.Where(squirrel.Eq{"email": req.Email})
	} else if req.Username != "" {
		query = query.Where(squirrel.Eq{"username": req.Username})
	} else {
		return nil, fmt.Errorf("id/email/username, one of them is required")
	}
	
	_, err := query.RunWith(u.db.DB).Exec()
	return &empty.Empty{}, err
}

func (u *userRepo) List(ctx context.Context, req *pb.GetListFilter) (*pb.Users, error) {
	var (
		users = pb.Users{Count: 0}
	)

	query := u.db.Builder.Select(
							`id, 
						username,
						email,
						password,
						first_name,
						last_name,
						bio,
						website,
						is_active,
						refresh_token,
						created_at,
						updated_at
	                `).From("users")

	query = query.Offset(uint64((req.Page - 1) * req.Limit)).Limit(uint64(req.Limit))
	query = query.OrderBy(req.OrderBy)

	query = query.Where(squirrel.Eq{"is_active": req.IsActive})

	rows, err := query.RunWith(u.db.DB).Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		resp := pb.UserModel{Posts: []*pb.Post{}}
		var (
			updatedAt sql.NullTime
		)
		err := query.RunWith(u.db.DB).QueryRow().Scan(
			&resp.Id,
			&resp.Username,
			&resp.Email,
			&resp.Password,
			&resp.FirstName,
			&resp.LastName,
			&resp.Bio,
			&resp.Website,
			&resp.IsActive,
			&resp.RefreshToken,
			&resp.CreatedAt,
			&updatedAt,
		)
		if err != nil {
			return nil, err
		}

		if updatedAt.Valid {
			resp.UpdatedAt = updatedAt.Time.Format("2006-01-02 15:04:05")
		}

		users.Users = append(users.Users, &resp)
		users.Count++
	}

	return &users, nil
}

func (u *userRepo) CheckField(ctx context.Context, req *pb.CheckFieldReq) (*pb.Status, error) {
	var (
		response = pb.Status{}
	)
	var ifexists int
	num := u.db.Builder.Select("count(1)").From("users").Where(squirrel.Eq{req.Field: req.Value})

	err := num.RunWith(u.db.DB).Scan(&ifexists)
	if err != nil {
		response.Status = false
		return &response, err
	}
	if ifexists == 1 {
		response.Status = true
	} else if ifexists == 0 {
		response.Status = false
	}

	return &response, nil
}

func (u *userRepo) UpdateRefresh(ctx context.Context, req *pb.UpdateRefreshReq) (*pb.User, error) {
	var (
		updateMap = make(map[string]interface{})
		where     = squirrel.And{squirrel.Eq{"id": req.UserId}}
		resp      = pb.User{}
	)

	updateMap["refresh_token"] = req.RefreshToken

	query := u.db.Builder.Update("users").SetMap(updateMap).
		Where(where).
		Suffix(`RETURNING id, 
						username,
						email,
						password,
						first_name,
						last_name,
						bio,
						website,
						is_active,
						refresh_token,
						created_at,
						updated_at`)

	var (
		updatedAt sql.NullTime
	)
	err := query.RunWith(u.db.DB).QueryRow().Scan(
		&resp.Id,
		&resp.Username,
		&resp.Email,
		&resp.Password,
		&resp.FirstName,
		&resp.LastName,
		&resp.Bio,
		&resp.Website,
		&resp.IsActive,
		&resp.RefreshToken,
		&resp.CreatedAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}

	if updatedAt.Valid {
		resp.UpdatedAt = updatedAt.Time.Format("2006-01-02 15:04:05")
	}

	return &resp, nil
}
