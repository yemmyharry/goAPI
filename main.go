package main

import (
	"github.com/gin-gonic/gin"
	"goAPI/config"
	"goAPI/controller"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.Dbsetup()
	authController controller.AuthController = controller.NewAuthController()
)

func main() {
	defer config.CloseDBConnection(db)
	r := gin.Default()

	authRoutes := r.Group("/api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
	r.Run()

}
