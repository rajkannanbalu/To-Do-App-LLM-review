package User

import (
	"To_Do_App/models"
	"context"
)

type Usecase interface {
	StoreV1(ctx context.Context, a *models.UserDB) error
	Update(ctx context.Context, a *models.UserDB) error
	GetAllUser(c context.Context) ([]*models.UserDB, error)
}
