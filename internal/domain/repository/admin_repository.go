package repository

import (
	"context"
	"database/sql"
	"github.com/linqcod/grade-tracker-backend/internal/domain/entity"
)

const (
	GetAdminByEmailQuery = `
		SELECT id, email, password, first_name, second_name, patronymic
		FROM admins
		WHERE email = $1 
		LIMIT 1;
	`
)

type AdminRepository interface {
	GetAdminByEmail(login entity.AdminLoginDTO) (*entity.Admin, error)
}

type AdminRepositoryImpl struct {
	ctx context.Context
	db  *sql.DB
}

func NewAdminRepository(ctx context.Context, db *sql.DB) AdminRepository {
	return &AdminRepositoryImpl{
		ctx: ctx,
		db:  db,
	}
}

func (a AdminRepositoryImpl) GetAdminByEmail(login entity.AdminLoginDTO) (*entity.Admin, error) {
	var admin entity.Admin
	if err := a.db.QueryRowContext(a.ctx, GetAdminByEmailQuery, login.Email).Scan(
		&admin.Id,
		&admin.Email,
		&admin.Password,
		&admin.FirstName,
		&admin.SecondName,
		&admin.Patronymic,
	); err != nil {
		return nil, err
	}

	return &admin, nil
}
