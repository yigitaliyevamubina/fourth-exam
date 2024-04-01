package postgres

import (
	"context"
	"time"

	pb "exam/user-service/genproto/user_service"
	"exam/user-service/storage/repo"

	"github.com/gocql/gocql"
	"github.com/golang/protobuf/ptypes/empty"
)

type userRepo struct {
	session *gocql.Session
}

// Constructor
func NewCassandraRepo(cluster *gocql.ClusterConfig) repo.UserServiceI {
	session, err := cluster.CreateSession()
	if err != nil {
		return nil
	}

	return &userRepo{session: session}
}

func (u *userRepo) Create(ctx context.Context, req *pb.User) (*pb.User, error) {
	if req.Id == "" {
		req.Id = gocql.TimeUUID().String()
	}
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)

	query := u.session.Query(`
		INSERT INTO users (
			id, username, email, password, first_name, last_name, bio, website, is_active, refresh_token, created_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, req.Id, req.Username, req.Email, req.Password, req.FirstName, req.LastName, req.Bio, req.Website, true, req.RefreshToken, timestamp)

	if err := query.Exec(); err != nil {
		return nil, err
	}

	return req, nil
}

func (u *userRepo) Get(ctx context.Context, req *pb.GetRequest) (*pb.UserModel, error) {
	resp := &pb.UserModel{}
	var createdAt time.Time
	var updatedAt time.Time

	query := u.session.Query(`
		SELECT id, username, email, password, first_name, last_name, bio, website, is_active, refresh_token, created_at, updated_at
		FROM users
		WHERE id = ?
	`, req.UserId)

	err := query.Scan(
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
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		if err == gocql.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	resp.CreatedAt = createdAt.Format("2006-01-02 15:04:05")

	resp.UpdatedAt = updatedAt.Format("2006-01-02 15:04:05")

	return resp, nil
}

func (u *userRepo) Update(ctx context.Context, req *pb.User) (*pb.User, error) {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)

	updateMap := map[string]interface{}{
		"first_name": req.FirstName,
		"last_name":  req.LastName,
		"username":   req.Username,
		"bio":        req.Bio,
		"website":    req.Website,
		"is_active":  req.IsActive,
		"updated_at": timestamp,
	}

	query := u.session.Query(`
		UPDATE users
		SET first_name = ?, last_name = ?, username = ?, bio = ?, website = ?, is_active = ?, updated_at = ?
		WHERE id = ?
	`, updateMap["first_name"], updateMap["last_name"], updateMap["username"], updateMap["bio"], updateMap["website"], updateMap["is_active"], updateMap["updated_at"], req.Id)

	if err := query.Exec(); err != nil {
		return nil, err
	}

	return req, nil
}

func (u *userRepo) Delete(ctx context.Context, req *pb.GetRequest) (*empty.Empty, error) {
	query := u.session.Query(`
		DELETE FROM users
		WHERE id = ?
	`, req.UserId)

	if err := query.Exec(); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (u *userRepo) List(ctx context.Context, req *pb.GetListFilter) (*pb.Users, error) {
	var (
		users = &pb.Users{Count: 0}
	)

	query := u.session.Query(`
		SELECT id, username, email, password, first_name, last_name, bio, website, is_active, refresh_token, created_at, updated_at
		FROM users
		WHERE is_active = ?
		LIMIT ? ALLOW FILTERING
	`, req.IsActive, req.Limit)

	iter := query.Iter()
	defer iter.Close()

	for {
		resp := &pb.UserModel{}
		var createdAt time.Time
		var updatedAt time.Time

		if !iter.Scan(
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
			&createdAt,
			&updatedAt,
		) {
			break
		}

		resp.CreatedAt = createdAt.Format("2006-01-02 15:04:05")
		resp.UpdatedAt = updatedAt.Format("2006-01-02 15:04:05")

		users.Users = append(users.Users, resp)
		users.Count++
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userRepo) CheckField(ctx context.Context, req *pb.CheckFieldReq) (*pb.Status, error) {
	var (
		response = &pb.Status{}
	)
	var ifexists int
	num := u.session.Query("SELECT COUNT(*) FROM users WHERE "+req.Field+" = ? ALLOW FILTERING", req.Value)

	err := num.Scan(&ifexists)
	if err != nil {
		response.Status = false
		return response, err
	}
	if ifexists == 1 {
		response.Status = true
	} else if ifexists == 0 {
		response.Status = false
	}

	return response, nil
}

func (u *userRepo) UpdateRefresh(ctx context.Context, req *pb.UpdateRefreshReq) (*pb.User, error) {
	updateMap := map[string]interface{}{
		"refresh_token": req.RefreshToken,
	}

	// Update the refresh token
	query := u.session.Query(`
        UPDATE users
        SET refresh_token = ?
        WHERE id = ?
    `, updateMap["refresh_token"], req.UserId)

	if err := query.Exec(); err != nil {
		return nil, err
	}

	// Get the user by user ID to get the updated details
	resp := &pb.User{}
	var createdAt time.Time
	var updatedAt time.Time

	err := u.session.Query(`
        SELECT id, username, email, password, first_name, last_name, bio, website, is_active, refresh_token, created_at, updated_at
        FROM users
        WHERE id = ?
    `, req.UserId).Scan(
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
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}

	resp.CreatedAt = createdAt.Format("2006-01-02 15:04:05")
	resp.UpdatedAt = updatedAt.Format("2006-01-02 15:04:05")

	return resp, nil
}
