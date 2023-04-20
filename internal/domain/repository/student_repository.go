package repository

import (
	"context"
	"database/sql"
	"github.com/linqcod/grade-tracker-backend/internal/domain/entity"
)

const (
	SaveStudentQuery       = "INSERT INTO students (email, password, first_name, second_name, patronymic, role, group_id) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	UpdateStudentQuery     = "UPDATE students SET email = $2, password = $3, first_name = $4, second_name = $5, patronymic = $6, group_id = $7 WHERE id = $1"
	DeleteStudentQuery     = "DELETE FROM students WHERE id = $1"
	GetStudentDetailsQuery = `
		SELECT groups.id, groups.title, students.id, students.email, students.password, students.first_name, students.second_name, students.patronymic 
		FROM groups 
		INNER JOIN students 
		ON students.group_id=groups.id 
		WHERE students.id = $1 
		LIMIT 1;
	`
	GetStudentByEmailQuery = `
		SELECT groups.id, groups.title, students.id, students.email, students.password, students.first_name, students.second_name, students.patronymic
		FROM groups 
		INNER JOIN students 
		ON students.group_id=groups.id 
		WHERE students.email = $1 
		LIMIT 1;
	`
	GetStudentGroupTitleQuery = "SELECT title FROM groups WHERE id = $1 LIMIT 1"
)

type StudentRepository interface {
	SaveStudent(student *entity.Student) (*entity.Student, error)
	UpdateStudent(student *entity.Student) (*entity.Student, error)
	DeleteStudent(studentId int64) error
	GetStudentDetails(studentId int64) (*entity.Student, error)
	GetStudentByEmail(login entity.StudentLoginDTO) (*entity.Student, error)
}

type StudentRepositoryImpl struct {
	ctx context.Context
	db  *sql.DB
}

func NewStudentRepository(ctx context.Context, db *sql.DB) StudentRepository {
	return &StudentRepositoryImpl{
		ctx: ctx,
		db:  db,
	}
}

func (s StudentRepositoryImpl) SaveStudent(student *entity.Student) (*entity.Student, error) {
	if err := s.db.QueryRowContext(
		s.ctx,
		SaveStudentQuery,
		student.Email,
		student.Password,
		student.FirstName,
		student.SecondName,
		student.Patronymic,
		student.Role,
		student.Group.Id,
	).Err(); err != nil {
		return nil, err
	}

	return student, nil
}

func (s StudentRepositoryImpl) UpdateStudent(student *entity.Student) (*entity.Student, error) {
	if err := s.db.QueryRowContext(
		s.ctx,
		UpdateStudentQuery,
		student.Id,
		student.Email,
		student.Password,
		student.FirstName,
		student.SecondName,
		student.Patronymic,
		student.Group.Id,
	).Err(); err != nil {
		return nil, err
	}

	if err := s.db.QueryRowContext(s.ctx, GetStudentGroupTitleQuery, student.Group.Id).
		Scan(&student.Group.Title); err != nil {
		return nil, err
	}

	return student, nil
}

func (s StudentRepositoryImpl) DeleteStudent(studentId int64) error {
	if err := s.db.QueryRowContext(s.ctx, DeleteStudentQuery, studentId).Err(); err != nil {
		return err
	}

	return nil
}

func (s StudentRepositoryImpl) GetStudentDetails(studentId int64) (*entity.Student, error) {
	var student entity.Student
	if err := s.db.QueryRowContext(s.ctx, GetStudentDetailsQuery, studentId).Scan(
		&student.Group.Id,
		&student.Group.Title,
		&student.Id,
		&student.Email,
		&student.Password,
		&student.FirstName,
		&student.SecondName,
		&student.Patronymic,
	); err != nil {
		return nil, err
	}

	return &student, nil
}

func (s StudentRepositoryImpl) GetStudentByEmail(login entity.StudentLoginDTO) (*entity.Student, error) {
	var student entity.Student
	if err := s.db.QueryRowContext(s.ctx, GetStudentByEmailQuery, login.Email).Scan(
		&student.Group.Id,
		&student.Group.Title,
		&student.Id,
		&student.Email,
		&student.Password,
		&student.FirstName,
		&student.SecondName,
		&student.Patronymic,
	); err != nil {
		return nil, err
	}

	return &student, nil
}
