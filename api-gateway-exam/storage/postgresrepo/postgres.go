package postgresrepo

import "exam/api-gateway/api/handlers/models"

type AdminStorageI interface {
	Create(admin *models.AdminResp) error
	Delete(userName, password string) error
	Get(userName string) (string, string, bool, error)
}
