package Task

import (

	"context"
	"To_Do_App/models"
)


type Repository interface {

	Delete(ctx context.Context, task_id int64) error
	GetByID(ctx context.Context, task_id int64)(*models.TaskDB, error)
	GetByUserID(ctx context.Context, user_id int64) ([]*models.TaskDB, error)
	GetAllTask(ctx context.Context) ([]*models.TaskDB, error)
	Store(ctx context.Context, a *models.TaskDB) error
	Update(ctx context.Context, a *models.TaskDB) error
	UpdateDone(ctx context.Context, task_id int64, task_status *models.TaskDB) error
	
}