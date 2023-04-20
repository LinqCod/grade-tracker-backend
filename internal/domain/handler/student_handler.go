package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/linqcod/grade-tracker-backend/internal/domain/entity"
	"github.com/linqcod/grade-tracker-backend/internal/domain/service"
	"github.com/linqcod/grade-tracker-backend/pkg/jwttoken"
	"github.com/linqcod/grade-tracker-backend/pkg/response"
	"net/http"
)

type StudentHandler struct {
	studentService service.StudentService
}

func NewStudentHandler(studentService service.StudentService) *StudentHandler {
	return &StudentHandler{
		studentService: studentService,
	}
}

func (h *StudentHandler) RegisterStudent(c *gin.Context) {
	var registerStudent entity.StudentRegistrationDTO
	err := c.ShouldBindJSON(&registerStudent)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.studentService.SaveStudent(&registerStudent)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	response.ResponseCreated(c, result)
}

func (h *StudentHandler) Login(c *gin.Context) {
	var loginDTO entity.StudentLoginDTO

	err := c.ShouldBindJSON(&loginDTO)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	validateStudent, err := h.studentService.GetStudentByEmail(loginDTO)
	if err != nil {
		// TODO: подумать, какую ошибку грамотней отдавать
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := jwttoken.CreateToken(validateStudent.Id)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	studentData := map[string]interface{}{
		"access_token":    token.AccessToken,
		"expiration_time": token.ExpirationTimeInUnix,
		"student_id":      validateStudent.Id,
		"email":           validateStudent.Email,
		"first_name":      validateStudent.FirstName,
		"second_name":     validateStudent.SecondName,
		"patronymic":      validateStudent.Patronymic,
		"group":           validateStudent.Group,
	}

	response.ResponseOKWithData(c, studentData)
}
