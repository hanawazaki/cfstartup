package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)
	user.Occupation = input.Occupation
	user.Role = "user"
	// input. diambil dari package input yang berisi struct RegisterUserInput
	// di proses ini struct input dimasukan ke struct user
	newUser, err := s.repository.Save(user)
	if err != nil {
		return user, err
	}
	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil

}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, nil
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil

}

// newbie notes

// interface RegisterUser dibuat dengan param input RegisterUserInput yang sudah dipass dari handler
// dengan nilai balikan User / error
// service mewakili fitur/logic sehingga nama service biasanya kata kerja seperti
// create/update/register dll

/*
service struct memiliki dependency terhadap repository
karena service akan melakukan save ke db "melalui" Repository
maka obj di dalam nya adalah repository dengan interface Repository
*/

// newService dibuat untuk mempassing nilai yang diproses di RegisterUser
// ke dalam struct service dengan nilai {repository}
