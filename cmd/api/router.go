package api

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/linqcod/grade-tracker-backend/cmd/api/middleware"
	"github.com/linqcod/grade-tracker-backend/internal/domain/entity"
	"github.com/linqcod/grade-tracker-backend/internal/domain/handler"
	"github.com/linqcod/grade-tracker-backend/internal/domain/repository"
	"github.com/linqcod/grade-tracker-backend/internal/domain/service"
)

func InitRouter(ctx context.Context, db *sql.DB) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	adminRepository := repository.NewAdminRepository(ctx, db)
	adminService := service.NewAdminService(adminRepository)
	adminHandler := handler.NewAdminHandler(adminService)

	studentRepository := repository.NewStudentRepository(ctx, db)
	studentService := service.NewStudentService(studentRepository)
	studentHandler := handler.NewStudentHandler(studentService)

	api := router.Group("/api/v1")
	{
		admins := api.Group("/admins")
		{
			admins.POST("/login", adminHandler.Login)
		}
		students := api.Group("/students")
		{
			students.POST("/", middleware.AuthMiddleware(), middleware.RoleCheckMiddleware(entity.AdminRole), studentHandler.RegisterStudent)
			students.POST("/login", studentHandler.Login)
		}
	}

	return router
}
