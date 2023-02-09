package repository

import (

	"context"
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"

	"To_Do_App/User"
	"To_Do_App/models"
	
)

type mysqlUserRepo struct {
	Conn *sql.DB
}

func NewMysqlUserRepo(db *sql.DB) User.Repository{

	return &mysqlUserRepo{
		Conn: db,
	}
}

// Store new data in the database
func (m *mysqlUserRepo) Store(ctx context.Context, user *models.UserDB) error{

	query := "INSERT INTO users (name) VALUES (?);"
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil{
		return err
	}

	res, err := stmt.ExecContext(ctx, user.Name)
	if err!= nil{
		return err
	}

	lastID, err := res.LastInsertId()
	if err != nil{
		return err
	}

	user.ID = lastID
	return nil

}	

// Update the existing user
func (m *mysqlUserRepo) Update(ctx context.Context, user *models.UserDB) error{

	query := "UPDATE users SET name=? WHERE id =?;"

	stmt, err:= m.Conn.PrepareContext(ctx, query)
	if err != nil{
		return err
	}

	res, err := stmt.ExecContext(ctx, user.Name, user.ID)
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

// Get all the user in the database
func (m *mysqlUserRepo) GetAllUser(ctx context.Context) ([]*models.UserDB, error) {

	rows, err := m.Conn.Query("SELECT id, name FROM users;")

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

	result := make([]*models.UserDB, 0)
	for rows.Next() {
		u := new(models.UserDB)
		err = rows.Scan(
			&u.ID,
			&u.Name,
		)
		result = append(result, u)
	}
	return result, nil
}




