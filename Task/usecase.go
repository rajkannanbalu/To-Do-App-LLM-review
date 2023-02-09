package Task

import (
	
	"context"
	"To_Do_App/models"
	"time"
)

// This is for UpdateDone Function
type TaskPatchReq struct {
	Status    string     `json:"task_status,omitempty"`
	UpdatedAt *time.Time `json:"updated_at"`
}


type Usecase interface {

	Delete(ctx context.Context, task_id int64) error
	GetByID(ctx context.Context, task_id int64)(*models.TaskDB, error)
	GetByUserID(ctx context.Context, user_id int64) ([]*models.TaskDB, error)
	GetAllTask(ctx context.Context) ([]*models.TaskDB, error)
	Store(ctx context.Context, a *models.TaskDB) error
	Update(ctx context.Context, a *models.TaskDB) error
	UpdateDone(ctx context.Context, task_id int64, statusReq *TaskPatchReq) error

}