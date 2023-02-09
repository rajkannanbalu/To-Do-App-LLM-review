package User

import (

	"context"
	"To_Do_App/models"
	
)




type Usecase interface {

	Store(ctx context.Context, a *models.UserDB) error
	Update(ctx context.Context, a *models.UserDB) error	
	GetAllUser(c context.Context) ([]*models.UserDB, error)

}