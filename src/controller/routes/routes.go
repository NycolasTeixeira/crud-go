package routes

import (
	"github.com/NycolasTeixeira/crud-go/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/getUserByid/:userId", controller.FindUserByEmail)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser/", controller.CreateUserById)
	r.PUT("/updateUser/:userId", controller.UpdateUserById)
	r.DELETE("/deleteUser/:userId", controller.DeleteUserById)
}
