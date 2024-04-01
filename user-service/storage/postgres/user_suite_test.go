package postgres

import (
	"context"
	"exam/user-service/config"
	pb "exam/user-service/genproto/user_service"
	db2 "exam/user-service/pkg/db"
	"exam/user-service/pkg/logger"
	"exam/user-service/storage/repo"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type UserTestSuite struct {
	suite.Suite
	CleanupFunc func()
	Repository  repo.UserServiceI
}

func (u *UserTestSuite) SetupSuite() {
	db, _ := db2.New(*config.Load())
	u.Repository = NewUserRepo(db, logger.New("", ""))
	u.CleanupFunc = db.Close
}

func (u *UserTestSuite) TestUserCRUD() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()
	//Create user
	id := uuid.New().String()
	user := &pb.User{
		Id:        id,
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Username: gofakeit.Username(),
		Email:     gofakeit.Email(),
		Password:  gofakeit.Phrase(),
		Bio: gofakeit.Sentence(15),
		Website: gofakeit.Phrase(),
		IsActive: true,
	}

	createResp, err := u.Repository.Create(ctx, user)
	u.Suite.NoError(err)
	u.Suite.NotNil(createResp)

	//Get user
	userId := &pb.GetRequest{
		UserId: id,
	}
	getResp, err := u.Repository.Get(ctx, userId)
	u.Suite.NoError(err)
	u.Suite.NotNil(getResp)
	u.Suite.Equal(getResp.LastName, user.LastName)
	u.Suite.Equal(getResp.FirstName, user.FirstName)
	u.Suite.Equal(getResp.Email, user.Email)
	u.Suite.Equal(getResp.Username, user.Username)
	u.Suite.Equal(getResp.Password, user.Password)
	u.Suite.Equal(getResp.Bio, user.Bio)
	u.Suite.Equal(getResp.Website, user.Website)

	//List users
	listResp, err := u.Repository.List(ctx, &pb.GetListFilter{
		Page:  1,
		Limit: 10,
		OrderBy: "created_at",
		IsActive: true,
	})
	u.Suite.NoError(err)
	u.Suite.NotNil(listResp)

	//Update user
	updatedName := gofakeit.FirstName()
	user.FirstName = updatedName
	updatedBio := gofakeit.Sentence(15)
	user.Bio = updatedBio
	user.Id = userId.UserId
	updateResp, err := u.Repository.Update(ctx, user)
	u.Suite.NoError(err)
	u.Suite.NotNil(updateResp)
	u.Suite.Equal(updatedName, updateResp.FirstName)
	u.Suite.Equal(updatedBio, updateResp.Bio)

	//Update refresh token
	updatedRefresh := gofakeit.PhraseNoun()
	resp, err := u.Repository.UpdateRefresh(context.Background(), &pb.UpdateRefreshReq{
		UserId: user.Id,
		RefreshToken: updatedRefresh,
	})
	u.Suite.NoError(err)
	u.Suite.Equal(resp.RefreshToken, updatedRefresh)

	//CheckField
	checkResp, err := u.Repository.CheckField(ctx, &pb.CheckFieldReq{
		Field: "email",
		Value:  user.Email,
	})
	u.Suite.NoError(err)
	u.Suite.NotNil(checkResp)
	u.Suite.Equal(checkResp.Status, true)

	//Delete user
	_, err = u.Repository.Delete(ctx, userId)
	u.Suite.NoError(err)
}

func (u *UserTestSuite) TearDownSuite() {
	u.CleanupFunc()
}

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
