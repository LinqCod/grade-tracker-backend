package service

import (
	"fmt"
	"github.com/linqcod/grade-tracker-backend/internal/domain/entity"
	"github.com/linqcod/grade-tracker-backend/internal/domain/repository"
	"github.com/linqcod/grade-tracker-backend/pkg/security"
	"golang.org/x/crypto/bcrypt"
)

type AdminService interface {
	GetAdminByEmail(login entity.AdminLoginDTO) (*entity.AdminDTO, error)
}

type AdminServiceImpl struct {
	adminRepo repository.AdminRepository
}

func NewAdminService(adminRepo repository.AdminRepository) AdminService {
	return &AdminServiceImpl{
		adminRepo: adminRepo,
	}
}

func (a AdminServiceImpl) GetAdminByEmail(login entity.AdminLoginDTO) (*entity.AdminDTO, error) {
	result, err := a.adminRepo.GetAdminByEmail(login)
	if err != nil {
		return nil, err
	}

	err = security.VerifyPassword(result.Password, login.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, fmt.Errorf("incorrect password. Error: %s", err.Error())
	}

	return &entity.AdminDTO{
		Id:         result.Id,
		Email:      result.Email,
		FirstName:  result.FirstName,
		SecondName: result.SecondName,
		Patronymic: result.Patronymic,
	}, nil
}
