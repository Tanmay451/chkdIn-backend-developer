package routes

import (
	"chkdIn-backend-developer/controllers"
	"os"

	"github.com/gin-gonic/gin"
)

// AddRoutes is responsible for adding all the routes so the server can handle
// new routes. this means that we can reuse this function for multiple prefixes.
// prefixes like job_portal are necessary for legacy url handling.
func AddRoutes(router *gin.RouterGroup) {
	api := router.Group("/api")
	{
		api.GET("/user-list", TokenAuth, controllers.GetUserList)
		api.POST("/register", controllers.Register)
		api.POST("/authenticate", controllers.Authenticate)
		api.PATCH("/update-user-status", TokenAuth, controllers.UpdateUserStatus) // need to refactor
		api.DELETE("delete-user", TokenAuth, controllers.DeleteUser)

		api.GET("/auth-failed", controllers.AuthenticateFailed)
	}

}

// SetupRouter sets up routes
func SetupRouter() *gin.Engine {
	router := gin.Default()
	gin.SetMode(os.Getenv("GIN_MODE"))

	AddRoutes(&router.RouterGroup)

	// router.NoRoute(controllers.PageNotFound)

	return router
}
