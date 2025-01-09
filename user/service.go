package user

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(request RegisterUserInput) (User, error)
	Login(request LoginInput) (User, error)
	IsEmailAvailable(email string) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(request RegisterUserInput) (User, error) {
	user := User{}
	user.Name = request.Name
	user.Email = request.Email
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err.Error())
		return user, err
	}
	user.PasswordHash = string(passwordHash)
	user.Role = request.Role

	newUser, err := s.repository.Save(user)
	if err != nil {
		fmt.Println(err.Error())
		return user, err
	}

	return newUser, nil
}

func (s *service) Login(request LoginInput) (User, error) {
	email := request.Email
	password := request.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		fmt.Println(err.Error())
		return user, errors.New("Faile found email!")
	}
	if user.Id == 0 {
		return user, errors.New("No user foud !")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		fmt.Println(err.Error())
		return user, errors.New("Password not match!")
	}

	return user, nil
}

func (s *service) IsEmailAvailable(email string) (bool, error) {
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	if user.Id == 0 {
		return true, nil
	}
	return false, nil
}
