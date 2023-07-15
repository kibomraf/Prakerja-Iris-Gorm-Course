package students

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Save(student Student) (Student, error)
	CheckEmail(email string) (Student, error)
}
type repository struct {
	db *gorm.DB
}

func StudentRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// repository connect to db, for save new user.
func (r *repository) Save(student Student) (Student, error) {
	err := r.db.Create(&student).Error
	if err != nil {
		return Student{}, err
	}
	return student, nil
}
func (r *repository) CheckEmail(email string) (Student, error) {
	var student Student
	err := r.db.Where("email = ?", email).Find(&student).Error
	if err != nil {
		return Student{}, errors.New("email has benn created")
	}
	return student, nil
}
