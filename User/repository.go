package User

import (
	
	"context"
	"To_Do_App/models"

)

type Repository interface {

	Store(ctx context.Context, a *models.UserDB) error
	Update(ctx context.Context, a *models.UserDB) error
	GetAllUser(ctx context.Context) ([]*models.UserDB, error)

}