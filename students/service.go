package students

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	CreateStudent(input Input) (Student, error)
	CheckEmailAvailibity(email string) (bool, error)
	LoginStudent(input Login) (Student, error)
}
type service struct {
	repository Repository
}

func StudentService(repository Repository) *service {
	return &service{repository}
}

// logic bussines : create new user.
func (s *service) CreateStudent(input Input) (Student, error) {
	//mapping input
	student := Student{
		Name:        input.Name,
		Email:       input.Email,
		Address:     input.Address,
		NoHandphone: input.NoHandphone,
		Created_at:  time.Now(),
		Updated_at:  time.Now(),
	}
	//hash password
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return Student{}, err
	}
	student.Password_hash = string(password)
	//save new user
	newStudent, err := s.repository.Save(student)
	if err != nil {
		return Student{}, err
	}
	return newStudent, nil
}

// check email
func (s *service) CheckEmailAvailibity(email string) (bool, error) {
	//mapping input
	//call repository input
	student, err := s.repository.CheckEmail(email)
	if student.Id > 0 || err != nil {
		return false, err
	}
	return true, nil
}

// login logic
func (s *service) LoginStudent(input Login) (Student, error) {
	//mapping input login
	email, password := input.Email, input.Password
	//calling logic bussines
	student, err := s.repository.FindByEmail(email)
	if err != nil {
		return Student{}, errors.New("email is wrong")
	}
	//compare input password and real password
	err = bcrypt.CompareHashAndPassword([]byte(student.Password_hash), []byte(password))
	if err != nil {
		return Student{}, errors.New("password is wrong")
	}
	return student, nil
}
