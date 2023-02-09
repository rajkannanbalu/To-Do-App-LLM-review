package repository

import (

	"context"
	"database/sql"
	"fmt"
	"To_Do_App/Task"
	"To_Do_App/models"
	"github.com/sirupsen/logrus"


)

const (
	timeFormat = "2006-01-02T15:04:05.999Z07:00" // reduce precision from RFC3339Nano as date format
)

type mysqlTaskRepo struct{
	Conn *sql.DB
}



func NewMysqlTaskRepo(db *sql.DB) Task.Repository{

	return &mysqlTaskRepo{
		Conn: db,
	}
}

// Delete the task for the specific task id
func (m *mysqlTaskRepo) Delete(ctx context.Context, task_id int64) error{

	query := "DELETE FROM tasks WHERE id = ?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil{
		return err
	}

	res, err := stmt.ExecContext(ctx, task_id)
	if err != nil {
		return nil
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAfected != 1{
		err = fmt.Errorf("Weird Behaviour. Total Affected %d", rowsAfected)
		return err
	}

	return nil
}

// Get a specific Task using task id
func(m *mysqlTaskRepo) GetByID(ctx context.Context, task_id int64) (res *models.TaskDB, err error){

	query := "SELECT id, name, status, comment, updated_at, created_at, user_id FROM tasks WHERE id = ?"

	list, err := m.fetch(ctx, query, task_id)
	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return nil, models.ErrNotFound
	}

	return res, nil
}

func (m *mysqlTaskRepo) GetAllTask(ctx context.Context) ([]*models.TaskDB, error) {

	rows, err := m.Conn.Query("SELECT id, name, status, comment, updated_at, created_at, user_id FROM tasks;")

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	result := make([]*models.TaskDB, 0)
	for rows.Next() {
		t := new(models.TaskDB)
		err = rows.Scan(
			&t.ID,
			&t.Name,
			&t.Status,
			&t.Comment,
			&t.UpdatedAt,
			&t.CreatedAt,
			&t.UserID,
		)
		result = append(result, t)
	}

	return result, nil
}

// Fetch all the data using user_id
 func (m *mysqlTaskRepo) GetByUserID(ctx context.Context, user_id int64) ([]*models.TaskDB, error){

	query := "SELECT id, name, status, comment, updated_at, created_at, user_id FROM tasks WHERE user_id = ?"

	rows, err := m.Conn.QueryContext(ctx, query, user_id)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	result := make([]*models.TaskDB, 0)
	for rows.Next() {
		t := new(models.TaskDB)
		err = rows.Scan(
			&t.ID,
			&t.Name,
			&t.Status,
			&t.Comment,
			&t.UpdatedAt,
			&t.CreatedAt,
			&t.UserID,
		)
		result = append(result, t)
	}

	return result, nil
}


// Store new data in the database
 func (m *mysqlTaskRepo) Store(ctx context.Context, task *models.TaskDB) error{

	query := "INSERT INTO tasks (name, status, comment, updated_at, created_at, user_id) VALUES (?, ?, ?, ?, ?, ?);"
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil{
		return err
	}

	res, err := stmt.ExecContext(ctx, task.Name, task.Status, task.Comment, task.UpdatedAt, task.CreatedAt, task.UserID)
	if err!= nil{
		return err
	}

	lastID, err := res.LastInsertId()
	if err != nil{
		return err
	}

	task.ID = lastID
	return nil

	}	

// Update the existing task 
func (m *mysqlTaskRepo) Update(ctx context.Context, task *models.TaskDB) error{

	query := "UPDATE tasks SET name=?, status=?, comment=?, updated_at=? WHERE id =?;"

	stmt, err:= m.Conn.PrepareContext(ctx, query)
	if err != nil{
		return err
	}

	res, err := stmt.ExecContext(ctx, task.Name, task.Status, task.Comment, task.UpdatedAt, task.ID)
	if err!= nil{
		return err
	}

	affect, err := res.RowsAffected()
	if err!= nil{
		return err
	}

	if affect != 1{
		err = fmt.Errorf("Weird Behaviour. Total Affected: %d", affect)
		return err
	}

	return nil
	

}

//Patch --> only status update
func (m *mysqlTaskRepo) UpdateDone(ctx context.Context, task_id int64, task *models.TaskDB) error {

	query := "UPDATE tasks SET status=?, updated_at=? WHERE id =?"

	stmt, err:= m.Conn.PrepareContext(ctx, query)
	if err != nil{
		return nil
	}

	res, err := stmt.ExecContext(ctx, task.Status, task.UpdatedAt, task_id)
	if err!= nil{
		return err
	}

	affect, err := res.RowsAffected()
	if err!= nil{
		return err
	}

	if affect != 1{
		err = fmt.Errorf("Weird Behaviour. Total Affected: %d", affect)
		return err
	}

	return nil

}

// Utility Function
func (m *mysqlTaskRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.TaskDB, error) {

	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	result := make([]*models.TaskDB, 0)
	for rows.Next() {
		t := new(models.TaskDB)
		err = rows.Scan(
			&t.ID,
			&t.Name,
			&t.Status,
			&t.Comment,
			&t.UpdatedAt,
			&t.CreatedAt,
			&t.UserID,
		)
		result = append(result, t)
	}
	
	return result, nil
}




