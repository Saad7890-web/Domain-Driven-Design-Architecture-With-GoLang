package auth

import (
	"context"
	"errors"

	"github.com/Saad7890-web/internal/domain/user"
	"github.com/Saad7890-web/pkg/hashPassword"
	"github.com/google/uuid"
)


var (
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrUserAlreadyExists = errors.New("user already exists")
)

type Service struct {
	userRepo user.Repository
}

func NewService(userRepo user.Repository)*Service{
	return &Service{userRepo: userRepo}
}

func (s *Service) Signup(ctx context.Context, email, password, name string)(*user.User, error){
	existing,err:=s.userRepo.FindByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	if existing != nil {
		return nil, ErrUserAlreadyExists
	}

	hashedPassword, err := hashPassword.HashPassword(password)

	if err != nil {
		return nil, err
	}

	u := &user.User{
		ID: uuid.New(),
		Email: email,
		Password: hashedPassword,
		Name: name,
	}

	if err:= s.userRepo.Create(ctx, u); err!=nil{
		return nil, err
	}

	return u, nil
}

func (s *Service) Login(
	ctx context.Context,
	email, password string,
) (*user.User, error) {

	u, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, ErrInvalidCredentials
	}

	if err := hashPassword.ComparePassword(u.Password, password); err != nil {
		return nil, ErrInvalidCredentials
	}

	return u, nil
}