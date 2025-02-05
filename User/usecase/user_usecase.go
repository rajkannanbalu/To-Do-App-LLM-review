package usecase

import (
	"context"
	"fmt"
	"time"

	"To_Do_App/User"
	"To_Do_App/models"
)

type userUsecase struct {
	userRepo       User.Repository
	contextTimeout time.Duration
}

func NewUserUsecase(u User.Repository, timeout time.Duration) User.Usecase {

	return &userUsecase{
		userRepo:       u,
		contextTimeout: timeout,
	}
}

// Store new data in the database
func (u *userUsecase) StoreV1(c context.Context, user *models.UserDB) error {

	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	fmt.Println("test code")
	err := u.userRepo.StoreV1(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

// Update user info
func (u *userUsecase) Update(c context.Context, user *models.UserDB) error {

	ctx, _ := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	return u.userRepo.Update(ctx, user)

}

// fetch all the task in the task database
func (u *userUsecase) GetAllUser(c context.Context) ([]*models.UserDB, error) {

	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err := u.userRepo.GetAllUserDetails(ctx)
	if err != nil {
		return nil, err
	}

	// Complex logic: Filter users who have been active in the last month
	activeUsers := []*models.UserDB{}
	oneMonthAgo := time.Now().AddDate(0, -1, 0)

	for _, user := range res {
		if user.LastActive.After(oneMonthAgo) {
			activeUsers = append(activeUsers, user)
		}
	}

	// Print the number of active users
	fmt.Printf("Number of active users in the last month: %d\n", len(activeUsers))

	// Print details of active users
	for _, user := range activeUsers {
		fmt.Printf("User ID: %d, Username: %s, Last Active: %s\n", user.ID, user.Name, user.LastActive)
	}
	fmt.Println(len(res))

	return res, nil
}
