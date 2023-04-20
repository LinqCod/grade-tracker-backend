package service

import (
	"fmt"
	"github.com/linqcod/grade-tracker-backend/internal/domain/entity"
	"github.com/linqcod/grade-tracker-backend/internal/domain/repository"
	"github.com/linqcod/grade-tracker-backend/pkg/security"
	"golang.org/x/crypto/bcrypt"
)

type StudentService interface {
	SaveStudent(registration *entity.StudentRegistrationDTO) (*entity.StudentDTO, error)
	UpdateStudent(student *entity.Student) (*entity.StudentDTO, error)
	DeleteStudent(studentId int64) error
	GetStudentDetails(studentId int64) (*entity.StudentDTO, error)
	GetStudentByEmail(login entity.StudentLoginDTO) (*entity.StudentDTO, error)
}

type StudentServiceImpl struct {
	studentRepo repository.StudentRepository
}

func NewStudentService(studentRepo repository.StudentRepository) StudentService {
	return &StudentServiceImpl{
		studentRepo: studentRepo,
	}
}

func (s StudentServiceImpl) SaveStudent(registration *entity.StudentRegistrationDTO) (*entity.StudentDTO, error) {
	student := entity.Student{
		User: entity.User{
			Email:      registration.Email,
			FirstName:  registration.FirstName,
			SecondName: registration.SecondName,
			Patronymic: registration.Patronymic,
			Role:       entity.StudentRole,
		},
		Group: registration.Group,
	}

	password, err := security.Hash(registration.Password)
	if err != nil {
		return nil, err
	}

	student.Password = password

	result, err := s.studentRepo.SaveStudent(&student)
	if err != nil {
		return nil, err
	}

	return &entity.StudentDTO{
		Id:         result.Id,
		Email:      result.Email,
		FirstName:  result.FirstName,
		SecondName: result.SecondName,
		Patronymic: result.Patronymic,
		Group:      result.Group,
	}, nil
}

func (s StudentServiceImpl) UpdateStudent(student *entity.Student) (*entity.StudentDTO, error) {
	password, err := security.Hash(student.Password)
	if err != nil {
		return nil, err
	}

	student.Password = password

	result, err := s.studentRepo.UpdateStudent(student)
	if err != nil {
		return nil, err
	}

	return &entity.StudentDTO{
		Id:         result.Id,
		Email:      result.Email,
		FirstName:  result.FirstName,
		SecondName: result.SecondName,
		Patronymic: result.Patronymic,
		Group:      result.Group,
	}, nil
}

func (s StudentServiceImpl) DeleteStudent(studentId int64) error {
	if err := s.studentRepo.DeleteStudent(studentId); err != nil {
		return err
	}

	return nil
}

func (s StudentServiceImpl) GetStudentDetails(studentId int64) (*entity.StudentDTO, error) {
	result, err := s.studentRepo.GetStudentDetails(studentId)
	if err != nil {
		return nil, err
	}

	return &entity.StudentDTO{
		Id:         result.Id,
		Email:      result.Email,
		FirstName:  result.FirstName,
		SecondName: result.SecondName,
		Patronymic: result.Patronymic,
		Group:      result.Group,
	}, nil
}

func (s StudentServiceImpl) GetStudentByEmail(login entity.StudentLoginDTO) (*entity.StudentDTO, error) {
	result, err := s.studentRepo.GetStudentByEmail(login)
	if err != nil {
		return nil, err
	}

	err = security.VerifyPassword(result.Password, login.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, fmt.Errorf("incorrect password. Error: %s", err.Error())
	}

	return &entity.StudentDTO{
		Id:         result.Id,
		Email:      result.Email,
		FirstName:  result.FirstName,
		SecondName: result.SecondName,
		Patronymic: result.Patronymic,
		Group:      result.Group,
	}, nil
}
