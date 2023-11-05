package services

import "registration/models"

var users []models.User

func RegisterUser(user models.User) error {
	users = append(users, user)
	return nil
}

func GetUserByUsername(username string) (models.User, error) {
	for _, user := range users {
		if user.Username == username {
			return user, nil
		}
	}
	return models.User{}, nil
}
