package user

import (
	"assessment2/pkg/user/controller"
	"assessment2/pkg/user/usecase"
	"assessment2/pkg/user/repository"

	"gorm.io/gorm"
)

func InitHttpUserController(db *gorm.DB) *controller.UserHTTPController {
	repo := repository.InitRepositoryUser(db)
	uc := usecase.InitUsecaseUser(repo)
	controller := controller.InitControllerUser(uc)

	return controller
}
