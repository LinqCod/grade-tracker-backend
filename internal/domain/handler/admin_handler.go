package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/linqcod/grade-tracker-backend/internal/domain/entity"
	"github.com/linqcod/grade-tracker-backend/internal/domain/service"
	"github.com/linqcod/grade-tracker-backend/pkg/jwttoken"
	"github.com/linqcod/grade-tracker-backend/pkg/response"
	"net/http"
)

type AdminHandler struct {
	adminService service.AdminService
}

func NewAdminHandler(adminService service.AdminService) *AdminHandler {
	return &AdminHandler{
		adminService: adminService,
	}
}

func (h *AdminHandler) Login(c *gin.Context) {
	var loginDTO entity.AdminLoginDTO

	err := c.ShouldBindJSON(&loginDTO)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	validateAdmin, err := h.adminService.GetAdminByEmail(loginDTO)
	if err != nil {
		// TODO: подумать, какую ошибку грамотней отдавать
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := jwttoken.CreateToken(validateAdmin.Id)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	adminData := map[string]interface{}{
		"access_token":    token.AccessToken,
		"expiration_time": token.ExpirationTimeInUnix,
		"admin_id":        validateAdmin.Id,
		"email":           validateAdmin.Email,
		"first_name":      validateAdmin.FirstName,
		"second_name":     validateAdmin.SecondName,
		"patronymic":      validateAdmin.Patronymic,
	}

	response.ResponseOKWithData(c, adminData)
}
