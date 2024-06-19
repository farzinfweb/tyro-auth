package auth

import (
	"authn/domain"
	"authn/repo"
	"authn/util/hasher"
	"context"
	"errors"
)

func AttemptLogin(username string, password string, repos repo.IUserRepo) (*domain.User, error) {
	user, err := repos.FindByUsername(context.TODO(), username)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := hasher.Generate(password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}
	if user.Password != hashedPassword {
		return nil, errors.New("username/password is incorrect")
	}

	return user, nil
}
